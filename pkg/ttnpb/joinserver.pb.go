// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/joinserver.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"

import context "context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SessionKeyRequest struct {
	// Join Server issued identifier for the session keys.
	SessionKeyID string `protobuf:"bytes,1,opt,name=session_key_id,json=sessionKeyId,proto3" json:"session_key_id,omitempty"`
	// LoRaWAN DevEUI.
	DevEUI               go_thethings_network_lorawan_stack_pkg_types.EUI64 `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3,customtype=go.thethings.network/lorawan-stack/pkg/types.EUI64" json:"dev_eui"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *SessionKeyRequest) Reset()      { *m = SessionKeyRequest{} }
func (*SessionKeyRequest) ProtoMessage() {}
func (*SessionKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_joinserver_83d5b3cfabd85c12, []int{0}
}
func (m *SessionKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SessionKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionKeyRequest.Merge(dst, src)
}
func (m *SessionKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *SessionKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionKeyRequest proto.InternalMessageInfo

func (m *SessionKeyRequest) GetSessionKeyID() string {
	if m != nil {
		return m.SessionKeyID
	}
	return ""
}

type NwkSKeysResponse struct {
	// The (encrypted) Forwarding Network Session Integrity Key (or Network Session Key in 1.0 compatibility mode).
	FNwkSIntKey KeyEnvelope `protobuf:"bytes,1,opt,name=f_nwk_s_int_key,json=fNwkSIntKey" json:"f_nwk_s_int_key"`
	// The (encrypted) Serving Network Session Integrity Key.
	SNwkSIntKey KeyEnvelope `protobuf:"bytes,2,opt,name=s_nwk_s_int_key,json=sNwkSIntKey" json:"s_nwk_s_int_key"`
	// The (encrypted) Network Session Encryption Key.
	NwkSEncKey           KeyEnvelope `protobuf:"bytes,3,opt,name=nwk_s_enc_key,json=nwkSEncKey" json:"nwk_s_enc_key"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NwkSKeysResponse) Reset()      { *m = NwkSKeysResponse{} }
