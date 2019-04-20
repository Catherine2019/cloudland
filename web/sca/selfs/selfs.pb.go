// Code generated by protoc-gen-go. DO NOT EDIT.
// source: selfs.proto

package selfs

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

type UpgradeRequest struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpgradeRequest) Reset()         { *m = UpgradeRequest{} }
func (m *UpgradeRequest) String() string { return proto.CompactTextString(m) }
func (*UpgradeRequest) ProtoMessage()    {}
func (*UpgradeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d7fad5e26180b, []int{0}
}

func (m *UpgradeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeRequest.Unmarshal(m, b)
}
func (m *UpgradeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeRequest.Marshal(b, m, deterministic)
}
func (m *UpgradeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeRequest.Merge(m, src)
}
func (m *UpgradeRequest) XXX_Size() int {
	return xxx_messageInfo_UpgradeRequest.Size(m)
}
func (m *UpgradeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeRequest proto.InternalMessageInfo

func (m *UpgradeRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type UpgradeReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpgradeReply) Reset()         { *m = UpgradeReply{} }
func (m *UpgradeReply) String() string { return proto.CompactTextString(m) }
func (*UpgradeReply) ProtoMessage()    {}
func (*UpgradeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d7fad5e26180b, []int{1}
}

func (m *UpgradeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeReply.Unmarshal(m, b)
}
func (m *UpgradeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeReply.Marshal(b, m, deterministic)
}
func (m *UpgradeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeReply.Merge(m, src)
}
func (m *UpgradeReply) XXX_Size() int {
	return xxx_messageInfo_UpgradeReply.Size(m)
}
func (m *UpgradeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeReply proto.InternalMessageInfo

type RuntimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeRequest) Reset()         { *m = RuntimeRequest{} }
func (m *RuntimeRequest) String() string { return proto.CompactTextString(m) }
func (*RuntimeRequest) ProtoMessage()    {}
func (*RuntimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d7fad5e26180b, []int{2}
}

func (m *RuntimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeRequest.Unmarshal(m, b)
}
func (m *RuntimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeRequest.Marshal(b, m, deterministic)
}
func (m *RuntimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeRequest.Merge(m, src)
}
func (m *RuntimeRequest) XXX_Size() int {
	return xxx_messageInfo_RuntimeRequest.Size(m)
}
func (m *RuntimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeRequest proto.InternalMessageInfo

type RuntimeReply struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Executable           string   `protobuf:"bytes,2,opt,name=executable,proto3" json:"executable,omitempty"`
	Pid                  string   `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Pwd                  string   `protobuf:"bytes,6,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Args                 []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	Environ              []string `protobuf:"bytes,5,rep,name=environ,proto3" json:"environ,omitempty"`
	Netrc                []string `protobuf:"bytes,7,rep,name=netrc,proto3" json:"netrc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeReply) Reset()         { *m = RuntimeReply{} }
func (m *RuntimeReply) String() string { return proto.CompactTextString(m) }
func (*RuntimeReply) ProtoMessage()    {}
func (*RuntimeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d7fad5e26180b, []int{3}
}

func (m *RuntimeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeReply.Unmarshal(m, b)
}
func (m *RuntimeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeReply.Marshal(b, m, deterministic)
}
func (m *RuntimeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeReply.Merge(m, src)
}
func (m *RuntimeReply) XXX_Size() int {
	return xxx_messageInfo_RuntimeReply.Size(m)
}
func (m *RuntimeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeReply.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeReply proto.InternalMessageInfo

func (m *RuntimeReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RuntimeReply) GetExecutable() string {
	if m != nil {
		return m.Executable
	}
	return ""
}

func (m *RuntimeReply) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *RuntimeReply) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *RuntimeReply) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *RuntimeReply) GetEnviron() []string {
	if m != nil {
		return m.Environ
	}
	return nil
}

func (m *RuntimeReply) GetNetrc() []string {
	if m != nil {
		return m.Netrc
	}
	return nil
}

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d7fad5e26180b, []int{4}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReply) Reset()         { *m = SetReply{} }
func (m *SetReply) String() string { return proto.CompactTextString(m) }
func (*SetReply) ProtoMessage()    {}
func (*SetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_501d7fad5e26180b, []int{5}
}

func (m *SetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReply.Unmarshal(m, b)
}
func (m *SetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReply.Marshal(b, m, deterministic)
}
func (m *SetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReply.Merge(m, src)
}
func (m *SetReply) XXX_Size() int {
	return xxx_messageInfo_SetReply.Size(m)
}
func (m *SetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpgradeRequest)(nil), "com.github.ibm.cloudland.sca.selfs.UpgradeRequest")
	proto.RegisterType((*UpgradeReply)(nil), "com.github.ibm.cloudland.sca.selfs.UpgradeReply")
	proto.RegisterType((*RuntimeRequest)(nil), "com.github.ibm.cloudland.sca.selfs.RuntimeRequest")
	proto.RegisterType((*RuntimeReply)(nil), "com.github.ibm.cloudland.sca.selfs.RuntimeReply")
	proto.RegisterType((*SetRequest)(nil), "com.github.ibm.cloudland.sca.selfs.SetRequest")
	proto.RegisterType((*SetReply)(nil), "com.github.ibm.cloudland.sca.selfs.SetReply")
}

func init() { proto.RegisterFile("selfs.proto", fileDescriptor_501d7fad5e26180b) }

var fileDescriptor_501d7fad5e26180b = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x15, 0x04, 0xc8, 0x32, 0x8b, 0x10, 0xb2, 0xf6, 0x60, 0x71, 0x58, 0xa1, 0x1c, 0x56, 0x68,
	0x55, 0x39, 0x15, 0xed, 0x0f, 0x94, 0x5b, 0x0f, 0xbd, 0x04, 0xf5, 0xd2, 0x5b, 0xe2, 0x0c, 0x34,
	0xaa, 0x63, 0xbb, 0xb1, 0x03, 0xcd, 0x5f, 0xf5, 0x5b, 0xfa, 0x45, 0x55, 0x4c, 0x42, 0xcb, 0xa1,
	0x6a, 0x7a, 0x9b, 0xf7, 0xe6, 0xbd, 0xcc, 0xe4, 0x8d, 0xe1, 0xb7, 0x41, 0xb1, 0x35, 0x4c, 0x17,
	0xca, 0x2a, 0x12, 0x70, 0x95, 0xb3, 0x5d, 0x66, 0x1f, 0xcb, 0x84, 0x65, 0x49, 0xce, 0xb8, 0x50,
	0x65, 0x2a, 0x62, 0x99, 0x32, 0xc3, 0x63, 0xe6, 0x94, 0xc1, 0x7f, 0x98, 0xde, 0xeb, 0x5d, 0x11,
	0xa7, 0x18, 0xe1, 0x73, 0x89, 0xc6, 0x12, 0x0a, 0xfe, 0x1e, 0x0b, 0x93, 0x29, 0x49, 0x7b, 0x8b,
	0xde, 0x72, 0x1c, 0xb5, 0x30, 0x98, 0xc2, 0xe4, 0xa4, 0xd5, 0xa2, 0x0a, 0x66, 0x30, 0x8d, 0x4a,
	0x69, 0xb3, 0xbc, 0xf5, 0x06, 0xaf, 0x3d, 0x98, 0x9c, 0x28, 0x2d, 0xaa, 0xaf, 0x3f, 0x46, 0xfe,
	0x02, 0xe0, 0x0b, 0xf2, 0xd2, 0xc6, 0x89, 0x40, 0xda, 0x77, 0xcd, 0x4f, 0x0c, 0x99, 0x81, 0xa7,
	0xb3, 0x94, 0x7a, 0xae, 0x51, 0x97, 0x8e, 0x39, 0xa4, 0x74, 0xd4, 0x30, 0x87, 0x94, 0x10, 0x18,
	0xc4, 0xc5, 0xce, 0xd0, 0xc1, 0xc2, 0x5b, 0x8e, 0x23, 0x57, 0xd7, 0x13, 0x51, 0xee, 0xb3, 0x42,
	0x49, 0x3a, 0x74, 0x74, 0x0b, 0xc9, 0x1f, 0x18, 0x4a, 0xb4, 0x05, 0xa7, 0xbe, 0xe3, 0x8f, 0x20,
	0xb8, 0x06, 0xd8, 0xa0, 0x6d, 0x7f, 0x7e, 0x06, 0xde, 0x13, 0x56, 0xcd, 0xae, 0x75, 0x59, 0xbb,
	0xf6, 0xb1, 0x28, 0xdb, 0x15, 0x8f, 0x20, 0x00, 0xf8, 0xe5, 0x5c, 0x5a, 0x54, 0xab, 0xb7, 0x3e,
	0x8c, 0x37, 0x28, 0xb6, 0x37, 0x69, 0x9e, 0x49, 0xa2, 0xc0, 0x6f, 0x42, 0x22, 0x2b, 0xf6, 0xfd,
	0x01, 0xd8, 0x79, 0xfa, 0xf3, 0xcb, 0x1f, 0x79, 0xea, 0x88, 0x15, 0xf8, 0x4d, 0xe4, 0xdd, 0x06,
	0x9e, 0x9f, 0xac, 0xdb, 0xc0, 0xb3, 0x9b, 0x72, 0xf0, 0x36, 0x68, 0x09, 0xeb, 0x62, 0xfc, 0x88,
	0x76, 0x7e, 0xd1, 0x59, 0xaf, 0x45, 0xb5, 0x5e, 0x3e, 0xfc, 0x6b, 0xa4, 0x5c, 0xe5, 0xe1, 0xed,
	0xfa, 0x2e, 0x3c, 0xc9, 0xc3, 0x03, 0x26, 0xa1, 0xe1, 0x71, 0xe8, 0x2c, 0xc9, 0xc8, 0x3d, 0xf6,
	0xab, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x07, 0x8f, 0xfa, 0xfb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SelfAdminClient is the client API for SelfAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SelfAdminClient interface {
	Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeReply, error)
	Runtime(ctx context.Context, in *RuntimeRequest, opts ...grpc.CallOption) (*RuntimeReply, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error)
}

type selfAdminClient struct {
	cc *grpc.ClientConn
}

func NewSelfAdminClient(cc *grpc.ClientConn) SelfAdminClient {
	return &selfAdminClient{cc}
}

func (c *selfAdminClient) Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeReply, error) {
	out := new(UpgradeReply)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.selfs.SelfAdmin/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selfAdminClient) Runtime(ctx context.Context, in *RuntimeRequest, opts ...grpc.CallOption) (*RuntimeReply, error) {
	out := new(RuntimeReply)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.selfs.SelfAdmin/Runtime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selfAdminClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error) {
	out := new(SetReply)
	err := c.cc.Invoke(ctx, "/com.github.ibm.cloudland.sca.selfs.SelfAdmin/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SelfAdminServer is the server API for SelfAdmin service.
type SelfAdminServer interface {
	Upgrade(context.Context, *UpgradeRequest) (*UpgradeReply, error)
	Runtime(context.Context, *RuntimeRequest) (*RuntimeReply, error)
	Set(context.Context, *SetRequest) (*SetReply, error)
}

// UnimplementedSelfAdminServer can be embedded to have forward compatible implementations.
type UnimplementedSelfAdminServer struct {
}

func (*UnimplementedSelfAdminServer) Upgrade(ctx context.Context, req *UpgradeRequest) (*UpgradeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upgrade not implemented")
}
func (*UnimplementedSelfAdminServer) Runtime(ctx context.Context, req *RuntimeRequest) (*RuntimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Runtime not implemented")
}
func (*UnimplementedSelfAdminServer) Set(ctx context.Context, req *SetRequest) (*SetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

func RegisterSelfAdminServer(s *grpc.Server, srv SelfAdminServer) {
	s.RegisterService(&_SelfAdmin_serviceDesc, srv)
}

func _SelfAdmin_Upgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfAdminServer).Upgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.selfs.SelfAdmin/Upgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfAdminServer).Upgrade(ctx, req.(*UpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SelfAdmin_Runtime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuntimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfAdminServer).Runtime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.selfs.SelfAdmin/Runtime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfAdminServer).Runtime(ctx, req.(*RuntimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SelfAdmin_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfAdminServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.ibm.cloudland.sca.selfs.SelfAdmin/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfAdminServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SelfAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.ibm.cloudland.sca.selfs.SelfAdmin",
	HandlerType: (*SelfAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upgrade",
			Handler:    _SelfAdmin_Upgrade_Handler,
		},
		{
			MethodName: "Runtime",
			Handler:    _SelfAdmin_Runtime_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _SelfAdmin_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "selfs.proto",
}
