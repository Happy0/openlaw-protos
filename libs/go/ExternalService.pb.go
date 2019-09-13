// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ExternalService.proto

package integration_framework_openlaw

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The possible statuses of the response after the execution is
// terminated.
type ExecuteResponse_Status int32

const (
	ExecuteResponse_FAILURE ExecuteResponse_Status = 0
	ExecuteResponse_SUCCESS ExecuteResponse_Status = 1
)

var ExecuteResponse_Status_name = map[int32]string{
	0: "FAILURE",
	1: "SUCCESS",
}

var ExecuteResponse_Status_value = map[string]int32{
	"FAILURE": 0,
	"SUCCESS": 1,
}

func (x ExecuteResponse_Status) String() string {
	return proto.EnumName(ExecuteResponse_Status_name, int32(x))
}

func (ExecuteResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cbc7f9d798314ef1, []int{3, 0}
}

// The Ethereum Address response message.
type EthereumAddressResponse struct {
	//The Ethereum public address generated by the External Service.
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumAddressResponse) Reset()         { *m = EthereumAddressResponse{} }
func (m *EthereumAddressResponse) String() string { return proto.CompactTextString(m) }
func (*EthereumAddressResponse) ProtoMessage()    {}
func (*EthereumAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbc7f9d798314ef1, []int{0}
}

