// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dbs.proto

package dbs

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

type StatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{0}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

type Table struct {
	Rows                 int64    `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`
	Deleted              int64    `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{1}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetRows() int64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func (m *Table) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func (m *Table) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Stats struct {
	MaxOpenConnections   int64    `protobuf:"varint,1,opt,name=max_open_connections,json=maxOpenConnections,proto3" json:"max_open_connections,omitempty"`
	OpenConnections      int64    `protobuf:"varint,2,opt,name=open_connections,json=openConnections,proto3" json:"open_connections,omitempty"`
	InUse                int64    `protobuf:"varint,3,opt,name=in_use,json=inUse,proto3" json:"in_use,omitempty"`
	Idle                 int64    `protobuf:"varint,4,opt,name=idle,proto3" json:"idle,omitempty"`
	WaitCount            int64    `protobuf:"varint,5,opt,name=wait_count,json=waitCount,proto3" json:"wait_count,omitempty"`
	WaitDuration         int64    `protobuf:"varint,6,opt,name=wait_duration,json=waitDuration,proto3" json:"wait_duration,omitempty"`
	MaxIdleClosed        int64    `protobuf:"varint,7,opt,name=max_idle_closed,json=maxIdleClosed,proto3" json:"max_idle_closed,omitempty"`
	MaxLifetimeClosed    int64    `protobuf:"varint,8,opt,name=max_lifetime_closed,json=maxLifetimeClosed,proto3" json:"max_lifetime_closed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{2}
}

func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetMaxOpenConnections() int64 {
	if m != nil {
		return m.MaxOpenConnections
	}
	return 0
}

func (m *Stats) GetOpenConnections() int64 {
	if m != nil {
		return m.OpenConnections
	}
	return 0
}

func (m *Stats) GetInUse() int64 {
	if m != nil {
		return m.InUse
	}
	return 0
}

func (m *Stats) GetIdle() int64 {
	if m != nil {
		return m.Idle
	}
	return 0
}

func (m *Stats) GetWaitCount() int64 {
	if m != nil {
		return m.WaitCount
	}
	return 0
}

func (m *Stats) GetWaitDuration() int64 {
	if m != nil {
		return m.WaitDuration
	}
	return 0
}

func (m *Stats) GetMaxIdleClosed() int64 {
	if m != nil {
		return m.MaxIdleClosed
	}
	return 0
}

func (m *Stats) GetMaxLifetimeClosed() int64 {
	if m != nil {
		return m.MaxLifetimeClosed
	}
	return 0
}

type StatsReply struct {
	Stats                *Stats   `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsReply) Reset()         { *m = StatsReply{} }
func (m *StatsReply) String() string { return proto.CompactTextString(m) }
func (*StatsReply) ProtoMessage()    {}
func (*StatsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{3}
}

func (m *StatsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsReply.Unmarshal(m, b)
}
func (m *StatsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsReply.Marshal(b, m, deterministic)
}
func (m *StatsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsReply.Merge(m, src)
}
func (m *StatsReply) XXX_Size() int {
	return xxx_messageInfo_StatsReply.Size(m)
}
func (m *StatsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatsReply proto.InternalMessageInfo

func (m *StatsReply) GetStats() *Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type TablesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TablesRequest) Reset()         { *m = TablesRequest{} }
func (m *TablesRequest) String() string { return proto.CompactTextString(m) }
func (*TablesRequest) ProtoMessage()    {}
func (*TablesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{4}
}

func (m *TablesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesRequest.Unmarshal(m, b)
}
func (m *TablesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesRequest.Marshal(b, m, deterministic)
}
func (m *TablesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesRequest.Merge(m, src)
}
func (m *TablesRequest) XXX_Size() int {
	return xxx_messageInfo_TablesRequest.Size(m)
}
func (m *TablesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TablesRequest proto.InternalMessageInfo

type TablesReply struct {
	Tables               map[string]*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TablesReply) Reset()         { *m = TablesReply{} }
func (m *TablesReply) String() string { return proto.CompactTextString(m) }
func (*TablesReply) ProtoMessage()    {}
func (*TablesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{5}
}

func (m *TablesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesReply.Unmarshal(m, b)
}
func (m *TablesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesReply.Marshal(b, m, deterministic)
}
func (m *TablesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesReply.Merge(m, src)
}
func (m *TablesReply) XXX_Size() int {
	return xxx_messageInfo_TablesReply.Size(m)
}
func (m *TablesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesReply.DiscardUnknown(m)
}

var xxx_messageInfo_TablesReply proto.InternalMessageInfo

func (m *TablesReply) GetTables() map[string]*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

type ExecRequest struct {
	Sql                  string   `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecRequest) Reset()         { *m = ExecRequest{} }
func (m *ExecRequest) String() string { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()    {}
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{6}
}

