// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/rights.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Right is the enum that defines all the different rights to do something in
// the network.
type Right int32

const (
	// RIGHT_INVALID is the invalid right and should not be used.
	// It can denote a parsing error.
	RIGHT_INVALID Right = 0
	// RIGHT_USER_PROFILE_READ is the right to read the user's profile.
	RIGHT_USER_PROFILE_READ Right = 1
	// RIGHT_USER_PROFILE_WRITE is the right to edit the user's profile.
	RIGHT_USER_PROFILE_WRITE Right = 2
	// RIGHT_USER_DELETE is the right to delete the user's account.
	RIGHT_USER_DELETE Right = 3
	// RIGHT_USER_AUTHORIZEDCLIENTS is the right to view and manage the authorized
	// clients of the user: list, authorize and de-authorize clients.
	RIGHT_USER_AUTHORIZEDCLIENTS Right = 4
	// RIGHT_USER_APPLICATIONS_LIST is the right to list all of the applications
	// the user is a collaborator of.
	RIGHT_USER_APPLICATIONS_LIST Right = 5
	// RIGHT_USER_APPLICATIONS_CREATE is the right to create an application in
	// the users account.
	RIGHT_USER_APPLICATIONS_CREATE Right = 6
	// RIGHT_USER_GATEWAYS_LIST is the right to list gateways the user is collaborator of.
	RIGHT_USER_GATEWAYS_LIST Right = 7
	// RIGHT_USER_GATEWAYS_CREATE is the right to register a gateway in the users account.
	RIGHT_USER_GATEWAYS_CREATE Right = 8
	// RIGHT_USER_CLIENTS is the right to create, list, view, update and delete
	// OAuth clients in the user's account.
	RIGHT_USER_CLIENTS Right = 9
	// RIGHT_USER_ADMIN is the right to performs actions that requires the user to be an admin.
	RIGHT_USER_ADMIN Right = 10
	// RIGHT_USER_KEYS is the right to view and manage user API keys.
	RIGHT_USER_KEYS Right = 11
	// RIGHT_APPLICATION_INFO is the right to see basic information about the application.
	// This does not include API keys.
	RIGHT_APPLICATION_INFO Right = 31
	// RIGHT_APPLICATION_SETTINGS_BASIC is the right to edit basic settings of the application.
	RIGHT_APPLICATION_SETTINGS_BASIC Right = 32
	// RIGHT_APPLICATION_SETTINGS_KEYS is the right to view and edit application
	// API keys.
	RIGHT_APPLICATION_SETTINGS_KEYS Right = 33
	// RIGHT_APPLICATION_SETTINGS_COLLABORATORS is the right to manage the
	// collaborators of the application.
	RIGHT_APPLICATION_SETTINGS_COLLABORATORS Right = 34
	// RIGHT_APPLICATION_DELETE is the right to delete the application.
	RIGHT_APPLICATION_DELETE Right = 35
	// RIGHT_APPLICATION_DEVICES_READ is the right to see the devices of an application.
	RIGHT_APPLICATION_DEVICES_READ Right = 36
	// RIGHT_APPLICATION_DEVICES_WRITE is the right to register devices to an application.
	RIGHT_APPLICATION_DEVICES_WRITE Right = 37
	// RIGHT_APPLICATION_TRAFFIC_READ is the right to read traffic of the
	// application (uplink and downlink).
	RIGHT_APPLICATION_TRAFFIC_READ Right = 38
	// RIGHT_APPLICATION_TRAFFIC_WRITE is the right to write traffic of the
	// application (uplink and downlink).
	RIGHT_APPLICATION_TRAFFIC_WRITE Right = 39
	// RIGHT_GATEWAY_INFO is the right to see basic information about the gateway.
	// This does not include API keys.
	RIGHT_GATEWAY_INFO Right = 51
	// RIGHT_GATEWAY_SETTINGS_BASIC is the right to manage basic settings of the gateway.
	RIGHT_GATEWAY_SETTINGS_BASIC Right = 52
	// RIGHT_GATEWAY_SETTINGS_KEYS is the right to view and edit API keys of the gateway.
	RIGHT_GATEWAY_SETTINGS_KEYS Right = 53
	// RIGHT_GATEWAY_SETTINGS_COLLABORATORS is the right to manage collaborators
	// of the gateway.
	RIGHT_GATEWAY_SETTINGS_COLLABORATORS Right = 54
	// RIGHT_GATEWAY_DELETE is the right to delete the gateway.
	RIGHT_GATEWAY_DELETE Right = 55
	// RIGHT_GATEWAY_TRAFFIC is the right to view traffic of the gateway.
	RIGHT_GATEWAY_TRAFFIC Right = 56
	// RIGHT_GATEWAY_STATUS is the right to view the status of the gateway.
	RIGHT_GATEWAY_STATUS Right = 57
	// RIGHT_GATEWAY_LOCATION is the right to view the location of the gateway.
	RIGHT_GATEWAY_LOCATION Right = 58
)

