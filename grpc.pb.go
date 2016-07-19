// Code generated by protoc-gen-go.
// source: grpc.proto
// DO NOT EDIT!

package gobeam

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

type StreamRequest struct {
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*StreamRequest)(nil), "gobeam.StreamRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for RobotService service

type RobotServiceClient interface {
	StreamReport(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (RobotService_StreamReportClient, error)
}

type robotServiceClient struct {
	cc *grpc.ClientConn
}

func NewRobotServiceClient(cc *grpc.ClientConn) RobotServiceClient {
	return &robotServiceClient{cc}
}

func (c *robotServiceClient) StreamReport(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (RobotService_StreamReportClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RobotService_serviceDesc.Streams[0], c.cc, "/gobeam.RobotService/StreamReport", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotServiceStreamReportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RobotService_StreamReportClient interface {
	Recv() (*Report, error)
	grpc.ClientStream
}

type robotServiceStreamReportClient struct {
	grpc.ClientStream
}

func (x *robotServiceStreamReportClient) Recv() (*Report, error) {
	m := new(Report)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RobotService service

type RobotServiceServer interface {
	StreamReport(*StreamRequest, RobotService_StreamReportServer) error
}

func RegisterRobotServiceServer(s *grpc.Server, srv RobotServiceServer) {
	s.RegisterService(&_RobotService_serviceDesc, srv)
}

func _RobotService_StreamReport_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RobotServiceServer).StreamReport(m, &robotServiceStreamReportServer{stream})
}

type RobotService_StreamReportServer interface {
	Send(*Report) error
	grpc.ServerStream
}

type robotServiceStreamReportServer struct {
	grpc.ServerStream
}

func (x *robotServiceStreamReportServer) Send(m *Report) error {
	return x.ServerStream.SendMsg(m)
}

var _RobotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gobeam.RobotService",
	HandlerType: (*RobotServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamReport",
			Handler:       _RobotService_StreamReport_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor1 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x4f, 0x4a, 0x4d, 0xcc, 0x95, 0xe2,
	0x02, 0x91, 0x10, 0x31, 0x25, 0x7e, 0x2e, 0xde, 0xe0, 0x92, 0x22, 0x20, 0x3f, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0xc4, 0xc8, 0x93, 0x8b, 0x27, 0x28, 0x3f, 0x29, 0xbf, 0x24, 0x38, 0xb5, 0xa8,
	0x2c, 0x33, 0x39, 0x55, 0xc8, 0x92, 0x8b, 0x07, 0xa6, 0xa0, 0x20, 0xbf, 0xa8, 0x44, 0x48, 0x54,
	0x0f, 0x62, 0x8a, 0x1e, 0x8a, 0x36, 0x29, 0x3e, 0x98, 0x30, 0x44, 0x99, 0x12, 0x83, 0x01, 0x63,
	0x12, 0x1b, 0xd8, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xec, 0xdb, 0xc9, 0x84,
	0x00, 0x00, 0x00,
}