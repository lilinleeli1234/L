// Code generated by protoc-gen-go. DO NOT EDIT.
// source: f3d.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	f3d.proto

It has these top-level messages:
	RoundInfo
	F3DGameAction
	StartGame
	StopGame
	BuyKeys
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoundInfo struct {
	// 游戏轮次
	Round int64 `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
	// 本轮游戏开始事件
	BeginTime int64 `protobuf:"varint,2,opt,name=beginTime" json:"beginTime,omitempty"`
	// 本轮游戏结束时间
	EndTime int64 `protobuf:"varint,3,opt,name=endTime" json:"endTime,omitempty"`
	// 本轮游戏目前为止最后一把钥匙持有人（游戏开奖时，这个就是中奖人）
	LastOwner string `protobuf:"bytes,4,opt,name=lastOwner" json:"lastOwner,omitempty"`
	// 最后一把钥匙的购买时间
	LastKeyTime int64 `protobuf:"varint,5,opt,name=lastKeyTime" json:"lastKeyTime,omitempty"`
	// 最后一把钥匙的价格
	LastKeyPrice int64 `protobuf:"varint,6,opt,name=lastKeyPrice" json:"lastKeyPrice,omitempty"`
	// 本轮游戏奖金池总额
	BonusPool int64 `protobuf:"varint,7,opt,name=bonusPool" json:"bonusPool,omitempty"`
	// 本轮游戏参与地址数
	UserCount int64 `protobuf:"varint,8,opt,name=userCount" json:"userCount,omitempty"`
	// 本轮游戏募集到的key个数
	KeyCount int64 `protobuf:"varint,9,opt,name=keyCount" json:"keyCount,omitempty"`
}

func (m *RoundInfo) Reset()                    { *m = RoundInfo{} }
func (m *RoundInfo) String() string            { return proto.CompactTextString(m) }
func (*RoundInfo) ProtoMessage()               {}
func (*RoundInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RoundInfo) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *RoundInfo) GetBeginTime() int64 {
	if m != nil {
		return m.BeginTime
	}
	return 0
}

func (m *RoundInfo) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *RoundInfo) GetLastOwner() string {
	if m != nil {
		return m.LastOwner
	}
	return ""
}

func (m *RoundInfo) GetLastKeyTime() int64 {
	if m != nil {
		return m.LastKeyTime
	}
	return 0
}

func (m *RoundInfo) GetLastKeyPrice() int64 {
	if m != nil {
		return m.LastKeyPrice
	}
	return 0
}

func (m *RoundInfo) GetBonusPool() int64 {
	if m != nil {
		return m.BonusPool
	}
	return 0
}

func (m *RoundInfo) GetUserCount() int64 {
	if m != nil {
		return m.UserCount
	}
	return 0
}

func (m *RoundInfo) GetKeyCount() int64 {
	if m != nil {
		return m.KeyCount
	}
	return 0
}

type F3DGameAction struct {
	// Types that are valid to be assigned to Value:
	//	*F3DGameAction_Start
	//	*F3DGameAction_Stop
	//	*F3DGameAction_Trade
	Value isF3DGameAction_Value `protobuf_oneof:"value"`
	Ty    int32                 `protobuf:"varint,4,opt,name=ty" json:"ty,omitempty"`
}

func (m *F3DGameAction) Reset()                    { *m = F3DGameAction{} }
func (m *F3DGameAction) String() string            { return proto.CompactTextString(m) }
func (*F3DGameAction) ProtoMessage()               {}
func (*F3DGameAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isF3DGameAction_Value interface {
	isF3DGameAction_Value()
}

type F3DGameAction_Start struct {
	Start *StartGame `protobuf:"bytes,1,opt,name=start,oneof"`
}
type F3DGameAction_Stop struct {
	Stop *StopGame `protobuf:"bytes,2,opt,name=stop,oneof"`
}
type F3DGameAction_Trade struct {
	Trade *BuyKeys `protobuf:"bytes,3,opt,name=trade,oneof"`
}

func (*F3DGameAction_Start) isF3DGameAction_Value() {}
func (*F3DGameAction_Stop) isF3DGameAction_Value()  {}
func (*F3DGameAction_Trade) isF3DGameAction_Value() {}

func (m *F3DGameAction) GetValue() isF3DGameAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *F3DGameAction) GetStart() *StartGame {
	if x, ok := m.GetValue().(*F3DGameAction_Start); ok {
		return x.Start
	}
	return nil
}

func (m *F3DGameAction) GetStop() *StopGame {
	if x, ok := m.GetValue().(*F3DGameAction_Stop); ok {
		return x.Stop
	}
	return nil
}

func (m *F3DGameAction) GetTrade() *BuyKeys {
	if x, ok := m.GetValue().(*F3DGameAction_Trade); ok {
		return x.Trade
	}
	return nil
}

func (m *F3DGameAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*F3DGameAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _F3DGameAction_OneofMarshaler, _F3DGameAction_OneofUnmarshaler, _F3DGameAction_OneofSizer, []interface{}{
		(*F3DGameAction_Start)(nil),
		(*F3DGameAction_Stop)(nil),
		(*F3DGameAction_Trade)(nil),
	}
}

func _F3DGameAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*F3DGameAction)
	// value
	switch x := m.Value.(type) {
	case *F3DGameAction_Start:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Start); err != nil {
			return err
		}
	case *F3DGameAction_Stop:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Stop); err != nil {
			return err
		}
	case *F3DGameAction_Trade:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Trade); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("F3DGameAction.Value has unexpected type %T", x)
	}
	return nil
}

func _F3DGameAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*F3DGameAction)
	switch tag {
	case 1: // value.start
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StartGame)
		err := b.DecodeMessage(msg)
		m.Value = &F3DGameAction_Start{msg}
		return true, err
	case 2: // value.stop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StopGame)
		err := b.DecodeMessage(msg)
		m.Value = &F3DGameAction_Stop{msg}
		return true, err
	case 3: // value.trade
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuyKeys)
		err := b.DecodeMessage(msg)
		m.Value = &F3DGameAction_Trade{msg}
		return true, err
	default:
		return false, nil
	}
}

func _F3DGameAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*F3DGameAction)
	// value
	switch x := m.Value.(type) {
	case *F3DGameAction_Start:
		s := proto.Size(x.Start)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *F3DGameAction_Stop:
		s := proto.Size(x.Stop)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *F3DGameAction_Trade:
		s := proto.Size(x.Trade)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StartGame struct {
	LastRound int64 `protobuf:"varint,1,opt,name=lastRound" json:"lastRound,omitempty"`
}

func (m *StartGame) Reset()                    { *m = StartGame{} }
func (m *StartGame) String() string            { return proto.CompactTextString(m) }
func (*StartGame) ProtoMessage()               {}
func (*StartGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StartGame) GetLastRound() int64 {
	if m != nil {
		return m.LastRound
	}
	return 0
}

type StopGame struct {
	Round int64 `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
}

