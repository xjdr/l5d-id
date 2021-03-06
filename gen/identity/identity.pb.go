// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity.proto

package identity // import "github.com/olix0r/l5d-id/gen/identity"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CertifyRequest struct {
	Token                     []byte   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	CertificateSigningRequest []byte   `protobuf:"bytes,2,opt,name=certificate_signing_request,json=certificateSigningRequest,proto3" json:"certificate_signing_request,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *CertifyRequest) Reset()         { *m = CertifyRequest{} }
func (m *CertifyRequest) String() string { return proto.CompactTextString(m) }
func (*CertifyRequest) ProtoMessage()    {}
func (*CertifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_b6ab172143c7462d, []int{0}
}
func (m *CertifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertifyRequest.Unmarshal(m, b)
}
func (m *CertifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertifyRequest.Marshal(b, m, deterministic)
}
func (dst *CertifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertifyRequest.Merge(dst, src)
}
func (m *CertifyRequest) XXX_Size() int {
	return xxx_messageInfo_CertifyRequest.Size(m)
}
func (m *CertifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertifyRequest proto.InternalMessageInfo

func (m *CertifyRequest) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *CertifyRequest) GetCertificateSigningRequest() []byte {
	if m != nil {
		return m.CertificateSigningRequest
	}
	return nil
}

type CertifyResponse struct {
	Certificate          []byte   `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertifyResponse) Reset()         { *m = CertifyResponse{} }
func (m *CertifyResponse) String() string { return proto.CompactTextString(m) }
func (*CertifyResponse) ProtoMessage()    {}
func (*CertifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_b6ab172143c7462d, []int{1}
}
func (m *CertifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertifyResponse.Unmarshal(m, b)
}
func (m *CertifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertifyResponse.Marshal(b, m, deterministic)
}
func (dst *CertifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertifyResponse.Merge(dst, src)
}
func (m *CertifyResponse) XXX_Size() int {
	return xxx_messageInfo_CertifyResponse.Size(m)
}
func (m *CertifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertifyResponse proto.InternalMessageInfo

func (m *CertifyResponse) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func init() {
	proto.RegisterType((*CertifyRequest)(nil), "io.linkerd.identity.CertifyRequest")
	proto.RegisterType((*CertifyResponse)(nil), "io.linkerd.identity.CertifyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	Certify(ctx context.Context, in *CertifyRequest, opts ...grpc.CallOption) (*CertifyResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) Certify(ctx context.Context, in *CertifyRequest, opts ...grpc.CallOption) (*CertifyResponse, error) {
	out := new(CertifyResponse)
	err := c.cc.Invoke(ctx, "/io.linkerd.identity.Identity/Certify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	Certify(context.Context, *CertifyRequest) (*CertifyResponse, error)
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_Certify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Certify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.linkerd.identity.Identity/Certify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Certify(ctx, req.(*CertifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.linkerd.identity.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certify",
			Handler:    _Identity_Certify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor_identity_b6ab172143c7462d) }

var fileDescriptor_identity_b6ab172143c7462d = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4c, 0x49, 0xcd,
	0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce, 0xcc, 0xd7, 0xcb,
	0xc9, 0xcc, 0xcb, 0x4e, 0x2d, 0x4a, 0xd1, 0x83, 0x49, 0x29, 0xa5, 0x71, 0xf1, 0x39, 0xa7, 0x16,
	0x95, 0x64, 0xa6, 0x55, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0x96,
	0xe4, 0x67, 0xa7, 0xe6, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x38, 0x42, 0x76, 0x5c,
	0xd2, 0xc9, 0x60, 0x75, 0x99, 0xc9, 0x89, 0x25, 0xa9, 0xf1, 0xc5, 0x99, 0xe9, 0x79, 0x99, 0x79,
	0xe9, 0xf1, 0x45, 0x10, 0x4d, 0x12, 0x4c, 0x60, 0xb5, 0x92, 0x48, 0x4a, 0x82, 0x21, 0x2a, 0xa0,
	0xa6, 0x2a, 0x19, 0x73, 0xf1, 0xc3, 0xed, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x52, 0xe0,
	0xe2, 0x46, 0x52, 0x0f, 0xb5, 0x0e, 0x59, 0xc8, 0x28, 0x89, 0x8b, 0xc3, 0x13, 0xea, 0x50, 0xa1,
	0x30, 0x2e, 0x76, 0xa8, 0x01, 0x42, 0xca, 0x7a, 0x58, 0x7c, 0xa2, 0x87, 0xea, 0x0d, 0x29, 0x15,
	0xfc, 0x8a, 0x20, 0x6e, 0x50, 0x62, 0x70, 0x52, 0x8f, 0x52, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcf, 0xc9, 0xac, 0x30, 0x28, 0xd2, 0xcf, 0x31, 0x4d, 0xd1,
	0xcd, 0x4c, 0xd1, 0x4f, 0x4f, 0xcd, 0xd3, 0x87, 0x69, 0x4d, 0x62, 0x03, 0x87, 0xa2, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xa7, 0xd5, 0x99, 0x86, 0x57, 0x01, 0x00, 0x00,
}
