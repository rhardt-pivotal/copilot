package handlers_test

import (
	"context"

	"code.cloudfoundry.org/copilot/api"
	"code.cloudfoundry.org/copilot/handlers"
	"code.cloudfoundry.org/copilot/handlers/fakes"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"

	"code.cloudfoundry.org/copilot/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Capi Handlers", func() {
	var (
		handler                              *handlers.CAPI
		logger                               lager.Logger
		fakeRoutesRepo                       *fakes.RoutesRepo
		fakeRouteMappingsRepo                *fakes.RouteMappingsRepo
		fakeCAPIDiegoProcessAssociationsRepo *fakes.CAPIDiegoProcessAssociationsRepo
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")

		fakeRoutesRepo = &fakes.RoutesRepo{}
		fakeRouteMappingsRepo = &fakes.RouteMappingsRepo{}
		fakeCAPIDiegoProcessAssociationsRepo = &fakes.CAPIDiegoProcessAssociationsRepo{}
		handler = &handlers.CAPI{
			Logger:                           logger,
			RoutesRepo:                       fakeRoutesRepo,
			RouteMappingsRepo:                fakeRouteMappingsRepo,
			CAPIDiegoProcessAssociationsRepo: fakeCAPIDiegoProcessAssociationsRepo,
		}
	})

	Describe("Health", func() {
		It("always returns healthy", func() {
			ctx := context.Background()
			resp, err := handler.Health(ctx, new(api.HealthRequest))
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).To(Equal(&api.HealthResponse{Healthy: true}))
		})
	})

	Describe("UpsertRoute", func() {
		It("validates the inputs", func() {
			ctx := context.Background()
			_, err := handler.UpsertRoute(ctx, &api.UpsertRouteRequest{
				Route: &api.Route{
					Guid: "some-route-guid",
				}})
			Expect(err.Error()).To(ContainSubstring("required"))
			_, err = handler.UpsertRoute(ctx, &api.UpsertRouteRequest{
				Route: &api.Route{
					Host: "some-hostname",
				}})
			Expect(err.Error()).To(ContainSubstring("required"))
		})

		It("adds the route if it is new", func() {
			ctx := context.Background()
			_, err := handler.UpsertRoute(ctx, &api.UpsertRouteRequest{
				Route: &api.Route{
					Guid:     "route-guid-a",
					Host:     "route-a.example.com",
					Internal: true,
				}})
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeRoutesRepo.UpsertCallCount()).To(Equal(1))
			Expect(fakeRoutesRepo.UpsertArgsForCall(0)).To(Equal(&models.Route{
				GUID:     "route-guid-a",
				Host:     "route-a.example.com",
				Internal: true,
			}))
		})
	})

	Describe("DeleteRoute", func() {
		It("calls Delete on the RoutesRepo using the provided guid", func() {
			fakeRoutesRepo := &fakes.RoutesRepo{}
			ctx := context.Background()
			handler.RoutesRepo = fakeRoutesRepo
			_, err := handler.DeleteRoute(ctx, &api.DeleteRouteRequest{Guid: "route-guid-a"})
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeRoutesRepo.DeleteCallCount()).To(Equal(1))
			Expect(fakeRoutesRepo.DeleteArgsForCall(0)).To(Equal(models.RouteGUID("route-guid-a")))
		})

		It("validates the inputs", func() {
			ctx := context.Background()
			_, err := handler.DeleteRoute(ctx, &api.DeleteRouteRequest{})
			Expect(err.Error()).To(ContainSubstring("required"))
		})
	})

	Describe("MapRoute", func() {
		BeforeEach(func() {
			handler.RoutesRepo.Upsert(&models.Route{
				GUID: "route-guid-a",
				Host: "route-a.example.com",
			})
		})

		It("validates the inputs", func() {
			ctx := context.Background()
			_, err := handler.MapRoute(ctx, &api.MapRouteRequest{})
			Expect(err.Error()).To(ContainSubstring("required"))
			_, err = handler.MapRoute(ctx, &api.MapRouteRequest{RouteMapping: &api.RouteMapping{
				RouteGuid: "some-route-guid",
			}})
			Expect(err.Error()).To(ContainSubstring("required"))
			_, err = handler.MapRoute(ctx, &api.MapRouteRequest{RouteMapping: &api.RouteMapping{
				CapiProcessGuid: "some-process-guid",
			}})
			Expect(err.Error()).To(ContainSubstring("required"))
		})

		It("maps the route", func() {
			ctx := context.Background()
			_, err := handler.MapRoute(ctx, &api.MapRouteRequest{
				RouteMapping: &api.RouteMapping{
					RouteGuid:       "route-guid-a",
					CapiProcessGuid: "some-capi-process-guid",
					RouteWeight:     1,
				}})
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeRouteMappingsRepo.MapCallCount()).To(Equal(1))
			Expect(fakeRouteMappingsRepo.MapArgsForCall(0)).To(Equal(&models.RouteMapping{
				RouteGUID:       "route-guid-a",
				CAPIProcessGUID: "some-capi-process-guid",
				RouteWeight:     1,
			}))
		})
	})

	Describe("UnmapRoute", func() {
		It("validates the inputs", func() {
			ctx := context.Background()
			_, err := handler.MapRoute(ctx, &api.MapRouteRequest{})
			Expect(err.Error()).To(ContainSubstring("required"))
			_, err = handler.UnmapRoute(ctx, &api.UnmapRouteRequest{RouteMapping: &api.RouteMapping{RouteGuid: "some-route-guid"}})
			Expect(err.Error()).To(ContainSubstring("required"))
			_, err = handler.UnmapRoute(ctx, &api.UnmapRouteRequest{RouteMapping: &api.RouteMapping{CapiProcessGuid: "some-process-guid"}})
			Expect(err.Error()).To(ContainSubstring("required"))
		})

		It("unmaps the routes", func() {
			ctx := context.Background()
			_, err := handler.UnmapRoute(ctx, &api.UnmapRouteRequest{
				RouteMapping: &api.RouteMapping{RouteGuid: "to-be-deleted-route-guid",
					CapiProcessGuid: "some-capi-process-guid",
					RouteWeight:     1,
				}})
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeRouteMappingsRepo.UnmapCallCount()).To(Equal(1))
			Expect(fakeRouteMappingsRepo.UnmapArgsForCall(0)).To(Equal(&models.RouteMapping{
				RouteGUID:       "to-be-deleted-route-guid",
				CAPIProcessGUID: "some-capi-process-guid",
				RouteWeight:     1,
			}))
		})
	})

	Describe("UpsertCapiDiegoProcessAssociation", func() {
		It("validates the inputs", func() {
			ctx := context.Background()
			_, err := handler.UpsertCapiDiegoProcessAssociation(ctx, &api.UpsertCapiDiegoProcessAssociationRequest{})
			Expect(err.Error()).To(ContainSubstring("required"))
			_, err = handler.UpsertCapiDiegoProcessAssociation(ctx, &api.UpsertCapiDiegoProcessAssociationRequest{
				CapiDiegoProcessAssociation: &api.CapiDiegoProcessAssociation{
					CapiProcessGuid: "some-capi-process-guid",
				},
			})
			Expect(err.Error()).To(ContainSubstring("required"))
			_, err = handler.UpsertCapiDiegoProcessAssociation(ctx, &api.UpsertCapiDiegoProcessAssociationRequest{
				CapiDiegoProcessAssociation: &api.CapiDiegoProcessAssociation{
					DiegoProcessGuids: []string{
						"some-diego-process-guid",
					},
				},
			})
			Expect(err.Error()).To(ContainSubstring("required"))
		})

		It("associates the capi and diego process guids", func() {
			ctx := context.Background()
			_, err := handler.UpsertCapiDiegoProcessAssociation(ctx, &api.UpsertCapiDiegoProcessAssociationRequest{
				CapiDiegoProcessAssociation: &api.CapiDiegoProcessAssociation{
					CapiProcessGuid: "some-capi-process-guid",
					DiegoProcessGuids: []string{
						"some-diego-process-guid-1",
						"some-diego-process-guid-2",
					},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeCAPIDiegoProcessAssociationsRepo.UpsertCallCount()).To(Equal(1))
			Expect(fakeCAPIDiegoProcessAssociationsRepo.UpsertArgsForCall(0)).To(Equal(&models.CAPIDiegoProcessAssociation{
				CAPIProcessGUID: "some-capi-process-guid",
				DiegoProcessGUIDs: models.DiegoProcessGUIDs{
					"some-diego-process-guid-1",
					"some-diego-process-guid-2",
				},
			}))
		})
	})

	Describe("DeleteCapiDiegoProcessAssociation", func() {
		It("validates the inputs", func() {
			ctx := context.Background()
			_, err := handler.DeleteCapiDiegoProcessAssociation(ctx, &api.DeleteCapiDiegoProcessAssociationRequest{})
			Expect(err.Error()).To(ContainSubstring("required"))
		})

		It("deletes the association", func() {
			ctx := context.Background()
			_, err := handler.DeleteCapiDiegoProcessAssociation(ctx, &api.DeleteCapiDiegoProcessAssociationRequest{
				CapiProcessGuid: "some-capi-process-guid",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeCAPIDiegoProcessAssociationsRepo.DeleteCallCount()).To(Equal(1))
			cpg := models.CAPIProcessGUID("some-capi-process-guid")
			Expect(fakeCAPIDiegoProcessAssociationsRepo.DeleteArgsForCall(0)).To(Equal(&cpg))
		})
	})

	Describe("BulkSync", func() {
		Context("when inputs are empty", func() {
			It("does not throw an error", func() {
				ctx := context.Background()
				_, err := handler.BulkSync(ctx, &api.BulkSyncRequest{})
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeRouteMappingsRepo.SyncCallCount()).To(Equal(1))
				Expect(fakeRoutesRepo.SyncCallCount()).To(Equal(1))
				Expect(fakeCAPIDiegoProcessAssociationsRepo.SyncCallCount()).To(Equal(1))
			})
		})

		It("syncs", func() {
			ctx := context.Background()
			_, err := handler.BulkSync(ctx, &api.BulkSyncRequest{
				RouteMappings: []*api.RouteMapping{{
					RouteGuid:       "route-guid-a",
					CapiProcessGuid: "some-capi-process-guid",
					RouteWeight:     1,
				}},
				Routes: []*api.Route{
					{
						Guid: "route-guid-a",
						Host: "example.host.com",
						Path: "/nothing/matters",
					},
				},
				CapiDiegoProcessAssociations: []*api.CapiDiegoProcessAssociation{{
					CapiProcessGuid: "some-capi-process-guid",
					DiegoProcessGuids: []string{
						"some-diego-process-guid-1",
						"some-diego-process-guid-2",
					},
				}},
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeRouteMappingsRepo.SyncCallCount()).To(Equal(1))
			Expect(fakeRouteMappingsRepo.SyncArgsForCall(0)).To(Equal([]*models.RouteMapping{{
				RouteGUID:       "route-guid-a",
				CAPIProcessGUID: "some-capi-process-guid",
				RouteWeight:     1,
			}}))

			Expect(fakeRoutesRepo.SyncCallCount()).To(Equal(1))
			Expect(fakeRoutesRepo.SyncArgsForCall(0)).To(Equal([]*models.Route{
				{
					GUID: "route-guid-a",
					Host: "example.host.com",
					Path: "/nothing/matters",
				},
			}))

			Expect(fakeCAPIDiegoProcessAssociationsRepo.SyncCallCount()).To(Equal(1))
			Expect(fakeCAPIDiegoProcessAssociationsRepo.SyncArgsForCall(0)).To(Equal([]*models.CAPIDiegoProcessAssociation{{
				CAPIProcessGUID: "some-capi-process-guid",
				DiegoProcessGUIDs: []models.DiegoProcessGUID{
					"some-diego-process-guid-1",
					"some-diego-process-guid-2",
				},
			}}))
		})
	})
})
