// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/copilot/models"
)

type CapiDiego struct {
	GetStub        func(capiProcessGUID *models.CAPIProcessGUID) *models.CAPIDiegoProcessAssociation
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		capiProcessGUID *models.CAPIProcessGUID
	}
	getReturns struct {
		result1 *models.CAPIDiegoProcessAssociation
	}
	getReturnsOnCall map[int]struct {
		result1 *models.CAPIDiegoProcessAssociation
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CapiDiego) Get(capiProcessGUID *models.CAPIProcessGUID) *models.CAPIDiegoProcessAssociation {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		capiProcessGUID *models.CAPIProcessGUID
	}{capiProcessGUID})
	fake.recordInvocation("Get", []interface{}{capiProcessGUID})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(capiProcessGUID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *CapiDiego) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *CapiDiego) GetArgsForCall(i int) *models.CAPIProcessGUID {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].capiProcessGUID
}

func (fake *CapiDiego) GetReturns(result1 *models.CAPIDiegoProcessAssociation) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *models.CAPIDiegoProcessAssociation
	}{result1}
}

func (fake *CapiDiego) GetReturnsOnCall(i int, result1 *models.CAPIDiegoProcessAssociation) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *models.CAPIDiegoProcessAssociation
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *models.CAPIDiegoProcessAssociation
	}{result1}
}

func (fake *CapiDiego) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CapiDiego) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}