func (m *ExecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecRequest.Unmarshal(m, b)
}
func (m *ExecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecRequest.Marshal(b, m, deterministic)
}
func (m *ExecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecRequest.Merge(m, src)
}
func (m *ExecRequest) XXX_Size() int {
	return xxx_messageInfo_ExecRequest.Size(m)
}
func (m *ExecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecRequest proto.InternalMessageInfo

func (m *ExecRequest) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

type ExecReply struct {
	Affected             int64    `protobuf:"varint,1,opt,name=affected,proto3" json:"affected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecReply) Reset()         { *m = ExecReply{} }
func (m *ExecReply) String() string { return proto.CompactTextString(m) }
func (*ExecReply) ProtoMessage()    {}
func (*ExecReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{7}
}

func (m *ExecReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecReply.Unmarshal(m, b)
}
func (m *ExecReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecReply.Marshal(b, m, deterministic)
}
func (m *ExecReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecReply.Merge(m, src)
}
func (m *ExecReply) XXX_Size() int {
	return xxx_messageInfo_ExecReply.Size(m)
}
func (m *ExecReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecReply.DiscardUnknown(m)
}

var xxx_messageInfo_ExecReply proto.InternalMessageInfo

func (m *ExecReply) GetAffected() int64 {
	if m != nil {
		return m.Affected
	}
	return 0
}

type QueryRequest struct {
	Sql                  string   `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{8}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

type Result struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{9}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type QueryReply struct {
	Result               []*Result `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}
func (*QueryReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e975d2518bdd2efa, []int{10}
}

func (m *QueryReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReply.Unmarshal(m, b)
}
func (m *QueryReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReply.Marshal(b, m, deterministic)
}
func (m *QueryReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReply.Merge(m, src)
}
func (m *QueryReply) XXX_Size() int {
	return xxx_messageInfo_QueryReply.Size(m)
}
func (m *QueryReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReply proto.InternalMessageInfo

func (m *QueryReply) GetResult() []*Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*StatsRequest)(nil), "com.ibm.cloudland.sca.dbs.StatsRequest")
	proto.RegisterType((*Table)(nil), "com.ibm.cloudland.sca.dbs.Table")
	proto.RegisterType((*Stats)(nil), "com.ibm.cloudland.sca.dbs.Stats")
	proto.RegisterType((*StatsReply)(nil), "com.ibm.cloudland.sca.dbs.StatsReply")
	proto.RegisterType((*TablesRequest)(nil), "com.ibm.cloudland.sca.dbs.TablesRequest")
	proto.RegisterType((*TablesReply)(nil), "com.ibm.cloudland.sca.dbs.TablesReply")
	proto.RegisterMapType((map[string]*Table)(nil), "com.ibm.cloudland.sca.dbs.TablesReply.TablesEntry")
	proto.RegisterType((*ExecRequest)(nil), "com.ibm.cloudland.sca.dbs.ExecRequest")
	proto.RegisterType((*ExecReply)(nil), "com.ibm.cloudland.sca.dbs.ExecReply")
	proto.RegisterType((*QueryRequest)(nil), "com.ibm.cloudland.sca.dbs.QueryRequest")
	proto.RegisterType((*Result)(nil), "com.ibm.cloudland.sca.dbs.Result")
	proto.RegisterType((*QueryReply)(nil), "com.ibm.cloudland.sca.dbs.QueryReply")
}

func init() { proto.RegisterFile("dbs.proto", fileDescriptor_e975d2518bdd2efa) }

var fileDescriptor_e975d2518bdd2efa = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xd7, 0x66, 0xc9, 0x96, 0xeb, 0xc6, 0x86, 0x19, 0x28, 0x54, 0x42, 0x04, 0x03, 0x5b,
	0x79, 0x89, 0x50, 0x91, 0x10, 0xf0, 0xc6, 0x7e, 0x08, 0xf1, 0x43, 0x42, 0x84, 0x1f, 0xd2, 0x00,
	0xa9, 0x72, 0x62, 0x4f, 0x8a, 0x70, 0xe2, 0x2e, 0x76, 0x58, 0xfb, 0x37, 0xf1, 0xce, 0x7f, 0xc6,
	0x3b, 0xf2, 0x39, 0xdd, 0x26, 0xd0, 0xb2, 0xbd, 0xf9, 0xce, 0xdf, 0xfb, 0xf4, 0xee, 0xbe, 0x6e,
	0x20, 0xe4, 0x99, 0x4e, 0xa6, 0xb5, 0x32, 0x8a, 0xdc, 0xce, 0x55, 0x99, 0x14, 0x59, 0x99, 0xe4,
	0x52, 0x35, 0x5c, 0xb2, 0x8a, 0x27, 0x3a, 0x67, 0x09, 0xcf, 0x34, 0xbd, 0x06, 0x6b, 0x1f, 0x0d,
	0x33, 0x3a, 0x15, 0xc7, 0x8d, 0xd0, 0x86, 0xbe, 0x05, 0xff, 0x13, 0xcb, 0xa4, 0x20, 0x04, 0x96,
	0x6b, 0x75, 0xa2, 0xa3, 0x7e, 0xdc, 0x1b, 0x79, 0x29, 0x9e, 0x49, 0x04, 0x2b, 0x5c, 0x48, 0x61,
	0x04, 0x8f, 0x3c, 0x4c, 0x2f, 0x42, 0xb2, 0x05, 0xbe, 0xa8, 0x6b, 0x55, 0x47, 0xcb, 0x71, 0x6f,
	0x14, 0xa6, 0x2e, 0xa0, 0xbf, 0xfa, 0xe0, 0x23, 0x9d, 0x3c, 0x86, 0xad, 0x92, 0xcd, 0x26, 0x6a,
	0x2a, 0xaa, 0x49, 0xae, 0xaa, 0x4a, 0xe4, 0xa6, 0x50, 0x95, 0x8e, 0x7a, 0x88, 0x21, 0x25, 0x9b,
	0xbd, 0x9f, 0x8a, 0x6a, 0xef, 0xec, 0x86, 0x3c, 0x82, 0xcd, 0xff, 0xd4, 0xae, 0x97, 0x0d, 0xf5,
	0x8f, 0xf4, 0x26, 0x04, 0x45, 0x35, 0x69, 0xb4, 0x68, 0xbb, 0xf2, 0x8b, 0xea, 0xb3, 0xc6, 0x09,
	0x0a, 0x2e, 0x05, 0xb6, 0xe4, 0xa5, 0x78, 0x26, 0x77, 0x00, 0x4e, 0x58, 0x61, 0x26, 0xb9, 0x6a,
	0x2a, 0x13, 0xf9, 0x78, 0x13, 0xda, 0xcc, 0x9e, 0x4d, 0x90, 0xfb, 0xb0, 0x8e, 0xd7, 0xbc, 0xa9,
	0x99, 0x65, 0x47, 0x01, 0x2a, 0xd6, 0x6c, 0x72, 0xbf, 0xcd, 0x91, 0x6d, 0xd8, 0xb0, 0xb3, 0x58,
	0xde, 0x24, 0x97, 0x4a, 0x0b, 0x1e, 0xad, 0xa0, 0x6c, 0xbd, 0x64, 0xb3, 0xd7, 0x5c, 0x8a, 0x3d,
	0x4c, 0x92, 0x04, 0x6e, 0x58, 0x9d, 0x2c, 0x8e, 0x84, 0x29, 0xca, 0x53, 0xed, 0x2a, 0x6a, 0xaf,
	0x97, 0x6c, 0xf6, 0xae, 0xbd, 0x71, 0x7a, 0xba, 0x0f, 0xd0, 0x5a, 0x31, 0x95, 0x73, 0xf2, 0x14,
	0x7c, 0x6d, 0x23, 0x5c, 0xd1, 0x60, 0x1c, 0x27, 0x17, 0x7a, 0x98, 0xb8, 0x2a, 0x27, 0xa7, 0x1b,
	0xb0, 0x8e, 0x06, 0x9e, 0x3a, 0xfa, 0xbb, 0x07, 0x83, 0x45, 0xc6, 0x82, 0xdf, 0x40, 0x60, 0x30,
	0x8c, 0x7a, 0xb1, 0x37, 0x1a, 0x8c, 0xc7, 0x1d, 0xe4, 0x73, 0x75, 0xed, 0xf9, 0xa0, 0x32, 0xf5,
	0x3c, 0x6d, 0x09, 0xc3, 0x6f, 0x0b, 0x34, 0xa6, 0xc9, 0x26, 0x78, 0x3f, 0xc4, 0x1c, 0x3b, 0x0e,
	0x53, 0x7b, 0xb4, 0x53, 0xfc, 0x64, 0xb2, 0x11, 0x68, 0x5d, 0xf7, 0x14, 0x08, 0x4a, 0x9d, 0xfc,
	0x45, 0xff, 0x59, 0x8f, 0xde, 0x85, 0xc1, 0xc1, 0x4c, 0xe4, 0xed, 0x1c, 0x16, 0xae, 0x8f, 0xe5,
	0x02, 0xae, 0x8f, 0x25, 0xdd, 0x81, 0xd0, 0x09, 0xec, 0x58, 0x43, 0x58, 0x65, 0x47, 0x47, 0x22,
	0xb7, 0x8f, 0xd3, 0xbd, 0xaa, 0xd3, 0x98, 0xc6, 0xb0, 0xf6, 0xa1, 0x11, 0xf5, 0xfc, 0x62, 0x54,
	0x0c, 0x41, 0x2a, 0x74, 0x23, 0x0d, 0xb9, 0x05, 0x01, 0xb6, 0xe0, 0xd6, 0x13, 0xa6, 0x6d, 0x44,
	0x5f, 0x01, 0xb4, 0x0c, 0xfb, 0x6b, 0xcf, 0x21, 0xa8, 0x51, 0xdf, 0x2e, 0xf1, 0x5e, 0xc7, 0x60,
	0x0e, 0x9c, 0xb6, 0x05, 0xe3, 0x3f, 0x7d, 0x58, 0xd9, 0xdf, 0x7d, 0xc9, 0xcb, 0xa2, 0x22, 0x87,
	0x8b, 0xff, 0xc7, 0xce, 0xa5, 0xf6, 0xba, 0xd6, 0x87, 0x0f, 0x2f, 0x17, 0x4e, 0xe5, 0x9c, 0x2e,
	0x91, 0xef, 0x10, 0x38, 0x6b, 0xc8, 0xe8, 0x0a, 0x06, 0x3b, 0xf8, 0xf6, 0xd5, 0x9e, 0x02, 0x5d,
	0x22, 0x5f, 0x60, 0xd9, 0xae, 0x9e, 0x74, 0x55, 0x9c, 0x33, 0x6f, 0xf8, 0xe0, 0x52, 0x9d, 0xe3,
	0x1e, 0x82, 0x8f, 0x5b, 0xee, 0x5c, 0xc8, 0x79, 0x2f, 0x3b, 0x17, 0x72, 0x66, 0x18, 0x5d, 0xda,
	0xf5, 0xbf, 0x7a, 0x3c, 0xd3, 0x59, 0x80, 0x9f, 0xc4, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xac, 0xce, 0x6c, 0xd8, 0x1f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DBAdminClient is the client API for DBAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBAdminClient interface {
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error)
	Tables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesReply, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecReply, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
}

type dBAdminClient struct {
	cc *grpc.ClientConn
}

func NewDBAdminClient(cc *grpc.ClientConn) DBAdminClient {
	return &dBAdminClient{cc}
}

func (c *dBAdminClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error) {
	out := new(StatsReply)
	err := c.cc.Invoke(ctx, "/com.ibm.cloudland.sca.dbs.DBAdmin/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBAdminClient) Tables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesReply, error) {
	out := new(TablesReply)
	err := c.cc.Invoke(ctx, "/com.ibm.cloudland.sca.dbs.DBAdmin/Tables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBAdminClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecReply, error) {
	out := new(ExecReply)
	err := c.cc.Invoke(ctx, "/com.ibm.cloudland.sca.dbs.DBAdmin/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBAdminClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := c.cc.Invoke(ctx, "/com.ibm.cloudland.sca.dbs.DBAdmin/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBAdminServer is the server API for DBAdmin service.
type DBAdminServer interface {
	Stats(context.Context, *StatsRequest) (*StatsReply, error)
	Tables(context.Context, *TablesRequest) (*TablesReply, error)
	Exec(context.Context, *ExecRequest) (*ExecReply, error)
	Query(context.Context, *QueryRequest) (*QueryReply, error)
}

// UnimplementedDBAdminServer can be embedded to have forward compatible implementations.
type UnimplementedDBAdminServer struct {
}

func (*UnimplementedDBAdminServer) Stats(ctx context.Context, req *StatsRequest) (*StatsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (*UnimplementedDBAdminServer) Tables(ctx context.Context, req *TablesRequest) (*TablesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tables not implemented")
}
func (*UnimplementedDBAdminServer) Exec(ctx context.Context, req *ExecRequest) (*ExecReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (*UnimplementedDBAdminServer) Query(ctx context.Context, req *QueryRequest) (*QueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterDBAdminServer(s *grpc.Server, srv DBAdminServer) {
	s.RegisterService(&_DBAdmin_serviceDesc, srv)
}

func _DBAdmin_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBAdminServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ibm.cloudland.sca.dbs.DBAdmin/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBAdminServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBAdmin_Tables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBAdminServer).Tables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ibm.cloudland.sca.dbs.DBAdmin/Tables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBAdminServer).Tables(ctx, req.(*TablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBAdmin_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBAdminServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ibm.cloudland.sca.dbs.DBAdmin/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBAdminServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBAdmin_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBAdminServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ibm.cloudland.sca.dbs.DBAdmin/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBAdminServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.ibm.cloudland.sca.dbs.DBAdmin",
	HandlerType: (*DBAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stats",
			Handler:    _DBAdmin_Stats_Handler,
		},
		{
			MethodName: "Tables",
			Handler:    _DBAdmin_Tables_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _DBAdmin_Exec_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _DBAdmin_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbs.proto",
}