func (m *StopGame) Reset()                    { *m = StopGame{} }
func (m *StopGame) String() string            { return proto.CompactTextString(m) }
func (*StopGame) ProtoMessage()               {}
func (*StopGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StopGame) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type BuyKeys struct {
	KeyNum int64 `protobuf:"varint,1,opt,name=keyNum" json:"keyNum,omitempty"`
}

func (m *BuyKeys) Reset()                    { *m = BuyKeys{} }
func (m *BuyKeys) String() string            { return proto.CompactTextString(m) }
func (*BuyKeys) ProtoMessage()               {}
func (*BuyKeys) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BuyKeys) GetKeyNum() int64 {
	if m != nil {
		return m.KeyNum
	}
	return 0
}

func init() {
	proto.RegisterType((*RoundInfo)(nil), "types.RoundInfo")
	proto.RegisterType((*F3DGameAction)(nil), "types.F3DGameAction")
	proto.RegisterType((*StartGame)(nil), "types.StartGame")
	proto.RegisterType((*StopGame)(nil), "types.StopGame")
	proto.RegisterType((*BuyKeys)(nil), "types.BuyKeys")
}

func init() { proto.RegisterFile("f3d.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x6d, 0xb7, 0xae, 0xeb, 0xa9, 0x4e, 0x09, 0x22, 0x45, 0xbc, 0xa8, 0x05, 0x65, 0xde,
	0xec, 0x62, 0x7b, 0x02, 0xa7, 0xf8, 0x87, 0x81, 0x8e, 0xea, 0x0b, 0x74, 0x6b, 0x26, 0x65, 0x5b,
	0x52, 0x9a, 0x44, 0xc9, 0x83, 0xf8, 0x00, 0xbe, 0xa9, 0xe4, 0x34, 0xed, 0x26, 0x78, 0xf9, 0x7d,
	0xbf, 0xef, 0xf4, 0x84, 0xef, 0x14, 0x82, 0xd5, 0x24, 0x1f, 0x95, 0x15, 0x97, 0x9c, 0x78, 0x52,
	0x97, 0x54, 0x24, 0xdf, 0x2e, 0x04, 0x29, 0x57, 0x2c, 0x7f, 0x66, 0x2b, 0x4e, 0x4e, 0xc1, 0xab,
	0x8c, 0x88, 0x9c, 0xd8, 0x19, 0x76, 0xd2, 0x5a, 0x90, 0x0b, 0x08, 0x16, 0xf4, 0xa3, 0x60, 0xef,
	0xc5, 0x96, 0x46, 0x2e, 0x92, 0x9d, 0x41, 0x22, 0xf0, 0x29, 0xcb, 0x91, 0x75, 0x90, 0x35, 0xd2,
	0xcc, 0x6d, 0x32, 0x21, 0x5f, 0xbf, 0x18, 0xad, 0xa2, 0x6e, 0xec, 0x0c, 0x83, 0x74, 0x67, 0x90,
	0x18, 0x42, 0x23, 0x66, 0x54, 0xe3, 0xac, 0x87, 0xb3, 0xfb, 0x16, 0x49, 0xe0, 0xd0, 0xca, 0x79,
	0x55, 0x2c, 0x69, 0xd4, 0xc3, 0xc8, 0x1f, 0x0f, 0xdf, 0xc6, 0x99, 0x12, 0x73, 0xce, 0x37, 0x91,
	0x6f, 0xdf, 0xd6, 0x18, 0x86, 0x2a, 0x41, 0xab, 0x3b, 0xae, 0x98, 0x8c, 0xfa, 0x35, 0x6d, 0x0d,
	0x72, 0x0e, 0xfd, 0x35, 0xd5, 0x35, 0x0c, 0x10, 0xb6, 0x3a, 0xf9, 0x71, 0xe0, 0xe8, 0x61, 0x72,
	0xff, 0x98, 0x6d, 0xe9, 0xed, 0x52, 0x16, 0x9c, 0x91, 0x21, 0x78, 0x42, 0x66, 0x95, 0xc4, 0x6e,
	0xc2, 0xf1, 0xc9, 0x08, 0x0b, 0x1c, 0xbd, 0x19, 0xcf, 0xc4, 0x9e, 0x0e, 0xd2, 0x3a, 0x40, 0xae,
	0xa0, 0x2b, 0x24, 0x2f, 0xb1, 0xaa, 0x70, 0x7c, 0xdc, 0x06, 0x79, 0x69, 0x73, 0x88, 0xc9, 0x35,
	0x78, 0xb2, 0xca, 0xf2, 0xba, 0xb6, 0x70, 0x3c, 0xb0, 0xb9, 0xa9, 0xd2, 0x33, 0xaa, 0x85, 0xf9,
	0x1c, 0x62, 0x32, 0x00, 0x57, 0x6a, 0xec, 0xcf, 0x4b, 0x5d, 0xa9, 0xa7, 0x3e, 0x78, 0x9f, 0xd9,
	0x46, 0xd1, 0xe4, 0x06, 0x82, 0x76, 0x7b, 0x53, 0x76, 0xba, 0x77, 0xbe, 0x9d, 0x91, 0xc4, 0xd0,
	0x6f, 0xf6, 0xff, 0x7f, 0xe4, 0xe4, 0x12, 0x7c, 0xbb, 0x99, 0x9c, 0x41, 0x6f, 0x4d, 0xf5, 0x8b,
	0xda, 0xda, 0x84, 0x55, 0x8b, 0x1e, 0xfe, 0x39, 0x93, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde,
	0x24, 0xc7, 0x2a, 0x46, 0x02, 0x00, 0x00,
}
