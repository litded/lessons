// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package consignment

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

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Command struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Command) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Command) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Command) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Command) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Response struct {
	Created              bool       `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Command              *Command   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Commands             []*Command `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
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

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Response) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*Command)(nil), "consignment.Command")
	proto.RegisterType((*Response)(nil), "consignment.Response")
	proto.RegisterType((*GetRequest)(nil), "consignment.GetRequest")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x7d, 0x85, 0x07, 0x94, 0x5b, 0xc5, 0xe4, 0x46, 0x81, 0xe8, 0x42, 0x52, 0x37, 0xac, 0xd0,
	0x60, 0xe2, 0xd2, 0x44, 0x59, 0x18, 0xb6, 0xc3, 0x07, 0x18, 0x6c, 0x6f, 0xca, 0x24, 0x74, 0xa6,
	0xce, 0x0c, 0xf8, 0x07, 0xfe, 0x81, 0xff, 0xe0, 0x67, 0x3a, 0x9d, 0xb6, 0x58, 0x54, 0x76, 0xf7,
	0x9e, 0x73, 0xe6, 0x9e, 0x73, 0x9a, 0xc2, 0x55, 0xa6, 0xa4, 0x91, 0xd7, 0x91, 0x14, 0x9a, 0x27,
	0x22, 0x25, 0x61, 0xea, 0xf3, 0xc4, 0xb1, 0x18, 0xd4, 0xa0, 0x30, 0x85, 0xee, 0x4c, 0x0a, 0xb3,
	0xe4, 0x82, 0x14, 0xf6, 0xa0, 0xc1, 0xe3, 0xa1, 0x37, 0xf2, 0xc6, 0x5d, 0x66, 0x27, 0xbc, 0x84,
	0x20, 0xda, 0x68, 0x23, 0x53, 0x52, 0xcf, 0x96, 0x68, 0x38, 0x02, 0x2a, 0x68, 0x1e, 0x63, 0x1f,
	0xda, 0x52, 0xf1, 0x84, 0x8b, 0x61, 0xd3, 0x71, 0xe5, 0x86, 0x03, 0xe8, 0x6c, 0x74, 0xf1, 0xe8,
	0x7f, 0x41, 0xe4, 0xeb, 0x3c, 0x0e, 0x3f, 0x3d, 0xe8, 0xcc, 0x64, 0x9a, 0x2e, 0x45, 0xfc, 0xcb,
	0x6d, 0x04, 0x41, 0x4c, 0x3a, 0x52, 0x3c, 0x33, 0x5c, 0x8a, 0xd2, 0xad, 0x0e, 0xe5, 0x76, 0x6f,
	0xc4, 0x93, 0x95, 0x71, 0x76, 0x2d, 0x56, 0x6e, 0x78, 0x07, 0x10, 0x55, 0x25, 0xb4, 0x75, 0x6c,
	0x8e, 0x83, 0x69, 0x7f, 0x52, 0x6f, 0xbe, 0xeb, 0xc8, 0x6a, 0x4a, 0xbc, 0x80, 0xee, 0x96, 0xb4,
	0xa6, 0x75, 0x1e, 0xb4, 0xe5, 0xfc, 0xfc, 0x02, 0xb0, 0x51, 0xdf, 0x3d, 0xf0, 0x19, 0xe9, 0xcc,
	0x5e, 0x21, 0x1c, 0x42, 0x27, 0x52, 0xb4, 0x34, 0x54, 0x04, 0xf6, 0x59, 0xb5, 0xe2, 0xc4, 0x32,
	0x45, 0x21, 0x97, 0x38, 0x98, 0x9e, 0xfe, 0x30, 0x76, 0x1c, 0xab, 0x44, 0x78, 0x03, 0x7e, 0x39,
	0x6a, 0xdb, 0xa2, 0x79, 0xf0, 0xc1, 0x4e, 0x15, 0x1e, 0x01, 0x3c, 0x91, 0x61, 0xf4, 0xba, 0x21,
	0x6d, 0xa6, 0x1f, 0x1e, 0x9c, 0x2c, 0x56, 0x3c, 0xcb, 0xb8, 0x48, 0x16, 0xa4, 0xb6, 0x3c, 0x22,
	0xbc, 0x87, 0xe3, 0x99, 0x8b, 0x53, 0x7d, 0xda, 0x3f, 0x4f, 0x9e, 0x9f, 0xed, 0xa1, 0x55, 0xb7,
	0xf0, 0x1f, 0x3e, 0x42, 0xcf, 0x3a, 0x3c, 0xac, 0xd7, 0xa5, 0x52, 0xe3, 0x60, 0x4f, 0xfa, 0x6d,
	0x7f, 0xf0, 0xc6, 0x4b, 0xdb, 0xfd, 0x5c, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x31,
	0x13, 0xba, 0x83, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
	//???????????????? ?????????? rpc ??????????
	GetAllCommands(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.ShippingService/CreateCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetAllCommands(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.ShippingService/GetAllCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateCommand(context.Context, *Command) (*Response, error)
	//???????????????? ?????????? rpc ??????????
	GetAllCommands(context.Context, *GetRequest) (*Response, error)
}

// UnimplementedShippingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShippingServiceServer struct {
}

func (*UnimplementedShippingServiceServer) CreateCommand(ctx context.Context, req *Command) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommand not implemented")
}
func (*UnimplementedShippingServiceServer) GetAllCommands(ctx context.Context, req *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommands not implemented")
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.ShippingService/CreateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetAllCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetAllCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.ShippingService/GetAllCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetAllCommands(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommand",
			Handler:    _ShippingService_CreateCommand_Handler,
		},
		{
			MethodName: "GetAllCommands",
			Handler:    _ShippingService_GetAllCommands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}
