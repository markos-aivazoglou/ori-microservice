// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/checkin.proto

package microservice

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

type CheckInRequest struct {
	UserId               string                   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Location             *CheckInRequest_Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Timestamp            string                   `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Venue                *CheckInRequest_Venue    `protobuf:"bytes,4,opt,name=venue,proto3" json:"venue,omitempty"`
	FriendsWith          []string                 `protobuf:"bytes,5,rep,name=friends_with,json=friendsWith,proto3" json:"friends_with,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CheckInRequest) Reset()         { *m = CheckInRequest{} }
func (m *CheckInRequest) String() string { return proto.CompactTextString(m) }
func (*CheckInRequest) ProtoMessage()    {}
func (*CheckInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7313afa335dc5440, []int{0}
}

func (m *CheckInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInRequest.Unmarshal(m, b)
}
func (m *CheckInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInRequest.Marshal(b, m, deterministic)
}
func (m *CheckInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInRequest.Merge(m, src)
}
func (m *CheckInRequest) XXX_Size() int {
	return xxx_messageInfo_CheckInRequest.Size(m)
}
func (m *CheckInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInRequest proto.InternalMessageInfo

func (m *CheckInRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CheckInRequest) GetLocation() *CheckInRequest_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *CheckInRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *CheckInRequest) GetVenue() *CheckInRequest_Venue {
	if m != nil {
		return m.Venue
	}
	return nil
}

func (m *CheckInRequest) GetFriendsWith() []string {
	if m != nil {
		return m.FriendsWith
	}
	return nil
}

type CheckInRequest_Location struct {
	Longitude            float32  `protobuf:"fixed32,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float32  `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckInRequest_Location) Reset()         { *m = CheckInRequest_Location{} }
func (m *CheckInRequest_Location) String() string { return proto.CompactTextString(m) }
func (*CheckInRequest_Location) ProtoMessage()    {}
func (*CheckInRequest_Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_7313afa335dc5440, []int{0, 0}
}

func (m *CheckInRequest_Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInRequest_Location.Unmarshal(m, b)
}
func (m *CheckInRequest_Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInRequest_Location.Marshal(b, m, deterministic)
}
func (m *CheckInRequest_Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInRequest_Location.Merge(m, src)
}
func (m *CheckInRequest_Location) XXX_Size() int {
	return xxx_messageInfo_CheckInRequest_Location.Size(m)
}
func (m *CheckInRequest_Location) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInRequest_Location.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInRequest_Location proto.InternalMessageInfo

func (m *CheckInRequest_Location) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *CheckInRequest_Location) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

type CheckInRequest_Venue struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckInRequest_Venue) Reset()         { *m = CheckInRequest_Venue{} }
func (m *CheckInRequest_Venue) String() string { return proto.CompactTextString(m) }
func (*CheckInRequest_Venue) ProtoMessage()    {}
func (*CheckInRequest_Venue) Descriptor() ([]byte, []int) {
	return fileDescriptor_7313afa335dc5440, []int{0, 1}
}

func (m *CheckInRequest_Venue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInRequest_Venue.Unmarshal(m, b)
}
func (m *CheckInRequest_Venue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInRequest_Venue.Marshal(b, m, deterministic)
}
func (m *CheckInRequest_Venue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInRequest_Venue.Merge(m, src)
}
func (m *CheckInRequest_Venue) XXX_Size() int {
	return xxx_messageInfo_CheckInRequest_Venue.Size(m)
}
func (m *CheckInRequest_Venue) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInRequest_Venue.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInRequest_Venue proto.InternalMessageInfo

func (m *CheckInRequest_Venue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CheckInResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckInResponse) Reset()         { *m = CheckInResponse{} }
func (m *CheckInResponse) String() string { return proto.CompactTextString(m) }
func (*CheckInResponse) ProtoMessage()    {}
func (*CheckInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7313afa335dc5440, []int{1}
}

func (m *CheckInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInResponse.Unmarshal(m, b)
}
func (m *CheckInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInResponse.Marshal(b, m, deterministic)
}
func (m *CheckInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInResponse.Merge(m, src)
}
func (m *CheckInResponse) XXX_Size() int {
	return xxx_messageInfo_CheckInResponse.Size(m)
}
func (m *CheckInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CheckInRequest)(nil), "checkin.CheckInRequest")
	proto.RegisterType((*CheckInRequest_Location)(nil), "checkin.CheckInRequest.Location")
	proto.RegisterType((*CheckInRequest_Venue)(nil), "checkin.CheckInRequest.Venue")
	proto.RegisterType((*CheckInResponse)(nil), "checkin.CheckInResponse")
}

func init() {
	proto.RegisterFile("v1/checkin.proto", fileDescriptor_7313afa335dc5440)
}

