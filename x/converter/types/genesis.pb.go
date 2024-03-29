// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: erc721_bridge/converter/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// class_collections is a slice of the registered token pairs at genesis
	ClassCollections []ClassCollection `protobuf:"bytes,1,rep,name=class_collections,json=classCollections,proto3" json:"class_collections"`
	// class_traces is a slice of the transfered on `erc-721` port token by nft-transfer
	ClassTraces []ClassTrace `protobuf:"bytes,2,rep,name=class_traces,json=classTraces,proto3" json:"class_traces"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f31bd2c96fa73a7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetClassCollections() []ClassCollection {
	if m != nil {
		return m.ClassCollections
	}
	return nil
}

func (m *GenesisState) GetClassTraces() []ClassTrace {
	if m != nil {
		return m.ClassTraces
	}
	return nil
}

// ClassTrace defines an instance that records a pairing consisting of a native
//  class and an ERC721 contract.
type ClassTrace struct {
	// class_id is the ibc class id
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// contract is the address of the erc721 contract
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// tokens is a array of the token in class or the erc721 contract
	Tokens []TokenTrace `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens"`
}

func (m *ClassTrace) Reset()         { *m = ClassTrace{} }
func (m *ClassTrace) String() string { return proto.CompactTextString(m) }
func (*ClassTrace) ProtoMessage()    {}
func (*ClassTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f31bd2c96fa73a7, []int{1}
}
func (m *ClassTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassTrace.Merge(m, src)
}
func (m *ClassTrace) XXX_Size() int {
	return m.Size()
}
func (m *ClassTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassTrace.DiscardUnknown(m)
}

var xxx_messageInfo_ClassTrace proto.InternalMessageInfo

func (m *ClassTrace) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *ClassTrace) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ClassTrace) GetTokens() []TokenTrace {
	if m != nil {
		return m.Tokens
	}
	return nil
}

// TokenTrace defines an instance that records a pairing consisting of a native
//  nft and an ERC721 token.
type TokenTrace struct {
	// token_id is the token id of the native nft module
	TokenId string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// erc721_id is the id of the erc721 contract token
	Erc721Id string `protobuf:"bytes,2,opt,name=erc721_id,json=erc721Id,proto3" json:"erc721_id,omitempty"`
}

func (m *TokenTrace) Reset()         { *m = TokenTrace{} }
func (m *TokenTrace) String() string { return proto.CompactTextString(m) }
func (*TokenTrace) ProtoMessage()    {}
func (*TokenTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f31bd2c96fa73a7, []int{2}
}
func (m *TokenTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTrace.Merge(m, src)
}
func (m *TokenTrace) XXX_Size() int {
	return m.Size()
}
func (m *TokenTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTrace.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTrace proto.InternalMessageInfo

func (m *TokenTrace) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *TokenTrace) GetErc721Id() string {
	if m != nil {
		return m.Erc721Id
	}
	return ""
}

type ClassCollection struct {
	ClassPair ClassPair `protobuf:"bytes,1,opt,name=class_pair,json=classPair,proto3" json:"class_pair"`
	// tokens is a array of the token in class or the erc721 contract
	Tokens []TokenTrace `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens"`
}

func (m *ClassCollection) Reset()         { *m = ClassCollection{} }
func (m *ClassCollection) String() string { return proto.CompactTextString(m) }
func (*ClassCollection) ProtoMessage()    {}
func (*ClassCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f31bd2c96fa73a7, []int{3}
}
func (m *ClassCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassCollection.Merge(m, src)
}
func (m *ClassCollection) XXX_Size() int {
	return m.Size()
}
func (m *ClassCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ClassCollection proto.InternalMessageInfo

func (m *ClassCollection) GetClassPair() ClassPair {
	if m != nil {
		return m.ClassPair
	}
	return ClassPair{}
}

func (m *ClassCollection) GetTokens() []TokenTrace {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "erc721_bridge.converter.v1.GenesisState")
	proto.RegisterType((*ClassTrace)(nil), "erc721_bridge.converter.v1.ClassTrace")
	proto.RegisterType((*TokenTrace)(nil), "erc721_bridge.converter.v1.TokenTrace")
	proto.RegisterType((*ClassCollection)(nil), "erc721_bridge.converter.v1.ClassCollection")
}

func init() {
	proto.RegisterFile("erc721_bridge/converter/v1/genesis.proto", fileDescriptor_9f31bd2c96fa73a7)
}

