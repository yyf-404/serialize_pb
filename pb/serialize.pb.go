// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/serialize.proto

package serialize

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

type ResultCode int32

const (
	ResultCode_SUCCESS       ResultCode = 0
	ResultCode_SERVER_ERR    ResultCode = -1
	ResultCode_INVALID_PARAM ResultCode = -2
)

var ResultCode_name = map[int32]string{
	0:  "SUCCESS",
	-1: "SERVER_ERR",
	-2: "INVALID_PARAM",
}

var ResultCode_value = map[string]int32{
	"SUCCESS":       0,
	"SERVER_ERR":    -1,
	"INVALID_PARAM": -2,
}

func (x ResultCode) String() string {
	return proto.EnumName(ResultCode_name, int32(x))
}

func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ed1a15a203eeb68, []int{0}
}

type MsgRequest struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsUser               bool     `protobuf:"varint,3,opt,name=isUser,proto3" json:"isUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRequest) Reset()         { *m = MsgRequest{} }
func (m *MsgRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRequest) ProtoMessage()    {}
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ed1a15a203eeb68, []int{0}
}

func (m *MsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRequest.Unmarshal(m, b)
}
func (m *MsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRequest.Marshal(b, m, deterministic)
}
func (m *MsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequest.Merge(m, src)
}
func (m *MsgRequest) XXX_Size() int {
	return xxx_messageInfo_MsgRequest.Size(m)
}
func (m *MsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequest proto.InternalMessageInfo

func (m *MsgRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgRequest) GetIsUser() bool {
	if m != nil {
		return m.IsUser
	}
	return false
}

type MsgReply struct {
	Code                 ResultCode `protobuf:"varint,1,opt,name=code,proto3,enum=serialize.ResultCode" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgReply) Reset()         { *m = MsgReply{} }
func (m *MsgReply) String() string { return proto.CompactTextString(m) }
func (*MsgReply) ProtoMessage()    {}
func (*MsgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ed1a15a203eeb68, []int{1}
}

func (m *MsgReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgReply.Unmarshal(m, b)
}
func (m *MsgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgReply.Marshal(b, m, deterministic)
}
func (m *MsgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReply.Merge(m, src)
}
func (m *MsgReply) XXX_Size() int {
	return xxx_messageInfo_MsgReply.Size(m)
}
func (m *MsgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReply.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReply proto.InternalMessageInfo

func (m *MsgReply) GetCode() ResultCode {
	if m != nil {
		return m.Code
	}
	return ResultCode_SUCCESS
}

func (m *MsgReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("serialize.ResultCode", ResultCode_name, ResultCode_value)
	proto.RegisterType((*MsgRequest)(nil), "serialize.MsgRequest")
	proto.RegisterType((*MsgReply)(nil), "serialize.MsgReply")
}

func init() { proto.RegisterFile("pb/serialize.proto", fileDescriptor_8ed1a15a203eeb68) }

var fileDescriptor_8ed1a15a203eeb68 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0x6a, 0x6d, 0xda, 0x11, 0x4b, 0x18, 0x51, 0x43, 0x4f, 0x21, 0xa7, 0xe8, 0xa1, 0x42,
	0x3d, 0x78, 0x0e, 0x31, 0x60, 0xd1, 0x54, 0x99, 0xa5, 0xbd, 0x96, 0xb4, 0x3b, 0x84, 0x85, 0xb4,
	0x89, 0xd9, 0x54, 0xa8, 0x1f, 0xaf, 0xc2, 0x62, 0x1a, 0x71, 0x4e, 0xf3, 0xde, 0xbc, 0x99, 0x79,
	0x3c, 0xc0, 0x72, 0x7d, 0xa7, 0xb9, 0x52, 0x69, 0xae, 0x3e, 0x79, 0x52, 0x56, 0x45, 0x5d, 0xe0,
	0xf0, 0x48, 0xf8, 0x4f, 0x00, 0x89, 0xce, 0x88, 0xdf, 0xf7, 0xac, 0x6b, 0x44, 0xe8, 0xed, 0xd2,
	0x2d, 0xbb, 0x5d, 0xcf, 0x0a, 0x86, 0x64, 0x7a, 0x1c, 0x41, 0x57, 0x49, 0xd7, 0xf2, 0xac, 0xe0,
	0x94, 0xba, 0x4a, 0xe2, 0x15, 0xf4, 0x95, 0x5e, 0x68, 0xae, 0xdc, 0x13, 0xcf, 0x0a, 0x06, 0xf4,
	0x8b, 0xfc, 0x57, 0x18, 0x98, 0x4b, 0x65, 0x7e, 0xc0, 0x1b, 0xe8, 0x6d, 0x0a, 0xc9, 0x66, 0x6b,
	0x34, 0xbd, 0x9c, 0xb4, 0x06, 0x88, 0xf5, 0x3e, 0xaf, 0xa3, 0x42, 0x32, 0x19, 0x09, 0xba, 0x60,
	0x6f, 0x59, 0xeb, 0x34, 0x6b, 0xbe, 0x36, 0xf0, 0x76, 0x0e, 0xd0, 0xaa, 0xf1, 0x0c, 0x6c, 0xb1,
	0x88, 0xa2, 0x58, 0x08, 0xa7, 0x83, 0xd7, 0x00, 0x22, 0xa6, 0x65, 0x4c, 0xab, 0x98, 0xc8, 0xf9,
	0x6e, 0xca, 0xc2, 0x31, 0x9c, 0xcf, 0xe6, 0xcb, 0xf0, 0x65, 0xf6, 0xb8, 0x7a, 0x0b, 0x29, 0x4c,
	0x9c, 0xaf, 0xe3, 0x6c, 0xfa, 0x0c, 0x8e, 0x68, 0x7c, 0x08, 0xae, 0x3e, 0xd4, 0x86, 0xf1, 0x01,
	0x6c, 0xc1, 0x3b, 0x99, 0xe8, 0x0c, 0xff, 0xba, 0x6c, 0x23, 0x19, 0x5f, 0xfc, 0xa7, 0xcb, 0xfc,
	0xe0, 0x77, 0xd6, 0x7d, 0x93, 0xe4, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x8c, 0xae,
	0xe5, 0x5f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SerializeServiceClient is the client API for SerializeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SerializeServiceClient interface {
	//  方法
	SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error)
}

type serializeServiceClient struct {
	cc *grpc.ClientConn
}

func NewSerializeServiceClient(cc *grpc.ClientConn) SerializeServiceClient {
	return &serializeServiceClient{cc}
}

func (c *serializeServiceClient) SendMsg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/serialize.SerializeService/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SerializeServiceServer is the server API for SerializeService service.
type SerializeServiceServer interface {
	//  方法
	SendMsg(context.Context, *MsgRequest) (*MsgReply, error)
}

// UnimplementedSerializeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSerializeServiceServer struct {
}

func (*UnimplementedSerializeServiceServer) SendMsg(ctx context.Context, req *MsgRequest) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}

func RegisterSerializeServiceServer(s *grpc.Server, srv SerializeServiceServer) {
	s.RegisterService(&_SerializeService_serviceDesc, srv)
}

func _SerializeService_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SerializeServiceServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serialize.SerializeService/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SerializeServiceServer).SendMsg(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SerializeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serialize.SerializeService",
	HandlerType: (*SerializeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _SerializeService_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/serialize.proto",
}
