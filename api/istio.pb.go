// Code generated by protoc-gen-go. DO NOT EDIT.
// source: istio.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InternalRoutesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalRoutesRequest) Reset()         { *m = InternalRoutesRequest{} }
func (m *InternalRoutesRequest) String() string { return proto.CompactTextString(m) }
func (*InternalRoutesRequest) ProtoMessage()    {}
func (*InternalRoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{0}
}
func (m *InternalRoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalRoutesRequest.Unmarshal(m, b)
}
func (m *InternalRoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalRoutesRequest.Marshal(b, m, deterministic)
}
func (dst *InternalRoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalRoutesRequest.Merge(dst, src)
}
func (m *InternalRoutesRequest) XXX_Size() int {
	return xxx_messageInfo_InternalRoutesRequest.Size(m)
}
func (m *InternalRoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalRoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalRoutesRequest proto.InternalMessageInfo

type InternalRoutesResponse struct {
	InternalRoutes       []*InternalRouteWithBackends `protobuf:"bytes,1,rep,name=internal_routes,json=internalRoutes" json:"internal_routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *InternalRoutesResponse) Reset()         { *m = InternalRoutesResponse{} }
func (m *InternalRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*InternalRoutesResponse) ProtoMessage()    {}
func (*InternalRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{1}
}
func (m *InternalRoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalRoutesResponse.Unmarshal(m, b)
}
func (m *InternalRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalRoutesResponse.Marshal(b, m, deterministic)
}
func (dst *InternalRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalRoutesResponse.Merge(dst, src)
}
func (m *InternalRoutesResponse) XXX_Size() int {
	return xxx_messageInfo_InternalRoutesResponse.Size(m)
}
func (m *InternalRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InternalRoutesResponse proto.InternalMessageInfo

func (m *InternalRoutesResponse) GetInternalRoutes() []*InternalRouteWithBackends {
	if m != nil {
		return m.InternalRoutes
	}
	return nil
}

type InternalRouteWithBackends struct {
	Hostname             string      `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Vip                  string      `protobuf:"bytes,2,opt,name=vip" json:"vip,omitempty"`
	Backends             *BackendSet `protobuf:"bytes,3,opt,name=backends" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *InternalRouteWithBackends) Reset()         { *m = InternalRouteWithBackends{} }
func (m *InternalRouteWithBackends) String() string { return proto.CompactTextString(m) }
func (*InternalRouteWithBackends) ProtoMessage()    {}
func (*InternalRouteWithBackends) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{2}
}
func (m *InternalRouteWithBackends) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalRouteWithBackends.Unmarshal(m, b)
}
func (m *InternalRouteWithBackends) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalRouteWithBackends.Marshal(b, m, deterministic)
}
func (dst *InternalRouteWithBackends) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalRouteWithBackends.Merge(dst, src)
}
func (m *InternalRouteWithBackends) XXX_Size() int {
	return xxx_messageInfo_InternalRouteWithBackends.Size(m)
}
func (m *InternalRouteWithBackends) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalRouteWithBackends.DiscardUnknown(m)
}

var xxx_messageInfo_InternalRouteWithBackends proto.InternalMessageInfo

func (m *InternalRouteWithBackends) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *InternalRouteWithBackends) GetVip() string {
	if m != nil {
		return m.Vip
	}
	return ""
}

func (m *InternalRouteWithBackends) GetBackends() *BackendSet {
	if m != nil {
		return m.Backends
	}
	return nil
}

type RoutesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutesRequest) Reset()         { *m = RoutesRequest{} }
func (m *RoutesRequest) String() string { return proto.CompactTextString(m) }
func (*RoutesRequest) ProtoMessage()    {}
func (*RoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{3}
}
func (m *RoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesRequest.Unmarshal(m, b)
}
func (m *RoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesRequest.Marshal(b, m, deterministic)
}
func (dst *RoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesRequest.Merge(dst, src)
}
func (m *RoutesRequest) XXX_Size() int {
	return xxx_messageInfo_RoutesRequest.Size(m)
}
func (m *RoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesRequest proto.InternalMessageInfo

type RoutesResponse struct {
	Routes               []*RouteWithBackends `protobuf:"bytes,2,rep,name=routes" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RoutesResponse) Reset()         { *m = RoutesResponse{} }
func (m *RoutesResponse) String() string { return proto.CompactTextString(m) }
func (*RoutesResponse) ProtoMessage()    {}
func (*RoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{4}
}
func (m *RoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesResponse.Unmarshal(m, b)
}
func (m *RoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesResponse.Marshal(b, m, deterministic)
}
func (dst *RoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesResponse.Merge(dst, src)
}
func (m *RoutesResponse) XXX_Size() int {
	return xxx_messageInfo_RoutesResponse.Size(m)
}
func (m *RoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesResponse proto.InternalMessageInfo

func (m *RoutesResponse) GetRoutes() []*RouteWithBackends {
	if m != nil {
		return m.Routes
	}
	return nil
}

type RouteWithBackends struct {
	Hostname             string      `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Path                 string      `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Backends             *BackendSet `protobuf:"bytes,3,opt,name=backends" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RouteWithBackends) Reset()         { *m = RouteWithBackends{} }
func (m *RouteWithBackends) String() string { return proto.CompactTextString(m) }
func (*RouteWithBackends) ProtoMessage()    {}
func (*RouteWithBackends) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{5}
}
func (m *RouteWithBackends) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteWithBackends.Unmarshal(m, b)
}
func (m *RouteWithBackends) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteWithBackends.Marshal(b, m, deterministic)
}
func (dst *RouteWithBackends) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteWithBackends.Merge(dst, src)
}
func (m *RouteWithBackends) XXX_Size() int {
	return xxx_messageInfo_RouteWithBackends.Size(m)
}
func (m *RouteWithBackends) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteWithBackends.DiscardUnknown(m)
}

var xxx_messageInfo_RouteWithBackends proto.InternalMessageInfo

func (m *RouteWithBackends) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *RouteWithBackends) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RouteWithBackends) GetBackends() *BackendSet {
	if m != nil {
		return m.Backends
	}
	return nil
}

type BackendSet struct {
	Backends             []*Backend `protobuf:"bytes,1,rep,name=backends" json:"backends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BackendSet) Reset()         { *m = BackendSet{} }
func (m *BackendSet) String() string { return proto.CompactTextString(m) }
func (*BackendSet) ProtoMessage()    {}
func (*BackendSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{6}
}
func (m *BackendSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendSet.Unmarshal(m, b)
}
func (m *BackendSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendSet.Marshal(b, m, deterministic)
}
func (dst *BackendSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendSet.Merge(dst, src)
}
func (m *BackendSet) XXX_Size() int {
	return xxx_messageInfo_BackendSet.Size(m)
}
func (m *BackendSet) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendSet.DiscardUnknown(m)
}

var xxx_messageInfo_BackendSet proto.InternalMessageInfo

func (m *BackendSet) GetBackends() []*Backend {
	if m != nil {
		return m.Backends
	}
	return nil
}

type Backend struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Backend) Reset()         { *m = Backend{} }
func (m *Backend) String() string { return proto.CompactTextString(m) }
func (*Backend) ProtoMessage()    {}
func (*Backend) Descriptor() ([]byte, []int) {
	return fileDescriptor_istio_9b3f5372a9b7e8a6, []int{7}
}
func (m *Backend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backend.Unmarshal(m, b)
}
func (m *Backend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backend.Marshal(b, m, deterministic)
}
func (dst *Backend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backend.Merge(dst, src)
}
func (m *Backend) XXX_Size() int {
	return xxx_messageInfo_Backend.Size(m)
}
func (m *Backend) XXX_DiscardUnknown() {
	xxx_messageInfo_Backend.DiscardUnknown(m)
}

var xxx_messageInfo_Backend proto.InternalMessageInfo

func (m *Backend) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Backend) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*InternalRoutesRequest)(nil), "api.InternalRoutesRequest")
	proto.RegisterType((*InternalRoutesResponse)(nil), "api.InternalRoutesResponse")
	proto.RegisterType((*InternalRouteWithBackends)(nil), "api.InternalRouteWithBackends")
	proto.RegisterType((*RoutesRequest)(nil), "api.RoutesRequest")
	proto.RegisterType((*RoutesResponse)(nil), "api.RoutesResponse")
	proto.RegisterType((*RouteWithBackends)(nil), "api.RouteWithBackends")
	proto.RegisterType((*BackendSet)(nil), "api.BackendSet")
	proto.RegisterType((*Backend)(nil), "api.Backend")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IstioCopilotClient is the client API for IstioCopilot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IstioCopilotClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	Routes(ctx context.Context, in *RoutesRequest, opts ...grpc.CallOption) (*RoutesResponse, error)
	InternalRoutes(ctx context.Context, in *InternalRoutesRequest, opts ...grpc.CallOption) (*InternalRoutesResponse, error)
}

