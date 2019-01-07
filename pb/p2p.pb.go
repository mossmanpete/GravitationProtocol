// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: p2p.proto

package protocols_p2p

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// designed to be shared between all app protocols
type MessageData struct {
	// shared between all requests
	ClientVersion        string   `protobuf:"bytes,1,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Gossip               bool     `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	NodePubKey           []byte   `protobuf:"bytes,6,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Sign                 []byte   `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageData) Reset()         { *m = MessageData{} }
func (m *MessageData) String() string { return proto.CompactTextString(m) }
func (*MessageData) ProtoMessage()    {}
func (*MessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{0}
}
func (m *MessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageData.Unmarshal(m, b)
}
func (m *MessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageData.Marshal(b, m, deterministic)
}
func (m *MessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageData.Merge(m, src)
}
func (m *MessageData) XXX_Size() int {
	return xxx_messageInfo_MessageData.Size(m)
}
func (m *MessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageData.DiscardUnknown(m)
}

var xxx_messageInfo_MessageData proto.InternalMessageInfo

func (m *MessageData) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *MessageData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MessageData) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *MessageData) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MessageData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *MessageData) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

// a protocol define a set of reuqest and responses
type GravitationRequest struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	Profile              []string                       `protobuf:"bytes,2,rep,name=profile,proto3" json:"profile,omitempty"`
	SubOrbit             []*GravitationRequest_SubOrbit `protobuf:"bytes,3,rep,name=sub_orbit,json=subOrbit,proto3" json:"sub_orbit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GravitationRequest) Reset()         { *m = GravitationRequest{} }
func (m *GravitationRequest) String() string { return proto.CompactTextString(m) }
func (*GravitationRequest) ProtoMessage()    {}
func (*GravitationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{1}
}
func (m *GravitationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GravitationRequest.Unmarshal(m, b)
}
func (m *GravitationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GravitationRequest.Marshal(b, m, deterministic)
}
func (m *GravitationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GravitationRequest.Merge(m, src)
}
func (m *GravitationRequest) XXX_Size() int {
	return xxx_messageInfo_GravitationRequest.Size(m)
}
func (m *GravitationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GravitationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GravitationRequest proto.InternalMessageInfo

func (m *GravitationRequest) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *GravitationRequest) GetProfile() []string {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *GravitationRequest) GetSubOrbit() []*GravitationRequest_SubOrbit {
	if m != nil {
		return m.SubOrbit
	}
	return nil
}

type GravitationRequest_SubOrbit struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Profile              []string `protobuf:"bytes,2,rep,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GravitationRequest_SubOrbit) Reset()         { *m = GravitationRequest_SubOrbit{} }
func (m *GravitationRequest_SubOrbit) String() string { return proto.CompactTextString(m) }
func (*GravitationRequest_SubOrbit) ProtoMessage()    {}
func (*GravitationRequest_SubOrbit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{1, 0}
}
func (m *GravitationRequest_SubOrbit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GravitationRequest_SubOrbit.Unmarshal(m, b)
}
func (m *GravitationRequest_SubOrbit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GravitationRequest_SubOrbit.Marshal(b, m, deterministic)
}
func (m *GravitationRequest_SubOrbit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GravitationRequest_SubOrbit.Merge(m, src)
}
func (m *GravitationRequest_SubOrbit) XXX_Size() int {
	return xxx_messageInfo_GravitationRequest_SubOrbit.Size(m)
}
func (m *GravitationRequest_SubOrbit) XXX_DiscardUnknown() {
	xxx_messageInfo_GravitationRequest_SubOrbit.DiscardUnknown(m)
}

var xxx_messageInfo_GravitationRequest_SubOrbit proto.InternalMessageInfo

func (m *GravitationRequest_SubOrbit) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *GravitationRequest_SubOrbit) GetProfile() []string {
	if m != nil {
		return m.Profile
	}
	return nil
}

type GravitationResponse struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	// method specific data
	Profile              []string                        `protobuf:"bytes,2,rep,name=profile,proto3" json:"profile,omitempty"`
	SubOrbit             []*GravitationResponse_SubOrbit `protobuf:"bytes,3,rep,name=sub_orbit,json=subOrbit,proto3" json:"sub_orbit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GravitationResponse) Reset()         { *m = GravitationResponse{} }
func (m *GravitationResponse) String() string { return proto.CompactTextString(m) }
func (*GravitationResponse) ProtoMessage()    {}
func (*GravitationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{2}
}
func (m *GravitationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GravitationResponse.Unmarshal(m, b)
}
func (m *GravitationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GravitationResponse.Marshal(b, m, deterministic)
}
func (m *GravitationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GravitationResponse.Merge(m, src)
}
func (m *GravitationResponse) XXX_Size() int {
	return xxx_messageInfo_GravitationResponse.Size(m)
}
func (m *GravitationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GravitationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GravitationResponse proto.InternalMessageInfo

func (m *GravitationResponse) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *GravitationResponse) GetProfile() []string {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *GravitationResponse) GetSubOrbit() []*GravitationResponse_SubOrbit {
	if m != nil {
		return m.SubOrbit
	}
	return nil
}

type GravitationResponse_SubOrbit struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Profile              []string `protobuf:"bytes,2,rep,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GravitationResponse_SubOrbit) Reset()         { *m = GravitationResponse_SubOrbit{} }
func (m *GravitationResponse_SubOrbit) String() string { return proto.CompactTextString(m) }
func (*GravitationResponse_SubOrbit) ProtoMessage()    {}
func (*GravitationResponse_SubOrbit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{2, 0}
}
func (m *GravitationResponse_SubOrbit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GravitationResponse_SubOrbit.Unmarshal(m, b)
}
func (m *GravitationResponse_SubOrbit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GravitationResponse_SubOrbit.Marshal(b, m, deterministic)
}
func (m *GravitationResponse_SubOrbit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GravitationResponse_SubOrbit.Merge(m, src)
}
func (m *GravitationResponse_SubOrbit) XXX_Size() int {
	return xxx_messageInfo_GravitationResponse_SubOrbit.Size(m)
}
func (m *GravitationResponse_SubOrbit) XXX_DiscardUnknown() {
	xxx_messageInfo_GravitationResponse_SubOrbit.DiscardUnknown(m)
}

