// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: amm/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSendCreatePool struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	PoolParam        string `protobuf:"bytes,5,opt,name=poolParam,proto3" json:"poolParam,omitempty"`
	PoolAssets       string `protobuf:"bytes,6,opt,name=poolAssets,proto3" json:"poolAssets,omitempty"`
}

func (m *MsgSendCreatePool) Reset()         { *m = MsgSendCreatePool{} }
func (m *MsgSendCreatePool) String() string { return proto.CompactTextString(m) }
func (*MsgSendCreatePool) ProtoMessage()    {}
func (*MsgSendCreatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_93040f60fa240c68, []int{0}
}
func (m *MsgSendCreatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendCreatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendCreatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendCreatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendCreatePool.Merge(m, src)
}
func (m *MsgSendCreatePool) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendCreatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendCreatePool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendCreatePool proto.InternalMessageInfo

func (m *MsgSendCreatePool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendCreatePool) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendCreatePool) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendCreatePool) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendCreatePool) GetPoolParam() string {
	if m != nil {
		return m.PoolParam
	}
	return ""
}

func (m *MsgSendCreatePool) GetPoolAssets() string {
	if m != nil {
		return m.PoolAssets
	}
	return ""
}

type MsgSendCreatePoolResponse struct {
}

func (m *MsgSendCreatePoolResponse) Reset()         { *m = MsgSendCreatePoolResponse{} }
func (m *MsgSendCreatePoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendCreatePoolResponse) ProtoMessage()    {}
func (*MsgSendCreatePoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_93040f60fa240c68, []int{1}
}
func (m *MsgSendCreatePoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendCreatePoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendCreatePoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendCreatePoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendCreatePoolResponse.Merge(m, src)
}
func (m *MsgSendCreatePoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendCreatePoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendCreatePoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendCreatePoolResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendCreatePool)(nil), "lostak.vela.amm.MsgSendCreatePool")
	proto.RegisterType((*MsgSendCreatePoolResponse)(nil), "lostak.vela.amm.MsgSendCreatePoolResponse")
}

func init() { proto.RegisterFile("amm/tx.proto", fileDescriptor_93040f60fa240c68) }

var fileDescriptor_93040f60fa240c68 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xeb, 0xaf, 0xfd, 0x8a, 0x6a, 0x21, 0x7e, 0x3c, 0x99, 0x1f, 0x59, 0x55, 0xa6, 0xaa,
	0x83, 0x23, 0xc1, 0xca, 0xc2, 0xcf, 0xc2, 0x50, 0xa9, 0x2a, 0x88, 0x81, 0x09, 0x37, 0x1c, 0xa5,
	0x11, 0x76, 0x8e, 0x15, 0xbb, 0xa8, 0xdc, 0x05, 0x97, 0x05, 0x5b, 0x47, 0x46, 0x94, 0xdc, 0x08,
	0x8a, 0xab, 0x02, 0x6a, 0x06, 0x36, 0xbf, 0xcf, 0xb1, 0x1e, 0xeb, 0xf8, 0xa5, 0xdb, 0xca, 0x98,
	0xd8, 0x2f, 0xa4, 0x2d, 0xd0, 0x23, 0xdb, 0xd5, 0xe8, 0xbc, 0x7a, 0x92, 0xcf, 0xa0, 0x95, 0x54,
	0xc6, 0x44, 0xef, 0x84, 0xee, 0x8f, 0x5c, 0x7a, 0x03, 0xf9, 0xe3, 0x65, 0x01, 0xca, 0xc3, 0x18,
	0x51, 0x33, 0x4e, 0xb7, 0x92, 0x3a, 0x61, 0xc1, 0x49, 0x9f, 0x0c, 0x7a, 0x93, 0x75, 0x64, 0x8c,
	0x76, 0x2c, 0x16, 0x9e, 0xff, 0x0b, 0x38, 0x9c, 0xd9, 0x31, 0xed, 0x25, 0x33, 0x95, 0xe7, 0xa0,
	0xaf, 0xaf, 0x78, 0x3b, 0x0c, 0x7e, 0x00, 0x1b, 0xd2, 0x3d, 0x9f, 0x19, 0xc0, 0xb9, 0xbf, 0xcd,
	0x0c, 0x38, 0xaf, 0x8c, 0xe5, 0x9d, 0x3e, 0x19, 0x74, 0x26, 0x0d, 0x5e, 0x9b, 0x2c, 0xa2, 0x1e,
	0xab, 0x42, 0x19, 0xfe, 0x7f, 0x65, 0xfa, 0x06, 0x4c, 0x50, 0x5a, 0x87, 0x73, 0xe7, 0xc0, 0x3b,
	0xde, 0x0d, 0xe3, 0x5f, 0x24, 0x3a, 0xa2, 0x07, 0x8d, 0x55, 0x26, 0xe0, 0x2c, 0xe6, 0x0e, 0x4e,
	0x52, 0xda, 0x1e, 0xb9, 0x94, 0x3d, 0xd0, 0x9d, 0x8d, 0x5d, 0x23, 0xb9, 0xf1, 0x27, 0xb2, 0x21,
	0x39, 0x1c, 0xfe, 0x7d, 0x67, 0xfd, 0xd0, 0xc5, 0xd9, 0x5b, 0x29, 0xc8, 0xb2, 0x14, 0xe4, 0xb3,
	0x14, 0xe4, 0xb5, 0x12, 0xad, 0x65, 0x25, 0x5a, 0x1f, 0x95, 0x68, 0xdd, 0x47, 0x69, 0xe6, 0x67,
	0xf3, 0xa9, 0x4c, 0xd0, 0xc4, 0x2b, 0x5f, 0x7c, 0x07, 0x5a, 0xc5, 0x8b, 0x38, 0x74, 0xf4, 0x62,
	0xc1, 0x4d, 0xbb, 0xa1, 0xa7, 0xd3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x94, 0x44, 0x9b,
	0xb7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SendCreatePool(ctx context.Context, in *MsgSendCreatePool, opts ...grpc.CallOption) (*MsgSendCreatePoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendCreatePool(ctx context.Context, in *MsgSendCreatePool, opts ...grpc.CallOption) (*MsgSendCreatePoolResponse, error) {
	out := new(MsgSendCreatePoolResponse)
	err := c.cc.Invoke(ctx, "/lostak.vela.amm.Msg/SendCreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendCreatePool(context.Context, *MsgSendCreatePool) (*MsgSendCreatePoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendCreatePool(ctx context.Context, req *MsgSendCreatePool) (*MsgSendCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCreatePool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendCreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendCreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lostak.vela.amm.Msg/SendCreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendCreatePool(ctx, req.(*MsgSendCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lostak.vela.amm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCreatePool",
			Handler:    _Msg_SendCreatePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "amm/tx.proto",
}

func (m *MsgSendCreatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendCreatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendCreatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolAssets) > 0 {
		i -= len(m.PoolAssets)
		copy(dAtA[i:], m.PoolAssets)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PoolAssets)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PoolParam) > 0 {
		i -= len(m.PoolParam)
		copy(dAtA[i:], m.PoolParam)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PoolParam)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendCreatePoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendCreatePoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendCreatePoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendCreatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.PoolParam)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PoolAssets)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendCreatePoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendCreatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendCreatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendCreatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParam", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolParam = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssets = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendCreatePoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendCreatePoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendCreatePoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