var Right_name = map[int32]string{
	0:  "RIGHT_INVALID",
	1:  "RIGHT_USER_PROFILE_READ",
	2:  "RIGHT_USER_PROFILE_WRITE",
	3:  "RIGHT_USER_DELETE",
	4:  "RIGHT_USER_AUTHORIZEDCLIENTS",
	5:  "RIGHT_USER_APPLICATIONS_LIST",
	6:  "RIGHT_USER_APPLICATIONS_CREATE",
	7:  "RIGHT_USER_GATEWAYS_LIST",
	8:  "RIGHT_USER_GATEWAYS_CREATE",
	9:  "RIGHT_USER_CLIENTS",
	10: "RIGHT_USER_ADMIN",
	11: "RIGHT_USER_KEYS",
	31: "RIGHT_APPLICATION_INFO",
	32: "RIGHT_APPLICATION_SETTINGS_BASIC",
	33: "RIGHT_APPLICATION_SETTINGS_KEYS",
	34: "RIGHT_APPLICATION_SETTINGS_COLLABORATORS",
	35: "RIGHT_APPLICATION_DELETE",
	36: "RIGHT_APPLICATION_DEVICES_READ",
	37: "RIGHT_APPLICATION_DEVICES_WRITE",
	38: "RIGHT_APPLICATION_TRAFFIC_READ",
	39: "RIGHT_APPLICATION_TRAFFIC_WRITE",
	51: "RIGHT_GATEWAY_INFO",
	52: "RIGHT_GATEWAY_SETTINGS_BASIC",
	53: "RIGHT_GATEWAY_SETTINGS_KEYS",
	54: "RIGHT_GATEWAY_SETTINGS_COLLABORATORS",
	55: "RIGHT_GATEWAY_DELETE",
	56: "RIGHT_GATEWAY_TRAFFIC",
	57: "RIGHT_GATEWAY_STATUS",
	58: "RIGHT_GATEWAY_LOCATION",
}
var Right_value = map[string]int32{
	"RIGHT_INVALID":                            0,
	"RIGHT_USER_PROFILE_READ":                  1,
	"RIGHT_USER_PROFILE_WRITE":                 2,
	"RIGHT_USER_DELETE":                        3,
	"RIGHT_USER_AUTHORIZEDCLIENTS":             4,
	"RIGHT_USER_APPLICATIONS_LIST":             5,
	"RIGHT_USER_APPLICATIONS_CREATE":           6,
	"RIGHT_USER_GATEWAYS_LIST":                 7,
	"RIGHT_USER_GATEWAYS_CREATE":               8,
	"RIGHT_USER_CLIENTS":                       9,
	"RIGHT_USER_ADMIN":                         10,
	"RIGHT_USER_KEYS":                          11,
	"RIGHT_APPLICATION_INFO":                   31,
	"RIGHT_APPLICATION_SETTINGS_BASIC":         32,
	"RIGHT_APPLICATION_SETTINGS_KEYS":          33,
	"RIGHT_APPLICATION_SETTINGS_COLLABORATORS": 34,
	"RIGHT_APPLICATION_DELETE":                 35,
	"RIGHT_APPLICATION_DEVICES_READ":           36,
	"RIGHT_APPLICATION_DEVICES_WRITE":          37,
	"RIGHT_APPLICATION_TRAFFIC_READ":           38,
	"RIGHT_APPLICATION_TRAFFIC_WRITE":          39,
	"RIGHT_GATEWAY_INFO":                       51,
	"RIGHT_GATEWAY_SETTINGS_BASIC":             52,
	"RIGHT_GATEWAY_SETTINGS_KEYS":              53,
	"RIGHT_GATEWAY_SETTINGS_COLLABORATORS":     54,
	"RIGHT_GATEWAY_DELETE":                     55,
	"RIGHT_GATEWAY_TRAFFIC":                    56,
	"RIGHT_GATEWAY_STATUS":                     57,
	"RIGHT_GATEWAY_LOCATION":                   58,
}

