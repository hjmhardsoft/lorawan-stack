// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/secrets.proto

package ttnpb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Secret contains a secret value. It also contains the ID of the Encryption key used to encrypt it.
type Secret struct {
	// ID of the Key used to encrypt the secret.
	KeyID                string   `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Secret) Reset()      { *m = Secret{} }
func (*Secret) ProtoMessage() {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c9d796b7f7ca235, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return m.Size()
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *Secret) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Secret)(nil), "ttn.lorawan.v3.Secret")
	golang_proto.RegisterType((*Secret)(nil), "ttn.lorawan.v3.Secret")
}

func init() { proto.RegisterFile("lorawan-stack/api/secrets.proto", fileDescriptor_8c9d796b7f7ca235) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/secrets.proto", fileDescriptor_8c9d796b7f7ca235)
}

var fileDescriptor_8c9d796b7f7ca235 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x6f, 0xda, 0x50,
	0x10, 0xc6, 0xef, 0x5a, 0x19, 0x15, 0xab, 0xaa, 0x2a, 0x26, 0xc4, 0x70, 0xa0, 0x4e, 0x2c, 0xf6,
	0x1b, 0xf8, 0x0b, 0x6a, 0x75, 0xa1, 0xdd, 0xe8, 0xd6, 0xa5, 0xb2, 0xcd, 0x8b, 0xb1, 0x4c, 0xfc,
	0x2c, 0xfb, 0x61, 0xe2, 0x4c, 0x8c, 0x8c, 0x19, 0x33, 0x66, 0x89, 0xc4, 0xc8, 0xc8, 0xc8, 0xc8,
	0xc8, 0xc8, 0x14, 0xe1, 0xf7, 0x16, 0x46, 0x46, 0x94, 0x29, 0x92, 0x9d, 0x21, 0x51, 0xb6, 0xef,
	0x3b, 0xfd, 0xbe, 0xd3, 0xdd, 0x67, 0x76, 0xa7, 0x22, 0x75, 0xe7, 0x6e, 0x6c, 0x65, 0xd2, 0xf5,
	0x23, 0xe6, 0x26, 0x21, 0xcb, 0xb8, 0x9f, 0x72, 0x99, 0xd9, 0x49, 0x2a, 0xa4, 0x68, 0x7d, 0x93,
	0x32, 0xb6, 0x5f, 0x21, 0x3b, 0x1f, 0x74, 0xac, 0x20, 0x94, 0x93, 0x99, 0x67, 0xfb, 0xe2, 0x9a,
	0x05, 0x22, 0x10, 0xac, 0xc2, 0xbc, 0xd9, 0x55, 0xe5, 0x2a, 0x53, 0xa9, 0x3a, 0xde, 0xf9, 0xf9,
	0x06, 0xe7, 0x71, 0x2e, 0x8a, 0x24, 0x15, 0x37, 0x45, 0x1d, 0xf2, 0xad, 0x80, 0xc7, 0x56, 0xee,
	0x4e, 0xc3, 0xb1, 0x2b, 0x39, 0xfb, 0x20, 0xea, 0x15, 0x3f, 0x7e, 0x9b, 0x8d, 0xbf, 0xd5, 0x49,
	0xad, 0x9e, 0xd9, 0x88, 0x78, 0xf1, 0x3f, 0x1c, 0xb7, 0xb1, 0x87, 0xfd, 0xa6, 0xd3, 0x54, 0x4f,
	0x5d, 0xe3, 0x0f, 0x2f, 0x86, 0xbf, 0x46, 0x46, 0xc4, 0x8b, 0xe1, 0xb8, 0x45, 0xa6, 0x91, 0xbb,
	0xd3, 0x19, 0x6f, 0x7f, 0xea, 0x61, 0xff, 0xab, 0xf3, 0xe5, 0xd9, 0x31, 0x6e, 0x3f, 0xb7, 0x17,
	0xdf, 0x47, 0xf5, 0xd8, 0x79, 0xc4, 0x5d, 0x49, 0xb8, 0x2f, 0x09, 0x0f, 0x25, 0xc1, 0xb1, 0x24,
	0x38, 0x95, 0x04, 0xe7, 0x92, 0xe0, 0x52, 0x12, 0x2e, 0x14, 0xe1, 0x52, 0x11, 0xac, 0x14, 0xe1,
	0x5a, 0x11, 0x6c, 0x14, 0xc1, 0x56, 0x11, 0xec, 0x14, 0xe1, 0x5e, 0x11, 0x1e, 0x14, 0xc1, 0x51,
	0x11, 0x9e, 0x14, 0xc1, 0x59, 0x11, 0x5e, 0x14, 0xc1, 0x42, 0x13, 0x2c, 0x35, 0xe1, 0x9d, 0x26,
	0xb8, 0xd7, 0x84, 0x0f, 0x9a, 0x60, 0xa5, 0x09, 0xd6, 0x9a, 0x70, 0xa3, 0x09, 0xb7, 0x9a, 0xf0,
	0x1f, 0x0b, 0x84, 0x2d, 0x27, 0x5c, 0x4e, 0xc2, 0x38, 0xc8, 0xec, 0x98, 0xcb, 0xb9, 0x48, 0x23,
	0xf6, 0xbe, 0xfa, 0x7c, 0xc0, 0x92, 0x28, 0x60, 0x52, 0xc6, 0x89, 0xe7, 0x35, 0xaa, 0xd7, 0x07,
	0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x9a, 0x18, 0xc7, 0x9f, 0x01, 0x00, 0x00,
}

func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
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
	if this.KeyID != that1.KeyID {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	return true
}
func (m *Secret) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Secret) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Secret) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintSecrets(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintSecrets(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSecrets(dAtA []byte, offset int, v uint64) int {
	offset -= sovSecrets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedSecret(r randySecrets, easy bool) *Secret {
	this := &Secret{}
	this.KeyID = randStringSecrets(r)
	v1 := r.Intn(100)
	this.Value = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Value[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySecrets interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSecrets(r randySecrets) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSecrets(r randySecrets) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneSecrets(r)
	}
	return string(tmps)
}
func randUnrecognizedSecrets(r randySecrets, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSecrets(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSecrets(dAtA []byte, r randySecrets, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSecrets(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateSecrets(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateSecrets(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSecrets(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSecrets(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSecrets(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSecrets(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Secret) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovSecrets(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSecrets(uint64(l))
	}
	return n
}

func sovSecrets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSecrets(x uint64) (n int) {
	return sovSecrets((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *Secret) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Secret{`,
		`KeyID:` + fmt.Sprintf("%v", this.KeyID) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSecrets(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Secret) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecrets
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
			return fmt.Errorf("proto: Secret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Secret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecrets
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
				return ErrInvalidLengthSecrets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSecrets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecrets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSecrets
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSecrets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSecrets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecrets
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSecrets
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
func skipSecrets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSecrets
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
					return 0, ErrIntOverflowSecrets
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
					return 0, ErrIntOverflowSecrets
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
				return 0, ErrInvalidLengthSecrets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSecrets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSecrets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSecrets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSecrets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSecrets = fmt.Errorf("proto: unexpected end of group")
)