package snapshot_test

import (
	"time"

	"code.cloudfoundry.org/copilot/certs"
	"code.cloudfoundry.org/copilot/certs/fakes"
	"code.cloudfoundry.org/copilot/models"
	"code.cloudfoundry.org/copilot/snapshot"
	"code.cloudfoundry.org/lager/lagertest"
	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	networking "istio.io/api/networking/v1alpha3"
)

var _ = Describe("Config", func() {
	var (
		config      *snapshot.Config
		fakeLocator *fakes.Locator
	)

	BeforeEach(func() {
		fakeLocator = &fakes.Locator{}
		config = snapshot.NewConfig(fakeLocator, lagertest.NewTestLogger("config"))
	})

	Describe("CreateSidecarResources", func() {
		It("creates a sidecar resource", func() {
			sidecars := config.CreateSidecarResources()

			Expect(sidecars).To(HaveLen(1))
			Expect(sidecars[0].Metadata.Name).To(Equal("cloudfoundry-sidecar"))

			sc := networking.Sidecar{}
			err := types.UnmarshalAny(sidecars[0].Body, &sc)
			Expect(err).NotTo(HaveOccurred())

			Expect(sc).To(Equal(networking.Sidecar{
				Egress: []*networking.IstioEgressListener{
					&networking.IstioEgressListener{
						Hosts: []string{"internal/*"},
					},
				},
			}))
		})
	})

	Describe("CreateGatewayResources", func() {
		It("creates gateway resources", func() {
			gateways := config.CreateGatewayResources()
			var ga networking.Gateway

			Expect(gateways).To(HaveLen(1))
			Expect(gateways[0].Metadata.Name).To(Equal("cloudfoundry-ingress"))
			Expect(gateways[0].Metadata.Version).To(Equal("1"))

			err := types.UnmarshalAny(gateways[0].Body, &ga)
			Expect(err).NotTo(HaveOccurred())

			Expect(ga).To(Equal(networking.Gateway{
				Servers: []*networking.Server{
					{
						Port: &networking.Port{
							Number:   80,
							Protocol: "http",
							Name:     "http",
						},
						Hosts: []string{"*"},
					},
				},
			}))
		})

		Context("When locator returns cert pair paths", func() {
			It("creates gateway resources with http and https servers", func() {
				certPairs := []certs.PemInfo{
					{
						Hosts:    []string{"example.com"},
						CertPath: "/some/path/not/important.crt",
						KeyPath:  "/some/path/not/important.key",
					},
				}
				fakeLocator.LocateReturns(certPairs, nil)

				gateways := config.CreateGatewayResources()
				var ga networking.Gateway

				Expect(gateways).To(HaveLen(1))
				Expect(gateways[0].Metadata.Name).To(Equal("cloudfoundry-ingress"))

				err := types.UnmarshalAny(gateways[0].Body, &ga)
				Expect(err).NotTo(HaveOccurred())

				Expect(ga).To(Equal(networking.Gateway{
					Servers: []*networking.Server{
						{
							Port: &networking.Port{
								Number:   80,
								Protocol: "http",
								Name:     "http",
							},
							Hosts: []string{"*"},
						},
						{
							Port: &networking.Port{
								Number:   443,
								Protocol: "https",
								Name:     "example.com",
							},
							Tls: &networking.Server_TLSOptions{
								Mode:              networking.Server_TLSOptions_SIMPLE,
								ServerCertificate: "/some/path/not/important.crt",
								PrivateKey:        "/some/path/not/important.key",
							},
							Hosts: []string{"example.com"},
						},
					},
				}))
			})

			It("does not create gateways with the same server port name", func() {
				certPairs := []certs.PemInfo{
					{
						Hosts:    []string{"example.com"},
						CertPath: "/some/path/not/important.crt",
						KeyPath:  "/some/path/not/important.key",
					},
					{
						Hosts:    []string{"example2.com", "example3.com"},
						CertPath: "/some/path/not/important.crt",
						KeyPath:  "/some/path/not/important.key",
					},
				}
				fakeLocator.LocateReturns(certPairs, nil)

				gateways := config.CreateGatewayResources()
				var ga networking.Gateway

				Expect(gateways).To(HaveLen(1))
				Expect(gateways[0].Metadata.Name).To(Equal("cloudfoundry-ingress"))

				err := types.UnmarshalAny(gateways[0].Body, &ga)
				Expect(err).NotTo(HaveOccurred())

				Expect(ga).To(Equal(networking.Gateway{
					Servers: []*networking.Server{
						{
							Port: &networking.Port{
								Number:   80,
								Protocol: "http",
								Name:     "http",
							},
							Hosts: []string{"*"},
						},
						{
							Port: &networking.Port{
								Number:   443,
								Protocol: "https",
								Name:     "example.com",
							},
							Tls: &networking.Server_TLSOptions{
								Mode:              networking.Server_TLSOptions_SIMPLE,
								ServerCertificate: "/some/path/not/important.crt",
								PrivateKey:        "/some/path/not/important.key",
							},
							Hosts: []string{"example.com"},
						},
						{
							Port: &networking.Port{
								Number:   443,
								Protocol: "https",
								Name:     "example2.com",
							},
							Tls: &networking.Server_TLSOptions{
								Mode:              networking.Server_TLSOptions_SIMPLE,
								ServerCertificate: "/some/path/not/important.crt",
								PrivateKey:        "/some/path/not/important.key",
							},
							Hosts: []string{"example2.com", "example3.com"},
						},
					},
				}))
			})
		})
	})

	Describe("CreateVirtualServiceResources", func() {
		It("creates virtualService resources", func() {
			virtualServices := config.CreateVirtualServiceResources(routesWithBackends(), "1")
			var vs networking.VirtualService

			Expect(virtualServices).To(HaveLen(1))
			Expect(virtualServices[0].Metadata.Name).To(Equal("copilot-service-for-foo.example.com"))

			err := types.UnmarshalAny(virtualServices[0].Body, &vs)
			Expect(err).NotTo(HaveOccurred())

			Expect(vs.Hosts).To(Equal([]string{"foo.example.com"}))
			Expect(vs.Gateways).To(Equal([]string{"cloudfoundry-ingress"}))
			Expect(vs.Http).To(ConsistOf([]*networking.HTTPRoute{
				{
					Match: []*networking.HTTPMatchRequest{
						{
							Uri: &networking.StringMatch{
								MatchType: &networking.StringMatch_Prefix{
									Prefix: "/something",
								},
							},
						},
					},
					Route: []*networking.HTTPRouteDestination{
						{
							Destination: &networking.Destination{
								Host: "foo.example.com",
								Port: &networking.PortSelector{
									Port: &networking.PortSelector_Number{
										Number: 7070,
									},
								},
								Subset: "x-capi-guid",
							},
							Weight: 50,
						},
						{
							Destination: &networking.Destination{
								Host: "foo.example.com",
								Port: &networking.PortSelector{
									Port: &networking.PortSelector_Number{
										Number: 9090,
									},
								},
								Subset: "y-capi-guid",
							},
							Weight: 50,
						},
					},
				},
				{
					Route: []*networking.HTTPRouteDestination{
						{
							Destination: &networking.Destination{
								Host: "foo.example.com",
								Port: &networking.PortSelector{
									Port: &networking.PortSelector_Number{
										Number: 8080,
									},
								},
								Subset: "a-capi-guid",
							},
							Weight: 100,
						},
					},
				},
			}))
		})

		Context("internal routes", func() {
			It("creates virtualService resources with retries", func() {
				virtualServices := config.CreateVirtualServiceResources(internalRoutesWithBackends(), "1")
				var vs networking.VirtualService

				Expect(virtualServices).To(HaveLen(1))
				Expect(virtualServices[0].Metadata.Name).To(Equal("internal/copilot-service-for-foo.bar.internal"))

				err := types.UnmarshalAny(virtualServices[0].Body, &vs)
				Expect(err).NotTo(HaveOccurred())

				Expect(vs.Hosts).To(Equal([]string{"foo.bar.internal"}))
				Expect(vs.Gateways).To(HaveLen(0))
				Expect(vs.Http).To(ConsistOf([]*networking.HTTPRoute{
					{
						Route: []*networking.HTTPRouteDestination{
							{
								Destination: &networking.Destination{
									Host: "foo.bar.internal",
									Port: &networking.PortSelector{
										Port: &networking.PortSelector_Number{
											Number: 8081,
										},
									},
									Subset: "x-capi-guid",
								},
								Weight: 100,
							},
						},
						Retries: &networking.HTTPRetry{
							Attempts: 3,
							RetryOn:  "5xx",
						},
						Timeout: types.DurationProto(15 * time.Second),
					},
				}))
			})
		})
	})

	Describe("CreateDestinationRuleResources", func() {
		It("creates destinationRule resources", func() {
			destinationRules := config.CreateDestinationRuleResources(routesWithBackends(), "1")
			var dr networking.DestinationRule

			Expect(destinationRules).To(HaveLen(1))
			Expect(destinationRules[0].Metadata.Name).To(Equal("copilot-rule-for-foo.example.com"))

			err := types.UnmarshalAny(destinationRules[0].Body, &dr)
			Expect(err).NotTo(HaveOccurred())

			Expect(dr.Host).To(Equal("foo.example.com"))
			Expect(dr.Subsets).To(ConsistOf([]*networking.Subset{
				{
					Name:   "a-capi-guid",
					Labels: map[string]string{"cfapp": "a-capi-guid"},
				},
				{
					Name:   "y-capi-guid",
					Labels: map[string]string{"cfapp": "y-capi-guid"},
				},
				{
					Name:   "x-capi-guid",
					Labels: map[string]string{"cfapp": "x-capi-guid"},
				},
			}))
		})

		Context("internal routes", func() {
			It("creates destination rules", func() {
				destinationRules := config.CreateDestinationRuleResources(internalRoutesWithBackends(), "1")
				var dr networking.DestinationRule

				Expect(destinationRules).To(HaveLen(1))
				Expect(destinationRules[0].Metadata.Name).To(Equal("internal/copilot-rule-for-foo.bar.internal"))

				err := types.UnmarshalAny(destinationRules[0].Body, &dr)
				Expect(err).NotTo(HaveOccurred())

				Expect(dr.Host).To(Equal("foo.bar.internal"))
				Expect(dr.Subsets).To(ConsistOf([]*networking.Subset{
					{
						Name:   "x-capi-guid",
						Labels: map[string]string{"cfapp": "x-capi-guid"},
					},
				}))
			})
		})
	})

	Describe("CreateServiceEntryResources", func() {
		It("creates serviceEntry resources", func() {
			serviceEntries := config.CreateServiceEntryResources(routesWithBackends(), "1")
			var se networking.ServiceEntry

			Expect(serviceEntries).To(HaveLen(1))
			Expect(serviceEntries[0].Metadata.Name).To(Equal("copilot-service-entry-for-foo.example.com"))

			err := types.UnmarshalAny(serviceEntries[0].Body, &se)
			Expect(err).NotTo(HaveOccurred())

			Expect(se.Hosts).To(Equal([]string{"foo.example.com"}))
			Expect(se.Addresses).To(BeNil())
			Expect(se.Ports).To(Equal([]*networking.Port{
				{
					Name:     "http",
					Number:   7070,
					Protocol: "http",
				},
			}))
			Expect(se.Location).To(Equal(networking.ServiceEntry_MESH_INTERNAL))
			Expect(se.Resolution).To(Equal(networking.ServiceEntry_STATIC))
			Expect(se.Endpoints).To(ConsistOf([]*networking.ServiceEntry_Endpoint{
				{
					Address: "10.10.10.1",
					Ports: map[string]uint32{
						"http": 65003,
					},
					Labels: map[string]string{"cfapp": "a-capi-guid"},
				},
				{
					Address: "10.0.0.0",
					Ports: map[string]uint32{
						"http": 65007,
					},
					Labels: map[string]string{"cfapp": "y-capi-guid"},
				},
				{
					Address: "10.0.0.1",
					Ports: map[string]uint32{
						"http": 65005,
					},
					Labels: map[string]string{"cfapp": "x-capi-guid"},
				},
			}))
		})

		Context("internal routes", func() {
			It("creates service entries resources", func() {
				serviceEntries := config.CreateServiceEntryResources(internalRoutesWithBackends(), "1")
				var se networking.ServiceEntry

				Expect(serviceEntries).To(HaveLen(1))
				Expect(serviceEntries[0].Metadata.Name).To(Equal("internal/copilot-service-entry-for-foo.bar.internal"))

				err := types.UnmarshalAny(serviceEntries[0].Body, &se)
				Expect(err).NotTo(HaveOccurred())

				Expect(se.Hosts).To(Equal([]string{"foo.bar.internal"}))
				Expect(se.Addresses).To(Equal([]string{"127.127.0.1"}))
				Expect(se.Ports).To(Equal([]*networking.Port{
					{
						Name:     "http",
						Number:   8081,
						Protocol: "http",
					}}))
				Expect(se.Location).To(Equal(networking.ServiceEntry_MESH_INTERNAL))
				Expect(se.Resolution).To(Equal(networking.ServiceEntry_STATIC))
				Expect(se.Endpoints).To(ConsistOf([]*networking.ServiceEntry_Endpoint{
					{
						Address: "10.255.0.1",
						Ports: map[string]uint32{
							"http": 8080,
						},
						Labels: map[string]string{"cfapp": "x-capi-guid"},
					},
				}))

			})
		})
	})
})

