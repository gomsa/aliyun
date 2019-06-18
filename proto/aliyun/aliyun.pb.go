// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/aliyun/aliyun.proto

package aliyun

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Request struct {
	Method               string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Scheme               string            `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Domain               string            `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Version              string            `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	ApiName              string            `protobuf:"bytes,6,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	QueryParams          map[string]string `protobuf:"bytes,7,rep,name=query_params,json=queryParams,proto3" json:"query_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8580126ab2a4da8, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *Request) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Request) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Request) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *Request) GetQueryParams() map[string]string {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

type Response struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8580126ab2a4da8, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "aliyun.Request")
	proto.RegisterMapType((map[string]string)(nil), "aliyun.Request.QueryParamsEntry")
	proto.RegisterType((*Response)(nil), "aliyun.Response")
}

func init() { proto.RegisterFile("proto/aliyun/aliyun.proto", fileDescriptor_b8580126ab2a4da8) }

var fileDescriptor_b8580126ab2a4da8 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6b, 0xc2, 0x30,
	0x14, 0xc7, 0xd7, 0x16, 0x5b, 0xf7, 0x1c, 0xac, 0x04, 0x19, 0xd1, 0x53, 0x29, 0x3b, 0x78, 0xea,
	0xc0, 0x5d, 0xc6, 0x06, 0x83, 0x21, 0x5e, 0x87, 0xeb, 0x17, 0x90, 0xac, 0x3e, 0xb0, 0xcc, 0x24,
	0x6d, 0x92, 0x0a, 0xfd, 0x0c, 0xfb, 0xd2, 0xd2, 0x24, 0x45, 0xf0, 0x94, 0xfc, 0x7e, 0xef, 0x25,
	0xfc, 0x5f, 0x02, 0x8b, 0x46, 0x49, 0x23, 0x5f, 0xd8, 0xa9, 0xee, 0x3b, 0xe1, 0x97, 0xc2, 0x3a,
	0x12, 0x3b, 0xca, 0xff, 0x43, 0x48, 0x4a, 0x6c, 0x3b, 0xd4, 0x86, 0x3c, 0x41, 0xcc, 0xd1, 0x1c,
	0xe5, 0x81, 0x06, 0x59, 0xb0, 0xba, 0x2f, 0x3d, 0x0d, 0x5e, 0x57, 0x47, 0xe4, 0x48, 0x43, 0xe7,
	0x1d, 0x0d, 0xfe, 0x20, 0x39, 0xab, 0x05, 0x8d, 0x9c, 0x77, 0x44, 0x28, 0x24, 0x67, 0x54, 0xba,
	0x96, 0x82, 0x4e, 0x6c, 0x61, 0x44, 0xb2, 0x80, 0x29, 0x6b, 0xea, 0xbd, 0x60, 0x1c, 0x69, 0xec,
	0x4a, 0xac, 0xa9, 0xbf, 0x19, 0x47, 0xb2, 0x81, 0x87, 0xb6, 0x43, 0xd5, 0xef, 0x1b, 0xa6, 0x18,
	0xd7, 0x34, 0xc9, 0xa2, 0xd5, 0x6c, 0x9d, 0x15, 0x3e, 0xb5, 0xcf, 0x58, 0xfc, 0x0c, 0x3d, 0x3b,
	0xdb, 0xb2, 0x15, 0x46, 0xf5, 0xe5, 0xac, 0xbd, 0x9a, 0xe5, 0x27, 0xa4, 0xb7, 0x0d, 0x24, 0x85,
	0xe8, 0x0f, 0x7b, 0x3f, 0xd2, 0xb0, 0x25, 0x73, 0x98, 0x9c, 0xd9, 0xa9, 0x1b, 0xc7, 0x71, 0xf0,
	0x1e, 0xbe, 0x05, 0xf9, 0x33, 0x4c, 0x4b, 0xd4, 0x8d, 0x14, 0x1a, 0x87, 0x29, 0x2a, 0x29, 0x0c,
	0x0a, 0xe3, 0xcf, 0x8e, 0xb8, 0xde, 0x42, 0xfc, 0x65, 0x53, 0x91, 0x0f, 0x98, 0xef, 0x94, 0xac,
	0x50, 0xeb, 0x8d, 0xe4, 0x5c, 0x8a, 0xf1, 0x25, 0x1f, 0x6f, 0x62, 0x2f, 0xd3, 0xab, 0x70, 0xd7,
	0xe7, 0x77, 0xbf, 0xb1, 0xfd, 0x89, 0xd7, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x50, 0x2f,
	0x8e, 0xa6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Aliyun service

type AliyunClient interface {
	// 处理请求
	ProcessCommonRequest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type aliyunClient struct {
	c           client.Client
	serviceName string
}

func NewAliyunClient(serviceName string, c client.Client) AliyunClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "aliyun"
	}
	return &aliyunClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *aliyunClient) ProcessCommonRequest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Aliyun.ProcessCommonRequest", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Aliyun service

type AliyunHandler interface {
	// 处理请求
	ProcessCommonRequest(context.Context, *Request, *Response) error
}

func RegisterAliyunHandler(s server.Server, hdlr AliyunHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Aliyun{hdlr}, opts...))
}

type Aliyun struct {
	AliyunHandler
}

func (h *Aliyun) ProcessCommonRequest(ctx context.Context, in *Request, out *Response) error {
	return h.AliyunHandler.ProcessCommonRequest(ctx, in, out)
}