package models

import (
	"fmt"
	"os"
	"sync"
	"time"

	"code.cloudfoundry.org/bbs/events"
	bbsmodels "code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/copilot/api"
	"code.cloudfoundry.org/lager"
)

type sets struct {
	External *api.BackendSet
	Internal *api.BackendSet
}

type store struct {
	sync.RWMutex
	content map[DiegoProcessGUID]sets
}

func (s *store) Insert(guid DiegoProcessGUID, isInternal bool, additionalBackend *api.Backend) {
	if additionalBackend == nil {
		return
	}

	s.Lock()
	if _, ok := s.content[guid]; !ok {
		s.content[guid] = sets{
			External: &api.BackendSet{},
			Internal: &api.BackendSet{},
		}
	}

	backends := s.content[guid].External.Backends
	if isInternal {
		backends = s.content[guid].Internal.Backends
	}
	s.Unlock()

	for _, backend := range backends {
		if fmt.Sprintf("%s:%d", backend.Address, backend.Port) == fmt.Sprintf("%s:%d", additionalBackend.Address, additionalBackend.Port) {
			return
		}
	}

	s.Lock()
	if isInternal {
		s.content[guid].Internal.Backends = append(s.content[guid].Internal.Backends, additionalBackend)
	} else {
		s.content[guid].External.Backends = append(s.content[guid].External.Backends, additionalBackend)
	}
	s.Unlock()
}

func (s *store) Remove(guid DiegoProcessGUID) {
	s.Lock()
	defer s.Unlock()

	delete(s.content, guid)
}

type BackendSetRepo struct {
	bbs    bbsEventer
	logger lager.Logger
	ticker <-chan time.Time
	store  store
}

//go:generate counterfeiter -o fakes/bbs_eventer.go --fake-name BBSEventer . bbsEventer
type bbsEventer interface {
	SubscribeToEvents(logger lager.Logger) (events.EventSource, error)
	ActualLRPGroups(lager.Logger, bbsmodels.ActualLRPFilter) ([]*bbsmodels.ActualLRPGroup, error)
}

func NewBackendSetRepo(bbs bbsEventer, logger lager.Logger, ticker <-chan time.Time) *BackendSetRepo {
	return &BackendSetRepo{
		bbs:    bbs,
		logger: logger,
		ticker: ticker,
		store: store{
			content: make(map[DiegoProcessGUID]sets),
		},
	}
}

func (b *BackendSetRepo) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	stop := make(chan struct{})

	eventSource, err := b.bbs.SubscribeToEvents(b.logger)
	if err != nil {
		return err
	}

	go b.collectEvents(stop, eventSource)
	go b.reconcileLRPs(stop, b.ticker)

	close(ready)

	for {
		select {
		case <-signals:
			close(stop)
			return nil
		}
	}
}

func (b *BackendSetRepo) Get(guid DiegoProcessGUID) *api.BackendSet {
	b.store.RLock()
	defer b.store.RUnlock()
	return b.store.content[guid].External
}

func (b *BackendSetRepo) GetInternalBackends(guid DiegoProcessGUID) *api.BackendSet {
	b.store.RLock()
	defer b.store.RUnlock()
	return b.store.content[guid].Internal
}

func (b *BackendSetRepo) collectEvents(stop <-chan struct{}, eventSource events.EventSource) {
	for {
		select {
		case <-stop:
			b.logger.Info("events-exit")
			return
		default:
			event, err := eventSource.Next()
			if err != nil {
				b.logger.Debug("events-next", lager.Data{"events-error": err.Error()})
				continue
			}

			switch event.EventType() {
			case bbsmodels.EventTypeActualLRPCreated:
				createdEvent := event.(*bbsmodels.ActualLRPCreatedEvent)
				instance := createdEvent.GetActualLrpGroup().GetInstance()
				ex, in := processInstance(instance)
				guid := DiegoProcessGUID(instance.ActualLRPKey.GetProcessGuid())
				b.store.Insert(guid, false, ex)
				b.store.Insert(guid, true, in)

			case bbsmodels.EventTypeActualLRPRemoved:
				removedEvent := event.(*bbsmodels.ActualLRPRemovedEvent)
				guid := DiegoProcessGUID(removedEvent.GetActualLrpGroup().GetInstance().ActualLRPKey.GetProcessGuid())
				b.store.Remove(guid)
			default:
				b.logger.Debug("unhandled-event-type")
				return
			}
		}
	}
}

func (b *BackendSetRepo) reconcileLRPs(stop <-chan struct{}, ticker <-chan time.Time) {
	for {
		select {
		case <-ticker:
			groups, err := b.bbs.ActualLRPGroups(b.logger, bbsmodels.ActualLRPFilter{})
			if err != nil {
				b.logger.Debug("lrp-groups-error", lager.Data{"lrp-groups-error": err.Error()})
				continue
			}

			// not locking replacement store - no other goroutine can update it
			replaceStore := store{content: make(map[DiegoProcessGUID]sets)}
			for _, group := range groups {
				ex, in := processInstance(group.Instance)
				guid := DiegoProcessGUID(group.GetInstance().ActualLRPKey.GetProcessGuid())
				replaceStore.Insert(guid, false, ex)
				replaceStore.Insert(guid, true, in)
			}

			b.store.Lock()
			b.store.content = replaceStore.content
			b.store.Unlock()
		case <-stop:
			b.logger.Info("lrp-groups-exit")
			return
		}
	}
}

func processInstance(instance *bbsmodels.ActualLRP) (*api.Backend, *api.Backend) {
	var (
		appHostPort      uint32
		appContainerPort uint32
		externalBackend  *api.Backend
		internalBackend  *api.Backend
	)

	if instance.State != bbsmodels.ActualLRPStateRunning {
		return externalBackend, internalBackend
	}

	for _, port := range instance.ActualLRPNetInfo.Ports {
		if port.ContainerPort != CF_APP_SSH_PORT {
			appHostPort = port.HostPort
			appContainerPort = port.ContainerPort
		}
	}

	if appHostPort != 0 {
		internalBackend = &api.Backend{Address: instance.ActualLRPNetInfo.Address, Port: appHostPort}
	}

	if appContainerPort != 0 {
		externalBackend = &api.Backend{Address: instance.ActualLRPNetInfo.InstanceAddress, Port: appContainerPort}
	}

	return internalBackend, externalBackend
}
