// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/join.proto

package ttnpb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_v3_pkg_types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type JoinRequest struct {
	RawPayload         []byte                                                  `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	Payload            *Message                                                `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	DevAddr            go_thethings_network_lorawan_stack_v3_pkg_types.DevAddr `protobuf:"bytes,3,opt,name=dev_addr,json=devAddr,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.DevAddr" json:"dev_addr"`
	SelectedMACVersion MACVersion                                              `protobuf:"varint,4,opt,name=selected_mac_version,json=selectedMacVersion,proto3,enum=ttn.lorawan.v3.MACVersion" json:"selected_mac_version,omitempty"`
	NetID              go_thethings_network_lorawan_stack_v3_pkg_types.NetID   `protobuf:"bytes,5,opt,name=net_id,json=netId,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.NetID" json:"net_id"`
	DownlinkSettings   DLSettings                                              `protobuf:"bytes,6,opt,name=downlink_settings,json=downlinkSettings,proto3" json:"downlink_settings"`
	RxDelay            RxDelay                                                 `protobuf:"varint,7,opt,name=rx_delay,json=rxDelay,proto3,enum=ttn.lorawan.v3.RxDelay" json:"rx_delay,omitempty"`
	// Optional CFList.
	CFList         *CFList  `protobuf:"bytes,8,opt,name=cf_list,json=cfList,proto3" json:"cf_list,omitempty"`
	CorrelationIDs []string `protobuf:"bytes,10,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	// Consumed airtime for the transmission of the join request. Calculated by Network Server using the RawPayload size and the transmission settings.
	ConsumedAirtime      *time.Duration `protobuf:"bytes,11,opt,name=consumed_airtime,json=consumedAirtime,proto3,stdduration" json:"consumed_airtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JoinRequest) Reset()      { *m = JoinRequest{} }
func (*JoinRequest) ProtoMessage() {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{0}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return m.Size()
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinRequest) GetPayload() *Message {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *JoinRequest) GetSelectedMACVersion() MACVersion {
	if m != nil {
		return m.SelectedMACVersion
	}
	return MAC_UNKNOWN
}

func (m *JoinRequest) GetDownlinkSettings() DLSettings {
	if m != nil {
		return m.DownlinkSettings
	}
	return DLSettings{}
}

func (m *JoinRequest) GetRxDelay() RxDelay {
	if m != nil {
		return m.RxDelay
	}
	return RX_DELAY_0
}

func (m *JoinRequest) GetCFList() *CFList {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *JoinRequest) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

func (m *JoinRequest) GetConsumedAirtime() *time.Duration {
	if m != nil {
		return m.ConsumedAirtime
	}
	return nil
}

type JoinResponse struct {
	RawPayload           []byte `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	SessionKeys          `protobuf:"bytes,2,opt,name=session_keys,json=sessionKeys,proto3,embedded=session_keys" json:"session_keys"`
	Lifetime             time.Duration `protobuf:"bytes,3,opt,name=lifetime,proto3,stdduration" json:"lifetime"`
	CorrelationIDs       []string      `protobuf:"bytes,4,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *JoinResponse) Reset()      { *m = JoinResponse{} }
func (*JoinResponse) ProtoMessage() {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{1}
}
func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinResponse) GetLifetime() time.Duration {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *JoinResponse) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	golang_proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
	golang_proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
}

func init() { proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14)
}