func (*NwkSKeysResponse) ProtoMessage() {}
func (*NwkSKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_joinserver_83d5b3cfabd85c12, []int{1}
}
func (m *NwkSKeysResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NwkSKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NwkSKeysResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NwkSKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NwkSKeysResponse.Merge(dst, src)
}
func (m *NwkSKeysResponse) XXX_Size() int {
	return m.Size()
}
func (m *NwkSKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NwkSKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NwkSKeysResponse proto.InternalMessageInfo

func (m *NwkSKeysResponse) GetFNwkSIntKey() KeyEnvelope {
	if m != nil {
		return m.FNwkSIntKey
	}
	return KeyEnvelope{}
}

func (m *NwkSKeysResponse) GetSNwkSIntKey() KeyEnvelope {
	if m != nil {
		return m.SNwkSIntKey
	}
	return KeyEnvelope{}
}

func (m *NwkSKeysResponse) GetNwkSEncKey() KeyEnvelope {
	if m != nil {
		return m.NwkSEncKey
	}
	return KeyEnvelope{}
}

type AppSKeyResponse struct {
	// The (encrypted) Application Session Key.
	AppSKey              KeyEnvelope `protobuf:"bytes,1,opt,name=app_s_key,json=appSKey" json:"app_s_key"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppSKeyResponse) Reset()      { *m = AppSKeyResponse{} }
func (*AppSKeyResponse) ProtoMessage() {}
func (*AppSKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_joinserver_83d5b3cfabd85c12, []int{2}
}
func (m *AppSKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppSKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppSKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AppSKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppSKeyResponse.Merge(dst, src)
}
func (m *AppSKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *AppSKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppSKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppSKeyResponse proto.InternalMessageInfo

func (m *AppSKeyResponse) GetAppSKey() KeyEnvelope {
	if m != nil {
		return m.AppSKey
	}
	return KeyEnvelope{}
}

func init() {
	proto.RegisterType((*SessionKeyRequest)(nil), "ttn.lorawan.v3.SessionKeyRequest")
	golang_proto.RegisterType((*SessionKeyRequest)(nil), "ttn.lorawan.v3.SessionKeyRequest")
	proto.RegisterType((*NwkSKeysResponse)(nil), "ttn.lorawan.v3.NwkSKeysResponse")
	golang_proto.RegisterType((*NwkSKeysResponse)(nil), "ttn.lorawan.v3.NwkSKeysResponse")
	proto.RegisterType((*AppSKeyResponse)(nil), "ttn.lorawan.v3.AppSKeyResponse")
	golang_proto.RegisterType((*AppSKeyResponse)(nil), "ttn.lorawan.v3.AppSKeyResponse")
}
func (this *SessionKeyRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SessionKeyRequest)
	if !ok {
		that2, ok := that.(SessionKeyRequest)
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
	if this.SessionKeyID != that1.SessionKeyID {
		return false
	}
	if !this.DevEUI.Equal(that1.DevEUI) {
		return false
	}
	return true
}
func (this *NwkSKeysResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NwkSKeysResponse)
	if !ok {
		that2, ok := that.(NwkSKeysResponse)
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
	if !this.FNwkSIntKey.Equal(&that1.FNwkSIntKey) {
		return false
	}
	if !this.SNwkSIntKey.Equal(&that1.SNwkSIntKey) {
		return false
	}
	if !this.NwkSEncKey.Equal(&that1.NwkSEncKey) {
		return false
	}
	return true
}
func (this *AppSKeyResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AppSKeyResponse)
	if !ok {
		that2, ok := that.(AppSKeyResponse)
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
	if !this.AppSKey.Equal(&that1.AppSKey) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NsJs service

type NsJsClient interface {
	HandleJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	GetNwkSKeys(ctx context.Context, in *SessionKeyRequest, opts ...grpc.CallOption) (*NwkSKeysResponse, error)
}

type nsJsClient struct {
	cc *grpc.ClientConn
}

func NewNsJsClient(cc *grpc.ClientConn) NsJsClient {
	return &nsJsClient{cc}
}

func (c *nsJsClient) HandleJoin(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsJs/HandleJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsJsClient) GetNwkSKeys(ctx context.Context, in *SessionKeyRequest, opts ...grpc.CallOption) (*NwkSKeysResponse, error) {
	out := new(NwkSKeysResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsJs/GetNwkSKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NsJs service

type NsJsServer interface {
	HandleJoin(context.Context, *JoinRequest) (*JoinResponse, error)
	GetNwkSKeys(context.Context, *SessionKeyRequest) (*NwkSKeysResponse, error)
}

func RegisterNsJsServer(s *grpc.Server, srv NsJsServer) {
	s.RegisterService(&_NsJs_serviceDesc, srv)
}

func _NsJs_HandleJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsJsServer).HandleJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsJs/HandleJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsJsServer).HandleJoin(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsJs_GetNwkSKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsJsServer).GetNwkSKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsJs/GetNwkSKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsJsServer).GetNwkSKeys(ctx, req.(*SessionKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsJs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsJs",
	HandlerType: (*NsJsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleJoin",
			Handler:    _NsJs_HandleJoin_Handler,
		},
		{
			MethodName: "GetNwkSKeys",
			Handler:    _NsJs_GetNwkSKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/joinserver.proto",
}

// Client API for AsJs service

type AsJsClient interface {
	GetAppSKey(ctx context.Context, in *SessionKeyRequest, opts ...grpc.CallOption) (*AppSKeyResponse, error)
}

type asJsClient struct {
	cc *grpc.ClientConn
}

func NewAsJsClient(cc *grpc.ClientConn) AsJsClient {
	return &asJsClient{cc}
}

func (c *asJsClient) GetAppSKey(ctx context.Context, in *SessionKeyRequest, opts ...grpc.CallOption) (*AppSKeyResponse, error) {
	out := new(AppSKeyResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.AsJs/GetAppSKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AsJs service

type AsJsServer interface {
	GetAppSKey(context.Context, *SessionKeyRequest) (*AppSKeyResponse, error)
}

func RegisterAsJsServer(s *grpc.Server, srv AsJsServer) {
	s.RegisterService(&_AsJs_serviceDesc, srv)
}

func _AsJs_GetAppSKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsJsServer).GetAppSKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.AsJs/GetAppSKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsJsServer).GetAppSKey(ctx, req.(*SessionKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AsJs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.AsJs",
	HandlerType: (*AsJsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppSKey",
			Handler:    _AsJs_GetAppSKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/joinserver.proto",
}

func (m *SessionKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SessionKeyID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoinserver(dAtA, i, uint64(len(m.SessionKeyID)))
		i += copy(dAtA[i:], m.SessionKeyID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintJoinserver(dAtA, i, uint64(m.DevEUI.Size()))
	n1, err := m.DevEUI.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *NwkSKeysResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NwkSKeysResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintJoinserver(dAtA, i, uint64(m.FNwkSIntKey.Size()))
	n2, err := m.FNwkSIntKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x12
	i++
	i = encodeVarintJoinserver(dAtA, i, uint64(m.SNwkSIntKey.Size()))
	n3, err := m.SNwkSIntKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x1a
	i++
	i = encodeVarintJoinserver(dAtA, i, uint64(m.NwkSEncKey.Size()))
	n4, err := m.NwkSEncKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *AppSKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppSKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintJoinserver(dAtA, i, uint64(m.AppSKey.Size()))
	n5, err := m.AppSKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func encodeVarintJoinserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedSessionKeyRequest(r randyJoinserver, easy bool) *SessionKeyRequest {
	this := &SessionKeyRequest{}
	this.SessionKeyID = randStringJoinserver(r)
	v1 := go_thethings_network_lorawan_stack_pkg_types.NewPopulatedEUI64(r)
	this.DevEUI = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedNwkSKeysResponse(r randyJoinserver, easy bool) *NwkSKeysResponse {
	this := &NwkSKeysResponse{}
	v2 := NewPopulatedKeyEnvelope(r, easy)
	this.FNwkSIntKey = *v2
	v3 := NewPopulatedKeyEnvelope(r, easy)
	this.SNwkSIntKey = *v3
	v4 := NewPopulatedKeyEnvelope(r, easy)
	this.NwkSEncKey = *v4
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAppSKeyResponse(r randyJoinserver, easy bool) *AppSKeyResponse {
	this := &AppSKeyResponse{}
	v5 := NewPopulatedKeyEnvelope(r, easy)
	this.AppSKey = *v5
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyJoinserver interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneJoinserver(r randyJoinserver) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringJoinserver(r randyJoinserver) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneJoinserver(r)
	}
	return string(tmps)
}
func randUnrecognizedJoinserver(r randyJoinserver, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldJoinserver(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldJoinserver(dAtA []byte, r randyJoinserver, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateJoinserver(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulateJoinserver(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulateJoinserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateJoinserver(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateJoinserver(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateJoinserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateJoinserver(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *SessionKeyRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.SessionKeyID)
	if l > 0 {
		n += 1 + l + sovJoinserver(uint64(l))
	}
	l = m.DevEUI.Size()
	n += 1 + l + sovJoinserver(uint64(l))
	return n
}

func (m *NwkSKeysResponse) Size() (n int) {
	var l int
	_ = l
	l = m.FNwkSIntKey.Size()
	n += 1 + l + sovJoinserver(uint64(l))
	l = m.SNwkSIntKey.Size()
	n += 1 + l + sovJoinserver(uint64(l))
	l = m.NwkSEncKey.Size()
	n += 1 + l + sovJoinserver(uint64(l))
	return n
}

func (m *AppSKeyResponse) Size() (n int) {
	var l int
	_ = l
	l = m.AppSKey.Size()
	n += 1 + l + sovJoinserver(uint64(l))
	return n
}

func sovJoinserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJoinserver(x uint64) (n int) {
	return sovJoinserver((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *SessionKeyRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SessionKeyRequest{`,
		`SessionKeyID:` + fmt.Sprintf("%v", this.SessionKeyID) + `,`,
		`DevEUI:` + fmt.Sprintf("%v", this.DevEUI) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NwkSKeysResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NwkSKeysResponse{`,
		`FNwkSIntKey:` + strings.Replace(strings.Replace(this.FNwkSIntKey.String(), "KeyEnvelope", "KeyEnvelope", 1), `&`, ``, 1) + `,`,
		`SNwkSIntKey:` + strings.Replace(strings.Replace(this.SNwkSIntKey.String(), "KeyEnvelope", "KeyEnvelope", 1), `&`, ``, 1) + `,`,
		`NwkSEncKey:` + strings.Replace(strings.Replace(this.NwkSEncKey.String(), "KeyEnvelope", "KeyEnvelope", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AppSKeyResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AppSKeyResponse{`,
		`AppSKey:` + strings.Replace(strings.Replace(this.AppSKey.String(), "KeyEnvelope", "KeyEnvelope", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJoinserver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SessionKeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoinserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJoinserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoinserver
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DevEUI.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoinserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoinserver
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
func (m *NwkSKeysResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoinserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NwkSKeysResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NwkSKeysResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoinserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoinserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSEncKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoinserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NwkSEncKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoinserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoinserver
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
func (m *AppSKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoinserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AppSKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppSKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppSKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoinserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoinserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AppSKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoinserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoinserver
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
func skipJoinserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJoinserver
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
					return 0, ErrIntOverflowJoinserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJoinserver
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthJoinserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJoinserver
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipJoinserver(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthJoinserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJoinserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("lorawan-stack/api/joinserver.proto", fileDescriptor_joinserver_83d5b3cfabd85c12)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/joinserver.proto", fileDescriptor_joinserver_83d5b3cfabd85c12)
}

var fileDescriptor_joinserver_83d5b3cfabd85c12 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x4c, 0xdb, 0x40,
	0x14, 0xc6, 0xef, 0x28, 0x02, 0x71, 0x50, 0xa0, 0x9e, 0x50, 0x8a, 0x5e, 0x68, 0x26, 0x86, 0xe2,
	0x48, 0xa1, 0x62, 0xeb, 0x40, 0x94, 0x28, 0x0d, 0x91, 0x10, 0x4a, 0x84, 0x54, 0x55, 0xaa, 0x2c,
	0x27, 0x79, 0x04, 0xd7, 0xf4, 0xce, 0xcd, 0x5d, 0x1c, 0x79, 0x63, 0x64, 0xec, 0xd8, 0xb1, 0x6a,
	0x17, 0xd4, 0x89, 0x91, 0x91, 0x91, 0x91, 0x11, 0x75, 0x88, 0xf0, 0x79, 0x41, 0x9d, 0x18, 0x19,
	0x2b, 0x3b, 0xa9, 0x28, 0x86, 0x4a, 0xe9, 0xe6, 0xf7, 0xde, 0xe7, 0x9f, 0xbf, 0xf7, 0xc7, 0x2c,
	0x77, 0x20, 0xba, 0x76, 0xdf, 0xe6, 0x6b, 0x52, 0xd9, 0x2d, 0x37, 0x6f, 0x7b, 0x4e, 0xfe, 0x83,
	0x70, 0xb8, 0xc4, 0xae, 0x8f, 0x5d, 0xd3, 0xeb, 0x0a, 0x25, 0x8c, 0x79, 0xa5, 0xb8, 0x39, 0xd2,
	0x99, 0xfe, 0x7a, 0x66, 0xad, 0xe3, 0xa8, 0xfd, 0x5e, 0xd3, 0x6c, 0x89, 0x8f, 0xf9, 0x8e, 0xe8,
	0x88, 0x7c, 0x22, 0x6b, 0xf6, 0xf6, 0x92, 0x28, 0x09, 0x92, 0xa7, 0xe1, 0xeb, 0x99, 0xe5, 0xc7,
	0x3f, 0xf1, 0xef, 0xaa, 0x8b, 0x81, 0x1c, 0x56, 0x73, 0x3f, 0x28, 0x7b, 0xd6, 0x40, 0x29, 0x1d,
	0xc1, 0x6b, 0x18, 0xd4, 0xf1, 0x53, 0x0f, 0xa5, 0x32, 0x36, 0xd8, 0xbc, 0x1c, 0x26, 0x2d, 0x17,
	0x03, 0xcb, 0x69, 0x2f, 0xd1, 0x15, 0xba, 0x3a, 0x53, 0x5c, 0xd4, 0x83, 0xec, 0xdc, 0x9d, 0xbc,
	0x5a, 0xaa, 0xcf, 0xc9, 0xbb, 0xa8, 0x6d, 0xbc, 0x67, 0xd3, 0x6d, 0xf4, 0x2d, 0xec, 0x39, 0x4b,
	0x13, 0x2b, 0x74, 0x75, 0xae, 0x58, 0x3a, 0x1f, 0x64, 0xc9, 0xcf, 0x41, 0xb6, 0xd0, 0x11, 0xa6,
	0xda, 0x47, 0xb5, 0xef, 0xf0, 0x8e, 0x34, 0x39, 0xaa, 0xbe, 0xe8, 0xba, 0xf9, 0xfb, 0xce, 0x3c,
	0xb7, 0x93, 0x57, 0x81, 0x87, 0xd2, 0x2c, 0xef, 0x56, 0x37, 0x5e, 0xe9, 0x41, 0x76, 0xaa, 0x84,
	0x7e, 0x79, 0xb7, 0x5a, 0x9f, 0x6a, 0xa3, 0x5f, 0xee, 0x39, 0xb9, 0x5f, 0x94, 0x2d, 0x6e, 0xf7,
	0xdd, 0x46, 0x0d, 0x03, 0x59, 0x47, 0xe9, 0x09, 0x2e, 0xd1, 0xa8, 0xb0, 0x85, 0x3d, 0x8b, 0xf7,
	0x5d, 0x4b, 0x5a, 0x0e, 0x57, 0xb1, 0xdf, 0xc4, 0xec, 0x6c, 0xe1, 0xb9, 0x79, 0x7f, 0xac, 0x66,
	0x0d, 0x83, 0x32, 0xf7, 0xf1, 0x40, 0x78, 0x58, 0x9c, 0x8c, 0x8d, 0xd5, 0x67, 0xf7, 0x62, 0x5c,
	0x95, 0xab, 0x1a, 0x06, 0x31, 0x48, 0xa6, 0x40, 0x13, 0x63, 0x83, 0xe4, 0x5f, 0xa0, 0x12, 0x7b,
	0x3a, 0xc4, 0x20, 0x6f, 0x25, 0x98, 0x27, 0xe3, 0x62, 0x18, 0xef, 0xbb, 0x8d, 0x32, 0x6f, 0xd5,
	0x30, 0xc8, 0xed, 0xb0, 0x85, 0x4d, 0xcf, 0x6b, 0x24, 0x5b, 0x19, 0xb5, 0xfa, 0x9a, 0xcd, 0xd8,
	0x9e, 0x67, 0xc9, 0xff, 0x6b, 0x72, 0xda, 0x1e, 0x62, 0x0a, 0xdf, 0x29, 0x9b, 0xdc, 0x96, 0x5b,
	0xd2, 0xa8, 0x30, 0xf6, 0xc6, 0xe6, 0xed, 0x03, 0xdc, 0x12, 0x0e, 0x37, 0x1e, 0x20, 0xe2, 0xec,
	0xe8, 0x12, 0x32, 0xcb, 0x8f, 0x17, 0x47, 0x86, 0xea, 0x6c, 0xb6, 0x82, 0xea, 0xcf, 0x4a, 0x8c,
	0x17, 0x69, 0xf1, 0x83, 0xcb, 0xca, 0xac, 0xa4, 0x25, 0xe9, 0x7d, 0x16, 0xde, 0xb2, 0xc9, 0xcd,
	0xd8, 0xe4, 0x0e, 0x63, 0x15, 0x54, 0xa3, 0x11, 0x8c, 0x83, 0xce, 0xa6, 0x25, 0xa9, 0xf1, 0x15,
	0xbf, 0xd1, 0xf3, 0x10, 0xe8, 0x45, 0x08, 0xf4, 0x32, 0x04, 0x72, 0x15, 0x02, 0xb9, 0x0e, 0x81,
	0xdc, 0x84, 0x40, 0x6e, 0x43, 0xa0, 0x87, 0x1a, 0xe8, 0x91, 0x06, 0x72, 0xac, 0x81, 0x9e, 0x68,
	0x20, 0xa7, 0x1a, 0xc8, 0x99, 0x06, 0x72, 0xae, 0x81, 0x5e, 0x68, 0xa0, 0x97, 0x1a, 0xc8, 0x95,
	0x06, 0x7a, 0xad, 0x81, 0xdc, 0x68, 0xa0, 0xb7, 0x1a, 0xc8, 0x61, 0x04, 0xe4, 0x28, 0x02, 0xfa,
	0x39, 0x02, 0xf2, 0x25, 0x02, 0xfa, 0x35, 0x02, 0x72, 0x1c, 0x01, 0x39, 0x89, 0x80, 0x9e, 0x46,
	0x40, 0xcf, 0x22, 0xa0, 0xef, 0x5e, 0x8e, 0x7b, 0xfb, 0x8a, 0x7b, 0xcd, 0xe6, 0x54, 0xf2, 0x5f,
	0xae, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x09, 0xea, 0xec, 0x38, 0x04, 0x00, 0x00,
}
