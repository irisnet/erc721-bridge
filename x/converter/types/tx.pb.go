// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: converter/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgConvertNFT defines a Msg to convert a native Cosmos coin to a ERC721 token
type MsgConvertNFT struct {
	// class_id of the native Cosmos Class to convert
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// token_id of the native Cosmos Class to convert
	TokenId string `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// receiver is the hex address to receive ERC721 token
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender is the cosmos bech32 address from the owner of the given Cosmos Class
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertNFT) Reset()         { *m = MsgConvertNFT{} }
func (m *MsgConvertNFT) String() string { return proto.CompactTextString(m) }
func (*MsgConvertNFT) ProtoMessage()    {}
func (*MsgConvertNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_c320cd564e01ee74, []int{0}
}
func (m *MsgConvertNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertNFT.Merge(m, src)
}
func (m *MsgConvertNFT) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertNFT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertNFT proto.InternalMessageInfo

func (m *MsgConvertNFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *MsgConvertNFT) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *MsgConvertNFT) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertNFT) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertNFTResponse returns no fields
type MsgConvertNFTResponse struct {
}

func (m *MsgConvertNFTResponse) Reset()         { *m = MsgConvertNFTResponse{} }
func (m *MsgConvertNFTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertNFTResponse) ProtoMessage()    {}
func (*MsgConvertNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c320cd564e01ee74, []int{1}
}
func (m *MsgConvertNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertNFTResponse.Merge(m, src)
}
func (m *MsgConvertNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertNFTResponse proto.InternalMessageInfo

// MsgConvertERC721 defines a Msg to convert a ERC721 token to a native Cosmos
// coin.
type MsgConvertERC721 struct {
	// contract_address of an ERC721 token contract, that is registered in a token pair
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// token_id of the ERC721 token to convert
	TokenId github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_id"`
	// receiver is the bech32 address to receive native Cosmos Class
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender is the hex address from the owner of the given ERC721 tokens
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertERC721) Reset()         { *m = MsgConvertERC721{} }
func (m *MsgConvertERC721) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC721) ProtoMessage()    {}
func (*MsgConvertERC721) Descriptor() ([]byte, []int) {
	return fileDescriptor_c320cd564e01ee74, []int{2}
}
func (m *MsgConvertERC721) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC721) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC721.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC721) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC721.Merge(m, src)
}
func (m *MsgConvertERC721) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC721) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC721.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC721 proto.InternalMessageInfo

func (m *MsgConvertERC721) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgConvertERC721) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertERC721) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertERC721Response returns no fields
type MsgConvertERC721Response struct {
}

func (m *MsgConvertERC721Response) Reset()         { *m = MsgConvertERC721Response{} }
func (m *MsgConvertERC721Response) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC721Response) ProtoMessage()    {}
func (*MsgConvertERC721Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c320cd564e01ee74, []int{3}
}
func (m *MsgConvertERC721Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC721Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC721Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC721Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC721Response.Merge(m, src)
}
func (m *MsgConvertERC721Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC721Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC721Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC721Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvertNFT)(nil), "erc721_bridge.converter.v1.MsgConvertNFT")
	proto.RegisterType((*MsgConvertNFTResponse)(nil), "erc721_bridge.converter.v1.MsgConvertNFTResponse")
	proto.RegisterType((*MsgConvertERC721)(nil), "erc721_bridge.converter.v1.MsgConvertERC721")
	proto.RegisterType((*MsgConvertERC721Response)(nil), "erc721_bridge.converter.v1.MsgConvertERC721Response")
}

func init() { proto.RegisterFile("converter/v1/tx.proto", fileDescriptor_c320cd564e01ee74) }