func routesWithBackends() []*models.RouteWithBackends {
	return []*models.RouteWithBackends{
		{
			Hostname: "foo.example.com",
			Path:     "/something",
			Backends: models.BackendSet{
				Backends: []*models.Backend{
					{
						Address:       "10.0.0.1",
						Port:          uint32(65005),
						ContainerPort: uint32(7070),
					},
				},
			},
			CapiProcessGUID: "x-capi-guid",
			RouteWeight:     int32(50),
		},
		{
			Hostname: "foo.example.com",
			Path:     "/something",
			Backends: models.BackendSet{
				Backends: []*models.Backend{
					{
						Address:       "10.0.0.0",
						Port:          uint32(65007),
						ContainerPort: uint32(9090),
					},
				},
			},
			CapiProcessGUID: "y-capi-guid",
			RouteWeight:     int32(50),
		},
		{
			Hostname: "foo.example.com",
			Path:     "",
			Backends: models.BackendSet{
				Backends: []*models.Backend{
					{
						Address:       "10.10.10.1",
						Port:          uint32(65003),
						ContainerPort: uint32(8080),
					},
				},
			},
			CapiProcessGUID: "a-capi-guid",
			RouteWeight:     int32(100),
		},
	}
}

func internalRoutesWithBackends() []*models.RouteWithBackends {
	return []*models.RouteWithBackends{
		{
			Hostname: "foo.bar.internal",
			VIP:      "127.127.0.1",
			Internal: true,
			Path:     "/something",
			Backends: models.BackendSet{
				Backends: []*models.Backend{
					{
						Address:       "10.255.0.1",
						Port:          uint32(8080),
						ContainerPort: uint32(8081),
					},
				},
			},
			CapiProcessGUID: "x-capi-guid",
			RouteWeight:     int32(100),
		},
	}
}
