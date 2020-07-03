// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drand/protocol.proto

package drand

import (
	fmt "fmt"
	dkg "github.com/drand/drand/protobuf/crypto/dkg"
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

type IdentityRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityRequest) Reset()         { *m = IdentityRequest{} }
func (m *IdentityRequest) String() string { return proto.CompactTextString(m) }
func (*IdentityRequest) ProtoMessage()    {}
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{0}
}

func (m *IdentityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityRequest.Unmarshal(m, b)
}
func (m *IdentityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityRequest.Marshal(b, m, deterministic)
}
func (m *IdentityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityRequest.Merge(m, src)
}
func (m *IdentityRequest) XXX_Size() int {
	return xxx_messageInfo_IdentityRequest.Size(m)
}
func (m *IdentityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityRequest proto.InternalMessageInfo

// SignalDKGPacket is the packet nodes send to a coordinator that collects all
// keys and setups the group and sends them back to the nodes such that they can
// start the DKG automatically.
type SignalDKGPacket struct {
	Node        *Identity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	SecretProof []byte    `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// In resharing cases, previous_group_hash is the hash of the previous group.
	// It is to make sure the nodes build on top of the correct previous group.
	PreviousGroupHash    []byte   `protobuf:"bytes,3,opt,name=previous_group_hash,json=previousGroupHash,proto3" json:"previous_group_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalDKGPacket) Reset()         { *m = SignalDKGPacket{} }
func (m *SignalDKGPacket) String() string { return proto.CompactTextString(m) }
func (*SignalDKGPacket) ProtoMessage()    {}
func (*SignalDKGPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{1}
}

func (m *SignalDKGPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalDKGPacket.Unmarshal(m, b)
}
func (m *SignalDKGPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalDKGPacket.Marshal(b, m, deterministic)
}
func (m *SignalDKGPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalDKGPacket.Merge(m, src)
}
func (m *SignalDKGPacket) XXX_Size() int {
	return xxx_messageInfo_SignalDKGPacket.Size(m)
}
func (m *SignalDKGPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalDKGPacket.DiscardUnknown(m)
}

var xxx_messageInfo_SignalDKGPacket proto.InternalMessageInfo

func (m *SignalDKGPacket) GetNode() *Identity {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *SignalDKGPacket) GetSecretProof() []byte {
	if m != nil {
		return m.SecretProof
	}
	return nil
}

func (m *SignalDKGPacket) GetPreviousGroupHash() []byte {
	if m != nil {
		return m.PreviousGroupHash
	}
	return nil
}

// PushDKGInfor is the packet the coordinator sends that contains the group over
// which to run the DKG on, the secret proof (to prove it's he's part of the
// expected group, and it's not a random packet) and as well the time at which
// every node should start the DKG.
type DKGInfoPacket struct {
	NewGroup    *GroupPacket `protobuf:"bytes,1,opt,name=new_group,json=newGroup,proto3" json:"new_group,omitempty"`
	SecretProof []byte       `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// timeout in seconds
	DkgTimeout uint32 `protobuf:"varint,3,opt,name=dkg_timeout,json=dkgTimeout,proto3" json:"dkg_timeout,omitempty"`
	// signature from the coordinator to prove he is the one sending that group
	// file.
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DKGInfoPacket) Reset()         { *m = DKGInfoPacket{} }
func (m *DKGInfoPacket) String() string { return proto.CompactTextString(m) }
func (*DKGInfoPacket) ProtoMessage()    {}
func (*DKGInfoPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{2}
}

func (m *DKGInfoPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DKGInfoPacket.Unmarshal(m, b)
}
func (m *DKGInfoPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DKGInfoPacket.Marshal(b, m, deterministic)
}
func (m *DKGInfoPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DKGInfoPacket.Merge(m, src)
}
func (m *DKGInfoPacket) XXX_Size() int {
	return xxx_messageInfo_DKGInfoPacket.Size(m)
}
func (m *DKGInfoPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_DKGInfoPacket.DiscardUnknown(m)
}

var xxx_messageInfo_DKGInfoPacket proto.InternalMessageInfo

func (m *DKGInfoPacket) GetNewGroup() *GroupPacket {
	if m != nil {
		return m.NewGroup
	}
	return nil
}

func (m *DKGInfoPacket) GetSecretProof() []byte {
	if m != nil {
		return m.SecretProof
	}
	return nil
}

func (m *DKGInfoPacket) GetDkgTimeout() uint32 {
	if m != nil {
		return m.DkgTimeout
	}
	return 0
}

func (m *DKGInfoPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type PartialBeaconPacket struct {
	// Round is the round for which the beacon will be created from the partial
	// signatures
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	// signature of the previous round - could be removed at some point but now
	// is used to verify the signature even before accessing the store
	PreviousSig []byte `protobuf:"bytes,2,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	// partial signature - a threshold of them needs to be aggregated to produce
	// the final beacon at the given round.Beautiful
	PartialSig           []byte   `protobuf:"bytes,3,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartialBeaconPacket) Reset()         { *m = PartialBeaconPacket{} }
func (m *PartialBeaconPacket) String() string { return proto.CompactTextString(m) }
func (*PartialBeaconPacket) ProtoMessage()    {}
func (*PartialBeaconPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{3}
}

func (m *PartialBeaconPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialBeaconPacket.Unmarshal(m, b)
}
func (m *PartialBeaconPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialBeaconPacket.Marshal(b, m, deterministic)
}
func (m *PartialBeaconPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialBeaconPacket.Merge(m, src)
}
func (m *PartialBeaconPacket) XXX_Size() int {
	return xxx_messageInfo_PartialBeaconPacket.Size(m)
}
func (m *PartialBeaconPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialBeaconPacket.DiscardUnknown(m)
}

var xxx_messageInfo_PartialBeaconPacket proto.InternalMessageInfo

func (m *PartialBeaconPacket) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *PartialBeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *PartialBeaconPacket) GetPartialSig() []byte {
	if m != nil {
		return m.PartialSig
	}
	return nil
}

// BroadcastPacket is the packet that nodes send to others nodes as part of the
// broadcasting protocol.
type BroadcastPacket struct {
	Dkg *dkg.Packet `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	// issuer is the address of the node participating to the broadcasting
	// channel - this is to find the identity of the signer of this packet.
	Issuer string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// dealer indicates if the issuer is a dealer or not during the broadcast
	Dealer bool `protobuf:"varint,3,opt,name=dealer,proto3" json:"dealer,omitempty"`
	// signature over hash of the DKG packet
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastPacket) Reset()         { *m = BroadcastPacket{} }
func (m *BroadcastPacket) String() string { return proto.CompactTextString(m) }
func (*BroadcastPacket) ProtoMessage()    {}
func (*BroadcastPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{4}
}

func (m *BroadcastPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastPacket.Unmarshal(m, b)
}
func (m *BroadcastPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastPacket.Marshal(b, m, deterministic)
}
func (m *BroadcastPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastPacket.Merge(m, src)
}
func (m *BroadcastPacket) XXX_Size() int {
	return xxx_messageInfo_BroadcastPacket.Size(m)
}
func (m *BroadcastPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastPacket.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastPacket proto.InternalMessageInfo

func (m *BroadcastPacket) GetDkg() *dkg.Packet {
	if m != nil {
		return m.Dkg
	}
	return nil
}

func (m *BroadcastPacket) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *BroadcastPacket) GetDealer() bool {
	if m != nil {
		return m.Dealer
	}
	return false
}

func (m *BroadcastPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// SyncRequest is from a node that needs to sync up with the current head of the
// chain
type SyncRequest struct {
	FromRound            uint64   `protobuf:"varint,1,opt,name=from_round,json=fromRound,proto3" json:"from_round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{5}
}

func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetFromRound() uint64 {
	if m != nil {
		return m.FromRound
	}
	return 0
}

type BeaconPacket struct {
	PreviousSig          []byte   `protobuf:"bytes,1,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	Round                uint64   `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconPacket) Reset()         { *m = BeaconPacket{} }
func (m *BeaconPacket) String() string { return proto.CompactTextString(m) }
func (*BeaconPacket) ProtoMessage()    {}
func (*BeaconPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{6}
}

func (m *BeaconPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconPacket.Unmarshal(m, b)
}
func (m *BeaconPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconPacket.Marshal(b, m, deterministic)
}
func (m *BeaconPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconPacket.Merge(m, src)
}
func (m *BeaconPacket) XXX_Size() int {
	return xxx_messageInfo_BeaconPacket.Size(m)
}
func (m *BeaconPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconPacket.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconPacket proto.InternalMessageInfo

func (m *BeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *BeaconPacket) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *BeaconPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*IdentityRequest)(nil), "drand.IdentityRequest")
	proto.RegisterType((*SignalDKGPacket)(nil), "drand.SignalDKGPacket")
	proto.RegisterType((*DKGInfoPacket)(nil), "drand.DKGInfoPacket")
	proto.RegisterType((*PartialBeaconPacket)(nil), "drand.PartialBeaconPacket")
	proto.RegisterType((*BroadcastPacket)(nil), "drand.BroadcastPacket")
	proto.RegisterType((*SyncRequest)(nil), "drand.SyncRequest")
	proto.RegisterType((*BeaconPacket)(nil), "drand.BeaconPacket")
}

func init() {
	proto.RegisterFile("drand/protocol.proto", fileDescriptor_e344a98fea1e2f3a)
}

var fileDescriptor_e344a98fea1e2f3a = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x45, 0xf9, 0x22, 0x1a, 0xd9, 0x98, 0xac, 0x4d, 0x30, 0xa2, 0xa1, 0xa9, 0x4a, 0x69, 0x0e,
	0x45, 0x6e, 0x53, 0x08, 0x14, 0x7a, 0x72, 0x53, 0xdc, 0x90, 0x8b, 0x91, 0x7b, 0xea, 0xc5, 0xac,
	0xa5, 0xb5, 0xb4, 0xd8, 0xda, 0x55, 0x57, 0xab, 0x06, 0x5f, 0x7a, 0xef, 0xdf, 0xe8, 0xbf, 0xea,
	0xbf, 0x29, 0xfb, 0x21, 0x55, 0x56, 0x0a, 0xed, 0xc1, 0xe0, 0x79, 0xb3, 0x33, 0xf3, 0xf4, 0xe6,
	0xed, 0xc2, 0x28, 0x11, 0x98, 0x25, 0x93, 0x42, 0x70, 0xc9, 0x63, 0xbe, 0x0d, 0xf5, 0x1f, 0x74,
	0xac, 0x51, 0x7f, 0x14, 0x8b, 0x5d, 0x21, 0xf9, 0x24, 0xd9, 0xa4, 0xea, 0x67, 0x92, 0x3e, 0x32,
	0x25, 0x31, 0xcf, 0x73, 0xce, 0x0c, 0x16, 0x9c, 0xc1, 0xe0, 0x2e, 0x21, 0x4c, 0x52, 0xb9, 0x8b,
	0xc8, 0xd7, 0x8a, 0x94, 0x32, 0xf8, 0xe1, 0xc0, 0x60, 0x41, 0x53, 0x86, 0xb7, 0xb7, 0xf7, 0xb3,
	0x39, 0x8e, 0x37, 0x44, 0xa2, 0xe7, 0x70, 0xc4, 0x78, 0x42, 0xc6, 0xce, 0xa5, 0x73, 0xe5, 0x5d,
	0x0f, 0x42, 0xdd, 0x29, 0x6c, 0x2a, 0x75, 0x12, 0x3d, 0x83, 0x5e, 0x49, 0x62, 0x41, 0xe4, 0xb2,
	0x10, 0x9c, 0xaf, 0xc7, 0x07, 0x97, 0xce, 0x55, 0x2f, 0xf2, 0x0c, 0x36, 0x57, 0x10, 0x0a, 0x61,
	0x58, 0x08, 0xf2, 0x8d, 0xf2, 0xaa, 0x5c, 0xa6, 0x82, 0x57, 0xc5, 0x32, 0xc3, 0x65, 0x36, 0x3e,
	0xd4, 0x27, 0xcf, 0xea, 0xd4, 0x4c, 0x65, 0x3e, 0xe1, 0x32, 0x0b, 0x7e, 0x3a, 0xd0, 0xbf, 0xbd,
	0x9f, 0xdd, 0xb1, 0x35, 0xb7, 0x4c, 0x26, 0xe0, 0x32, 0xf2, 0x60, 0x8a, 0x2d, 0x1d, 0x64, 0xe9,
	0xe8, 0x32, 0x73, 0x2c, 0x3a, 0x65, 0xe4, 0x41, 0xc7, 0xff, 0xc3, 0xea, 0x29, 0x78, 0xc9, 0x26,
	0x5d, 0x4a, 0x9a, 0x13, 0x5e, 0x49, 0xcd, 0xa6, 0x1f, 0x41, 0xb2, 0x49, 0x3f, 0x1b, 0x04, 0x3d,
	0x01, 0xb7, 0x54, 0x8a, 0xc8, 0x4a, 0x90, 0xf1, 0x91, 0x6e, 0xf0, 0x07, 0x08, 0x38, 0x0c, 0xe7,
	0x58, 0x48, 0x8a, 0xb7, 0x53, 0x82, 0x63, 0xce, 0x2c, 0xd3, 0x11, 0x1c, 0x0b, 0x5e, 0xb1, 0x44,
	0xb3, 0x3c, 0x8a, 0x4c, 0xa0, 0xe8, 0x34, 0x0a, 0x94, 0x34, 0xad, 0xe9, 0xd4, 0xd8, 0x82, 0xa6,
	0x8a, 0x4e, 0x61, 0xfa, 0xe9, 0x13, 0x46, 0x1c, 0xb0, 0xd0, 0x82, 0xa6, 0xc1, 0x77, 0x18, 0x4c,
	0x05, 0xc7, 0x49, 0x8c, 0x4b, 0x69, 0x87, 0x5d, 0xc0, 0x61, 0xb2, 0x49, 0xad, 0x20, 0x5e, 0xa8,
	0x96, 0x6e, 0x95, 0x50, 0x38, 0x3a, 0x87, 0x13, 0x5a, 0x96, 0x15, 0x11, 0x7a, 0x9e, 0x1b, 0xd9,
	0x48, 0xe1, 0x09, 0xc1, 0x5b, 0x22, 0xf4, 0x94, 0xd3, 0xc8, 0x46, 0xff, 0xf8, 0xe0, 0x57, 0xe0,
	0x2d, 0x76, 0x2c, 0xb6, 0x86, 0x41, 0x17, 0x00, 0x6b, 0xc1, 0xf3, 0x65, 0xfb, 0x6b, 0x5d, 0x85,
	0x44, 0x0a, 0x08, 0x08, 0xf4, 0xf6, 0x74, 0xe9, 0x2a, 0xe0, 0x3c, 0x56, 0xa0, 0x91, 0xee, 0xa0,
	0x2d, 0xdd, 0x1e, 0xa9, 0xc3, 0x0e, 0xa9, 0xeb, 0x5f, 0x07, 0x70, 0x3a, 0xb7, 0xb7, 0x01, 0xdd,
	0x80, 0x37, 0x23, 0xb2, 0xf6, 0x27, 0x3a, 0xef, 0x1a, 0xd6, 0x30, 0xf7, 0xbb, 0x46, 0x46, 0xef,
	0x61, 0xd4, 0xb2, 0xbe, 0x90, 0x34, 0xa6, 0x05, 0x66, 0xb2, 0x69, 0xd0, 0xb9, 0x17, 0x7e, 0xcf,
	0xe2, 0x1f, 0xf3, 0x42, 0xee, 0xd0, 0x1b, 0xf0, 0xe6, 0x55, 0x99, 0x59, 0xc3, 0xa2, 0x91, 0x4d,
	0xee, 0x19, 0xf8, 0x51, 0x89, 0xdb, 0xac, 0xb2, 0x99, 0xd2, 0x59, 0x6e, 0xa7, 0xe4, 0x1d, 0xf4,
	0xf7, 0xec, 0x86, 0x7c, 0x9b, 0xfe, 0x8b, 0x09, 0x3b, 0xa5, 0x37, 0xe0, 0xaa, 0xc5, 0x7d, 0xc8,
	0x30, 0x65, 0xa8, 0xbe, 0x36, 0xad, 0x55, 0xfa, 0xc3, 0x9a, 0x41, 0xab, 0xc7, 0x6b, 0x67, 0xfa,
	0xf2, 0xcb, 0x8b, 0x94, 0xca, 0xac, 0x5a, 0x85, 0x31, 0xcf, 0x27, 0xe6, 0x19, 0x69, 0xbd, 0x3f,
	0xab, 0x6a, 0x6d, 0xc2, 0xd5, 0x89, 0x8e, 0xdf, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x70, 0xd6,
	0xa7, 0xe6, 0x9e, 0x04, 0x00, 0x00,
}
