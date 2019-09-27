// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: manager.proto

package managerproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// 请求用户信息
type GroupToInfo struct {
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid                  uint64   `protobuf:"varint,3,opt,name=gid,proto3" json:"gid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupToInfo) Reset()         { *m = GroupToInfo{} }
func (m *GroupToInfo) String() string { return proto.CompactTextString(m) }
func (*GroupToInfo) ProtoMessage()    {}
func (*GroupToInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{0}
}
func (m *GroupToInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupToInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupToInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupToInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupToInfo.Merge(m, src)
}
func (m *GroupToInfo) XXX_Size() int {
	return m.Size()
}
func (m *GroupToInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupToInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupToInfo proto.InternalMessageInfo

func (m *GroupToInfo) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GroupToInfo) GetGid() uint64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

type GroupToAckInfo struct {
	Id                   int32    `protobuf:"zigzag32,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupToAckInfo) Reset()         { *m = GroupToAckInfo{} }
func (m *GroupToAckInfo) String() string { return proto.CompactTextString(m) }
func (*GroupToAckInfo) ProtoMessage()    {}
func (*GroupToAckInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{1}
}
func (m *GroupToAckInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupToAckInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupToAckInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupToAckInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupToAckInfo.Merge(m, src)
}
func (m *GroupToAckInfo) XXX_Size() int {
	return m.Size()
}
func (m *GroupToAckInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupToAckInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GroupToAckInfo proto.InternalMessageInfo

func (m *GroupToAckInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 请求用户信息
type MicToInfo struct {
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid                  uint64   `protobuf:"varint,3,opt,name=gid,proto3" json:"gid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MicToInfo) Reset()         { *m = MicToInfo{} }
func (m *MicToInfo) String() string { return proto.CompactTextString(m) }
func (*MicToInfo) ProtoMessage()    {}
func (*MicToInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{2}
}
func (m *MicToInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MicToInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MicToInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MicToInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicToInfo.Merge(m, src)
}
func (m *MicToInfo) XXX_Size() int {
	return m.Size()
}
func (m *MicToInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MicToInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MicToInfo proto.InternalMessageInfo

func (m *MicToInfo) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MicToInfo) GetGid() uint64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

type MicToAckInfo struct {
	Id                   int32    `protobuf:"zigzag32,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MicToAckInfo) Reset()         { *m = MicToAckInfo{} }
func (m *MicToAckInfo) String() string { return proto.CompactTextString(m) }
func (*MicToAckInfo) ProtoMessage()    {}
func (*MicToAckInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{3}
}
func (m *MicToAckInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MicToAckInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MicToAckInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MicToAckInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicToAckInfo.Merge(m, src)
}
func (m *MicToAckInfo) XXX_Size() int {
	return m.Size()
}
func (m *MicToAckInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MicToAckInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MicToAckInfo proto.InternalMessageInfo

func (m *MicToAckInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BroadcastRegiste struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastRegiste) Reset()         { *m = BroadcastRegiste{} }
func (m *BroadcastRegiste) String() string { return proto.CompactTextString(m) }
func (*BroadcastRegiste) ProtoMessage()    {}
func (*BroadcastRegiste) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{4}
}
func (m *BroadcastRegiste) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BroadcastRegiste) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BroadcastRegiste.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BroadcastRegiste) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastRegiste.Merge(m, src)
}
func (m *BroadcastRegiste) XXX_Size() int {
	return m.Size()
}
func (m *BroadcastRegiste) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastRegiste.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastRegiste proto.InternalMessageInfo

func (m *BroadcastRegiste) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 请求用户信息的结果
type BroadcastInfo struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastInfo) Reset()         { *m = BroadcastInfo{} }
func (m *BroadcastInfo) String() string { return proto.CompactTextString(m) }
func (*BroadcastInfo) ProtoMessage()    {}
func (*BroadcastInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cde9ec64f0d2c859, []int{5}
}
func (m *BroadcastInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BroadcastInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BroadcastInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BroadcastInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastInfo.Merge(m, src)
}
func (m *BroadcastInfo) XXX_Size() int {
	return m.Size()
}
func (m *BroadcastInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastInfo proto.InternalMessageInfo

func (m *BroadcastInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*GroupToInfo)(nil), "managerproto.GroupToInfo")
	proto.RegisterType((*GroupToAckInfo)(nil), "managerproto.GroupToAckInfo")
	proto.RegisterType((*MicToInfo)(nil), "managerproto.MicToInfo")
	proto.RegisterType((*MicToAckInfo)(nil), "managerproto.MicToAckInfo")
	proto.RegisterType((*BroadcastRegiste)(nil), "managerproto.BroadcastRegiste")
	proto.RegisterType((*BroadcastInfo)(nil), "managerproto.BroadcastInfo")
}

func init() { proto.RegisterFile("manager.proto", fileDescriptor_cde9ec64f0d2c859) }