func (m *EthereumAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthereumAddressResponse.Unmarshal(m, b)
}
func (m *EthereumAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthereumAddressResponse.Marshal(b, m, deterministic)
}
func (m *EthereumAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumAddressResponse.Merge(m, src)
}
func (m *EthereumAddressResponse) XXX_Size() int {
	return xxx_messageInfo_EthereumAddressResponse.Size(m)
}
func (m *EthereumAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumAddressResponse proto.InternalMessageInfo

func (m *EthereumAddressResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// The Markup Interface response message.
type MarkupInterfaceResponse struct {
	// The Markup Interface definition that matches the
	// OpenLaw Input/Output structures.
	Definition           string   `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkupInterfaceResponse) Reset()         { *m = MarkupInterfaceResponse{} }
func (m *MarkupInterfaceResponse) String() string { return proto.CompactTextString(m) }
func (*MarkupInterfaceResponse) ProtoMessage()    {}
func (*MarkupInterfaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbc7f9d798314ef1, []int{1}
}

func (m *MarkupInterfaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkupInterfaceResponse.Unmarshal(m, b)
}
func (m *MarkupInterfaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkupInterfaceResponse.Marshal(b, m, deterministic)
}
func (m *MarkupInterfaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkupInterfaceResponse.Merge(m, src)
}
func (m *MarkupInterfaceResponse) XXX_Size() int {
	return xxx_messageInfo_MarkupInterfaceResponse.Size(m)
}
func (m *MarkupInterfaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkupInterfaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarkupInterfaceResponse proto.InternalMessageInfo

func (m *MarkupInterfaceResponse) GetDefinition() string {
	if m != nil {
		return m.Definition
	}
	return ""
}

// The Request message to be processed by the External Service.
type ExecuteRequest struct {
	// The contractId or flowId string value generated by OpenLaw.
	// e.g: 1a86c0ab-7895-497c-babb-a3c089df1203
	CallerId string `protobuf:"bytes,1,opt,name=callerId,proto3" json:"callerId,omitempty"`
	// The actionId string value which is an arbitrary content that
	// represents the action that happened in OpenLaw VM.
	ActionId string `protobuf:"bytes,2,opt,name=actionId,proto3" json:"actionId,omitempty"`
	// The requestId string value that represents the request in
	// Integrator Framework.
	RequestId string `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	// The inputJson string value that matches the Input type defined
	// in the Markup Interface. It must be converted from string to
	// json value in order to access its properties.
	InputJson string `protobuf:"bytes,4,opt,name=inputJson,proto3" json:"inputJson,omitempty"`
	// The ethereum public address from the External Service as
	// provided in the EthreumAddressResponse.
	ServicePublicAddress string   `protobuf:"bytes,5,opt,name=servicePublicAddress,proto3" json:"servicePublicAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteRequest) Reset()         { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()    {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbc7f9d798314ef1, []int{2}
}

func (m *ExecuteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteRequest.Unmarshal(m, b)
}
func (m *ExecuteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteRequest.Merge(m, src)
}
func (m *ExecuteRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteRequest.Size(m)
}
func (m *ExecuteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteRequest proto.InternalMessageInfo

func (m *ExecuteRequest) GetCallerId() string {
	if m != nil {
		return m.CallerId
	}
	return ""
}

func (m *ExecuteRequest) GetActionId() string {
	if m != nil {
		return m.ActionId
	}
	return ""
}

func (m *ExecuteRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *ExecuteRequest) GetInputJson() string {
	if m != nil {
		return m.InputJson
	}
	return ""
}

func (m *ExecuteRequest) GetServicePublicAddress() string {
	if m != nil {
		return m.ServicePublicAddress
	}
	return ""
}

// The Response message obtained from the executed call.
type ExecuteResponse struct {
	// The contractId or flowId string value generated by OpenLaw.
	// e.g: 1a86c0ab-7895-497c-babb-a3c089df1203
	CallerId string `protobuf:"bytes,1,opt,name=callerId,proto3" json:"callerId,omitempty"`
	// The actionId string value which is an arbitrary content that
	// represents the action that happened in OpenLaw VM.
	ActionId string `protobuf:"bytes,2,opt,name=actionId,proto3" json:"actionId,omitempty"`
	// The requestId string value that represents the request in
	// Integrator Framework.
	RequestId string `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	// The outputJson as string value that matches the Output type
	// defined in the Markup Interface. All values of the json must
	// be String and the json must have no spaces.
	// e.g:
	// - {"currency":"BRL","price":"29.03","update":"2019-09-10T16:31:00.000Z"}
	// - {"param1":"value1","param2":"value2","paramN":"valueN"}
	OutputJson string `protobuf:"bytes,4,opt,name=outputJson,proto3" json:"outputJson,omitempty"`
	// The output signature string which represents the signature of
	// the outputJson field. The stringified version of the outputJson
	// must be used for signature.
	OutputSignature string `protobuf:"bytes,5,opt,name=outputSignature,proto3" json:"outputSignature,omitempty"`
	// The status of the execution. If not provided, the default is FAILURE.
	Status ExecuteResponse_Status `protobuf:"varint,6,opt,name=status,proto3,enum=integration.framework.openlaw.ExecuteResponse_Status" json:"status,omitempty"`
	// The message returned by the External Service to indicate if the
	// execution was completed with success or not. It can be an error
	// message as well.
	Message string `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	// The ethereum public address from the External Service as provided
	// in the EthereumAddressResponse.
	ServicePublicAddress string   `protobuf:"bytes,8,opt,name=servicePublicAddress,proto3" json:"servicePublicAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteResponse) Reset()         { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()    {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbc7f9d798314ef1, []int{3}
}

func (m *ExecuteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteResponse.Unmarshal(m, b)
}
func (m *ExecuteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteResponse.Merge(m, src)
}
func (m *ExecuteResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteResponse.Size(m)
}
func (m *ExecuteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteResponse proto.InternalMessageInfo

func (m *ExecuteResponse) GetCallerId() string {
	if m != nil {
		return m.CallerId
	}
	return ""
}

func (m *ExecuteResponse) GetActionId() string {
	if m != nil {
		return m.ActionId
	}
	return ""
}

func (m *ExecuteResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *ExecuteResponse) GetOutputJson() string {
	if m != nil {
		return m.OutputJson
	}
	return ""
}

func (m *ExecuteResponse) GetOutputSignature() string {
	if m != nil {
		return m.OutputSignature
	}
	return ""
}

func (m *ExecuteResponse) GetStatus() ExecuteResponse_Status {
	if m != nil {
		return m.Status
	}
	return ExecuteResponse_FAILURE
}

func (m *ExecuteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ExecuteResponse) GetServicePublicAddress() string {
	if m != nil {
		return m.ServicePublicAddress
	}
	return ""
}

// The Empty request message to indicate no data in the request.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbc7f9d798314ef1, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("integration.framework.openlaw.ExecuteResponse_Status", ExecuteResponse_Status_name, ExecuteResponse_Status_value)
	proto.RegisterType((*EthereumAddressResponse)(nil), "integration.framework.openlaw.EthereumAddressResponse")
	proto.RegisterType((*MarkupInterfaceResponse)(nil), "integration.framework.openlaw.MarkupInterfaceResponse")
	proto.RegisterType((*ExecuteRequest)(nil), "integration.framework.openlaw.ExecuteRequest")
	proto.RegisterType((*ExecuteResponse)(nil), "integration.framework.openlaw.ExecuteResponse")
	proto.RegisterType((*Empty)(nil), "integration.framework.openlaw.Empty")
}

func init() { proto.RegisterFile("ExternalService.proto", fileDescriptor_cbc7f9d798314ef1) }

var fileDescriptor_cbc7f9d798314ef1 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xde, 0x54, 0xbb, 0x69, 0x8f, 0xd0, 0x96, 0x41, 0x69, 0x58, 0xb4, 0x94, 0xc1, 0x8b, 0xbd,
	0x31, 0x17, 0x5b, 0x14, 0xbc, 0x2c, 0x25, 0x4a, 0xc4, 0x82, 0x24, 0xf4, 0x01, 0xa6, 0xc9, 0xd9,
	0x75, 0x6c, 0x76, 0x26, 0xce, 0x9c, 0xb1, 0xf5, 0xcd, 0x7c, 0x25, 0x6f, 0x7d, 0x02, 0x49, 0x66,
	0xba, 0xb6, 0xa1, 0x3f, 0x0a, 0xf6, 0xf2, 0xfb, 0xbe, 0x9c, 0x33, 0xe7, 0xe7, 0x3b, 0x81, 0x67,
	0xd9, 0x05, 0xa1, 0x51, 0xa2, 0x29, 0xd1, 0x7c, 0x93, 0x15, 0xa6, 0xad, 0xd1, 0xa4, 0xd9, 0x0b,
	0xa9, 0x08, 0x17, 0x46, 0x90, 0xd4, 0x2a, 0x9d, 0x1b, 0xb1, 0xc4, 0x73, 0x6d, 0xce, 0x52, 0xdd,
	0xa2, 0x6a, 0xc4, 0x39, 0x3f, 0x80, 0xdd, 0x8c, 0x3e, 0xa3, 0x41, 0xb7, 0x3c, 0xac, 0x6b, 0x83,
	0xd6, 0x16, 0x68, 0x5b, 0xad, 0x2c, 0xb2, 0x04, 0x62, 0xe1, 0xa9, 0x24, 0xda, 0x8f, 0xa6, 0x9b,
	0xc5, 0x25, 0xe4, 0x6f, 0x61, 0xf7, 0x58, 0x98, 0x33, 0xd7, 0xe6, 0x8a, 0xd0, 0xcc, 0x45, 0x85,
	0xab, 0xa0, 0x3d, 0x80, 0x1a, 0xe7, 0x52, 0xc9, 0xee, 0xbd, 0x10, 0x77, 0x85, 0xe1, 0x3f, 0x22,
	0xd8, 0xca, 0x2e, 0xb0, 0x72, 0x84, 0x05, 0x7e, 0x75, 0x68, 0x89, 0x4d, 0x60, 0xa3, 0x12, 0x4d,
	0x83, 0x26, 0xaf, 0x43, 0xc0, 0x0a, 0x77, 0x9a, 0xa8, 0xba, 0xc0, 0xbc, 0x4e, 0xd6, 0xbc, 0x76,
	0x89, 0xd9, 0x73, 0xd8, 0x34, 0x3e, 0x45, 0x5e, 0x27, 0x8f, 0x7a, 0xf1, 0x0f, 0xd1, 0xa9, 0x52,
	0xb5, 0x8e, 0x3e, 0x58, 0xad, 0x92, 0xc7, 0x5e, 0x5d, 0x11, 0x6c, 0x06, 0x4f, 0xad, 0x1f, 0xd3,
	0x27, 0x77, 0xda, 0xc8, 0x2a, 0xf4, 0x9e, 0xac, 0xf7, 0x1f, 0xde, 0xa8, 0xf1, 0x5f, 0x6b, 0xb0,
	0xbd, 0x2a, 0x3d, 0xb4, 0xfb, 0x30, 0xb5, 0xef, 0x01, 0x68, 0x47, 0xd7, 0x8b, 0xbf, 0xc2, 0xb0,
	0x29, 0x6c, 0x7b, 0x54, 0xca, 0x85, 0x12, 0xe4, 0x0c, 0x86, 0xc2, 0x87, 0x34, 0x3b, 0x86, 0xb1,
	0x25, 0x41, 0xce, 0x26, 0xe3, 0xfd, 0x68, 0xba, 0x35, 0x7b, 0x9d, 0xde, 0x69, 0x87, 0x74, 0xd0,
	0x5f, 0x5a, 0xf6, 0xc1, 0x45, 0x48, 0xd2, 0x59, 0x62, 0x89, 0xd6, 0x8a, 0x05, 0x26, 0xb1, 0xb7,
	0x44, 0x80, 0xb7, 0x0e, 0x74, 0xe3, 0x8e, 0x81, 0x72, 0x18, 0xfb, 0xfc, 0xec, 0x09, 0xc4, 0xef,
	0x0e, 0xf3, 0x8f, 0x27, 0x45, 0xb6, 0x33, 0xea, 0x40, 0x79, 0x72, 0x74, 0x94, 0x95, 0xe5, 0x4e,
	0xc4, 0x63, 0x58, 0xcf, 0x96, 0x2d, 0x7d, 0x9f, 0xfd, 0xec, 0xa7, 0x7f, 0xcd, 0xe1, 0x8c, 0x80,
	0xbd, 0x47, 0x1a, 0xf8, 0x97, 0xbd, 0xbc, 0xaf, 0xc7, 0x2e, 0xdf, 0xe4, 0xcd, 0x7d, 0x5f, 0xdd,
	0x7c, 0x15, 0x7c, 0x14, 0x5e, 0x1d, 0x1c, 0xc0, 0x7f, 0x7a, 0xf5, 0x96, 0xb3, 0xe2, 0x23, 0xf6,
	0x05, 0xe2, 0xb0, 0x1c, 0xf6, 0xea, 0x6f, 0x97, 0xd8, 0xfb, 0x69, 0x92, 0xfe, 0xdb, 0xce, 0xf9,
	0xe8, 0x74, 0xdc, 0xff, 0x3a, 0x0e, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x60, 0xde, 0x2b, 0x51,
	0x53, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExternalServiceClient is the client API for ExternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExternalServiceClient interface {
	// Gets the ethereum public address from the service which is
	// used to verify signatures from the service.
	GetEthereumAddress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EthereumAddressResponse, error)
	// Gets the server Markup Interface definition to be used
	// by OpenLaw template editor.
	GetMarkupInterface(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MarkupInterfaceResponse, error)
	// Executes the request from OpenLaw Integrator Framework
	// and waits for the External Service response.
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type externalServiceClient struct {
	cc *grpc.ClientConn
}

func NewExternalServiceClient(cc *grpc.ClientConn) ExternalServiceClient {
	return &externalServiceClient{cc}
}

func (c *externalServiceClient) GetEthereumAddress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EthereumAddressResponse, error) {
	out := new(EthereumAddressResponse)
	err := c.cc.Invoke(ctx, "/integration.framework.openlaw.ExternalService/GetEthereumAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) GetMarkupInterface(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MarkupInterfaceResponse, error) {
	out := new(MarkupInterfaceResponse)
	err := c.cc.Invoke(ctx, "/integration.framework.openlaw.ExternalService/GetMarkupInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalServiceClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/integration.framework.openlaw.ExternalService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalServiceServer is the server API for ExternalService service.
type ExternalServiceServer interface {
	// Gets the ethereum public address from the service which is
	// used to verify signatures from the service.
	GetEthereumAddress(context.Context, *Empty) (*EthereumAddressResponse, error)
	// Gets the server Markup Interface definition to be used
	// by OpenLaw template editor.
	GetMarkupInterface(context.Context, *Empty) (*MarkupInterfaceResponse, error)
	// Executes the request from OpenLaw Integrator Framework
	// and waits for the External Service response.
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

// UnimplementedExternalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExternalServiceServer struct {
}

func (*UnimplementedExternalServiceServer) GetEthereumAddress(ctx context.Context, req *Empty) (*EthereumAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEthereumAddress not implemented")
}
func (*UnimplementedExternalServiceServer) GetMarkupInterface(ctx context.Context, req *Empty) (*MarkupInterfaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarkupInterface not implemented")
}
func (*UnimplementedExternalServiceServer) Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterExternalServiceServer(s *grpc.Server, srv ExternalServiceServer) {
	s.RegisterService(&_ExternalService_serviceDesc, srv)
}

func _ExternalService_GetEthereumAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).GetEthereumAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration.framework.openlaw.ExternalService/GetEthereumAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).GetEthereumAddress(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_GetMarkupInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).GetMarkupInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration.framework.openlaw.ExternalService/GetMarkupInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).GetMarkupInterface(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration.framework.openlaw.ExternalService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServiceServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExternalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "integration.framework.openlaw.ExternalService",
	HandlerType: (*ExternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEthereumAddress",
			Handler:    _ExternalService_GetEthereumAddress_Handler,
		},
		{
			MethodName: "GetMarkupInterface",
			Handler:    _ExternalService_GetMarkupInterface_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _ExternalService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ExternalService.proto",
}
