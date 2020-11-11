// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/feeds.proto

package feeds

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Feed struct {
	// rss feed name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// rss feed address
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feed) Reset()         { *m = Feed{} }
func (m *Feed) String() string { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()    {}
func (*Feed) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd517c38176c13bf, []int{0}
}

func (m *Feed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feed.Unmarshal(m, b)
}
func (m *Feed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feed.Marshal(b, m, deterministic)
}
func (m *Feed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed.Merge(m, src)
}
func (m *Feed) XXX_Size() int {
	return xxx_messageInfo_Feed.Size(m)
}
func (m *Feed) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed.DiscardUnknown(m)
}

var xxx_messageInfo_Feed proto.InternalMessageInfo

func (m *Feed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feed) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Entry struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Date                 int64    `protobuf:"varint,6,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd517c38176c13bf, []int{1}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entry) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Entry) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Entry) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Entry) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Entry) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type NewRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd517c38176c13bf, []int{2}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type NewResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResponse) Reset()         { *m = NewResponse{} }
func (m *NewResponse) String() string { return proto.CompactTextString(m) }
func (*NewResponse) ProtoMessage()    {}
func (*NewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd517c38176c13bf, []int{3}
}

func (m *NewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResponse.Unmarshal(m, b)
}
func (m *NewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResponse.Marshal(b, m, deterministic)
}
func (m *NewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResponse.Merge(m, src)
}
func (m *NewResponse) XXX_Size() int {
	return xxx_messageInfo_NewResponse.Size(m)
}
func (m *NewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Feed)(nil), "feeds.Feed")
	proto.RegisterType((*Entry)(nil), "feeds.Entry")
	proto.RegisterType((*NewRequest)(nil), "feeds.NewRequest")
	proto.RegisterType((*NewResponse)(nil), "feeds.NewResponse")
}

func init() {
	proto.RegisterFile("proto/feeds.proto", fileDescriptor_dd517c38176c13bf)
}

var fileDescriptor_dd517c38176c13bf = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xdd, 0xdb, 0xcd, 0x8a, 0x73, 0x9c, 0x78, 0x83, 0x48, 0xb0, 0x92, 0xad, 0xac, 0x56,
	0x50, 0x41, 0xd0, 0x4e, 0xd0, 0xd2, 0x62, 0x4b, 0xbb, 0xd5, 0x8c, 0xb0, 0x70, 0x97, 0x9c, 0xc9,
	0x2c, 0xe2, 0x03, 0xf8, 0xde, 0x26, 0x93, 0x88, 0xb6, 0x76, 0xff, 0xf7, 0x85, 0x3f, 0x93, 0x0c,
	0xac, 0x77, 0xde, 0xb1, 0xbb, 0x78, 0x23, 0x32, 0xa1, 0x97, 0x8c, 0x4a, 0xa0, 0xbb, 0x86, 0xe6,
	0x31, 0x06, 0x44, 0x68, 0xec, 0xb8, 0x25, 0x5d, 0x9d, 0x55, 0xe7, 0x07, 0x83, 0x64, 0xd4, 0xb0,
	0x3f, 0x1a, 0xe3, 0x29, 0x04, 0xbd, 0x10, 0xfd, 0x83, 0xdd, 0x57, 0x05, 0xea, 0xc1, 0xb2, 0xff,
	0xc4, 0x43, 0x58, 0x4c, 0xa6, 0xb4, 0x62, 0xc2, 0x13, 0x68, 0x8d, 0xdb, 0x8e, 0x93, 0x2d, 0x95,
	0x42, 0x78, 0x04, 0xf5, 0xec, 0x37, 0xba, 0x16, 0x99, 0x22, 0x1e, 0x83, 0xe2, 0x89, 0x37, 0xa4,
	0x1b, 0x71, 0x19, 0xd2, 0xcc, 0x57, 0x67, 0x99, 0x2c, 0x6b, 0x95, 0x67, 0x16, 0x4c, 0x2f, 0x34,
	0x23, 0x93, 0x6e, 0xa3, 0xae, 0x07, 0xc9, 0xdd, 0x2d, 0xc0, 0x13, 0x7d, 0x0c, 0xf4, 0x3e, 0x53,
	0xe0, 0x7f, 0xfe, 0x61, 0x05, 0x4b, 0xe9, 0x86, 0x9d, 0xb3, 0x81, 0x2e, 0x6f, 0x40, 0xa5, 0x45,
	0x04, 0xec, 0xa1, 0x8e, 0x1e, 0xd7, 0x7d, 0xde, 0xd6, 0xef, 0xfd, 0xa7, 0xf8, 0x57, 0xe5, 0x5a,
	0xb7, 0x77, 0xbf, 0x7a, 0x5e, 0xca, 0x46, 0xef, 0xe4, 0xf0, 0xa5, 0x15, 0xb8, 0xfa, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x04, 0x4b, 0x80, 0xda, 0x73, 0x01, 0x00, 0x00,
}
