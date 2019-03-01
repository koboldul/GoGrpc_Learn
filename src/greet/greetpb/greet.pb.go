// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greetpb/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName           string   `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetSecondName() string {
	if m != nil {
		return m.SecondName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManyTimesRequest) Reset()         { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()    {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{3}
}

func (m *GreetManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesRequest.Unmarshal(m, b)
}
func (m *GreetManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesRequest.Merge(m, src)
}
func (m *GreetManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesRequest.Size(m)
}
func (m *GreetManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesRequest proto.InternalMessageInfo

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManyTimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetManyTimesResponse) Reset()         { *m = GreetManyTimesResponse{} }
func (m *GreetManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesResponse) ProtoMessage()    {}
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{4}
}

func (m *GreetManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesResponse.Unmarshal(m, b)
}
func (m *GreetManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesResponse.Merge(m, src)
}
func (m *GreetManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesResponse.Size(m)
}
func (m *GreetManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesResponse proto.InternalMessageInfo

func (m *GreetManyTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongGreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LongGreetRequest) Reset()         { *m = LongGreetRequest{} }
func (m *LongGreetRequest) String() string { return proto.CompactTextString(m) }
func (*LongGreetRequest) ProtoMessage()    {}
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{5}
}

func (m *LongGreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetRequest.Unmarshal(m, b)
}
func (m *LongGreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetRequest.Marshal(b, m, deterministic)
}
func (m *LongGreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetRequest.Merge(m, src)
}
func (m *LongGreetRequest) XXX_Size() int {
	return xxx_messageInfo_LongGreetRequest.Size(m)
}
func (m *LongGreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetRequest proto.InternalMessageInfo

func (m *LongGreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongGreetResponse) Reset()         { *m = LongGreetResponse{} }
func (m *LongGreetResponse) String() string { return proto.CompactTextString(m) }
func (*LongGreetResponse) ProtoMessage()    {}
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{6}
}