var fileDescriptor_c320cd564e01ee74 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0xb3, 0xa9, 0xd4, 0x3a, 0x50, 0x2c, 0x83, 0xd5, 0xb8, 0xc8, 0x56, 0x16, 0xd4, 0x56,
	0xcc, 0x8c, 0x1b, 0x85, 0x9e, 0x6d, 0x51, 0x08, 0x58, 0x0f, 0xc1, 0x93, 0x97, 0xb0, 0x99, 0xf9,
	0x1c, 0x87, 0x36, 0xf3, 0x85, 0x99, 0x31, 0xb4, 0x57, 0x9f, 0x40, 0xf0, 0xe4, 0x03, 0xf8, 0x16,
	0x3e, 0x40, 0xbd, 0x15, 0xbc, 0x88, 0x87, 0x22, 0x89, 0x0f, 0x22, 0x99, 0x49, 0x36, 0x4d, 0x41,
	0xac, 0xf4, 0xb4, 0xf3, 0xcd, 0xef, 0xbf, 0xff, 0xf9, 0xcf, 0xf7, 0xed, 0x92, 0x75, 0x81, 0x66,
	0x08, 0xd6, 0x83, 0xe5, 0xc3, 0x82, 0xfb, 0x43, 0x36, 0xb0, 0xe8, 0x91, 0xa6, 0x60, 0xc5, 0x76,
	0xab, 0xe8, 0xf6, 0xac, 0x96, 0x0a, 0x58, 0x25, 0x62, 0xc3, 0x22, 0xbd, 0xa1, 0x50, 0x61, 0x90,
	0xf1, 0xc9, 0x2a, 0xbe, 0x91, 0xde, 0x51, 0x88, 0xea, 0x00, 0x78, 0x39, 0xd0, 0xbc, 0x34, 0x06,
	0x7d, 0xe9, 0x35, 0x1a, 0x17, 0x69, 0x7e, 0x44, 0x56, 0xf7, 0x9c, 0xda, 0x8d, 0x36, 0xaf, 0x5e,
	0xbc, 0xa6, 0xb7, 0xc9, 0x8a, 0x38, 0x28, 0x9d, 0xeb, 0x6a, 0xd9, 0x48, 0xee, 0x26, 0x9b, 0xd7,
	0x3a, 0x57, 0x43, 0xdd, 0x96, 0x13, 0xe4, 0x71, 0x1f, 0xcc, 0x04, 0xd5, 0x23, 0x0a, 0x75, 0x5b,
	0xd2, 0x94, 0xac, 0x58, 0x10, 0xa0, 0x87, 0x60, 0x1b, 0x4b, 0x01, 0x55, 0x35, 0xbd, 0x49, 0x96,
	0x1d, 0x18, 0x09, 0xb6, 0x71, 0x25, 0x90, 0x69, 0x95, 0xdf, 0x22, 0xeb, 0x0b, 0x47, 0x77, 0xc0,
	0x0d, 0xd0, 0x38, 0xc8, 0xbf, 0x26, 0x64, 0x6d, 0x4e, 0x9e, 0x77, 0x76, 0xb7, 0x5b, 0x05, 0xdd,
	0x22, 0x6b, 0x02, 0x8d, 0xb7, 0xa5, 0xf0, 0xdd, 0x52, 0x4a, 0x0b, 0xce, 0x4d, 0xf3, 0x5d, 0x9f,
	0xed, 0x3f, 0x8b, 0xdb, 0xb4, 0x7d, 0x3e, 0xe7, 0x0e, 0x3b, 0x3e, 0xdd, 0xa8, 0xfd, 0x3c, 0xdd,
	0xb8, 0xaf, 0xb4, 0x7f, 0xf7, 0xbe, 0xc7, 0x04, 0xf6, 0xb9, 0x40, 0xd7, 0x47, 0x37, 0x7d, 0x34,
	0x9d, 0xdc, 0xe7, 0xfe, 0x68, 0x00, 0x8e, 0xb5, 0x8d, 0xbf, 0xdc, 0xbd, 0x52, 0xd2, 0x38, 0x9f,
	0x7e, 0x76, 0xb5, 0xd6, 0xb7, 0x3a, 0x59, 0xda, 0x73, 0x8a, 0x7e, 0x4e, 0x08, 0x39, 0xd3, 0xf4,
	0x2d, 0xf6, 0xf7, 0xb1, 0xb2, 0x85, 0x26, 0xa5, 0xc5, 0x85, 0xa5, 0x55, 0x3f, 0x9b, 0x1f, 0xbe,
	0xff, 0xfe, 0x54, 0x7f, 0x90, 0xdf, 0xe3, 0xda, 0x6a, 0xd7, 0x47, 0xc9, 0xa3, 0x45, 0x73, 0xfe,
	0x89, 0x4d, 0x57, 0x4d, 0xf3, 0xd6, 0xd3, 0x2f, 0x09, 0x59, 0x5d, 0xec, 0xfd, 0xa3, 0x8b, 0x9d,
	0x19, 0xd5, 0xe9, 0xd3, 0xff, 0x51, 0x57, 0x21, 0x1f, 0x87, 0x90, 0x0f, 0xf3, 0xcd, 0x7f, 0x87,
	0x8c, 0x60, 0xe7, 0xe5, 0xf1, 0x28, 0x4b, 0x4e, 0x46, 0x59, 0xf2, 0x6b, 0x94, 0x25, 0x1f, 0xc7,
	0x59, 0xed, 0x64, 0x9c, 0xd5, 0x7e, 0x8c, 0xb3, 0xda, 0x9b, 0xd6, 0x99, 0x31, 0x4f, 0xdc, 0x0c,
	0xf8, 0x99, 0x5b, 0xcc, 0xc4, 0x0f, 0xf9, 0xdc, 0x36, 0x8c, 0xbd, 0xb7, 0x1c, 0xfe, 0x87, 0x27,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x02, 0x0b, 0x8b, 0x78, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertNFT mints a ERC721 representation of the native Cosmos Class
	// that is registered on the token mapping.
	ConvertNFT(ctx context.Context, in *MsgConvertNFT, opts ...grpc.CallOption) (*MsgConvertNFTResponse, error)
	// ConvertERC721 mints a native Cosmos Class representation of the x/nft token
	// contract that is registered on the token mapping.
	ConvertERC721(ctx context.Context, in *MsgConvertERC721, opts ...grpc.CallOption) (*MsgConvertERC721Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertNFT(ctx context.Context, in *MsgConvertNFT, opts ...grpc.CallOption) (*MsgConvertNFTResponse, error) {
	out := new(MsgConvertNFTResponse)
	err := c.cc.Invoke(ctx, "/erc721_bridge.converter.v1.Msg/ConvertNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConvertERC721(ctx context.Context, in *MsgConvertERC721, opts ...grpc.CallOption) (*MsgConvertERC721Response, error) {
	out := new(MsgConvertERC721Response)
	err := c.cc.Invoke(ctx, "/erc721_bridge.converter.v1.Msg/ConvertERC721", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConvertNFT mints a ERC721 representation of the native Cosmos Class
	// that is registered on the token mapping.
	ConvertNFT(context.Context, *MsgConvertNFT) (*MsgConvertNFTResponse, error)
	// ConvertERC721 mints a native Cosmos Class representation of the x/nft token
	// contract that is registered on the token mapping.
	ConvertERC721(context.Context, *MsgConvertERC721) (*MsgConvertERC721Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ConvertNFT(ctx context.Context, req *MsgConvertNFT) (*MsgConvertNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertNFT not implemented")
}
func (*UnimplementedMsgServer) ConvertERC721(ctx context.Context, req *MsgConvertERC721) (*MsgConvertERC721Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertERC721 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ConvertNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/erc721_bridge.converter.v1.Msg/ConvertNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertNFT(ctx, req.(*MsgConvertNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConvertERC721_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertERC721)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertERC721(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/erc721_bridge.converter.v1.Msg/ConvertERC721",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertERC721(ctx, req.(*MsgConvertERC721))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erc721_bridge.converter.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertNFT",
			Handler:    _Msg_ConvertNFT_Handler,
		},
		{
			MethodName: "ConvertERC721",
			Handler:    _Msg_ConvertERC721_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "converter/v1/tx.proto",
}

func (m *MsgConvertNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC721) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC721) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC721) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.TokenId.Size()
		i -= size
		if _, err := m.TokenId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC721Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC721Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC721Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvertNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgConvertERC721) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TokenId.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertERC721Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvertNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertERC721) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertERC721: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC721: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertERC721Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertERC721Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC721Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