func (Right) EnumDescriptor() ([]byte, []int) { return fileDescriptorRights, []int{0} }

func init() {
	proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
	golang_proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
}
func (x Right) String() string {
	s, ok := Right_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}

var fileDescriptorRights = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x3d, 0x4c, 0x14, 0x41,
	0x18, 0x9d, 0x51, 0x7e, 0x74, 0x8c, 0x71, 0x18, 0x01, 0xf1, 0x20, 0x1f, 0x08, 0xa8, 0x68, 0x94,
	0x33, 0xc1, 0xff, 0x6e, 0xd8, 0x9b, 0x83, 0x89, 0xeb, 0x2e, 0xd9, 0x19, 0x20, 0xd0, 0x5c, 0x3c,
	0x83, 0x77, 0x17, 0x22, 0x77, 0x81, 0x45, 0x5b, 0x4a, 0x4a, 0x4a, 0x3b, 0x4d, 0x6c, 0x28, 0x29,
	0x29, 0x29, 0x29, 0x29, 0x29, 0xd9, 0xd9, 0x86, 0x92, 0x92, 0xd2, 0xb0, 0xbb, 0x97, 0xdb, 0x83,
	0x3b, 0xac, 0x76, 0xe7, 0x7b, 0xef, 0x7b, 0xf3, 0xcd, 0x7b, 0x93, 0x21, 0xaf, 0x4a, 0x15, 0xbf,
	0xbc, 0x59, 0x9c, 0xfc, 0x5a, 0xfd, 0x9e, 0xd5, 0xe5, 0x15, 0x5d, 0xae, 0xac, 0x95, 0x36, 0x9c,
	0x15, 0xff, 0x67, 0x75, 0x7d, 0x35, 0xeb, 0xfb, 0x6b, 0xd9, 0x2f, 0xb5, 0x4a, 0x76, 0xbd, 0x52,
	0x2a, 0xfb, 0x1b, 0x93, 0xb5, 0xf5, 0xaa, 0x5f, 0x65, 0x5d, 0xbe, 0xbf, 0x36, 0xf9, 0x63, 0x2a,
	0xf3, 0x32, 0xd5, 0x59, 0xaa, 0x96, 0xaa, 0xd9, 0x08, 0x2e, 0x6e, 0x7e, 0x8b, 0x56, 0xd1, 0x22,
	0xfa, 0x8b, 0xdb, 0x9e, 0xef, 0x74, 0x93, 0x4e, 0xef, 0x42, 0x87, 0xf5, 0x90, 0xbb, 0x9e, 0x9c,
	0x99, 0xd5, 0x05, 0xe9, 0x2c, 0x70, 0x5b, 0xe6, 0x28, 0x62, 0x83, 0xe4, 0x41, 0x5c, 0x9a, 0x57,
	0xc2, 0x2b, 0xcc, 0x79, 0x6e, 0x5e, 0xda, 0xa2, 0xe0, 0x09, 0x9e, 0xa3, 0x98, 0x0d, 0x91, 0x81,
	0x16, 0xe0, 0xa2, 0x27, 0xb5, 0xa0, 0x37, 0x58, 0x1f, 0xe9, 0x49, 0xa1, 0x39, 0x61, 0x0b, 0x2d,
	0xe8, 0x4d, 0x36, 0x42, 0x86, 0x52, 0x65, 0x3e, 0xaf, 0x67, 0x5d, 0x4f, 0x2e, 0x8b, 0x9c, 0x65,
	0x4b, 0xe1, 0x68, 0x45, 0x3b, 0x2e, 0x33, 0xe6, 0xe6, 0x6c, 0x69, 0x71, 0x2d, 0x5d, 0x47, 0x15,
	0x6c, 0xa9, 0x34, 0xed, 0x64, 0xa3, 0x04, 0xda, 0x31, 0x2c, 0x4f, 0x70, 0x2d, 0x68, 0xd7, 0xa5,
	0xe1, 0x66, 0xb8, 0x16, 0x8b, 0x7c, 0x29, 0x51, 0xe8, 0x66, 0x40, 0x32, 0xad, 0xd0, 0xa4, 0xfb,
	0x16, 0xeb, 0x27, 0x2c, 0x85, 0xd7, 0x67, 0xbb, 0xcd, 0x7a, 0x09, 0x4d, 0xef, 0x9c, 0xfb, 0x2c,
	0x1d, 0x4a, 0xd8, 0x7d, 0x72, 0x2f, 0x55, 0xfd, 0x24, 0x96, 0x14, 0xbd, 0xc3, 0x32, 0xa4, 0x3f,
	0x2e, 0xa6, 0xe6, 0x2b, 0x48, 0x27, 0xef, 0xd2, 0x61, 0x36, 0x4e, 0x46, 0xae, 0x62, 0x4a, 0x68,
	0x2d, 0x9d, 0x19, 0x55, 0x98, 0xe6, 0x4a, 0x5a, 0x74, 0x84, 0x8d, 0x91, 0xe1, 0x6b, 0x58, 0xd1,
	0x36, 0x8f, 0xd8, 0x0b, 0x32, 0x71, 0x0d, 0xc9, 0x72, 0x6d, 0x9b, 0x4f, 0xbb, 0x1e, 0xd7, 0xae,
	0xa7, 0xe8, 0x68, 0xc3, 0x95, 0x34, 0x3b, 0xc9, 0x66, 0xac, 0xe1, 0x6b, 0x33, 0xba, 0x20, 0x2d,
	0xa1, 0xe2, 0xd0, 0xc7, 0x5b, 0x0f, 0x55, 0xe7, 0xc4, 0xd9, 0x3f, 0x6e, 0x2d, 0xa4, 0x3d, 0x9e,
	0xcf, 0x4b, 0x2b, 0x16, 0x7a, 0xd2, 0x5a, 0xa8, 0xce, 0x89, 0x85, 0x9e, 0x36, 0x72, 0x48, 0x22,
	0x8a, 0x0d, 0x9c, 0x6a, 0xdc, 0x91, 0x7a, 0xfd, 0x92, 0x79, 0xaf, 0xd9, 0x30, 0x19, 0x6c, 0xc3,
	0x88, 0x8c, 0x7b, 0xc3, 0x26, 0xc8, 0x78, 0x1b, 0x42, 0xb3, 0x69, 0x6f, 0xd9, 0x00, 0xe9, 0x6d,
	0x66, 0x26, 0x86, 0xbd, 0x63, 0x0f, 0x49, 0x5f, 0x33, 0x92, 0xcc, 0x4f, 0xdf, 0x5f, 0x6d, 0x52,
	0x9a, 0xeb, 0x79, 0x45, 0x3f, 0x34, 0x2e, 0x46, 0x1d, 0xb1, 0xdd, 0xf8, 0xf4, 0xf4, 0x63, 0xa6,
	0x63, 0xfb, 0x2f, 0xa0, 0xe9, 0xdf, 0xf8, 0x30, 0x00, 0x7c, 0x14, 0x00, 0x3e, 0x0e, 0x00, 0x9f,
	0x04, 0x80, 0x4f, 0x03, 0x40, 0x67, 0x01, 0xa0, 0xf3, 0x00, 0xf0, 0x96, 0x01, 0xb4, 0x6d, 0x00,
	0xed, 0x1a, 0xc0, 0x7b, 0x06, 0xd0, 0xbe, 0x01, 0x7c, 0x60, 0x00, 0x1f, 0x1a, 0xc0, 0x47, 0x06,
	0xf0, 0xb1, 0x01, 0x74, 0x62, 0x00, 0x9f, 0x1a, 0x40, 0x67, 0x06, 0xf0, 0xb9, 0x01, 0xb4, 0x15,
	0x02, 0xda, 0x0e, 0x01, 0xef, 0x84, 0x80, 0x7e, 0x85, 0x80, 0xff, 0x84, 0x80, 0x76, 0x43, 0x40,
	0x7b, 0x21, 0xe0, 0xfd, 0x10, 0xf0, 0x41, 0x08, 0x78, 0xf9, 0xd9, 0xff, 0x9e, 0x9c, 0xda, 0x6a,
	0xe9, 0xe2, 0x5b, 0x2b, 0x16, 0xbb, 0xa2, 0xb7, 0x63, 0xea, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x01, 0x18, 0xbd, 0xa6, 0x04, 0x00, 0x00,
}
