// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/curating/v1beta1/genesis.proto

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

type GenesisState struct {
	Params        Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Posts         Posts         `protobuf:"bytes,3,rep,name=posts,proto3,castrepeated=Posts" json:"posts"`
	Upvotes       Upvotes       `protobuf:"bytes,4,rep,name=upvotes,proto3,castrepeated=Upvotes" json:"upvotes"`
	CuratingQueue CuratingQueue `protobuf:"bytes,5,rep,name=curating_queue,json=curatingQueue,proto3,castrepeated=CuratingQueue" json:"curating_queue"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa25beab0641b9fc, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPosts() Posts {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *GenesisState) GetUpvotes() Upvotes {
	if m != nil {
		return m.Upvotes
	}
	return nil
}

func (m *GenesisState) GetCuratingQueue() CuratingQueue {
	if m != nil {
		return m.CuratingQueue
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "stargaze.curating.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("stargaze/curating/v1beta1/genesis.proto", fileDescriptor_aa25beab0641b9fc)
}

var fileDescriptor_aa25beab0641b9fc = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0x80, 0x13, 0xfb, 0x07, 0xa9, 0x55, 0x08, 0x0a, 0xb1, 0x8b, 0x69, 0x75, 0x63, 0x37, 0xce,
	0xd0, 0xf6, 0x00, 0x42, 0x14, 0xdc, 0x28, 0xd4, 0x88, 0x2e, 0xdc, 0x94, 0x49, 0x18, 0xc6, 0x80,
	0xed, 0x8c, 0x79, 0x33, 0xf5, 0xe7, 0x14, 0x9e, 0xc3, 0x73, 0xb8, 0xe8, 0xb2, 0x4b, 0x57, 0x2a,
	0xed, 0x45, 0x24, 0x93, 0xa4, 0xba, 0x69, 0xdd, 0xcd, 0x83, 0xef, 0xfb, 0xde, 0xc0, 0x73, 0x0e,
	0x41, 0xd1, 0x84, 0xd3, 0x17, 0x46, 0x22, 0x9d, 0x50, 0x15, 0x8f, 0x39, 0x99, 0x74, 0x43, 0xa6,
	0x68, 0x97, 0x70, 0x36, 0x66, 0x10, 0x03, 0x96, 0x89, 0x50, 0xc2, 0xdd, 0x2b, 0x40, 0x5c, 0x80,
	0x38, 0x07, 0x9b, 0x3b, 0x5c, 0x70, 0x61, 0x28, 0x92, 0xbe, 0x32, 0xa1, 0xd9, 0x59, 0x5d, 0x5e,
	0x16, 0x0c, 0x79, 0xf0, 0xbe, 0xe1, 0x6c, 0x9e, 0x65, 0xcb, 0xae, 0x14, 0x55, 0xcc, 0x3d, 0x76,
	0xaa, 0x92, 0x26, 0x74, 0x04, 0x9e, 0xdd, 0xb6, 0x3b, 0xf5, 0xde, 0x3e, 0x5e, 0xb9, 0x1c, 0x0f,
	0x0c, 0xe8, 0x97, 0xa7, 0x9f, 0x2d, 0x2b, 0xc8, 0x35, 0xf7, 0xd4, 0xa9, 0x48, 0x01, 0x0a, 0xbc,
	0x52, 0xbb, 0xd4, 0xa9, 0xf7, 0x5a, 0xeb, 0x7c, 0x01, 0xca, 0x6f, 0xa4, 0xf6, 0xdb, 0x57, 0xab,
	0x92, 0x4e, 0x10, 0x64, 0xb2, 0x7b, 0xee, 0xd4, 0xb4, 0x9c, 0x08, 0xc5, 0xc0, 0x2b, 0x9b, 0xce,
	0xba, 0x7f, 0x5c, 0x1b, 0xd2, 0xdf, 0xce, 0x4b, 0xb5, 0x6c, 0x86, 0xa0, 0x48, 0xb8, 0x43, 0x67,
	0xab, 0x90, 0x86, 0x0f, 0x9a, 0x69, 0xe6, 0x55, 0xfe, 0x8d, 0xde, 0x0c, 0x06, 0x34, 0x4e, 0xfc,
	0xdd, 0x3c, 0xda, 0x38, 0xc9, 0x81, 0xcb, 0xd4, 0x0f, 0x1a, 0xd1, 0xdf, 0xd1, 0xbf, 0x98, 0xce,
	0x91, 0x3d, 0x9b, 0x23, 0xfb, 0x7b, 0x8e, 0xec, 0xd7, 0x05, 0xb2, 0x66, 0x0b, 0x64, 0x7d, 0x2c,
	0x90, 0x75, 0xdb, 0xe7, 0xb1, 0xba, 0xd3, 0x21, 0x8e, 0xc4, 0x88, 0x48, 0x1d, 0xde, 0xc7, 0xd1,
	0x11, 0x7d, 0x64, 0x20, 0x46, 0x8c, 0x2c, 0x8f, 0xf4, 0xf4, 0x7b, 0x26, 0xf5, 0x2c, 0x19, 0x84,
	0x55, 0x73, 0x9c, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x6e, 0x8d, 0xe6, 0x22, 0x02,
	0x00, 0x00,
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
	if len(m.CuratingQueue) > 0 {
		for iNdEx := len(m.CuratingQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CuratingQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Upvotes) > 0 {
		for iNdEx := len(m.Upvotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Upvotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Posts) > 0 {
		for iNdEx := len(m.Posts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Posts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Posts) > 0 {
		for _, e := range m.Posts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Upvotes) > 0 {
		for _, e := range m.Upvotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CuratingQueue) > 0 {
		for _, e := range m.CuratingQueue {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Posts", wireType)
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
			m.Posts = append(m.Posts, Post{})
			if err := m.Posts[len(m.Posts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upvotes", wireType)
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
			m.Upvotes = append(m.Upvotes, Upvote{})
			if err := m.Upvotes[len(m.Upvotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuratingQueue", wireType)
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
			m.CuratingQueue = append(m.CuratingQueue, VPPair{})
			if err := m.CuratingQueue[len(m.CuratingQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