var fileDescriptor_dd69b88666e72e14 = []byte{
	// 831 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x6c, 0xdb, 0x46,
	0x18, 0xbd, 0x93, 0xf5, 0xe7, 0x93, 0xe1, 0x28, 0x44, 0x91, 0xb0, 0x6e, 0x71, 0x74, 0x3d, 0x09,
	0x05, 0x4c, 0xa2, 0x31, 0x8a, 0x02, 0x6d, 0x81, 0xc0, 0xb4, 0x90, 0xc2, 0x69, 0x12, 0x04, 0x34,
	0xda, 0x21, 0x40, 0x41, 0x9c, 0x79, 0x67, 0x9a, 0x15, 0xcd, 0x53, 0x79, 0x27, 0xc9, 0xca, 0x14,
	0x64, 0x0a, 0x3a, 0x15, 0x1d, 0x8a, 0x8c, 0x59, 0x0a, 0x64, 0xcc, 0xe8, 0x31, 0xa3, 0x47, 0x8f,
	0x41, 0x07, 0x35, 0x3a, 0x2e, 0x19, 0x33, 0x06, 0x9a, 0x0a, 0x1d, 0xa9, 0xca, 0x8e, 0x3c, 0x34,
	0x9d, 0xf8, 0xf1, 0xbe, 0xf7, 0x3d, 0xbc, 0xef, 0xdd, 0x3d, 0xf4, 0x69, 0xcc, 0x53, 0x32, 0x20,
	0xc9, 0xa6, 0x90, 0x24, 0xe8, 0x38, 0xa4, 0x1b, 0x39, 0x3f, 0xf3, 0x28, 0xb1, 0xbb, 0x29, 0x97,
	0xdc, 0x58, 0x95, 0x32, 0xb1, 0x0b, 0x84, 0xdd, 0xdf, 0x5a, 0xdb, 0x0e, 0x23, 0x79, 0xd8, 0xdb,
	0xb7, 0x03, 0x7e, 0xe4, 0xb0, 0xa4, 0xcf, 0x87, 0xdd, 0x94, 0x1f, 0x0f, 0x1d, 0x0d, 0x0e, 0x36,
	0x43, 0x96, 0x6c, 0xf6, 0x49, 0x1c, 0x51, 0x22, 0x99, 0xb3, 0x50, 0xe4, 0x94, 0x6b, 0x9b, 0xe7,
	0x28, 0x42, 0x1e, 0xf2, 0x7c, 0x78, 0xbf, 0x77, 0xa0, 0xff, 0xf4, 0x8f, 0xae, 0x0a, 0x38, 0x0e,
	0x39, 0x0f, 0x63, 0x36, 0x47, 0xd1, 0x5e, 0x4a, 0x64, 0xc4, 0x0b, 0x85, 0x6b, 0x97, 0xe8, 0xef,
	0xb0, 0xa1, 0x28, 0xba, 0xd6, 0x62, 0x77, 0xb6, 0x8d, 0x06, 0x6c, 0x3c, 0xae, 0xa2, 0xc6, 0x6d,
	0x1e, 0x25, 0x1e, 0xfb, 0xa5, 0xc7, 0x84, 0x34, 0x5a, 0xa8, 0x91, 0x92, 0x81, 0xdf, 0x25, 0xc3,
	0x98, 0x13, 0x6a, 0xc2, 0x75, 0xd8, 0x5a, 0x71, 0x6b, 0x13, 0xb7, 0xfc, 0xb0, 0x74, 0x78, 0xdd,
	0x43, 0x29, 0x19, 0xdc, 0xcf, 0x5b, 0xc6, 0x17, 0xa8, 0x36, 0x43, 0x95, 0xd6, 0x61, 0xab, 0x71,
	0xe3, 0xba, 0x7d, 0xd1, 0x2c, 0xfb, 0x2e, 0x13, 0x82, 0x84, 0xcc, 0x9b, 0xe1, 0x8c, 0x07, 0xa8,
	0x4e, 0x59, 0xdf, 0x27, 0x94, 0xa6, 0xe6, 0x92, 0x66, 0xbe, 0x79, 0x3a, 0xb2, 0xc0, 0x5f, 0x23,
	0xeb, 0xab, 0x90, 0xdb, 0xf2, 0x90, 0xc9, 0xc3, 0x28, 0x09, 0x85, 0x9d, 0x30, 0x39, 0xe0, 0x69,
	0xc7, 0xb9, 0x28, 0xbe, 0xbf, 0xe5, 0x74, 0x3b, 0xa1, 0x23, 0x87, 0x5d, 0x26, 0xec, 0x36, 0xeb,
	0x6f, 0x53, 0x9a, 0x7a, 0x35, 0x9a, 0x17, 0x06, 0x45, 0x1f, 0x09, 0x16, 0xb3, 0x40, 0x32, 0xea,
	0x1f, 0x91, 0xc0, 0xef, 0xb3, 0x54, 0x44, 0x3c, 0x31, 0xcb, 0xeb, 0xb0, 0xb5, 0x7a, 0x63, 0x6d,
	0x41, 0xdb, 0xf6, 0xce, 0x8f, 0x39, 0xc2, 0xbd, 0xa6, 0x46, 0x96, 0xb1, 0x57, 0xcc, 0xce, 0xcf,
	0x3d, 0x63, 0xc6, 0x77, 0x97, 0x04, 0xc5, 0x99, 0xf1, 0x13, 0xaa, 0x26, 0x4c, 0xfa, 0x11, 0x35,
	0x2b, 0x5a, 0xff, 0xad, 0x42, 0xff, 0x97, 0x1f, 0xaa, 0xff, 0x1e, 0x93, 0xbb, 0x6d, 0x35, 0xb2,
	0x2a, 0xba, 0xf0, 0x2a, 0x09, 0x93, 0xbb, 0xd4, 0xf8, 0x01, 0x5d, 0xa5, 0x7c, 0x90, 0xc4, 0x51,
	0xd2, 0xf1, 0x05, 0x93, 0x72, 0xca, 0x66, 0x56, 0xb5, 0xbb, 0x0b, 0x1b, 0xb4, 0xef, 0xec, 0x15,
	0x08, 0x77, 0x65, 0xe2, 0x56, 0x7e, 0x85, 0xa5, 0x26, 0x9c, 0xaa, 0xf1, 0x9a, 0x33, 0x8a, 0x59,
	0xdf, 0xf8, 0x16, 0xd5, 0xd3, 0x63, 0x9f, 0xb2, 0x98, 0x0c, 0xcd, 0x9a, 0xf6, 0x63, 0xe1, 0xae,
	0xbc, 0xe3, 0xf6, 0xb4, 0xed, 0xd6, 0x27, 0x6e, 0xe5, 0xf1, 0x94, 0xca, 0xab, 0xa5, 0xf9, 0x91,
	0xf1, 0x0d, 0xaa, 0x05, 0x07, 0x7e, 0x1c, 0x09, 0x69, 0xd6, 0xb5, 0x94, 0x6b, 0xef, 0x0f, 0xef,
	0xdc, 0xba, 0x13, 0x09, 0xe9, 0x22, 0x35, 0xb2, 0xaa, 0x79, 0xed, 0x55, 0x83, 0x83, 0xe9, 0xd7,
	0xf8, 0x0e, 0x5d, 0x09, 0x78, 0x9a, 0xb2, 0x58, 0xbf, 0x59, 0x3f, 0xa2, 0xc2, 0x44, 0xeb, 0x4b,
	0xad, 0x65, 0x17, 0x4f, 0xdc, 0xe5, 0xdf, 0x61, 0x75, 0xa3, 0x9c, 0x96, 0x4c, 0xaa, 0x46, 0xd6,
	0xea, 0xce, 0x1c, 0xb6, 0xdb, 0x16, 0xde, 0xea, 0xb9, 0xb1, 0x5d, 0x2a, 0x8c, 0x7b, 0xa8, 0x19,
	0xf0, 0x44, 0xf4, 0x8e, 0x18, 0xf5, 0x49, 0x94, 0xca, 0xe8, 0x88, 0x99, 0x0d, 0x2d, 0xe7, 0x63,
	0x3b, 0x8f, 0x88, 0x3d, 0x8b, 0x88, 0xdd, 0x2e, 0x22, 0xe2, 0xd6, 0x4f, 0x47, 0x16, 0x7c, 0xfa,
	0xb7, 0x05, 0xbd, 0x2b, 0xb3, 0xe1, 0xed, 0x7c, 0xf6, 0xeb, 0xf2, 0xc9, 0x33, 0x0b, 0xdc, 0x2e,
	0xd7, 0x97, 0x9b, 0x68, 0xe3, 0x8f, 0x12, 0x5a, 0xc9, 0x43, 0x20, 0xba, 0x3c, 0x11, 0xcc, 0xf8,
	0xfc, 0xb2, 0x14, 0x2c, 0x4f, 0xdc, 0xea, 0xc3, 0x72, 0xf3, 0xaa, 0xf9, 0xd9, 0x85, 0x1c, 0xdc,
	0x47, 0x2b, 0x82, 0x89, 0xe9, 0xeb, 0xf0, 0xa7, 0xc1, 0x2b, 0xc2, 0xf0, 0xc9, 0xfb, 0x1e, 0xed,
	0xe5, 0x98, 0xef, 0xd9, 0x50, 0xb8, 0xcd, 0xf3, 0xf7, 0x75, 0x36, 0xb2, 0xa0, 0xd7, 0x10, 0xf3,
	0xb6, 0x71, 0x13, 0xd5, 0xe3, 0xe8, 0x80, 0xe9, 0x15, 0x97, 0xfe, 0xcb, 0x8a, 0x40, 0xaf, 0xf8,
	0xef, 0xd0, 0x65, 0xa6, 0x97, 0xff, 0x8f, 0xe9, 0xee, 0x9f, 0xf0, 0x74, 0x8c, 0xe1, 0xd9, 0x18,
	0xc3, 0x57, 0x63, 0x0c, 0x5e, 0x8f, 0x31, 0x78, 0x33, 0xc6, 0xe0, 0xed, 0x18, 0x83, 0x77, 0x63,
	0x0c, 0x1f, 0x29, 0x0c, 0x9f, 0x28, 0x0c, 0x9e, 0x2b, 0x0c, 0x5f, 0x28, 0x0c, 0x4e, 0x14, 0x06,
	0x2f, 0x15, 0x06, 0xa7, 0x0a, 0xc3, 0x33, 0x85, 0xe1, 0x2b, 0x85, 0xc1, 0x6b, 0x85, 0xe1, 0x1b,
	0x85, 0xc1, 0x5b, 0x85, 0xe1, 0x3b, 0x85, 0xc1, 0xa3, 0x0c, 0x83, 0x27, 0x19, 0x86, 0xbf, 0x65,
	0x18, 0x3c, 0xcd, 0x30, 0x7c, 0x96, 0x61, 0xf0, 0x3c, 0xc3, 0xe0, 0x45, 0x86, 0xe1, 0x49, 0x86,
	0xe1, 0xcb, 0x0c, 0xc3, 0x07, 0xce, 0x07, 0xa4, 0x49, 0x26, 0xdd, 0xfd, 0xfd, 0xaa, 0xf6, 0x65,
	0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x9f, 0xcb, 0x41, 0xcd, 0x05, 0x00, 0x00,
}

