// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	PluginInfoRequest
	PluginInfoReply
	StopRequest
	StopReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto2 "github.com/spiffe/node-agent/plugins/common/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type ConfigureRequest proto2.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*proto2.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*proto2.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*proto2.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type ConfigureResponse proto2.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*proto2.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*proto2.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*proto2.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type GetPluginInfoRequest proto2.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*proto2.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*proto2.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type GetPluginInfoResponse proto2.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()          { (*proto2.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string  { return (*proto2.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()     {}
func (m *GetPluginInfoResponse) GetName() string { return (*proto2.GetPluginInfoResponse)(m).GetName() }
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*proto2.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string { return (*proto2.GetPluginInfoResponse)(m).GetType() }
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*proto2.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*proto2.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*proto2.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*proto2.GetPluginInfoResponse)(m).GetCompany()
}

type PluginInfoRequest struct {
}

func (m *PluginInfoRequest) Reset()                    { *m = PluginInfoRequest{} }
func (m *PluginInfoRequest) String() string            { return proto1.CompactTextString(m) }
func (*PluginInfoRequest) ProtoMessage()               {}
func (*PluginInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PluginInfoReply struct {
	PluginInfo []*proto2.GetPluginInfoResponse `protobuf:"bytes,1,rep,name=pluginInfo" json:"pluginInfo,omitempty"`
}

func (m *PluginInfoReply) Reset()                    { *m = PluginInfoReply{} }
func (m *PluginInfoReply) String() string            { return proto1.CompactTextString(m) }
func (*PluginInfoReply) ProtoMessage()               {}
func (*PluginInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PluginInfoReply) GetPluginInfo() []*proto2.GetPluginInfoResponse {
	if m != nil {
		return m.PluginInfo
	}
	return nil
}

type StopRequest struct {
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto1.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StopReply struct {
}

func (m *StopReply) Reset()                    { *m = StopReply{} }
func (m *StopReply) String() string            { return proto1.CompactTextString(m) }
func (*StopReply) ProtoMessage()               {}
func (*StopReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto1.RegisterType((*PluginInfoRequest)(nil), "proto.PluginInfoRequest")
	proto1.RegisterType((*PluginInfoReply)(nil), "proto.PluginInfoReply")
	proto1.RegisterType((*StopRequest)(nil), "proto.StopRequest")
	proto1.RegisterType((*StopReply)(nil), "proto.StopReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Server service

type ServerClient interface {
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error)
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoReply, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := grpc.Invoke(ctx, "/proto.Server/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoReply, error) {
	out := new(PluginInfoReply)
	err := grpc.Invoke(ctx, "/proto.Server/PluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerServer interface {
	Stop(context.Context, *StopRequest) (*StopReply, error)
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoReply, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stop",
			Handler:    _Server_Stop_Handler,
		},
		{
			MethodName: "PluginInfo",
			Handler:    _Server_PluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto1.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0x87, 0x40,
	0x10, 0xc5, 0x8b, 0xea, 0x0f, 0x8d, 0x45, 0xb5, 0x41, 0x88, 0x74, 0x08, 0x4f, 0x1d, 0xca, 0x05,
	0xbb, 0x86, 0xd7, 0xe8, 0x94, 0xe8, 0x27, 0x48, 0x1b, 0x4d, 0x58, 0x77, 0x26, 0x77, 0x15, 0xfc,
	0xf6, 0xe1, 0x6a, 0xb5, 0x51, 0xa7, 0x65, 0xde, 0x7b, 0xf3, 0xdb, 0x79, 0x70, 0x62, 0x70, 0x98,
	0x70, 0x48, 0x78, 0x20, 0x4b, 0xe2, 0xc8, 0x3d, 0x51, 0xd6, 0x76, 0xf6, 0x7d, 0xac, 0x92, 0x9a,
	0x7a, 0x69, 0xb8, 0x6b, 0x1a, 0x94, 0x9a, 0xde, 0xf0, 0xfe, 0xb5, 0x45, 0x6d, 0x25, 0xab, 0xb1,
	0xed, 0xb4, 0x91, 0x35, 0xf5, 0x3d, 0x69, 0xe9, 0x56, 0xb6, 0x61, 0xc5, 0xc4, 0x97, 0x70, 0x91,
	0xbb, 0xd0, 0xb3, 0x6e, 0xa8, 0xc0, 0x8f, 0x11, 0x8d, 0x8d, 0x5f, 0xe0, 0xcc, 0x17, 0x59, 0xcd,
	0xe2, 0x11, 0x80, 0xbf, 0xa5, 0x70, 0xff, 0xe6, 0xe0, 0x36, 0x48, 0xaf, 0x57, 0x46, 0xf2, 0x84,
	0xd6, 0x8f, 0x1b, 0x26, 0x6d, 0xb0, 0xf0, 0xf2, 0xf1, 0x29, 0x04, 0xa5, 0x25, 0xfe, 0xe2, 0x07,
	0x70, 0xbc, 0x8e, 0xac, 0xe6, 0x74, 0x82, 0x5d, 0xe9, 0x8a, 0x89, 0x3b, 0x38, 0x5c, 0x64, 0x21,
	0x36, 0xae, 0xb7, 0x12, 0x9d, 0xff, 0xd2, 0x96, 0x8b, 0x32, 0x80, 0x9f, 0x5f, 0x45, 0xb8, 0xf9,
	0x7f, 0xca, 0x44, 0x57, 0xff, 0x38, 0xac, 0xe6, 0x7c, 0xaf, 0xda, 0x39, 0xe3, 0xe1, 0x33, 0x00,
	0x00, 0xff, 0xff, 0xc2, 0x2a, 0xcd, 0x6c, 0x59, 0x01, 0x00, 0x00,
}