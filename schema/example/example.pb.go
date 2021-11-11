// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example.proto

package example

import (
	fmt "fmt"
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

// DeepEnum is one or zero.
type Example_Nested_DeepEnum int32

const (
	Example_Nested_FALSE Example_Nested_DeepEnum = 0
	Example_Nested_TRUE  Example_Nested_DeepEnum = 1
)

var Example_Nested_DeepEnum_name = map[int32]string{
	0: "FALSE",
	1: "TRUE",
}

var Example_Nested_DeepEnum_value = map[string]int32{
	"FALSE": 0,
	"TRUE":  1,
}

func (x Example_Nested_DeepEnum) String() string {
	return proto.EnumName(Example_Nested_DeepEnum_name, int32(x))
}

func (Example_Nested_DeepEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0, 0, 0}
}

type Example struct {
	SingleNested         *Example_Nested   `protobuf:"bytes,25,opt,name=single_nested,json=singleNested,proto3" json:"single_nested,omitempty"`
	Uuid                 string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Nested               []*Example_Nested `protobuf:"bytes,2,rep,name=nested,proto3" json:"nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}
func (*Example) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}
func (m *Example) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Example) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Example.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Example) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example.Merge(m, src)
}
func (m *Example) XXX_Size() int {
	return m.Size()
}
func (m *Example) XXX_DiscardUnknown() {
	xxx_messageInfo_Example.DiscardUnknown(m)
}

var xxx_messageInfo_Example proto.InternalMessageInfo

func (m *Example) GetSingleNested() *Example_Nested {
	if m != nil {
		return m.SingleNested
	}
	return nil
}

func (m *Example) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Example) GetNested() []*Example_Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

// Nested is nested type.
type Example_Nested struct {
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Ok                   Example_Nested_DeepEnum `protobuf:"varint,3,opt,name=ok,proto3,enum=example.Example_Nested_DeepEnum" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Example_Nested) Reset()         { *m = Example_Nested{} }
func (m *Example_Nested) String() string { return proto.CompactTextString(m) }
func (*Example_Nested) ProtoMessage()    {}
func (*Example_Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0, 0}
}
func (m *Example_Nested) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Example_Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Example_Nested.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Example_Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example_Nested.Merge(m, src)
}
func (m *Example_Nested) XXX_Size() int {
	return m.Size()
}
func (m *Example_Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Example_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Example_Nested proto.InternalMessageInfo

func (m *Example_Nested) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Example_Nested) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Example_Nested) GetOk() Example_Nested_DeepEnum {
	if m != nil {
		return m.Ok
	}
	return Example_Nested_FALSE
}

func init() {
	proto.RegisterEnum("example.Example_Nested_DeepEnum", Example_Nested_DeepEnum_name, Example_Nested_DeepEnum_value)
	proto.RegisterType((*Example)(nil), "example.Example")
	proto.RegisterType((*Example_Nested)(nil), "example.Example.Nested")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x26, 0x30,
	0x71, 0xb1, 0xbb, 0x42, 0xd8, 0x42, 0x36, 0x5c, 0xbc, 0xc5, 0x99, 0x79, 0xe9, 0x39, 0xa9, 0xf1,
	0x79, 0xa9, 0xc5, 0x25, 0xa9, 0x29, 0x12, 0x92, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xe2, 0x7a, 0x30,
	0xbd, 0x50, 0x85, 0x7a, 0x7e, 0x60, 0xe9, 0x20, 0x1e, 0x88, 0x6a, 0x08, 0x4f, 0x48, 0x88, 0x8b,
	0xa5, 0xb4, 0x34, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0xd2, 0xe7,
	0x62, 0x83, 0x1a, 0xc5, 0xa4, 0xc0, 0x8c, 0xcf, 0x28, 0xa8, 0x32, 0xa9, 0x46, 0x46, 0x2e, 0x36,
	0x84, 0x79, 0x79, 0x89, 0xb9, 0xa9, 0x30, 0xf3, 0x40, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0xc4, 0xf4,
	0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x10, 0x53, 0xc8, 0x80, 0x8b, 0x29, 0x3f, 0x5b,
	0x82, 0x59, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x01, 0x87, 0xe9, 0x7a, 0x2e, 0xa9, 0xa9, 0x05, 0xae,
	0x79, 0xa5, 0xb9, 0x41, 0x4c, 0xf9, 0xd9, 0x4a, 0xf2, 0x5c, 0x1c, 0x30, 0xbe, 0x10, 0x27, 0x17,
	0xab, 0x9b, 0xa3, 0x4f, 0xb0, 0xab, 0x00, 0x83, 0x10, 0x07, 0x17, 0x4b, 0x48, 0x50, 0xa8, 0xab,
	0x00, 0xa3, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x18, 0x65, 0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9e, 0x9f,
	0x5d, 0x6a, 0x6c, 0x64, 0xa8, 0x5f, 0x50, 0x9a, 0x53, 0x9c, 0x58, 0xa4, 0x0b, 0x0e, 0x4a, 0xdd,
	0x0a, 0xfd, 0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0x44, 0x7d, 0xa8, 0xe5, 0xd6, 0x50, 0x3a, 0x89, 0x0d,
	0x2c, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x54, 0x07, 0x47, 0x36, 0x7b, 0x01, 0x00, 0x00,
}

func (m *Example) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Example) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Example) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SingleNested != nil {
		{
			size, err := m.SingleNested.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExample(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xca
	}
	if len(m.Nested) > 0 {
		for iNdEx := len(m.Nested) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Nested[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExample(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Example_Nested) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Example_Nested) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Example_Nested) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Ok != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Ok))
		i--
		dAtA[i] = 0x18
	}
	if m.Age != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExample(dAtA []byte, offset int, v uint64) int {
	offset -= sovExample(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Example) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if len(m.Nested) > 0 {
		for _, e := range m.Nested {
			l = e.Size()
			n += 1 + l + sovExample(uint64(l))
		}
	}
	if m.SingleNested != nil {
		l = m.SingleNested.Size()
		n += 2 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Example_Nested) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovExample(uint64(m.Age))
	}
	if m.Ok != 0 {
		n += 1 + sovExample(uint64(m.Ok))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExample(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExample(x uint64) (n int) {
	return sovExample(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Example) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
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
			return fmt.Errorf("proto: Example: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Example: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nested = append(m.Nested, &Example_Nested{})
			if err := m.Nested[len(m.Nested)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 25:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SingleNested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SingleNested == nil {
				m.SingleNested = &Example_Nested{}
			}
			if err := m.SingleNested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Example_Nested) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
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
			return fmt.Errorf("proto: Nested: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nested: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
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
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			m.Ok = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ok |= Example_Nested_DeepEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExample(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExample
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
					return 0, ErrIntOverflowExample
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
					return 0, ErrIntOverflowExample
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
				return 0, ErrInvalidLengthExample
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExample
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExample
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExample        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExample          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExample = fmt.Errorf("proto: unexpected end of group")
)