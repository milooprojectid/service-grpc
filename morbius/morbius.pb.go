// Code generated by protoc-gen-go. DO NOT EDIT.
// source: morbius/morbius.proto

package morbius

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SentimentRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SentimentRequest) Reset()         { *m = SentimentRequest{} }
func (m *SentimentRequest) String() string { return proto.CompactTextString(m) }
func (*SentimentRequest) ProtoMessage()    {}
func (*SentimentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab01c99392c64236, []int{0}
}

func (m *SentimentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SentimentRequest.Unmarshal(m, b)
}
func (m *SentimentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SentimentRequest.Marshal(b, m, deterministic)
}
func (m *SentimentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentimentRequest.Merge(m, src)
}
func (m *SentimentRequest) XXX_Size() int {
	return xxx_messageInfo_SentimentRequest.Size(m)
}
func (m *SentimentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SentimentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SentimentRequest proto.InternalMessageInfo

func (m *SentimentRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SentimentResponse struct {
	ModelResult          int32    `protobuf:"varint,1,opt,name=model_result,json=modelResult,proto3" json:"model_result,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SentimentResponse) Reset()         { *m = SentimentResponse{} }
func (m *SentimentResponse) String() string { return proto.CompactTextString(m) }
func (*SentimentResponse) ProtoMessage()    {}
func (*SentimentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab01c99392c64236, []int{1}
}

func (m *SentimentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SentimentResponse.Unmarshal(m, b)
}
func (m *SentimentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SentimentResponse.Marshal(b, m, deterministic)
}
func (m *SentimentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentimentResponse.Merge(m, src)
}
func (m *SentimentResponse) XXX_Size() int {
	return xxx_messageInfo_SentimentResponse.Size(m)
}
func (m *SentimentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SentimentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SentimentResponse proto.InternalMessageInfo

func (m *SentimentResponse) GetModelResult() int32 {
	if m != nil {
		return m.ModelResult
	}
	return 0
}

func (m *SentimentResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SentimentResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*SentimentRequest)(nil), "SentimentRequest")
	proto.RegisterType((*SentimentResponse)(nil), "SentimentResponse")
}

func init() {
	proto.RegisterFile("morbius/morbius.proto", fileDescriptor_ab01c99392c64236)
}

var fileDescriptor_ab01c99392c64236 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcd, 0x2f, 0x4a,
	0xca, 0x2c, 0x2d, 0xd6, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0x6a, 0x5c, 0x02,
	0xc1, 0xa9, 0x79, 0x25, 0x99, 0xb9, 0xa9, 0x79, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60,
	0xb6, 0x52, 0x0e, 0x97, 0x20, 0x92, 0xba, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x45, 0x2e,
	0x9e, 0xdc, 0xfc, 0x94, 0xd4, 0x9c, 0xf8, 0xa2, 0xd4, 0xe2, 0xd2, 0x1c, 0x88, 0x06, 0xd6, 0x20,
	0x6e, 0xb0, 0x58, 0x10, 0x58, 0x48, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3,
	0xa0, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x09, 0x6c, 0x24, 0xb2, 0x10, 0xdc, 0x36, 0x66, 0x84, 0x6d,
	0x46, 0x2e, 0x5c, 0x7c, 0xbe, 0x10, 0x67, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x19,
	0x71, 0x71, 0xc2, 0xed, 0x17, 0x12, 0xd4, 0x43, 0x77, 0xb3, 0x94, 0x90, 0x1e, 0x86, 0xf3, 0x92,
	0xd8, 0xc0, 0x5e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xa7, 0x85, 0xa9, 0xfb, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MorbiusServiceClient is the client API for MorbiusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MorbiusServiceClient interface {
	Sentiment(ctx context.Context, in *SentimentRequest, opts ...grpc.CallOption) (*SentimentResponse, error)
}

type morbiusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMorbiusServiceClient(cc grpc.ClientConnInterface) MorbiusServiceClient {
	return &morbiusServiceClient{cc}
}

func (c *morbiusServiceClient) Sentiment(ctx context.Context, in *SentimentRequest, opts ...grpc.CallOption) (*SentimentResponse, error) {
	out := new(SentimentResponse)
	err := c.cc.Invoke(ctx, "/MorbiusService/Sentiment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MorbiusServiceServer is the server API for MorbiusService service.
type MorbiusServiceServer interface {
	Sentiment(context.Context, *SentimentRequest) (*SentimentResponse, error)
}

// UnimplementedMorbiusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMorbiusServiceServer struct {
}

func (*UnimplementedMorbiusServiceServer) Sentiment(ctx context.Context, req *SentimentRequest) (*SentimentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sentiment not implemented")
}

func RegisterMorbiusServiceServer(s *grpc.Server, srv MorbiusServiceServer) {
	s.RegisterService(&_MorbiusService_serviceDesc, srv)
}

func _MorbiusService_Sentiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MorbiusServiceServer).Sentiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MorbiusService/Sentiment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MorbiusServiceServer).Sentiment(ctx, req.(*SentimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MorbiusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MorbiusService",
	HandlerType: (*MorbiusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sentiment",
			Handler:    _MorbiusService_Sentiment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "morbius/morbius.proto",
}