func (this *JoinRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinRequest)
	if !ok {
		that2, ok := that.(JoinRequest)
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
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.Payload.Equal(that1.Payload) {
		return false
	}
	if !this.DevAddr.Equal(that1.DevAddr) {
		return false
	}
	if this.SelectedMACVersion != that1.SelectedMACVersion {
		return false
	}
	if !this.NetID.Equal(that1.NetID) {
		return false
	}
	if !this.DownlinkSettings.Equal(&that1.DownlinkSettings) {
		return false
	}
	if this.RxDelay != that1.RxDelay {
		return false
	}
	if !this.CFList.Equal(that1.CFList) {
		return false
	}
	if len(this.CorrelationIDs) != len(that1.CorrelationIDs) {
		return false
	}
	for i := range this.CorrelationIDs {
		if this.CorrelationIDs[i] != that1.CorrelationIDs[i] {
			return false
		}
	}
	if this.ConsumedAirtime != nil && that1.ConsumedAirtime != nil {
		if *this.ConsumedAirtime != *that1.ConsumedAirtime {
			return false
		}
	} else if this.ConsumedAirtime != nil {
		return false
	} else if that1.ConsumedAirtime != nil {
		return false
	}
	return true
}
func (this *JoinResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinResponse)
	if !ok {
		that2, ok := that.(JoinResponse)
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
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.SessionKeys.Equal(&that1.SessionKeys) {
		return false
	}
	if this.Lifetime != that1.Lifetime {
		return false
	}
	if len(this.CorrelationIDs) != len(that1.CorrelationIDs) {
		return false
	}
	for i := range this.CorrelationIDs {
		if this.CorrelationIDs[i] != that1.CorrelationIDs[i] {
			return false
		}
	}
	return true
}
func (m *JoinRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConsumedAirtime != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.ConsumedAirtime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ConsumedAirtime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintJoin(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.CorrelationIDs) > 0 {
		for iNdEx := len(m.CorrelationIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CorrelationIDs[iNdEx])
			copy(dAtA[i:], m.CorrelationIDs[iNdEx])
			i = encodeVarintJoin(dAtA, i, uint64(len(m.CorrelationIDs[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.CFList != nil {
		{
			size, err := m.CFList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintJoin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.RxDelay != 0 {
		i = encodeVarintJoin(dAtA, i, uint64(m.RxDelay))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.DownlinkSettings.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintJoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.NetID.Size()
		i -= size
		if _, err := m.NetID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintJoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.SelectedMACVersion != 0 {
		i = encodeVarintJoin(dAtA, i, uint64(m.SelectedMACVersion))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.DevAddr.Size()
		i -= size
		if _, err := m.DevAddr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintJoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintJoin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RawPayload) > 0 {
		i -= len(m.RawPayload)
		copy(dAtA[i:], m.RawPayload)
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CorrelationIDs) > 0 {
		for iNdEx := len(m.CorrelationIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CorrelationIDs[iNdEx])
			copy(dAtA[i:], m.CorrelationIDs[iNdEx])
			i = encodeVarintJoin(dAtA, i, uint64(len(m.CorrelationIDs[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	n5, err5 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Lifetime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Lifetime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintJoin(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SessionKeys.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintJoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RawPayload) > 0 {
		i -= len(m.RawPayload)
		copy(dAtA[i:], m.RawPayload)
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintJoin(dAtA []byte, offset int, v uint64) int {
	offset -= sovJoin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedJoinResponse(r randyJoin, easy bool) *JoinResponse {
	this := &JoinResponse{}
	v1 := r.Intn(100)
	this.RawPayload = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.RawPayload[i] = byte(r.Intn(256))
	}
	v2 := NewPopulatedSessionKeys(r, easy)
	this.SessionKeys = *v2
	v3 := github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.Lifetime = *v3
	v4 := r.Intn(10)
	this.CorrelationIDs = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.CorrelationIDs[i] = randStringJoin(r)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyJoin interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneJoin(r randyJoin) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringJoin(r randyJoin) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneJoin(r)
	}
	return string(tmps)
}
func randUnrecognizedJoin(r randyJoin, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldJoin(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldJoin(dAtA []byte, r randyJoin, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateJoin(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *JoinRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.DevAddr.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.SelectedMACVersion != 0 {
		n += 1 + sovJoin(uint64(m.SelectedMACVersion))
	}
	l = m.NetID.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = m.DownlinkSettings.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.RxDelay != 0 {
		n += 1 + sovJoin(uint64(m.RxDelay))
	}
	if m.CFList != nil {
		l = m.CFList.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	if m.ConsumedAirtime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ConsumedAirtime)
		n += 1 + l + sovJoin(uint64(l))
	}
	return n
}

func (m *JoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.SessionKeys.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Lifetime)
	n += 1 + l + sovJoin(uint64(l))
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	return n
}

func sovJoin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJoin(x uint64) (n int) {
	return sovJoin((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *JoinRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinRequest{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`Payload:` + strings.Replace(fmt.Sprintf("%v", this.Payload), "Message", "Message", 1) + `,`,
		`DevAddr:` + fmt.Sprintf("%v", this.DevAddr) + `,`,
		`SelectedMACVersion:` + fmt.Sprintf("%v", this.SelectedMACVersion) + `,`,
		`NetID:` + fmt.Sprintf("%v", this.NetID) + `,`,
		`DownlinkSettings:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DownlinkSettings), "DLSettings", "DLSettings", 1), `&`, ``, 1) + `,`,
		`RxDelay:` + fmt.Sprintf("%v", this.RxDelay) + `,`,
		`CFList:` + strings.Replace(fmt.Sprintf("%v", this.CFList), "CFList", "CFList", 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`ConsumedAirtime:` + strings.Replace(fmt.Sprintf("%v", this.ConsumedAirtime), "Duration", "types.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JoinResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinResponse{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`SessionKeys:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.SessionKeys), "SessionKeys", "SessionKeys", 1), `&`, ``, 1) + `,`,
		`Lifetime:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Lifetime), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJoin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *JoinRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: JoinRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &Message{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedMACVersion", wireType)
			}
			m.SelectedMACVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SelectedMACVersion |= MACVersion(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownlinkSettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DownlinkSettings.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxDelay", wireType)
			}
			m.RxDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RxDelay |= RxDelay(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CFList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CFList == nil {
				m.CFList = &CFList{}
			}
			if err := m.CFList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CorrelationIDs = append(m.CorrelationIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumedAirtime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsumedAirtime == nil {
				m.ConsumedAirtime = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.ConsumedAirtime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJoin
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
func (m *JoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: JoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SessionKeys.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lifetime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Lifetime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CorrelationIDs = append(m.CorrelationIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJoin
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
func skipJoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJoin
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
					return 0, ErrIntOverflowJoin
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
					return 0, ErrIntOverflowJoin
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
				return 0, ErrInvalidLengthJoin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJoin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJoin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJoin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJoin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJoin = fmt.Errorf("proto: unexpected end of group")
)
