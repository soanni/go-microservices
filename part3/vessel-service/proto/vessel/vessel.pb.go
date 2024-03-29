// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_0675a6e741cbe00d, []int{0}
}
func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (dst *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(dst, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_0675a6e741cbe00d, []int{1}
}
func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (dst *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(dst, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_0675a6e741cbe00d, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	CreateVessel(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) CreateVessel(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.CreateVessel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
	CreateVessel(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) CreateVessel(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.CreateVessel(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_vessel_0675a6e741cbe00d) }

var fileDescriptor_vessel_0675a6e741cbe00d = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x69, 0x9b, 0xa6, 0xa3, 0x2d, 0x32, 0x20, 0x6c, 0x8b, 0x42, 0xc8, 0x41, 0x72, 0xaa,
	0x50, 0x6f, 0xde, 0x44, 0x10, 0xf4, 0xb8, 0x05, 0x3d, 0x96, 0xed, 0xee, 0xa8, 0x0b, 0x6d, 0x12,
	0x92, 0x90, 0xd6, 0x93, 0xbf, 0xe2, 0xa7, 0x0a, 0x93, 0x6c, 0x25, 0x52, 0x3c, 0xed, 0xbc, 0x37,
	0x33, 0x8f, 0xb7, 0x6f, 0x60, 0x9a, 0x17, 0x59, 0x95, 0xdd, 0xd4, 0x54, 0x96, 0xb4, 0x69, 0x9f,
	0x39, 0x73, 0x18, 0x34, 0x28, 0xfe, 0xf6, 0x20, 0x78, 0xe1, 0x12, 0x27, 0xe0, 0x5b, 0x23, 0xbc,
	0xc8, 0x4b, 0x46, 0xd2, 0xb7, 0x06, 0x67, 0x10, 0x6a, 0x95, 0x2b, 0x6d, 0xab, 0x4f, 0xe1, 0x47,
	0x5e, 0x32, 0x90, 0x07, 0x8c, 0x57, 0x00, 0x5b, 0xb5, 0x5f, 0xed, 0xc8, 0xbe, 0x7f, 0x54, 0xa2,
	0xc7, 0xdd, 0xd1, 0x56, 0xed, 0x5f, 0x99, 0x40, 0x84, 0x7e, 0xaa, 0xb6, 0x24, 0xfa, 0x2c, 0xc6,
	0x35, 0x5e, 0xc2, 0x48, 0xd5, 0xca, 0x6e, 0xd4, 0x7a, 0x43, 0x62, 0x10, 0x79, 0x49, 0x28, 0x7f,
	0x09, 0x9c, 0x42, 0x98, 0xed, 0x52, 0x2a, 0x56, 0xd6, 0x88, 0x80, 0xb7, 0x86, 0x8c, 0x9f, 0x4c,
	0xfc, 0x0c, 0xe3, 0x65, 0x4e, 0xda, 0xbe, 0x59, 0xad, 0x2a, 0x9b, 0xa5, 0x1d, 0x63, 0xde, 0xbf,
	0xc6, 0xfc, 0x3f, 0xc6, 0xe2, 0x1a, 0x42, 0x49, 0x65, 0x9e, 0xa5, 0x25, 0xe1, 0x35, 0xb4, 0x21,
	0xb0, 0xc8, 0xe9, 0x62, 0x32, 0x6f, 0x13, 0x6a, 0xf2, 0x90, 0x6d, 0x17, 0x13, 0x18, 0x36, 0x55,
	0x29, 0xfc, 0xa8, 0x77, 0x64, 0xd0, 0xb5, 0x51, 0xc0, 0x50, 0x17, 0xa4, 0x2a, 0x32, 0x1c, 0x49,
	0x28, 0x1d, 0x5c, 0x7c, 0xc1, 0xb8, 0x19, 0x5e, 0x52, 0x51, 0x5b, 0x4d, 0x78, 0x07, 0xe3, 0x47,
	0x9b, 0x9a, 0xfb, 0x43, 0x00, 0x17, 0x4e, 0xb4, 0xf3, 0xd7, 0xd9, 0xb9, 0xa3, 0x9d, 0xed, 0xf8,
	0x04, 0x17, 0x70, 0xf6, 0xc0, 0xba, 0xee, 0x70, 0x5d, 0x3f, 0xc7, 0x76, 0xd6, 0x01, 0x9f, 0xfd,
	0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x79, 0x14, 0xe1, 0x13, 0x02, 0x00, 0x00,
}