var fileDescriptor_9f31bd2c96fa73a7 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x2d, 0x4a, 0x36,
	0x37, 0x32, 0x8c, 0x4f, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x4b, 0x2d,
	0x2a, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x42, 0x51, 0xa9, 0x07, 0x57, 0xa9, 0x57, 0x66, 0x28, 0xa5,
	0x8e, 0xc7, 0x14, 0x88, 0x14, 0xc4, 0x10, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f,
	0xc4, 0x82, 0x88, 0x2a, 0xed, 0x67, 0xe4, 0xe2, 0x71, 0x87, 0x58, 0x16, 0x5c, 0x92, 0x58, 0x92,
	0x2a, 0x14, 0xc7, 0x25, 0x98, 0x9c, 0x93, 0x58, 0x5c, 0x1c, 0x9f, 0x9c, 0x9f, 0x93, 0x93, 0x9a,
	0x5c, 0x92, 0x99, 0x9f, 0x57, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xad, 0x87, 0xdb,
	0x1d, 0x7a, 0xce, 0x20, 0x4d, 0xce, 0x70, 0x3d, 0x4e, 0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0x09,
	0x24, 0xa3, 0x0a, 0x17, 0x0b, 0xf9, 0x73, 0xf1, 0x40, 0xcc, 0x2f, 0x29, 0x4a, 0x4c, 0x4e, 0x2d,
	0x96, 0x60, 0x02, 0x1b, 0xad, 0x46, 0xd0, 0xe8, 0x10, 0x90, 0x72, 0xa8, 0xa9, 0xdc, 0xc9, 0x70,
	0x91, 0x62, 0xa5, 0x7e, 0x46, 0x2e, 0x2e, 0x84, 0x0a, 0x21, 0x49, 0x2e, 0x0e, 0x88, 0xf9, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xec, 0x60, 0xbe, 0x67, 0x8a, 0x90, 0x14, 0x17,
	0x47, 0x72, 0x7e, 0x1e, 0xc8, 0xde, 0x12, 0x09, 0x26, 0xb0, 0x14, 0x9c, 0x2f, 0xe4, 0xc2, 0xc5,
	0x56, 0x92, 0x9f, 0x9d, 0x9a, 0x57, 0x2c, 0xc1, 0x4c, 0xd8, 0x41, 0x21, 0x20, 0x95, 0xc8, 0x0e,
	0x82, 0xea, 0xb5, 0x62, 0x79, 0xb1, 0x40, 0x9e, 0x51, 0xc9, 0x8b, 0x8b, 0x0b, 0xa1, 0x02, 0xe4,
	0x20, 0xb0, 0x2c, 0x92, 0x83, 0xc0, 0x7c, 0xcf, 0x14, 0x21, 0x69, 0x2e, 0x4e, 0xa8, 0x2d, 0x99,
	0x29, 0x30, 0x17, 0x41, 0x04, 0x3c, 0x53, 0xa0, 0x66, 0xad, 0x64, 0xe4, 0xe2, 0x47, 0x0b, 0x5a,
	0x21, 0x2f, 0x2e, 0x2e, 0x88, 0x17, 0x0b, 0x12, 0x33, 0x8b, 0xc0, 0x66, 0x72, 0x1b, 0xa9, 0x12,
	0x0c, 0xc0, 0x80, 0xc4, 0xcc, 0x22, 0xa8, 0x73, 0x39, 0x93, 0x61, 0x02, 0x48, 0xfe, 0x66, 0xa2,
	0xd4, 0xdf, 0x4e, 0x3e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x94,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x59, 0x94, 0x59, 0x9c, 0x97,
	0x5a, 0x02, 0x4d, 0x9c, 0xba, 0xd0, 0x74, 0x5b, 0x81, 0x94, 0x72, 0x4b, 0x2a, 0x0b, 0x52, 0x8b,
	0x93, 0xd8, 0xc0, 0x09, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xd9, 0xf1, 0x4b, 0x27,
	0x03, 0x00, 0x00,
}

func (this *ClassTrace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClassTrace)
	if !ok {
		that2, ok := that.(ClassTrace)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ClassId != that1.ClassId {
		return false
	}
	if this.Contract != that1.Contract {
		return false
	}
	if len(this.Tokens) != len(that1.Tokens) {
		return false
	}
	for i := range this.Tokens {
		if !this.Tokens[i].Equal(&that1.Tokens[i]) {
			return false
		}
	}
	return true
}
func (this *TokenTrace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenTrace)
	if !ok {
		that2, ok := that.(TokenTrace)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TokenId != that1.TokenId {
		return false
	}
	if this.Erc721Id != that1.Erc721Id {
		return false
	}
	return true
}
func (this *ClassCollection) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClassCollection)
	if !ok {
		that2, ok := that.(ClassCollection)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ClassPair.Equal(&that1.ClassPair) {
		return false
	}
	if len(this.Tokens) != len(that1.Tokens) {
		return false
	}
	for i := range this.Tokens {
		if !this.Tokens[i].Equal(&that1.Tokens[i]) {
			return false
		}
	}
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassTraces) > 0 {
		for iNdEx := len(m.ClassTraces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassTraces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClassCollections) > 0 {
		for iNdEx := len(m.ClassCollections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassCollections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClassTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Erc721Id) > 0 {
		i -= len(m.Erc721Id)
		copy(dAtA[i:], m.Erc721Id)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Erc721Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClassCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ClassPair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClassCollections) > 0 {
		for _, e := range m.ClassCollections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClassTraces) > 0 {
		for _, e := range m.ClassTraces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ClassTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *TokenTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Erc721Id)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *ClassCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClassPair.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassCollections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassCollections = append(m.ClassCollections, ClassCollection{})
			if err := m.ClassCollections[len(m.ClassCollections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassTraces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassTraces = append(m.ClassTraces, ClassTrace{})
			if err := m.ClassTraces[len(m.ClassTraces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClassTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ClassTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, TokenTrace{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: TokenTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc721Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc721Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClassCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ClassCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClassPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, TokenTrace{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
