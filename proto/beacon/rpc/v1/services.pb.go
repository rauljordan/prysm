// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/beacon/rpc/v1/services.proto

package ethereum_beacon_rpc_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import v1 "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
import v11 "github.com/prysmaticlabs/prysm/proto/sharding/p2p/v1"

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

type ShuffleRequest struct {
	ValidatorCount       uint64   `protobuf:"varint,1,opt,name=validator_count,json=validatorCount,proto3" json:"validator_count,omitempty"`
	ValidatorIndex       uint64   `protobuf:"varint,2,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShuffleRequest) Reset()         { *m = ShuffleRequest{} }
func (m *ShuffleRequest) String() string { return proto.CompactTextString(m) }
func (*ShuffleRequest) ProtoMessage()    {}
func (*ShuffleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_828ddb27b483f4a2, []int{0}
}
func (m *ShuffleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShuffleRequest.Unmarshal(m, b)
}
func (m *ShuffleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShuffleRequest.Marshal(b, m, deterministic)
}
func (dst *ShuffleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShuffleRequest.Merge(dst, src)
}
func (m *ShuffleRequest) XXX_Size() int {
	return xxx_messageInfo_ShuffleRequest.Size(m)
}
func (m *ShuffleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShuffleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShuffleRequest proto.InternalMessageInfo

func (m *ShuffleRequest) GetValidatorCount() uint64 {
	if m != nil {
		return m.ValidatorCount
	}
	return 0
}

func (m *ShuffleRequest) GetValidatorIndex() uint64 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

type ShuffleResponse struct {
	AttesterIndices            []uint64 `protobuf:"varint,1,rep,packed,name=attester_indices,json=attesterIndices,proto3" json:"attester_indices,omitempty"`
	CutoffIndices              []uint64 `protobuf:"varint,2,rep,packed,name=cutoff_indices,json=cutoffIndices,proto3" json:"cutoff_indices,omitempty"`
	ProposerIndices            []uint64 `protobuf:"varint,3,rep,packed,name=proposer_indices,json=proposerIndices,proto3" json:"proposer_indices,omitempty"`
	AssignedAttestationHeights []uint64 `protobuf:"varint,4,rep,packed,name=assigned_attestation_heights,json=assignedAttestationHeights,proto3" json:"assigned_attestation_heights,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *ShuffleResponse) Reset()         { *m = ShuffleResponse{} }
func (m *ShuffleResponse) String() string { return proto.CompactTextString(m) }
func (*ShuffleResponse) ProtoMessage()    {}
func (*ShuffleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_828ddb27b483f4a2, []int{1}
}
func (m *ShuffleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShuffleResponse.Unmarshal(m, b)
}
func (m *ShuffleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShuffleResponse.Marshal(b, m, deterministic)
}
func (dst *ShuffleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShuffleResponse.Merge(dst, src)
}
func (m *ShuffleResponse) XXX_Size() int {
	return xxx_messageInfo_ShuffleResponse.Size(m)
}
func (m *ShuffleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShuffleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShuffleResponse proto.InternalMessageInfo

func (m *ShuffleResponse) GetAttesterIndices() []uint64 {
	if m != nil {
		return m.AttesterIndices
	}
	return nil
}

func (m *ShuffleResponse) GetCutoffIndices() []uint64 {
	if m != nil {
		return m.CutoffIndices
	}
	return nil
}

func (m *ShuffleResponse) GetProposerIndices() []uint64 {
	if m != nil {
		return m.ProposerIndices
	}
	return nil
}

func (m *ShuffleResponse) GetAssignedAttestationHeights() []uint64 {
	if m != nil {
		return m.AssignedAttestationHeights
	}
	return nil
}

type ProposeRequest struct {
	RandaoReveal            []byte              `protobuf:"bytes,1,opt,name=randao_reveal,json=randaoReveal,proto3" json:"randao_reveal,omitempty"`
	AttestationBitmask      []byte              `protobuf:"bytes,2,opt,name=attestation_bitmask,json=attestationBitmask,proto3" json:"attestation_bitmask,omitempty"`
	AttestationAggregateSig []uint32            `protobuf:"varint,3,rep,packed,name=attestation_aggregate_sig,json=attestationAggregateSig,proto3" json:"attestation_aggregate_sig,omitempty"`
	ShardAggregateVotes     []*v1.AggregateVote `protobuf:"bytes,5,rep,name=shard_aggregate_votes,json=shardAggregateVotes,proto3" json:"shard_aggregate_votes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}            `json:"-"`
	XXX_unrecognized        []byte              `json:"-"`
	XXX_sizecache           int32               `json:"-"`
}

func (m *ProposeRequest) Reset()         { *m = ProposeRequest{} }
func (m *ProposeRequest) String() string { return proto.CompactTextString(m) }
func (*ProposeRequest) ProtoMessage()    {}
func (*ProposeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_828ddb27b483f4a2, []int{2}
}
func (m *ProposeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposeRequest.Unmarshal(m, b)
}
func (m *ProposeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposeRequest.Marshal(b, m, deterministic)
}
func (dst *ProposeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposeRequest.Merge(dst, src)
}
func (m *ProposeRequest) XXX_Size() int {
	return xxx_messageInfo_ProposeRequest.Size(m)
}
func (m *ProposeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProposeRequest proto.InternalMessageInfo

func (m *ProposeRequest) GetRandaoReveal() []byte {
	if m != nil {
		return m.RandaoReveal
	}
	return nil
}

func (m *ProposeRequest) GetAttestationBitmask() []byte {
	if m != nil {
		return m.AttestationBitmask
	}
	return nil
}

func (m *ProposeRequest) GetAttestationAggregateSig() []uint32 {
	if m != nil {
		return m.AttestationAggregateSig
	}
	return nil
}

func (m *ProposeRequest) GetShardAggregateVotes() []*v1.AggregateVote {
	if m != nil {
		return m.ShardAggregateVotes
	}
	return nil
}

type ProposeResponse struct {
	BlockHash            []byte   `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposeResponse) Reset()         { *m = ProposeResponse{} }
func (m *ProposeResponse) String() string { return proto.CompactTextString(m) }
func (*ProposeResponse) ProtoMessage()    {}
func (*ProposeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_828ddb27b483f4a2, []int{3}
}
func (m *ProposeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposeResponse.Unmarshal(m, b)
}
func (m *ProposeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposeResponse.Marshal(b, m, deterministic)
}
func (dst *ProposeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposeResponse.Merge(dst, src)
}
func (m *ProposeResponse) XXX_Size() int {
	return xxx_messageInfo_ProposeResponse.Size(m)
}
func (m *ProposeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProposeResponse proto.InternalMessageInfo

func (m *ProposeResponse) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

type SignRequest struct {
	BlockHash            []byte         `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Signature            *v11.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_828ddb27b483f4a2, []int{4}
}
func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRequest.Unmarshal(m, b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
}
func (dst *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(dst, src)
}
func (m *SignRequest) XXX_Size() int {
	return xxx_messageInfo_SignRequest.Size(m)
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

func (m *SignRequest) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *SignRequest) GetSignature() *v11.Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignResponse struct {
	Signed               bool     `protobuf:"varint,1,opt,name=signed,proto3" json:"signed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignResponse) Reset()         { *m = SignResponse{} }
func (m *SignResponse) String() string { return proto.CompactTextString(m) }
func (*SignResponse) ProtoMessage()    {}
func (*SignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_828ddb27b483f4a2, []int{5}
}
func (m *SignResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignResponse.Unmarshal(m, b)
}
func (m *SignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignResponse.Marshal(b, m, deterministic)
}
func (dst *SignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResponse.Merge(dst, src)
}
func (m *SignResponse) XXX_Size() int {
	return xxx_messageInfo_SignResponse.Size(m)
}
func (m *SignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignResponse proto.InternalMessageInfo

func (m *SignResponse) GetSigned() bool {
	if m != nil {
		return m.Signed
	}
	return false
}

func init() {
	proto.RegisterType((*ShuffleRequest)(nil), "ethereum.beacon.rpc.v1.ShuffleRequest")
	proto.RegisterType((*ShuffleResponse)(nil), "ethereum.beacon.rpc.v1.ShuffleResponse")
	proto.RegisterType((*ProposeRequest)(nil), "ethereum.beacon.rpc.v1.ProposeRequest")
	proto.RegisterType((*ProposeResponse)(nil), "ethereum.beacon.rpc.v1.ProposeResponse")
	proto.RegisterType((*SignRequest)(nil), "ethereum.beacon.rpc.v1.SignRequest")
	proto.RegisterType((*SignResponse)(nil), "ethereum.beacon.rpc.v1.SignResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeaconServiceClient is the client API for BeaconService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconServiceClient interface {
	LatestBeaconBlock(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (BeaconService_LatestBeaconBlockClient, error)
	LatestCrystallizedState(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (BeaconService_LatestCrystallizedStateClient, error)
	ShuffleValidators(ctx context.Context, in *ShuffleRequest, opts ...grpc.CallOption) (*ShuffleResponse, error)
	SignBlock(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	ProposeBlock(ctx context.Context, in *ProposeRequest, opts ...grpc.CallOption) (*ProposeResponse, error)
}

type beaconServiceClient struct {
	cc *grpc.ClientConn
}

func NewBeaconServiceClient(cc *grpc.ClientConn) BeaconServiceClient {
	return &beaconServiceClient{cc}
}

func (c *beaconServiceClient) LatestBeaconBlock(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (BeaconService_LatestBeaconBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BeaconService_serviceDesc.Streams[0], "/ethereum.beacon.rpc.v1.BeaconService/LatestBeaconBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &beaconServiceLatestBeaconBlockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BeaconService_LatestBeaconBlockClient interface {
	Recv() (*v1.BeaconBlockResponse, error)
	grpc.ClientStream
}

type beaconServiceLatestBeaconBlockClient struct {
	grpc.ClientStream
}

func (x *beaconServiceLatestBeaconBlockClient) Recv() (*v1.BeaconBlockResponse, error) {
	m := new(v1.BeaconBlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *beaconServiceClient) LatestCrystallizedState(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (BeaconService_LatestCrystallizedStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BeaconService_serviceDesc.Streams[1], "/ethereum.beacon.rpc.v1.BeaconService/LatestCrystallizedState", opts...)
	if err != nil {
		return nil, err
	}
	x := &beaconServiceLatestCrystallizedStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BeaconService_LatestCrystallizedStateClient interface {
	Recv() (*v1.CrystallizedStateResponse, error)
	grpc.ClientStream
}

type beaconServiceLatestCrystallizedStateClient struct {
	grpc.ClientStream
}

func (x *beaconServiceLatestCrystallizedStateClient) Recv() (*v1.CrystallizedStateResponse, error) {
	m := new(v1.CrystallizedStateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *beaconServiceClient) ShuffleValidators(ctx context.Context, in *ShuffleRequest, opts ...grpc.CallOption) (*ShuffleResponse, error) {
	out := new(ShuffleResponse)
	err := c.cc.Invoke(ctx, "/ethereum.beacon.rpc.v1.BeaconService/ShuffleValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconServiceClient) SignBlock(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/ethereum.beacon.rpc.v1.BeaconService/SignBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconServiceClient) ProposeBlock(ctx context.Context, in *ProposeRequest, opts ...grpc.CallOption) (*ProposeResponse, error) {
	out := new(ProposeResponse)
	err := c.cc.Invoke(ctx, "/ethereum.beacon.rpc.v1.BeaconService/ProposeBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconServiceServer is the server API for BeaconService service.
type BeaconServiceServer interface {
	LatestBeaconBlock(*empty.Empty, BeaconService_LatestBeaconBlockServer) error
	LatestCrystallizedState(*empty.Empty, BeaconService_LatestCrystallizedStateServer) error
	ShuffleValidators(context.Context, *ShuffleRequest) (*ShuffleResponse, error)
	SignBlock(context.Context, *SignRequest) (*SignResponse, error)
	ProposeBlock(context.Context, *ProposeRequest) (*ProposeResponse, error)
}

func RegisterBeaconServiceServer(s *grpc.Server, srv BeaconServiceServer) {
	s.RegisterService(&_BeaconService_serviceDesc, srv)
}

func _BeaconService_LatestBeaconBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BeaconServiceServer).LatestBeaconBlock(m, &beaconServiceLatestBeaconBlockServer{stream})
}

type BeaconService_LatestBeaconBlockServer interface {
	Send(*v1.BeaconBlockResponse) error
	grpc.ServerStream
}

type beaconServiceLatestBeaconBlockServer struct {
	grpc.ServerStream
}

func (x *beaconServiceLatestBeaconBlockServer) Send(m *v1.BeaconBlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BeaconService_LatestCrystallizedState_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BeaconServiceServer).LatestCrystallizedState(m, &beaconServiceLatestCrystallizedStateServer{stream})
}

type BeaconService_LatestCrystallizedStateServer interface {
	Send(*v1.CrystallizedStateResponse) error
	grpc.ServerStream
}

type beaconServiceLatestCrystallizedStateServer struct {
	grpc.ServerStream
}

func (x *beaconServiceLatestCrystallizedStateServer) Send(m *v1.CrystallizedStateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BeaconService_ShuffleValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShuffleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconServiceServer).ShuffleValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.beacon.rpc.v1.BeaconService/ShuffleValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconServiceServer).ShuffleValidators(ctx, req.(*ShuffleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconService_SignBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconServiceServer).SignBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.beacon.rpc.v1.BeaconService/SignBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconServiceServer).SignBlock(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconService_ProposeBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconServiceServer).ProposeBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.beacon.rpc.v1.BeaconService/ProposeBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconServiceServer).ProposeBlock(ctx, req.(*ProposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.beacon.rpc.v1.BeaconService",
	HandlerType: (*BeaconServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShuffleValidators",
			Handler:    _BeaconService_ShuffleValidators_Handler,
		},
		{
			MethodName: "SignBlock",
			Handler:    _BeaconService_SignBlock_Handler,
		},
		{
			MethodName: "ProposeBlock",
			Handler:    _BeaconService_ProposeBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LatestBeaconBlock",
			Handler:       _BeaconService_LatestBeaconBlock_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LatestCrystallizedState",
			Handler:       _BeaconService_LatestCrystallizedState_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/beacon/rpc/v1/services.proto",
}

func init() {
	proto.RegisterFile("proto/beacon/rpc/v1/services.proto", fileDescriptor_services_828ddb27b483f4a2)
}

var fileDescriptor_services_828ddb27b483f4a2 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6b, 0xdb, 0x3e,
	0x14, 0x25, 0x4d, 0x7f, 0xe5, 0xd7, 0xdb, 0xfc, 0x69, 0x55, 0xd6, 0x66, 0xd9, 0x06, 0xc5, 0xfd,
	0xcf, 0xc0, 0x6e, 0xb2, 0xb7, 0x3d, 0xad, 0x2d, 0x83, 0x16, 0xf6, 0x30, 0x1c, 0x28, 0x6c, 0x2f,
	0x46, 0xb1, 0x6f, 0x6c, 0x53, 0xc7, 0xf2, 0x24, 0xd9, 0xac, 0xfb, 0x94, 0xfb, 0x2c, 0x7b, 0x1f,
	0x0c, 0x4b, 0x56, 0xe2, 0x74, 0x75, 0xf7, 0xe8, 0x73, 0xcf, 0xb9, 0xe7, 0xea, 0x48, 0xd7, 0x60,
	0x65, 0x9c, 0x49, 0xe6, 0x4c, 0x91, 0xfa, 0x2c, 0x75, 0x78, 0xe6, 0x3b, 0xc5, 0xc8, 0x11, 0xc8,
	0x8b, 0xd8, 0x47, 0x61, 0xab, 0x22, 0xd9, 0x43, 0x19, 0x21, 0xc7, 0x7c, 0x6e, 0x6b, 0x9a, 0xcd,
	0x33, 0xdf, 0x2e, 0x46, 0xc3, 0x23, 0xad, 0x15, 0x11, 0xe5, 0x41, 0x9c, 0x86, 0x4e, 0x36, 0xce,
	0x4a, 0xf5, 0x1c, 0x85, 0xa0, 0xa1, 0x51, 0x0f, 0x57, 0x1d, 0x9e, 0xe6, 0xbc, 0x0a, 0x19, 0x0b,
	0x13, 0x74, 0xd4, 0xd7, 0x34, 0x9f, 0x39, 0x38, 0xcf, 0xe4, 0x83, 0x2e, 0x5a, 0x53, 0xe8, 0x4d,
	0xa2, 0x7c, 0x36, 0x4b, 0xd0, 0xc5, 0x6f, 0x39, 0x0a, 0x49, 0x4e, 0xa1, 0x5f, 0xd0, 0x24, 0x0e,
	0xa8, 0x64, 0xdc, 0xf3, 0x59, 0x9e, 0xca, 0x41, 0xeb, 0xa0, 0x75, 0xb6, 0xee, 0xf6, 0x16, 0xf0,
	0x75, 0x89, 0xae, 0x12, 0xe3, 0x34, 0xc0, 0xef, 0x83, 0xb5, 0x47, 0xc4, 0xdb, 0x12, 0xb5, 0x7e,
	0xb6, 0xa0, 0xbf, 0x30, 0x11, 0x19, 0x4b, 0x05, 0x92, 0x73, 0xd8, 0xa6, 0x52, 0xa2, 0x90, 0xa8,
	0xb4, 0x65, 0x20, 0x83, 0xd6, 0x41, 0xfb, 0x6c, 0xdd, 0xed, 0x1b, 0xfc, 0x56, 0xc3, 0xe4, 0x18,
	0x7a, 0x7e, 0x2e, 0xd9, 0x6c, 0xb6, 0x20, 0xae, 0x29, 0x62, 0x57, 0xa3, 0x86, 0x76, 0x0e, 0xdb,
	0x19, 0x67, 0x19, 0x13, 0xb5, 0x8e, 0x6d, 0xdd, 0xd1, 0xe0, 0x86, 0xfa, 0x01, 0x5e, 0x53, 0x21,
	0xe2, 0x30, 0xc5, 0xc0, 0xd3, 0x6e, 0x54, 0xc6, 0x2c, 0xf5, 0x22, 0x8c, 0xc3, 0x48, 0x8a, 0xc1,
	0xba, 0x92, 0x0d, 0x0d, 0xe7, 0x72, 0x49, 0xb9, 0xd1, 0x0c, 0xeb, 0x77, 0x0b, 0x7a, 0x9f, 0x75,
	0x57, 0x93, 0xdb, 0x21, 0x74, 0x39, 0x4d, 0x03, 0xca, 0x3c, 0x8e, 0x05, 0xd2, 0x44, 0xa5, 0xd6,
	0x71, 0x3b, 0x1a, 0x74, 0x15, 0x46, 0x1c, 0xd8, 0xad, 0x1b, 0x4e, 0x63, 0x39, 0xa7, 0xe2, 0x5e,
	0xe5, 0xd6, 0x71, 0x49, 0xad, 0x74, 0xa5, 0x2b, 0xe4, 0x3d, 0xbc, 0xac, 0x0b, 0x68, 0x18, 0x72,
	0x0c, 0xa9, 0x44, 0x4f, 0xc4, 0xa1, 0x3a, 0x5e, 0xd7, 0xdd, 0xaf, 0x11, 0x2e, 0x4d, 0x7d, 0x12,
	0x87, 0xe4, 0x0b, 0xbc, 0x50, 0xcf, 0xa7, 0xa6, 0x2a, 0x98, 0x44, 0x31, 0xf8, 0xef, 0xa0, 0x7d,
	0xb6, 0x35, 0x3e, 0xb6, 0x1f, 0x3f, 0xbd, 0x6c, 0x9c, 0xd9, 0xc5, 0xc8, 0x5e, 0x34, 0xb9, 0x63,
	0x12, 0xdd, 0x5d, 0xd5, 0x63, 0x05, 0x13, 0xd6, 0x05, 0xf4, 0x17, 0xc7, 0xaf, 0x6e, 0xf4, 0x0d,
	0xc0, 0x34, 0x61, 0xfe, 0xbd, 0x17, 0x51, 0x11, 0x55, 0x87, 0xdf, 0x54, 0xc8, 0x0d, 0x15, 0x91,
	0xc5, 0x60, 0x6b, 0x12, 0x87, 0xa9, 0x49, 0xeb, 0x79, 0x36, 0xb9, 0x84, 0xcd, 0x32, 0x7b, 0x2a,
	0x73, 0x8e, 0x2a, 0x9d, 0xad, 0xf1, 0xe1, 0x72, 0x5c, 0xb3, 0x14, 0x66, 0xe0, 0x89, 0xa1, 0xba,
	0x4b, 0x95, 0x75, 0x02, 0x1d, 0x6d, 0x58, 0xcd, 0xb7, 0x07, 0x1b, 0xfa, 0x3a, 0x95, 0xdb, 0xff,
	0x6e, 0xf5, 0x35, 0xfe, 0xd5, 0x86, 0xee, 0x95, 0x3a, 0xff, 0x44, 0x6f, 0x26, 0xf9, 0x0a, 0x3b,
	0x9f, 0x68, 0x19, 0xa9, 0x86, 0xaf, 0xca, 0xa9, 0xc8, 0x9e, 0xad, 0xd7, 0xc8, 0x36, 0x6b, 0x64,
	0x7f, 0x2c, 0xd7, 0x68, 0xf8, 0xb6, 0x29, 0xc5, 0x9a, 0xd8, 0xcc, 0x70, 0xd1, 0x22, 0x33, 0xd8,
	0xd7, 0xbd, 0xaf, 0xf9, 0x83, 0x90, 0x34, 0x49, 0xe2, 0x1f, 0x18, 0x4c, 0x24, 0x95, 0xd8, 0xe8,
	0x30, 0x6a, 0x72, 0xf8, 0xab, 0x45, 0xcd, 0x27, 0x80, 0x9d, 0x6a, 0xe5, 0xee, 0xcc, 0x32, 0x0a,
	0x72, 0x62, 0x3f, 0xfd, 0xb3, 0xb1, 0x57, 0x7f, 0x01, 0xc3, 0xd3, 0x7f, 0xf2, 0xaa, 0x4c, 0xef,
	0x60, 0xb3, 0xcc, 0x58, 0x27, 0x74, 0xd8, 0xa8, 0x5a, 0xde, 0xfb, 0xf0, 0xe8, 0x79, 0x52, 0xd5,
	0xd7, 0x83, 0x4e, 0xf5, 0xbc, 0x74, 0xeb, 0xc6, 0xc1, 0x57, 0x77, 0xb0, 0x79, 0xf0, 0x47, 0x8f,
	0x75, 0xba, 0xa1, 0x32, 0x7e, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x34, 0x39, 0xf2, 0x54, 0xa2,
	0x05, 0x00, 0x00,
}
