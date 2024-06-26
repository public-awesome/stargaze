// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/globalfee/v1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// SetCodeAuthorizationProposal ...
//
// Deprecated: Do not use.
type SetCodeAuthorizationProposal struct {
	Title             string             `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description       string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	CodeAuthorization *CodeAuthorization `protobuf:"bytes,3,opt,name=code_authorization,json=codeAuthorization,proto3" json:"code_authorization,omitempty" yaml:"code_authorization"`
}

func (m *SetCodeAuthorizationProposal) Reset()         { *m = SetCodeAuthorizationProposal{} }
func (m *SetCodeAuthorizationProposal) String() string { return proto.CompactTextString(m) }
func (*SetCodeAuthorizationProposal) ProtoMessage()    {}
func (*SetCodeAuthorizationProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b49a370918ffdc3b, []int{0}
}
func (m *SetCodeAuthorizationProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetCodeAuthorizationProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetCodeAuthorizationProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetCodeAuthorizationProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetCodeAuthorizationProposal.Merge(m, src)
}
func (m *SetCodeAuthorizationProposal) XXX_Size() int {
	return m.Size()
}
func (m *SetCodeAuthorizationProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetCodeAuthorizationProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetCodeAuthorizationProposal proto.InternalMessageInfo

func (m *SetCodeAuthorizationProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SetCodeAuthorizationProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SetCodeAuthorizationProposal) GetCodeAuthorization() *CodeAuthorization {
	if m != nil {
		return m.CodeAuthorization
	}
	return nil
}

// RemoveCodeAuthorizationProposal
//
// Deprecated: Do not use.
type RemoveCodeAuthorizationProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	CodeID      uint64 `protobuf:"varint,3,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
}

