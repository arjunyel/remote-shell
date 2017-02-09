// Code generated by protoc-gen-go.
// source: services.proto
// DO NOT EDIT!

/*
Package Lab1 is a generated protocol buffer package.

Maybe this needs to be changed if we run into problems.

It is generated from these files:
        services.proto

It has these top-level messages:
        CommandRequest
        CommandReply
*/
package Lab1

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

// Request message
type CommandRequest struct {
	CmdName string   `protobuf:"bytes,1,opt,name=cmdName" json:"cmdName,omitempty"`
	CmdArgs []string `protobuf:"bytes,2,rep,name=cmdArgs" json:"cmdArgs,omitempty"`
}

func (m *CommandRequest) Reset()                    { *m = CommandRequest{} }
func (m *CommandRequest) String() string            { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()               {}
func (*CommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommandRequest) GetCmdName() string {
	if m != nil {
		return m.CmdName
	}
	return ""
}

func (m *CommandRequest) GetCmdArgs() []string {
	if m != nil {
		return m.CmdArgs
	}
	return nil
}

// Response message
type CommandReply struct {
	Output string `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
}

func (m *CommandReply) Reset()                    { *m = CommandReply{} }
func (m *CommandReply) String() string            { return proto.CompactTextString(m) }
func (*CommandReply) ProtoMessage()               {}
func (*CommandReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommandReply) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*CommandRequest)(nil), "Lab1.CommandRequest")
	proto.RegisterType((*CommandReply)(nil), "Lab1.CommandReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RemoteCommand service

type RemoteCommandClient interface {
	// Sends a remote command from a client to a server.
	SendCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error)
}

type remoteCommandClient struct {
	cc *grpc.ClientConn
}

func NewRemoteCommandClient(cc *grpc.ClientConn) RemoteCommandClient {
	return &remoteCommandClient{cc}
}

func (c *remoteCommandClient) SendCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandReply, error) {
	out := new(CommandReply)
	err := grpc.Invoke(ctx, "/Lab1.RemoteCommand/SendCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RemoteCommand service

type RemoteCommandServer interface {
	// Sends a remote command from a client to a server.
	SendCommand(context.Context, *CommandRequest) (*CommandReply, error)
}

func RegisterRemoteCommandServer(s *grpc.Server, srv RemoteCommandServer) {
	s.RegisterService(&_RemoteCommand_serviceDesc, srv)
}

func _RemoteCommand_SendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCommandServer).SendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Lab1.RemoteCommand/SendCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCommandServer).SendCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Lab1.RemoteCommand",
	HandlerType: (*RemoteCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommand",
			Handler:    _RemoteCommand_SendCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

func init() { proto.RegisterFile("services.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xf1, 0x49, 0x4c, 0x32,
	0x54, 0x72, 0xe1, 0xe2, 0x73, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x09, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0xce, 0x4d, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0xa1, 0x32, 0x8e, 0x45, 0xe9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc,
	0x50, 0x19, 0x10, 0x57, 0x49, 0x8d, 0x8b, 0x07, 0x6e, 0x4a, 0x41, 0x4e, 0xa5, 0x90, 0x18, 0x17,
	0x5b, 0x7e, 0x69, 0x49, 0x41, 0x69, 0x09, 0xd4, 0x08, 0x28, 0xcf, 0xc8, 0x8b, 0x8b, 0x37, 0x28,
	0x35, 0x37, 0xbf, 0x24, 0x15, 0xaa, 0x5a, 0xc8, 0x92, 0x8b, 0x3b, 0x38, 0x35, 0x2f, 0x05, 0xc6,
	0x15, 0xd1, 0x03, 0x39, 0x4a, 0x0f, 0xd5, 0x45, 0x52, 0x42, 0x68, 0xa2, 0x05, 0x39, 0x95, 0x4a,
	0x0c, 0x49, 0x6c, 0x60, 0x6f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x65, 0xf8, 0xb4, 0xd1,
	0xd8, 0x00, 0x00, 0x00,
}
