// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/lifecycle/chaincode_definition.proto

package lifecycle

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ChaincodeEndorsementInfo is (most) everything the peer needs to know in order
// to execute a chaincode
type ChaincodeEndorsementInfo struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	InitRequired         bool     `protobuf:"varint,2,opt,name=init_required,json=initRequired,proto3" json:"init_required,omitempty"`
	EndorsementPlugin    string   `protobuf:"bytes,3,opt,name=endorsement_plugin,json=endorsementPlugin,proto3" json:"endorsement_plugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeEndorsementInfo) Reset()         { *m = ChaincodeEndorsementInfo{} }
func (m *ChaincodeEndorsementInfo) String() string { return proto.CompactTextString(m) }
func (*ChaincodeEndorsementInfo) ProtoMessage()    {}
func (*ChaincodeEndorsementInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0faa93bbd697c66, []int{0}
}

func (m *ChaincodeEndorsementInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeEndorsementInfo.Unmarshal(m, b)
}
func (m *ChaincodeEndorsementInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeEndorsementInfo.Marshal(b, m, deterministic)
}
func (m *ChaincodeEndorsementInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeEndorsementInfo.Merge(m, src)
}
func (m *ChaincodeEndorsementInfo) XXX_Size() int {
	return xxx_messageInfo_ChaincodeEndorsementInfo.Size(m)
}
func (m *ChaincodeEndorsementInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeEndorsementInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeEndorsementInfo proto.InternalMessageInfo

func (m *ChaincodeEndorsementInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChaincodeEndorsementInfo) GetInitRequired() bool {
	if m != nil {
		return m.InitRequired
	}
	return false
}

func (m *ChaincodeEndorsementInfo) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}

// ValidationInfo is (most) everything the peer needs to know in order
// to validate a transaction
type ChaincodeValidationInfo struct {
	ValidationPlugin     string   `protobuf:"bytes,1,opt,name=validation_plugin,json=validationPlugin,proto3" json:"validation_plugin,omitempty"`
	ValidationParameter  []byte   `protobuf:"bytes,2,opt,name=validation_parameter,json=validationParameter,proto3" json:"validation_parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaincodeValidationInfo) Reset()         { *m = ChaincodeValidationInfo{} }
func (m *ChaincodeValidationInfo) String() string { return proto.CompactTextString(m) }
func (*ChaincodeValidationInfo) ProtoMessage()    {}
func (*ChaincodeValidationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0faa93bbd697c66, []int{1}
}

func (m *ChaincodeValidationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeValidationInfo.Unmarshal(m, b)
}
func (m *ChaincodeValidationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeValidationInfo.Marshal(b, m, deterministic)
}
func (m *ChaincodeValidationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeValidationInfo.Merge(m, src)
}
func (m *ChaincodeValidationInfo) XXX_Size() int {
	return xxx_messageInfo_ChaincodeValidationInfo.Size(m)
}
func (m *ChaincodeValidationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeValidationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeValidationInfo proto.InternalMessageInfo

func (m *ChaincodeValidationInfo) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}

func (m *ChaincodeValidationInfo) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeEndorsementInfo)(nil), "lifecycle.ChaincodeEndorsementInfo")
	proto.RegisterType((*ChaincodeValidationInfo)(nil), "lifecycle.ChaincodeValidationInfo")
}

func init() {
	proto.RegisterFile("peer/lifecycle/chaincode_definition.proto", fileDescriptor_f0faa93bbd697c66)
}

var fileDescriptor_f0faa93bbd697c66 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x4b, 0xb5, 0x40,
	0x14, 0xc6, 0xf1, 0x7d, 0xa1, 0xba, 0xc3, 0x0d, 0xba, 0x53, 0x90, 0xcb, 0xcb, 0x6d, 0x73, 0xa3,
	0x1c, 0x89, 0xa0, 0x0f, 0x50, 0xb4, 0x68, 0x17, 0x2e, 0x5a, 0xb4, 0x91, 0x71, 0xe6, 0xa8, 0x07,
	0x74, 0x8e, 0x1d, 0xf5, 0x82, 0xdf, 0xa0, 0x8f, 0x1d, 0x6a, 0x7a, 0x6d, 0x39, 0xe7, 0xf9, 0xc3,
	0x8f, 0x67, 0xc4, 0x6d, 0x05, 0xc0, 0x61, 0x81, 0x29, 0x98, 0xce, 0x14, 0x10, 0x9a, 0x5c, 0xa3,
	0x33, 0x64, 0x21, 0xb6, 0x90, 0xa2, 0xc3, 0x06, 0xc9, 0xa9, 0x8a, 0xa9, 0x21, 0xb9, 0x9a, 0x5d,
	0xbb, 0x6f, 0x4f, 0xf8, 0x2f, 0x93, 0xf3, 0xd5, 0x59, 0xe2, 0x1a, 0x4a, 0x70, 0xcd, 0x9b, 0x4b,
	0x49, 0xfa, 0xe2, 0xf4, 0x00, 0x5c, 0x23, 0x39, 0xdf, 0xdb, 0x7a, 0xfb, 0x55, 0x34, 0x3d, 0xe5,
	0x8d, 0x38, 0xef, 0x2b, 0x63, 0x86, 0xaf, 0x16, 0x19, 0xac, 0xff, 0x6f, 0xeb, 0xed, 0xcf, 0xa2,
	0x75, 0x7f, 0x8c, 0x7e, 0x6f, 0x32, 0x10, 0x12, 0x8e, 0x8d, 0x71, 0x55, 0xb4, 0x19, 0x3a, 0xff,
	0xff, 0xd0, 0xb4, 0x59, 0x28, 0xef, 0x83, 0xb0, 0xeb, 0xc4, 0xf5, 0x4c, 0xf2, 0xa1, 0x0b, 0xb4,
	0xba, 0x47, 0x1e, 0x40, 0xee, 0xc4, 0xe6, 0x30, 0x5f, 0xa6, 0xa2, 0x11, 0xe9, 0xe2, 0x28, 0x8c,
	0x3d, 0xf2, 0x41, 0x5c, 0x2d, 0xcd, 0x9a, 0x75, 0x09, 0x0d, 0xf0, 0x80, 0xb8, 0x8e, 0x2e, 0x17,
	0xfe, 0x49, 0x7a, 0x4e, 0xc5, 0x3d, 0x71, 0xa6, 0xf2, 0xae, 0x02, 0x2e, 0xc0, 0x66, 0xc0, 0x2a,
	0xd5, 0x09, 0xa3, 0x19, 0x07, 0xab, 0x55, 0xbf, 0xad, 0x9a, 0x57, 0xfb, 0x7c, 0xca, 0xb0, 0xc9,
	0xdb, 0x44, 0x19, 0x2a, 0xc3, 0x45, 0x28, 0x1c, 0x43, 0xc1, 0x18, 0x0a, 0x32, 0x0a, 0xff, 0xfe,
	0x49, 0x72, 0x32, 0x28, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x94, 0x68, 0xee, 0xac,
	0x01, 0x00, 0x00,
}