var fileDescriptor_cde9ec64f0d2c859 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0x72, 0xc1, 0x3c, 0x25,
	0x43, 0x2e, 0x6e, 0xf7, 0xa2, 0xfc, 0xd2, 0x82, 0x90, 0x7c, 0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x01,
	0x2e, 0xe6, 0xd2, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x10, 0x13, 0x24, 0x92,
	0x9e, 0x99, 0x22, 0xc1, 0x0c, 0x11, 0x49, 0xcf, 0x4c, 0x51, 0x52, 0xe0, 0xe2, 0x83, 0x6a, 0x71,
	0x4c, 0xce, 0x06, 0xeb, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x10, 0x0c,
	0x62, 0xca, 0x4c, 0x51, 0xd2, 0xe7, 0xe2, 0xf4, 0xcd, 0x4c, 0x26, 0xc1, 0x48, 0x39, 0x2e, 0x1e,
	0xb0, 0x06, 0x5c, 0x06, 0x2a, 0x71, 0x09, 0x38, 0x15, 0xe5, 0x27, 0xa6, 0x24, 0x27, 0x16, 0x97,
	0x04, 0xa5, 0xa6, 0x67, 0x16, 0x97, 0xa4, 0x22, 0xa9, 0xe1, 0x05, 0xab, 0x91, 0xe7, 0xe2, 0x85,
	0xab, 0x41, 0x33, 0x04, 0xac, 0xc0, 0xe8, 0x11, 0x13, 0x17, 0x8b, 0x4b, 0x62, 0x49, 0xa2, 0x90,
	0x0b, 0x17, 0xa7, 0x57, 0x7e, 0x66, 0x1e, 0xd8, 0x13, 0x42, 0x92, 0x7a, 0xc8, 0xe1, 0xa1, 0x87,
	0x14, 0x18, 0x52, 0x32, 0x58, 0xa5, 0x60, 0x6e, 0x74, 0xe5, 0xe2, 0xf2, 0x49, 0x4d, 0x2c, 0x4b,
	0xa5, 0xd0, 0x18, 0x3b, 0x2e, 0x76, 0x8f, 0xfc, 0x9c, 0x14, 0xdf, 0xcc, 0x64, 0x21, 0x71, 0x54,
	0x85, 0xf0, 0x20, 0x94, 0x92, 0xc2, 0x22, 0x01, 0xd3, 0xef, 0xc8, 0xc5, 0x15, 0x94, 0x9a, 0x93,
	0x9a, 0x58, 0x9c, 0x4a, 0xb6, 0x11, 0x5e, 0x5c, 0x9c, 0xf0, 0x90, 0x13, 0x92, 0x43, 0x55, 0x88,
	0x1e, 0xec, 0x52, 0xd2, 0x38, 0xe4, 0x41, 0x26, 0x19, 0x30, 0x3a, 0x09, 0x9c, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81,
	0x15, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05, 0x6c, 0x86, 0x46, 0x87, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataClient interface {
	JoinGroup(ctx context.Context, in *GroupToInfo, opts ...grpc.CallOption) (*GroupToAckInfo, error)
	LeaveGroup(ctx context.Context, in *GroupToInfo, opts ...grpc.CallOption) (*GroupToAckInfo, error)
	HoldMic(ctx context.Context, in *MicToInfo, opts ...grpc.CallOption) (*MicToAckInfo, error)
	ReleaseMic(ctx context.Context, in *MicToInfo, opts ...grpc.CallOption) (*MicToAckInfo, error)
	Broadcast(ctx context.Context, in *BroadcastRegiste, opts ...grpc.CallOption) (Data_BroadcastClient, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) JoinGroup(ctx context.Context, in *GroupToInfo, opts ...grpc.CallOption) (*GroupToAckInfo, error) {
	out := new(GroupToAckInfo)
	err := c.cc.Invoke(ctx, "/managerproto.Data/JoinGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) LeaveGroup(ctx context.Context, in *GroupToInfo, opts ...grpc.CallOption) (*GroupToAckInfo, error) {
	out := new(GroupToAckInfo)
	err := c.cc.Invoke(ctx, "/managerproto.Data/LeaveGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) HoldMic(ctx context.Context, in *MicToInfo, opts ...grpc.CallOption) (*MicToAckInfo, error) {
	out := new(MicToAckInfo)
	err := c.cc.Invoke(ctx, "/managerproto.Data/HoldMic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) ReleaseMic(ctx context.Context, in *MicToInfo, opts ...grpc.CallOption) (*MicToAckInfo, error) {
	out := new(MicToAckInfo)
	err := c.cc.Invoke(ctx, "/managerproto.Data/ReleaseMic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Broadcast(ctx context.Context, in *BroadcastRegiste, opts ...grpc.CallOption) (Data_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Data_serviceDesc.Streams[0], "/managerproto.Data/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataBroadcastClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Data_BroadcastClient interface {
	Recv() (*BroadcastInfo, error)
	grpc.ClientStream
}

type dataBroadcastClient struct {
	grpc.ClientStream
}

func (x *dataBroadcastClient) Recv() (*BroadcastInfo, error) {
	m := new(BroadcastInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServer is the server API for Data service.
type DataServer interface {
	JoinGroup(context.Context, *GroupToInfo) (*GroupToAckInfo, error)
	LeaveGroup(context.Context, *GroupToInfo) (*GroupToAckInfo, error)
	HoldMic(context.Context, *MicToInfo) (*MicToAckInfo, error)
	ReleaseMic(context.Context, *MicToInfo) (*MicToAckInfo, error)
	Broadcast(*BroadcastRegiste, Data_BroadcastServer) error
}

// UnimplementedDataServer can be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (*UnimplementedDataServer) JoinGroup(ctx context.Context, req *GroupToInfo) (*GroupToAckInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (*UnimplementedDataServer) LeaveGroup(ctx context.Context, req *GroupToInfo) (*GroupToAckInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroup not implemented")
}
func (*UnimplementedDataServer) HoldMic(ctx context.Context, req *MicToInfo) (*MicToAckInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HoldMic not implemented")
}
func (*UnimplementedDataServer) ReleaseMic(ctx context.Context, req *MicToInfo) (*MicToAckInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseMic not implemented")
}
func (*UnimplementedDataServer) Broadcast(req *BroadcastRegiste, srv Data_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupToInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/managerproto.Data/JoinGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).JoinGroup(ctx, req.(*GroupToInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_LeaveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupToInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).LeaveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/managerproto.Data/LeaveGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).LeaveGroup(ctx, req.(*GroupToInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_HoldMic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicToInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).HoldMic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/managerproto.Data/HoldMic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).HoldMic(ctx, req.(*MicToInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_ReleaseMic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicToInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).ReleaseMic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/managerproto.Data/ReleaseMic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).ReleaseMic(ctx, req.(*MicToInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BroadcastRegiste)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServer).Broadcast(m, &dataBroadcastServer{stream})
}

type Data_BroadcastServer interface {
	Send(*BroadcastInfo) error
	grpc.ServerStream
}

type dataBroadcastServer struct {
	grpc.ServerStream
}

func (x *dataBroadcastServer) Send(m *BroadcastInfo) error {
	return x.ServerStream.SendMsg(m)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "managerproto.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinGroup",
			Handler:    _Data_JoinGroup_Handler,
		},
		{
			MethodName: "LeaveGroup",
			Handler:    _Data_LeaveGroup_Handler,
		},
		{
			MethodName: "HoldMic",
			Handler:    _Data_HoldMic_Handler,
		},
		{
			MethodName: "ReleaseMic",
			Handler:    _Data_ReleaseMic_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _Data_Broadcast_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manager.proto",
}

func (m *GroupToInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupToInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupToInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Gid != 0 {
		i = encodeVarintManager(dAtA, i, uint64(m.Gid))
		i--
		dAtA[i] = 0x18
	}
	if m.Uid != 0 {
		i = encodeVarintManager(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func (m *GroupToAckInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupToAckInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupToAckInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintManager(dAtA, i, uint64((uint32(m.Id)<<1)^uint32((m.Id>>31))))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MicToInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MicToInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MicToInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Gid != 0 {
		i = encodeVarintManager(dAtA, i, uint64(m.Gid))
		i--
		dAtA[i] = 0x18
	}
	if m.Uid != 0 {
		i = encodeVarintManager(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func (m *MicToAckInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MicToAckInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MicToAckInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintManager(dAtA, i, uint64((uint32(m.Id)<<1)^uint32((m.Id>>31))))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BroadcastRegiste) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BroadcastRegiste) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BroadcastRegiste) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintManager(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BroadcastInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BroadcastInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BroadcastInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintManager(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintManager(dAtA []byte, offset int, v uint64) int {
	offset -= sovManager(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GroupToInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovManager(uint64(m.Uid))
	}
	if m.Gid != 0 {
		n += 1 + sovManager(uint64(m.Gid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GroupToAckInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sozManager(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MicToInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovManager(uint64(m.Uid))
	}
	if m.Gid != 0 {
		n += 1 + sovManager(uint64(m.Gid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MicToAckInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sozManager(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BroadcastRegiste) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovManager(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BroadcastInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovManager(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovManager(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozManager(x uint64) (n int) {
	return sovManager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GroupToInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupToInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupToInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gid", wireType)
			}
			m.Gid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupToAckInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupToAckInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupToAckInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Id = v
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MicToInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MicToInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MicToInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gid", wireType)
			}
			m.Gid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MicToAckInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MicToAckInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MicToAckInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Id = v
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BroadcastRegiste) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BroadcastRegiste: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BroadcastRegiste: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BroadcastInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BroadcastInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BroadcastInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipManager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowManager
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthManager
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthManager
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowManager
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipManager(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthManager
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthManager = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowManager   = fmt.Errorf("proto: integer overflow")
)