var fileDescriptor_7313afa335dc5440 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0xd5, 0xf4, 0xff, 0xf5, 0xd5, 0x4b, 0xb9, 0xa5, 0x56, 0x00, 0x29, 0x74, 0xea, 0xd4,
	0x8a, 0x76, 0x65, 0x40, 0x94, 0xa5, 0xa2, 0x62, 0xb0, 0x04, 0x48, 0x2c, 0x55, 0x48, 0x0e, 0x6a,
	0xd1, 0xd8, 0xc1, 0x76, 0xc2, 0xd7, 0xe3, 0xa3, 0xa1, 0x38, 0x21, 0x05, 0xa1, 0x6e, 0x77, 0xf7,
	0x3c, 0xfe, 0xf9, 0x39, 0x1b, 0x86, 0xf9, 0xc5, 0x2c, 0xda, 0x52, 0xf4, 0x26, 0xe4, 0x34, 0xd5,
	0xca, 0x2a, 0xec, 0x56, 0xed, 0xf8, 0xd3, 0x83, 0xff, 0xcb, 0xa2, 0x5e, 0x49, 0x4e, 0xef, 0x19,
	0x19, 0x8b, 0x23, 0xe8, 0x66, 0x86, 0xf4, 0x46, 0xc4, 0xac, 0x11, 0x34, 0x26, 0x7d, 0xde, 0x29,
	0xda, 0x55, 0x8c, 0x97, 0xd0, 0xdb, 0xa9, 0x28, 0xb4, 0x42, 0x49, 0xe6, 0x05, 0x8d, 0xc9, 0x60,
	0x1e, 0x4c, 0xbf, 0xb1, 0xbf, 0x19, 0xd3, 0x75, 0xe5, 0xe3, 0xf5, 0x09, 0x3c, 0x85, 0xbe, 0x15,
	0x09, 0x19, 0x1b, 0x26, 0x29, 0x6b, 0x3a, 0xf0, 0x7e, 0x80, 0x0b, 0x68, 0xe7, 0x24, 0x33, 0x62,
	0x2d, 0x07, 0x3e, 0x3b, 0x04, 0x7e, 0x28, 0x4c, 0xbc, 0xf4, 0xe2, 0x39, 0xfc, 0x7b, 0xd1, 0x82,
	0x64, 0x6c, 0x36, 0x1f, 0xc2, 0x6e, 0x59, 0x3b, 0x68, 0x4e, 0xfa, 0x7c, 0x50, 0xcd, 0x1e, 0x85,
	0xdd, 0xfa, 0x37, 0xd0, 0x5b, 0xff, 0x48, 0xb0, 0x53, 0xf2, 0x55, 0xd8, 0x2c, 0x26, 0xb7, 0x9a,
	0xc7, 0xf7, 0x03, 0xf4, 0xa1, 0xb7, 0x0b, 0x6d, 0x29, 0x7a, 0x4e, 0xac, 0x7b, 0xff, 0x04, 0xda,
	0xee, 0x62, 0x44, 0x68, 0xdd, 0x85, 0x09, 0x55, 0x0f, 0xe3, 0xea, 0xf1, 0x31, 0x1c, 0xd5, 0x21,
	0x4d, 0xaa, 0xa4, 0xa1, 0xf9, 0x2d, 0x74, 0x97, 0x65, 0x7e, 0xbc, 0x82, 0xc1, 0xbd, 0x21, 0x5d,
	0x39, 0x70, 0x74, 0x60, 0x31, 0x9f, 0xfd, 0x15, 0x4a, 0xd8, 0x35, 0x3e, 0x0d, 0x95, 0x16, 0xb3,
	0x44, 0x44, 0x5a, 0x19, 0xd2, 0xb9, 0x88, 0xe8, 0xb9, 0xe3, 0xbe, 0x71, 0xf1, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x11, 0x59, 0x5a, 0x49, 0xda, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CheckinClient is the client API for Checkin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckinClient interface {
	UserCheckIn(ctx context.Context, in *CheckInRequest, opts ...grpc.CallOption) (*CheckInResponse, error)
}

type checkinClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckinClient(cc grpc.ClientConnInterface) CheckinClient {
	return &checkinClient{cc}
}

func (c *checkinClient) UserCheckIn(ctx context.Context, in *CheckInRequest, opts ...grpc.CallOption) (*CheckInResponse, error) {
	out := new(CheckInResponse)
	err := c.cc.Invoke(ctx, "/checkin.Checkin/UserCheckIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckinServer is the server API for Checkin service.
type CheckinServer interface {
	UserCheckIn(context.Context, *CheckInRequest) (*CheckInResponse, error)
}

// UnimplementedCheckinServer can be embedded to have forward compatible implementations.
type UnimplementedCheckinServer struct {
}

func (*UnimplementedCheckinServer) UserCheckIn(ctx context.Context, req *CheckInRequest) (*CheckInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCheckIn not implemented")
}

func RegisterCheckinServer(s *grpc.Server, srv CheckinServer) {
	s.RegisterService(&_Checkin_serviceDesc, srv)
}

func _Checkin_UserCheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinServer).UserCheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkin.Checkin/UserCheckIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinServer).UserCheckIn(ctx, req.(*CheckInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Checkin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checkin.Checkin",
	HandlerType: (*CheckinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserCheckIn",
			Handler:    _Checkin_UserCheckIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/checkin.proto",
}