func (m *RemoveCodeAuthorizationProposal) Reset()         { *m = RemoveCodeAuthorizationProposal{} }
func (m *RemoveCodeAuthorizationProposal) String() string { return proto.CompactTextString(m) }
func (*RemoveCodeAuthorizationProposal) ProtoMessage()    {}
func (*RemoveCodeAuthorizationProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b49a370918ffdc3b, []int{1}
}
func (m *RemoveCodeAuthorizationProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveCodeAuthorizationProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveCodeAuthorizationProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveCodeAuthorizationProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveCodeAuthorizationProposal.Merge(m, src)
}
func (m *RemoveCodeAuthorizationProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveCodeAuthorizationProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveCodeAuthorizationProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveCodeAuthorizationProposal proto.InternalMessageInfo

func (m *RemoveCodeAuthorizationProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RemoveCodeAuthorizationProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoveCodeAuthorizationProposal) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

// RemoveCodeAuthorizationProposal ...
//
// Deprecated: Do not use.
type SetContractAuthorizationProposal struct {
	Title                 string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description           string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	ContractAuthorization *ContractAuthorization `protobuf:"bytes,3,opt,name=contract_authorization,json=contractAuthorization,proto3" json:"contract_authorization,omitempty" yaml:"contract_authorization"`
}

func (m *SetContractAuthorizationProposal) Reset()         { *m = SetContractAuthorizationProposal{} }
func (m *SetContractAuthorizationProposal) String() string { return proto.CompactTextString(m) }
func (*SetContractAuthorizationProposal) ProtoMessage()    {}
func (*SetContractAuthorizationProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b49a370918ffdc3b, []int{2}
}
func (m *SetContractAuthorizationProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetContractAuthorizationProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetContractAuthorizationProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetContractAuthorizationProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetContractAuthorizationProposal.Merge(m, src)
}
func (m *SetContractAuthorizationProposal) XXX_Size() int {
	return m.Size()
}
func (m *SetContractAuthorizationProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetContractAuthorizationProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetContractAuthorizationProposal proto.InternalMessageInfo

func (m *SetContractAuthorizationProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SetContractAuthorizationProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SetContractAuthorizationProposal) GetContractAuthorization() *ContractAuthorization {
	if m != nil {
		return m.ContractAuthorization
	}
	return nil
}

// RemoveCodeAuthorizationProposal ...
//
// Deprecated: Do not use.
type RemoveContractAuthorizationProposal struct {
	Title           string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description     string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	ContractAddress string `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
}

func (m *RemoveContractAuthorizationProposal) Reset()         { *m = RemoveContractAuthorizationProposal{} }
func (m *RemoveContractAuthorizationProposal) String() string { return proto.CompactTextString(m) }
func (*RemoveContractAuthorizationProposal) ProtoMessage()    {}
func (*RemoveContractAuthorizationProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b49a370918ffdc3b, []int{3}
}
func (m *RemoveContractAuthorizationProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveContractAuthorizationProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveContractAuthorizationProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveContractAuthorizationProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveContractAuthorizationProposal.Merge(m, src)
}
func (m *RemoveContractAuthorizationProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveContractAuthorizationProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveContractAuthorizationProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveContractAuthorizationProposal proto.InternalMessageInfo

func (m *RemoveContractAuthorizationProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RemoveContractAuthorizationProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoveContractAuthorizationProposal) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*SetCodeAuthorizationProposal)(nil), "publicawesome.stargaze.globalfee.v1.SetCodeAuthorizationProposal")
	proto.RegisterType((*RemoveCodeAuthorizationProposal)(nil), "publicawesome.stargaze.globalfee.v1.RemoveCodeAuthorizationProposal")
	proto.RegisterType((*SetContractAuthorizationProposal)(nil), "publicawesome.stargaze.globalfee.v1.SetContractAuthorizationProposal")
	proto.RegisterType((*RemoveContractAuthorizationProposal)(nil), "publicawesome.stargaze.globalfee.v1.RemoveContractAuthorizationProposal")
}

func init() {
	proto.RegisterFile("publicawesome/stargaze/globalfee/v1/proposal.proto", fileDescriptor_b49a370918ffdc3b)
}

var fileDescriptor_b49a370918ffdc3b = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7b, 0x06, 0x82, 0x7a, 0x45, 0xd0, 0x5a, 0x50, 0xdc, 0xd0, 0xda, 0xe1, 0x2a, 0x55,
	0xa5, 0x52, 0x6c, 0xa5, 0x55, 0x51, 0xe5, 0x0d, 0x83, 0x90, 0x2a, 0x24, 0x40, 0x66, 0x63, 0xa9,
	0xce, 0xf6, 0xe1, 0x5a, 0xb2, 0x7d, 0x96, 0x7d, 0x31, 0xb4, 0x33, 0x53, 0x27, 0x06, 0xbe, 0x01,
	0x5f, 0x80, 0x81, 0x0f, 0x81, 0x98, 0x3a, 0x30, 0x30, 0x59, 0x28, 0x19, 0x58, 0x98, 0xfc, 0x09,
	0x90, 0xcf, 0x4e, 0x63, 0x92, 0x28, 0xc9, 0x96, 0xc5, 0xba, 0x7b, 0xef, 0xfd, 0xef, 0xbd, 0xf7,
	0xf3, 0xd3, 0x83, 0xfb, 0x51, 0xd7, 0xf2, 0x3d, 0x1b, 0xbf, 0x27, 0x09, 0x0d, 0x88, 0x96, 0x30,
	0x1c, 0xbb, 0xf8, 0x9c, 0x68, 0xae, 0x4f, 0x2d, 0xec, 0xbf, 0x23, 0x44, 0x4b, 0x3b, 0x5a, 0x14,
	0xd3, 0x88, 0x26, 0xd8, 0x57, 0xa3, 0x98, 0x32, 0x2a, 0x6e, 0xff, 0xa7, 0x51, 0x07, 0x1a, 0xf5,
	0x4a, 0xa3, 0xa6, 0x9d, 0xe6, 0x1a, 0x0e, 0xbc, 0x90, 0x6a, 0xfc, 0x5b, 0xea, 0x9a, 0x1b, 0x36,
	0x4d, 0x02, 0x9a, 0x9c, 0xf0, 0x9b, 0x56, 0x5e, 0x2a, 0xd7, 0x5d, 0x97, 0xba, 0xb4, 0xb4, 0x17,
	0xa7, 0xca, 0x7a, 0x30, 0x4f, 0x71, 0xc3, 0xac, 0x5c, 0x84, 0x7e, 0x0a, 0x70, 0xf3, 0x0d, 0x61,
	0x4f, 0xa9, 0x43, 0x9e, 0x74, 0xd9, 0x29, 0x8d, 0xbd, 0x73, 0xcc, 0x3c, 0x1a, 0xbe, 0xae, 0x9a,
	0x10, 0x77, 0xe0, 0x0d, 0xe6, 0x31, 0x9f, 0x48, 0xa0, 0x05, 0x76, 0x97, 0x8d, 0xd5, 0x3c, 0x53,
	0x6e, 0x9d, 0xe1, 0xc0, 0xd7, 0x11, 0x37, 0x23, 0xb3, 0x74, 0x8b, 0x47, 0x70, 0xc5, 0x21, 0x89,
	0x1d, 0x7b, 0x51, 0x21, 0x97, 0x04, 0x1e, 0xbd, 0x9e, 0x67, 0x8a, 0x58, 0x46, 0xd7, 0x9c, 0xc8,
	0xac, 0x87, 0x8a, 0x1f, 0x01, 0x14, 0x6d, 0xea, 0x90, 0x13, 0x5c, 0x2f, 0x40, 0xba, 0xd6, 0x02,
	0xbb, 0x2b, 0xfb, 0x8f, 0xd5, 0x39, 0xf0, 0xa9, 0x63, 0xe5, 0x1b, 0x5b, 0x79, 0xa6, 0x6c, 0x94,
	0x99, 0xc7, 0xdf, 0x46, 0xe6, 0x9a, 0x3d, 0xaa, 0xd0, 0x5f, 0xfc, 0xf8, 0xd6, 0x6e, 0x56, 0x98,
	0x5d, 0x9a, 0xaa, 0x69, 0xc7, 0x22, 0x0c, 0x17, 0x6f, 0x87, 0x8c, 0x84, 0xec, 0xe2, 0xcf, 0xd7,
	0xbd, 0x9d, 0x21, 0xca, 0x69, 0xcc, 0x24, 0x80, 0x2e, 0x04, 0xa8, 0x98, 0x24, 0xa0, 0x29, 0x59,
	0x24, 0xd9, 0x43, 0x78, 0x93, 0x37, 0xef, 0x39, 0x9c, 0xe6, 0x75, 0x63, 0xb3, 0x97, 0x29, 0x8d,
	0xa2, 0xa2, 0xe3, 0x67, 0x79, 0xa6, 0xdc, 0xae, 0xf1, 0xf1, 0x1c, 0x64, 0x36, 0x8a, 0xd3, 0xb1,
	0xa3, 0xbf, 0x9c, 0x4d, 0xe2, 0xd1, 0x90, 0xc4, 0x8c, 0x36, 0x25, 0x80, 0xfe, 0x0a, 0xb0, 0xc5,
	0x79, 0x85, 0x2c, 0xc6, 0x36, 0x5b, 0x14, 0x8d, 0xcf, 0x00, 0xae, 0xdb, 0x55, 0x0d, 0x13, 0x67,
	0x4d, 0x9f, 0x73, 0xd6, 0x26, 0xb4, 0x61, 0x3c, 0xcc, 0x33, 0x65, 0x6b, 0xc0, 0x73, 0x52, 0x0e,
	0x64, 0xde, 0xb3, 0x27, 0x29, 0xf5, 0x57, 0xb3, 0x69, 0xef, 0x8d, 0xcc, 0xdd, 0x14, 0x8e, 0x12,
	0x40, 0x5f, 0x04, 0xb8, 0x3d, 0xf8, 0x29, 0x8b, 0x25, 0xfe, 0x1c, 0xae, 0x0e, 0x61, 0x38, 0x4e,
	0x4c, 0x92, 0x84, 0xa3, 0x5e, 0x36, 0x1e, 0xe4, 0x99, 0x72, 0x7f, 0x14, 0x57, 0x19, 0x81, 0xcc,
	0x3b, 0x57, 0xa0, 0x4a, 0x8b, 0x6e, 0xce, 0x46, 0xd4, 0x1e, 0x1f, 0xc8, 0xa9, 0x94, 0x0c, 0xf3,
	0x7b, 0x4f, 0x06, 0x97, 0x3d, 0x19, 0xfc, 0xee, 0xc9, 0xe0, 0x53, 0x5f, 0x5e, 0xba, 0xec, 0xcb,
	0x4b, 0xbf, 0xfa, 0xf2, 0xd2, 0xdb, 0x23, 0xd7, 0x63, 0xa7, 0x5d, 0x4b, 0xb5, 0x69, 0xa0, 0x95,
	0x03, 0xd1, 0x1e, 0xdb, 0xa9, 0x69, 0xe7, 0x50, 0xfb, 0x50, 0xdb, 0xac, 0xec, 0x2c, 0x22, 0x89,
	0xd5, 0xe0, 0x3b, 0xf5, 0xe0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x12, 0x41, 0xf7, 0x27,
	0x06, 0x00, 0x00,
}

func (m *SetCodeAuthorizationProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetCodeAuthorizationProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetCodeAuthorizationProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CodeAuthorization != nil {
		{
			size, err := m.CodeAuthorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveCodeAuthorizationProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveCodeAuthorizationProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveCodeAuthorizationProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CodeID != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetContractAuthorizationProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetContractAuthorizationProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetContractAuthorizationProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ContractAuthorization != nil {
		{
			size, err := m.ContractAuthorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveContractAuthorizationProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveContractAuthorizationProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveContractAuthorizationProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SetCodeAuthorizationProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.CodeAuthorization != nil {
		l = m.CodeAuthorization.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *RemoveCodeAuthorizationProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.CodeID != 0 {
		n += 1 + sovProposal(uint64(m.CodeID))
	}
	return n
}

func (m *SetContractAuthorizationProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.ContractAuthorization != nil {
		l = m.ContractAuthorization.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *RemoveContractAuthorizationProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SetCodeAuthorizationProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: SetCodeAuthorizationProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetCodeAuthorizationProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeAuthorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CodeAuthorization == nil {
				m.CodeAuthorization = &CodeAuthorization{}
			}
			if err := m.CodeAuthorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *RemoveCodeAuthorizationProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: RemoveCodeAuthorizationProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveCodeAuthorizationProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *SetContractAuthorizationProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: SetContractAuthorizationProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetContractAuthorizationProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAuthorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContractAuthorization == nil {
				m.ContractAuthorization = &ContractAuthorization{}
			}
			if err := m.ContractAuthorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *RemoveContractAuthorizationProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: RemoveContractAuthorizationProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveContractAuthorizationProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
