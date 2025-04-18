// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/cron/v1/cron.proto

package types

import (
	fmt "fmt"
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

// Params holds parameters for the cron module.
type Params struct {
	// Addresses which act as admins of the module. They can promote and demote
	// contracts without having to go via governance.
	AdminAddresses []string `protobuf:"bytes,1,rep,name=admin_addresses,json=adminAddresses,proto3" json:"admin_addresses,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b37a39eca21e75, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAdminAddresses() []string {
	if m != nil {
		return m.AdminAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "publicawesome.stargaze.cron.v1.Params")
}

func init() {
	proto.RegisterFile("publicawesome/stargaze/cron/v1/cron.proto", fileDescriptor_72b37a39eca21e75)
}

var fileDescriptor_72b37a39eca21e75 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x28, 0x4d, 0xca,
	0xc9, 0x4c, 0x4e, 0x2c, 0x4f, 0x2d, 0xce, 0xcf, 0x4d, 0xd5, 0x2f, 0x2e, 0x49, 0x2c, 0x4a, 0x4f,
	0xac, 0x4a, 0xd5, 0x4f, 0x2e, 0xca, 0xcf, 0xd3, 0x2f, 0x33, 0x04, 0xd3, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x72, 0x28, 0x4a, 0xf5, 0x60, 0x4a, 0xf5, 0xc0, 0x4a, 0xca, 0x0c, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4a, 0xf5, 0x41, 0x2c, 0x88, 0x2e, 0x25, 0x73, 0x2e, 0xb6, 0x80,
	0xc4, 0xa2, 0xc4, 0xdc, 0x62, 0x21, 0x75, 0x2e, 0xfe, 0xc4, 0x94, 0xdc, 0xcc, 0xbc, 0xf8, 0xc4,
	0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x3e,
	0xb0, 0xb0, 0x23, 0x4c, 0xd4, 0x8a, 0x65, 0xc6, 0x02, 0x79, 0x06, 0x27, 0xdf, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x87, 0xb8, 0x49, 0x17, 0xc3, 0xfd, 0x65, 0x86, 0xe6, 0xfa, 0x15, 0x10, 0x5f,
	0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x9d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x59, 0xa2, 0x9b, 0xf1, 0x00, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdminAddresses) > 0 {
		for iNdEx := len(m.AdminAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdminAddresses[iNdEx])
			copy(dAtA[i:], m.AdminAddresses[iNdEx])
			i = encodeVarintCron(dAtA, i, uint64(len(m.AdminAddresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCron(dAtA []byte, offset int, v uint64) int {
	offset -= sovCron(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AdminAddresses) > 0 {
		for _, s := range m.AdminAddresses {
			l = len(s)
			n += 1 + l + sovCron(uint64(l))
		}
	}
	return n
}

func sovCron(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCron(x uint64) (n int) {
	return sovCron(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCron
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCron
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
				return ErrInvalidLengthCron
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCron
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddresses = append(m.AdminAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCron(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCron
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
func skipCron(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCron
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
					return 0, ErrIntOverflowCron
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
					return 0, ErrIntOverflowCron
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
				return 0, ErrInvalidLengthCron
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCron
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCron
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCron        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCron          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCron = fmt.Errorf("proto: unexpected end of group")
)
