// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cpi.proto

package cpi

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BaseRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	StemcellApiVersion   int32    `protobuf:"varint,2,opt,name=stemcell_api_version,json=stemcellApiVersion,proto3" json:"stemcell_api_version,omitempty"`
	DirectorUuid         string   `protobuf:"bytes,3,opt,name=director_uuid,json=directorUuid,proto3" json:"director_uuid,omitempty"`
	Properties           []byte   `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{0}
}

func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (m *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(m, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

func (m *BaseRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BaseRequest) GetStemcellApiVersion() int32 {
	if m != nil {
		return m.StemcellApiVersion
	}
	return 0
}

func (m *BaseRequest) GetDirectorUuid() string {
	if m != nil {
		return m.DirectorUuid
	}
	return ""
}

func (m *BaseRequest) GetProperties() []byte {
	if m != nil {
		return m.Properties
	}
	return nil
}

type BaseResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Log                  string   `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{1}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BaseResponse) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

type InfoResponse struct {
	Base                 *BaseResponse `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ApiVersion           int32         `protobuf:"varint,2,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	StemcellFormats      []string      `protobuf:"bytes,3,rep,name=stemcell_formats,json=stemcellFormats,proto3" json:"stemcell_formats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27dcbb49f4ec00bf, []int{2}
}

func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetBase() *BaseResponse {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *InfoResponse) GetApiVersion() int32 {
	if m != nil {
		return m.ApiVersion
	}
	return 0
}

func (m *InfoResponse) GetStemcellFormats() []string {
	if m != nil {
		return m.StemcellFormats
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "cpi.BaseRequest")
	proto.RegisterType((*BaseResponse)(nil), "cpi.BaseResponse")
	proto.RegisterType((*InfoResponse)(nil), "cpi.InfoResponse")
}

func init() { proto.RegisterFile("cpi.proto", fileDescriptor_27dcbb49f4ec00bf) }

var fileDescriptor_27dcbb49f4ec00bf = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0xe9, 0x84, 0xbe, 0x55, 0x9c, 0x8f, 0x1d, 0x8a, 0x07, 0x2d, 0x13, 0xa1, 0x22,
	0x0c, 0x99, 0xe0, 0x5d, 0x05, 0x61, 0x37, 0x09, 0xe8, 0xb5, 0x74, 0xed, 0x9b, 0x04, 0xba, 0x26,
	0x26, 0xa9, 0xe0, 0xd1, 0xff, 0xc2, 0x3f, 0x57, 0x9a, 0xd8, 0x59, 0xf0, 0xf6, 0xf2, 0x7d, 0xe4,
	0xe5, 0xf7, 0x7d, 0x81, 0xb8, 0xd2, 0x72, 0xa9, 0x8d, 0x72, 0x0a, 0x79, 0xa5, 0xe5, 0xe2, 0x9b,
	0xc1, 0xf4, 0xa1, 0xb4, 0x24, 0xe8, 0xbd, 0x23, 0xeb, 0x10, 0x21, 0x72, 0x9f, 0x9a, 0x52, 0x96,
	0xb1, 0x3c, 0x16, 0x7e, 0xc6, 0x1b, 0x98, 0x5b, 0x47, 0xbb, 0x8a, 0x9a, 0xa6, 0x28, 0xb5, 0x2c,
	0x3e, 0xc8, 0x58, 0xa9, 0xda, 0xf4, 0x20, 0x63, 0xf9, 0x44, 0xe0, 0xe0, 0xdd, 0x6b, 0xf9, 0x1a,
	0x1c, 0xbc, 0x80, 0xa3, 0x5a, 0x1a, 0xaa, 0x9c, 0x32, 0x45, 0xd7, 0xc9, 0x3a, 0xe5, 0x7e, 0x5d,
	0x32, 0x88, 0x2f, 0x9d, 0xac, 0xf1, 0x0c, 0x40, 0x1b, 0xa5, 0xc9, 0x38, 0x49, 0x36, 0x8d, 0x32,
	0x96, 0x27, 0x62, 0xa4, 0x2c, 0xee, 0x20, 0x09, 0x64, 0x56, 0xab, 0xd6, 0x12, 0xce, 0x61, 0x42,
	0xc6, 0x28, 0xf3, 0xcb, 0x16, 0x0e, 0x38, 0x03, 0xde, 0xa8, 0x37, 0xcf, 0x12, 0x8b, 0x7e, 0x5c,
	0x7c, 0x31, 0x48, 0xd6, 0xed, 0x56, 0xed, 0x2f, 0x5e, 0x42, 0xb4, 0x29, 0x6d, 0xc8, 0x34, 0x5d,
	0x9d, 0x2c, 0xfb, 0x0a, 0xc6, 0x9b, 0x85, 0xb7, 0xf1, 0x1c, 0xa6, 0xff, 0xd3, 0x41, 0xf9, 0x97,
	0xea, 0x0a, 0x66, 0xfb, 0x1e, 0xb6, 0xca, 0xec, 0x4a, 0x67, 0x53, 0x9e, 0xf1, 0x3c, 0x16, 0xc7,
	0x83, 0xfe, 0x14, 0xe4, 0xd5, 0x0a, 0xf8, 0xe3, 0xf3, 0x1a, 0xaf, 0x21, 0xea, 0x49, 0x70, 0x36,
	0x7a, 0xd3, 0xf7, 0x7c, 0x1a, 0x28, 0xc6, 0x98, 0x9b, 0x43, 0xff, 0x2d, 0xb7, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x01, 0x70, 0xc3, 0xad, 0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CPIClient is the client API for CPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CPIClient interface {
	Info(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type cPIClient struct {
	cc *grpc.ClientConn
}

func NewCPIClient(cc *grpc.ClientConn) CPIClient {
	return &cPIClient{cc}
}

func (c *cPIClient) Info(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/cpi.CPI/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CPIServer is the server API for CPI service.
type CPIServer interface {
	Info(context.Context, *BaseRequest) (*InfoResponse, error)
}

func RegisterCPIServer(s *grpc.Server, srv CPIServer) {
	s.RegisterService(&_CPI_serviceDesc, srv)
}

func _CPI_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPIServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpi.CPI/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPIServer).Info(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cpi.CPI",
	HandlerType: (*CPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _CPI_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cpi.proto",
}