type istioCopilotClient struct {
	cc *grpc.ClientConn
}

func NewIstioCopilotClient(cc *grpc.ClientConn) IstioCopilotClient {
	return &istioCopilotClient{cc}
}

func (c *istioCopilotClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/api.IstioCopilot/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioCopilotClient) Routes(ctx context.Context, in *RoutesRequest, opts ...grpc.CallOption) (*RoutesResponse, error) {
	out := new(RoutesResponse)
	err := c.cc.Invoke(ctx, "/api.IstioCopilot/Routes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *istioCopilotClient) InternalRoutes(ctx context.Context, in *InternalRoutesRequest, opts ...grpc.CallOption) (*InternalRoutesResponse, error) {
	out := new(InternalRoutesResponse)
	err := c.cc.Invoke(ctx, "/api.IstioCopilot/InternalRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IstioCopilot service

type IstioCopilotServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	Routes(context.Context, *RoutesRequest) (*RoutesResponse, error)
	InternalRoutes(context.Context, *InternalRoutesRequest) (*InternalRoutesResponse, error)
}

func RegisterIstioCopilotServer(s *grpc.Server, srv IstioCopilotServer) {
	s.RegisterService(&_IstioCopilot_serviceDesc, srv)
}

func _IstioCopilot_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioCopilotServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IstioCopilot/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioCopilotServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioCopilot_Routes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioCopilotServer).Routes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IstioCopilot/Routes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioCopilotServer).Routes(ctx, req.(*RoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IstioCopilot_InternalRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioCopilotServer).InternalRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.IstioCopilot/InternalRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioCopilotServer).InternalRoutes(ctx, req.(*InternalRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IstioCopilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.IstioCopilot",
	HandlerType: (*IstioCopilotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _IstioCopilot_Health_Handler,
		},
		{
			MethodName: "Routes",
			Handler:    _IstioCopilot_Routes_Handler,
		},
		{
			MethodName: "InternalRoutes",
			Handler:    _IstioCopilot_InternalRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "istio.proto",
}

func init() { proto.RegisterFile("istio.proto", fileDescriptor_istio_9b3f5372a9b7e8a6) }

var fileDescriptor_istio_9b3f5372a9b7e8a6 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4f, 0x02, 0x31,
	0x10, 0x65, 0x59, 0x02, 0x38, 0x7c, 0x3a, 0x46, 0x5c, 0xd7, 0xc4, 0x90, 0x3d, 0x6d, 0x62, 0xb2,
	0x07, 0x48, 0xf4, 0xae, 0x07, 0x45, 0x2f, 0x66, 0x3d, 0x78, 0x34, 0x05, 0x9a, 0x6c, 0x23, 0xb4,
	0x75, 0x5b, 0xf8, 0x87, 0xfe, 0x2f, 0x43, 0xb7, 0xe0, 0x16, 0xf4, 0xc0, 0x6d, 0xbe, 0xde, 0xeb,
	0xbc, 0x99, 0x29, 0xb4, 0x98, 0xd2, 0x4c, 0x24, 0x32, 0x17, 0x5a, 0xa0, 0x4f, 0x24, 0x0b, 0xdb,
	0x33, 0xb1, 0x5c, 0x0a, 0x5e, 0x84, 0xa2, 0x0b, 0x38, 0x9f, 0x70, 0x4d, 0x73, 0x4e, 0x16, 0xa9,
	0x58, 0x69, 0xaa, 0x52, 0xfa, 0xb5, 0xa2, 0x4a, 0x47, 0x04, 0x06, 0xfb, 0x09, 0x25, 0x05, 0x57,
	0x14, 0x1f, 0xa1, 0xc7, 0x6c, 0xe6, 0x23, 0x37, 0xa9, 0xc0, 0x1b, 0xfa, 0x71, 0x6b, 0x74, 0x9d,
	0x10, 0xc9, 0x12, 0x07, 0xf5, 0xce, 0x74, 0x76, 0x4f, 0x66, 0x9f, 0x94, 0xcf, 0x55, 0xda, 0x65,
	0x0e, 0x61, 0xb4, 0x86, 0xcb, 0x7f, 0x8b, 0x31, 0x84, 0x66, 0x26, 0x94, 0xe6, 0x64, 0x49, 0x03,
	0x6f, 0xe8, 0xc5, 0x27, 0xe9, 0xce, 0xc7, 0x3e, 0xf8, 0x6b, 0x26, 0x83, 0xaa, 0x09, 0x6f, 0x4c,
	0xbc, 0x81, 0xe6, 0xd4, 0x22, 0x03, 0x7f, 0xe8, 0xc5, 0xad, 0x51, 0xcf, 0x34, 0x63, 0xe9, 0xde,
	0xa8, 0x4e, 0x77, 0x05, 0x51, 0x0f, 0x3a, 0xae, 0xd6, 0x57, 0xe8, 0xee, 0x69, 0x4c, 0xa0, 0x6e,
	0xa5, 0x55, 0x8d, 0xb4, 0x81, 0x61, 0x3b, 0x94, 0x64, 0xab, 0x9e, 0x6b, 0x4d, 0xaf, 0x5f, 0x2d,
	0x3d, 0x21, 0xe1, 0xf4, 0x38, 0x49, 0x08, 0x35, 0x49, 0x74, 0x66, 0x35, 0x19, 0xfb, 0x38, 0x51,
	0xb7, 0x00, 0xbf, 0x71, 0x8c, 0x4b, 0xd0, 0x62, 0x39, 0xed, 0x32, 0xb4, 0x84, 0xbb, 0x83, 0x86,
	0x0d, 0x62, 0x00, 0x0d, 0x32, 0x9f, 0xe7, 0x54, 0x29, 0xdb, 0xde, 0xd6, 0x35, 0xdd, 0x89, 0x5c,
	0x9b, 0xee, 0x3a, 0xa9, 0xb1, 0x47, 0xdf, 0x1e, 0xb4, 0x27, 0x9b, 0xe3, 0x7a, 0x10, 0x92, 0x2d,
	0x84, 0xc6, 0x31, 0xd4, 0x9f, 0x28, 0x59, 0xe8, 0x0c, 0xd1, 0xbc, 0x55, 0x38, 0x76, 0xc6, 0xe1,
	0x99, 0x13, 0x2b, 0xc6, 0x1c, 0x55, 0x36, 0xa0, 0x62, 0xf4, 0x16, 0xe4, 0x2c, 0xc6, 0x82, 0xdc,
	0xdd, 0x44, 0x15, 0x7c, 0x81, 0xae, 0x7b, 0x9b, 0x18, 0x1e, 0x9e, 0xde, 0x8e, 0xe4, 0xea, 0xcf,
	0xdc, 0x96, 0x6c, 0x5a, 0x37, 0x1f, 0x61, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xfb, 0xf9,
	0x10, 0x2a, 0x03, 0x00, 0x00,
}