var xxx_messageInfo_GravitationResponse_SubOrbit proto.InternalMessageInfo

func (m *GravitationResponse_SubOrbit) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *GravitationResponse_SubOrbit) GetProfile() []string {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageData)(nil), "protocols.p2p.MessageData")
	proto.RegisterType((*GravitationRequest)(nil), "protocols.p2p.GravitationRequest")
	proto.RegisterType((*GravitationRequest_SubOrbit)(nil), "protocols.p2p.GravitationRequest.SubOrbit")
	proto.RegisterType((*GravitationResponse)(nil), "protocols.p2p.GravitationResponse")
	proto.RegisterType((*GravitationResponse_SubOrbit)(nil), "protocols.p2p.GravitationResponse.SubOrbit")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor_e7fdddb109e6467a) }

var fileDescriptor_e7fdddb109e6467a = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0xa4, 0xa6, 0xcd, 0xc4, 0x7a, 0x58, 0x41, 0x97, 0x22, 0x12, 0x8a, 0x87, 0xa0,
	0x90, 0x43, 0xbc, 0xea, 0x4d, 0xa8, 0x45, 0x44, 0x59, 0xc1, 0x6b, 0x49, 0x9a, 0xb1, 0x2c, 0xb4,
	0xd9, 0x35, 0xb3, 0x11, 0xfc, 0x81, 0xfe, 0x23, 0xf1, 0x2c, 0x89, 0x29, 0x6d, 0x2d, 0x7a, 0xd2,
	0x53, 0xe6, 0xbd, 0xcc, 0xec, 0xbc, 0x6f, 0xc0, 0x37, 0x89, 0x89, 0x4d, 0xa9, 0xad, 0xe6, 0xfd,
	0xe6, 0x33, 0xd5, 0x73, 0x8a, 0x4d, 0x62, 0x86, 0x6f, 0x0c, 0x82, 0x5b, 0x24, 0x4a, 0x67, 0x78,
	0x95, 0xda, 0x94, 0x9f, 0x40, 0x7f, 0x3a, 0x57, 0x58, 0xd8, 0x47, 0x2c, 0x49, 0xe9, 0x42, 0xb0,
	0x90, 0x45, 0xbe, 0xdc, 0x34, 0xf9, 0x11, 0xf8, 0x56, 0x2d, 0x90, 0x6c, 0xba, 0x30, 0xc2, 0x09,
	0x59, 0xe4, 0xca, 0x95, 0xc1, 0xf7, 0xc0, 0x51, 0xb9, 0x70, 0x9b, 0x41, 0x47, 0xe5, 0xfc, 0x00,
	0xbc, 0x99, 0x26, 0x52, 0x46, 0x74, 0x42, 0x16, 0xf5, 0x64, 0xab, 0x6a, 0xbf, 0xd0, 0x39, 0x8e,
	0x73, 0xb1, 0xd3, 0xf4, 0xb6, 0x8a, 0x1f, 0x03, 0xd4, 0xd5, 0x7d, 0x95, 0xdd, 0xe0, 0xab, 0xf0,
	0x42, 0x16, 0xed, 0xca, 0x35, 0x87, 0x73, 0xe8, 0x90, 0x9a, 0x15, 0xa2, 0xdb, 0xfc, 0x69, 0xea,
	0xe1, 0x3b, 0x03, 0x3e, 0x2a, 0xd3, 0x17, 0x65, 0x53, 0xab, 0x74, 0x21, 0xf1, 0xb9, 0x42, 0xb2,
	0xfc, 0x02, 0x82, 0xc5, 0x8a, 0xae, 0x81, 0x09, 0x92, 0x41, 0xbc, 0x71, 0x83, 0x78, 0x8d, 0x5f,
	0xae, 0xb7, 0x73, 0x01, 0x5d, 0x53, 0xea, 0x27, 0x35, 0x47, 0xe1, 0x84, 0x6e, 0xe4, 0xcb, 0xa5,
	0xe4, 0x23, 0xf0, 0xa9, 0xca, 0x26, 0xba, 0xcc, 0x94, 0x15, 0x6e, 0xe8, 0x46, 0x41, 0x72, 0xfa,
	0xed, 0xd5, 0xed, 0x34, 0xf1, 0x43, 0x95, 0xdd, 0xd5, 0x13, 0xb2, 0x47, 0x6d, 0x35, 0xb8, 0x84,
	0xde, 0xd2, 0xe5, 0x87, 0xd0, 0x35, 0x88, 0xe5, 0x44, 0xe5, 0xed, 0xd5, 0xbd, 0x5a, 0x8e, 0xf3,
	0x9f, 0x73, 0x0c, 0x3f, 0x18, 0xec, 0x6f, 0x2c, 0x22, 0xa3, 0x0b, 0xc2, 0x7f, 0xe3, 0xbe, 0xde,
	0xe6, 0x3e, 0xfb, 0x8d, 0xfb, 0x2b, 0xce, 0xdf, 0x83, 0x67, 0x5e, 0xb3, 0xf4, 0xfc, 0x33, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0x90, 0x12, 0x2e, 0xda, 0x02, 0x00, 0x00,
}