func (m *LongGreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetResponse.Unmarshal(m, b)
}
func (m *LongGreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetResponse.Marshal(b, m, deterministic)
}
func (m *LongGreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetResponse.Merge(m, src)
}
func (m *LongGreetResponse) XXX_Size() int {
	return xxx_messageInfo_LongGreetResponse.Size(m)
}
func (m *LongGreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetResponse proto.InternalMessageInfo

func (m *LongGreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetAllRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetAllRequest) Reset()         { *m = GreetAllRequest{} }
func (m *GreetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GreetAllRequest) ProtoMessage()    {}
func (*GreetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{7}
}

func (m *GreetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetAllRequest.Unmarshal(m, b)
}
func (m *GreetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetAllRequest.Marshal(b, m, deterministic)
}
func (m *GreetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetAllRequest.Merge(m, src)
}
func (m *GreetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GreetAllRequest.Size(m)
}
func (m *GreetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetAllRequest proto.InternalMessageInfo

func (m *GreetAllRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetAllResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetAllResponse) Reset()         { *m = GreetAllResponse{} }
func (m *GreetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GreetAllResponse) ProtoMessage()    {}
func (*GreetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{8}
}

func (m *GreetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetAllResponse.Unmarshal(m, b)
}
func (m *GreetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetAllResponse.Marshal(b, m, deterministic)
}
func (m *GreetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetAllResponse.Merge(m, src)
}
func (m *GreetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GreetAllResponse.Size(m)
}
func (m *GreetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetAllResponse proto.InternalMessageInfo

func (m *GreetAllResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetWithDeadlineRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetWithDeadlineRequest) Reset()         { *m = GreetWithDeadlineRequest{} }
func (m *GreetWithDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineRequest) ProtoMessage()    {}
func (*GreetWithDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{9}
}

func (m *GreetWithDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineRequest.Unmarshal(m, b)
}
func (m *GreetWithDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineRequest.Merge(m, src)
}
func (m *GreetWithDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineRequest.Size(m)
}
func (m *GreetWithDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineRequest proto.InternalMessageInfo

func (m *GreetWithDeadlineRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetWithDeadlineResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetWithDeadlineResponse) Reset()         { *m = GreetWithDeadlineResponse{} }
func (m *GreetWithDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineResponse) ProtoMessage()    {}
func (*GreetWithDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd67c47c0cf51822, []int{10}
}

func (m *GreetWithDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineResponse.Unmarshal(m, b)
}
func (m *GreetWithDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineResponse.Merge(m, src)
}
func (m *GreetWithDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineResponse.Size(m)
}
func (m *GreetWithDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineResponse proto.InternalMessageInfo

func (m *GreetWithDeadlineResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreetManyTimesResponse)(nil), "greet.GreetManyTimesResponse")
	proto.RegisterType((*LongGreetRequest)(nil), "greet.LongGreetRequest")
	proto.RegisterType((*LongGreetResponse)(nil), "greet.LongGreetResponse")
	proto.RegisterType((*GreetAllRequest)(nil), "greet.GreetAllRequest")
	proto.RegisterType((*GreetAllResponse)(nil), "greet.GreetAllResponse")
	proto.RegisterType((*GreetWithDeadlineRequest)(nil), "greet.GreetWithDeadlineRequest")
	proto.RegisterType((*GreetWithDeadlineResponse)(nil), "greet.GreetWithDeadlineResponse")
}

func init() { proto.RegisterFile("greetpb/greet.proto", fileDescriptor_cd67c47c0cf51822) }

var fileDescriptor_cd67c47c0cf51822 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xef, 0x4b, 0xfa, 0x50,
	0x14, 0xc6, 0xbf, 0x13, 0xf4, 0xeb, 0x8e, 0x95, 0x7a, 0x2c, 0xb5, 0x91, 0x28, 0x7b, 0x93, 0x24,
	0x98, 0x68, 0xef, 0x82, 0x42, 0x11, 0x84, 0xa8, 0x20, 0x0b, 0x8a, 0xde, 0xc4, 0xd4, 0xd3, 0x1a,
	0xcc, 0x3b, 0xdb, 0x9d, 0x41, 0xff, 0x79, 0x2f, 0xc3, 0xbb, 0xab, 0xce, 0x5f, 0x0d, 0xf6, 0x4a,
	0xef, 0x39, 0xcf, 0x79, 0x3e, 0xe7, 0xc2, 0xb3, 0x0b, 0x39, 0xd3, 0x25, 0xf2, 0x26, 0x83, 0x73,
	0xf1, 0x5b, 0x9f, 0xb8, 0x8e, 0xe7, 0x60, 0x5c, 0x1c, 0xf4, 0x1b, 0x48, 0xf6, 0x66, 0x7f, 0x2c,
	0x66, 0x62, 0x09, 0xe0, 0xdd, 0x72, 0xb9, 0xf7, 0xc6, 0x8c, 0x31, 0x15, 0x95, 0x8a, 0x52, 0x55,
	0xfb, 0xaa, 0xa8, 0xdc, 0x1b, 0x63, 0xc2, 0x32, 0xa4, 0x38, 0x0d, 0x1d, 0x36, 0xf2, 0xfb, 0x31,
	0xd1, 0x07, 0xbf, 0x34, 0x13, 0xe8, 0x97, 0xb0, 0x27, 0xbc, 0xfa, 0xf4, 0x39, 0x25, 0xee, 0x61,
	0x0d, 0x92, 0xa6, 0xf4, 0x16, 0x6e, 0xa9, 0x66, 0xba, 0xee, 0xaf, 0x30, 0x47, 0xf6, 0x17, 0x02,
	0xfd, 0x14, 0xf6, 0xe5, 0x30, 0x9f, 0x38, 0x8c, 0x13, 0xe6, 0x21, 0xe1, 0x12, 0x9f, 0xda, 0x9e,
	0xdc, 0x44, 0x9e, 0xf4, 0x2e, 0x1c, 0x09, 0xe1, 0x9d, 0xc1, 0xbe, 0x9f, 0xac, 0x31, 0xf1, 0x48,
	0xb8, 0x06, 0xe4, 0xd7, 0x5d, 0x42, 0xb8, 0xd7, 0x90, 0xb9, 0x75, 0x98, 0x19, 0xfd, 0x86, 0x35,
	0xc8, 0x06, 0x0c, 0x42, 0x68, 0x57, 0x90, 0x16, 0xc2, 0xb6, 0x6d, 0x47, 0x82, 0x9d, 0x41, 0x66,
	0x39, 0x1f, 0xc2, 0xea, 0x41, 0x51, 0x68, 0x9f, 0x2d, 0xef, 0xa3, 0x4b, 0xc6, 0xc8, 0xb6, 0x18,
	0x45, 0x82, 0xb6, 0xe0, 0x78, 0x8b, 0xd1, 0xdf, 0xf4, 0xe6, 0x4f, 0x4c, 0xc6, 0xe6, 0x91, 0xdc,
	0x2f, 0x6b, 0x48, 0x78, 0x01, 0x71, 0x71, 0xc6, 0x5c, 0x90, 0x24, 0x17, 0xd2, 0x0e, 0x57, 0x8b,
	0xbe, 0xb9, 0xfe, 0x0f, 0x5f, 0x20, 0xbb, 0xc1, 0xc6, 0x72, 0x50, 0xbc, 0xe5, 0x7a, 0x5a, 0x65,
	0xb7, 0x60, 0xe1, 0xfc, 0x00, 0x07, 0xab, 0x51, 0xc1, 0x93, 0xe0, 0xd4, 0x7a, 0x0e, 0xb5, 0xd2,
	0x8e, 0xee, 0xdc, 0xb0, 0xa1, 0x60, 0x07, 0xd4, 0x45, 0x14, 0xb0, 0x20, 0xf5, 0xeb, 0xe9, 0xd2,
	0x8a, 0x9b, 0x8d, 0xb9, 0x47, 0x55, 0xc1, 0xb6, 0xfc, 0x72, 0xdb, 0xb6, 0x8d, 0xf9, 0x20, 0x72,
	0x19, 0x19, 0xad, 0xb0, 0x51, 0x5f, 0x1a, 0x34, 0x94, 0x8e, 0xfa, 0xfa, 0x5f, 0x3e, 0x0d, 0x83,
	0x84, 0x78, 0x15, 0x5a, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x65, 0xa2, 0x80, 0x2c, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	//Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error)
	//Server Streaming
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	//Client Streaming
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
	//Client Streaming
	GreetAll(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetAllClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error) {
	out := new(GreetWithDeadlineResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/GreetWithDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetManyTimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetManyTimesResponse, error) {
	m := new(GreetManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetAll(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[2], "/greet.GreetService/GreetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetAllClient{stream}
	return x, nil
}

type GreetService_GreetAllClient interface {
	Send(*GreetAllRequest) error
	Recv() (*GreetAllResponse, error)
	grpc.ClientStream
}

type greetServiceGreetAllClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetAllClient) Send(m *GreetAllRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetAllClient) Recv() (*GreetAllResponse, error) {
	m := new(GreetAllResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	//Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetWithDeadline(context.Context, *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error)
	//Server Streaming
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
	//Client Streaming
	LongGreet(GreetService_LongGreetServer) error
	//Client Streaming
	GreetAll(GreetService_GreetAllServer) error
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/GreetWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, req.(*GreetWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetManyTimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetAll(&greetServiceGreetAllServer{stream})
}

type GreetService_GreetAllServer interface {
	Send(*GreetAllResponse) error
	Recv() (*GreetAllRequest, error)
	grpc.ServerStream
}

type greetServiceGreetAllServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetAllServer) Send(m *GreetAllResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetAllServer) Recv() (*GreetAllRequest, error) {
	m := new(GreetAllRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
		{
			MethodName: "GreetWithDeadline",
			Handler:    _GreetService_GreetWithDeadline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetAll",
			Handler:       _GreetService_GreetAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greetpb/greet.proto",
}