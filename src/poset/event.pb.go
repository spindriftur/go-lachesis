// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package poset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import peers "github.com/Fantom-foundation/go-lachesis/src/peers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TransactionType int32

const (
	TransactionType_PEER_ADD     TransactionType = 0
	TransactionType_PEER_REMOVE  TransactionType = 1
	TransactionType_POS_TRANSFER TransactionType = 2
)

var TransactionType_name = map[int32]string{
	0: "PEER_ADD",
	1: "PEER_REMOVE",
	2: "POS_TRANSFER",
}
var TransactionType_value = map[string]int32{
	"PEER_ADD":     0,
	"PEER_REMOVE":  1,
	"POS_TRANSFER": 2,
}

func (x TransactionType) String() string {
	return proto.EnumName(TransactionType_name, int32(x))
}
func (TransactionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type InternalTransaction struct {
	Type   TransactionType    `protobuf:"varint,1,opt,name=Type,json=type,enum=poset.TransactionType" json:"Type,omitempty"`
	Peer   *peers.PeerMessage `protobuf:"bytes,2,opt,name=peer" json:"peer,omitempty"`
	Amount uint64             `protobuf:"varint,3,opt,name=Amount,json=amount" json:"Amount,omitempty"`
}

func (m *InternalTransaction) Reset()                    { *m = InternalTransaction{} }
func (m *InternalTransaction) String() string            { return proto.CompactTextString(m) }
func (*InternalTransaction) ProtoMessage()               {}
func (*InternalTransaction) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *InternalTransaction) GetType() TransactionType {
	if m != nil {
		return m.Type
	}
	return TransactionType_PEER_ADD
}

func (m *InternalTransaction) GetPeer() *peers.PeerMessage {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *InternalTransaction) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type BlockSignature struct {
	Validator []byte `protobuf:"bytes,1,opt,name=Validator,json=validator,proto3" json:"Validator,omitempty"`
	Index     int64  `protobuf:"varint,2,opt,name=Index,json=index" json:"Index,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=Signature,json=signature" json:"Signature,omitempty"`
}

func (m *BlockSignature) Reset()                    { *m = BlockSignature{} }
func (m *BlockSignature) String() string            { return proto.CompactTextString(m) }
func (*BlockSignature) ProtoMessage()               {}
func (*BlockSignature) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *BlockSignature) GetValidator() []byte {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *BlockSignature) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BlockSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type EventBody struct {
	Transactions         [][]byte               `protobuf:"bytes,1,rep,name=Transactions,json=transactions,proto3" json:"Transactions,omitempty"`
	InternalTransactions []*InternalTransaction `protobuf:"bytes,2,rep,name=InternalTransactions,json=internalTransactions" json:"InternalTransactions,omitempty"`
	Parents              [][]byte               `protobuf:"bytes,3,rep,name=Parents,json=parents,proto3" json:"Parents,omitempty"`
	Creator              []byte                 `protobuf:"bytes,4,opt,name=Creator,json=creator,proto3" json:"Creator,omitempty"`
	Index                int64                  `protobuf:"varint,5,opt,name=Index,json=index" json:"Index,omitempty"`
	BlockSignatures      []*BlockSignature      `protobuf:"bytes,6,rep,name=BlockSignatures,json=blockSignatures" json:"BlockSignatures,omitempty"`
}

func (m *EventBody) Reset()                    { *m = EventBody{} }
func (m *EventBody) String() string            { return proto.CompactTextString(m) }
func (*EventBody) ProtoMessage()               {}
func (*EventBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *EventBody) GetTransactions() [][]byte {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *EventBody) GetInternalTransactions() []*InternalTransaction {
	if m != nil {
		return m.InternalTransactions
	}
	return nil
}

func (m *EventBody) GetParents() [][]byte {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *EventBody) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *EventBody) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *EventBody) GetBlockSignatures() []*BlockSignature {
	if m != nil {
		return m.BlockSignatures
	}
	return nil
}

type EventMessage struct {
	Body                 *EventBody `protobuf:"bytes,1,opt,name=Body,json=body" json:"Body,omitempty"`
	Signature            string     `protobuf:"bytes,2,opt,name=Signature,json=signature" json:"Signature,omitempty"`
	SelfParentIndex      int64      `protobuf:"varint,3,opt,name=SelfParentIndex,json=selfParentIndex" json:"SelfParentIndex,omitempty"`
	OtherParentCreatorID uint64     `protobuf:"varint,4,opt,name=OtherParentCreatorID,json=otherParentCreatorID" json:"OtherParentCreatorID,omitempty"`
	OtherParentIndex     int64      `protobuf:"varint,5,opt,name=OtherParentIndex,json=otherParentIndex" json:"OtherParentIndex,omitempty"`
	CreatorID            uint64     `protobuf:"varint,6,opt,name=CreatorID,json=creatorID" json:"CreatorID,omitempty"`
	TopologicalIndex     int64      `protobuf:"varint,7,opt,name=TopologicalIndex,json=topologicalIndex" json:"TopologicalIndex,omitempty"`
	Hash                 []byte     `protobuf:"bytes,8,opt,name=Hash,json=hash,proto3" json:"Hash,omitempty"`
}

func (m *EventMessage) Reset()                    { *m = EventMessage{} }
func (m *EventMessage) String() string            { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()               {}
func (*EventMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *EventMessage) GetBody() *EventBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *EventMessage) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *EventMessage) GetSelfParentIndex() int64 {
	if m != nil {
		return m.SelfParentIndex
	}
	return 0
}

func (m *EventMessage) GetOtherParentCreatorID() uint64 {
	if m != nil {
		return m.OtherParentCreatorID
	}
	return 0
}

func (m *EventMessage) GetOtherParentIndex() int64 {
	if m != nil {
		return m.OtherParentIndex
	}
	return 0
}

func (m *EventMessage) GetCreatorID() uint64 {
	if m != nil {
		return m.CreatorID
	}
	return 0
}

func (m *EventMessage) GetTopologicalIndex() int64 {
	if m != nil {
		return m.TopologicalIndex
	}
	return 0
}

func (m *EventMessage) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Event struct {
	Message          *EventMessage `protobuf:"bytes,1,opt,name=Message,json=message" json:"Message,omitempty"`
	LamportTimestamp int64         `protobuf:"varint,2,opt,name=LamportTimestamp,json=lamportTimestamp" json:"LamportTimestamp,omitempty"`
	Frame            int64         `protobuf:"varint,3,opt,name=Frame,json=frame" json:"Frame,omitempty"`
	FlagTableBytes   []byte        `protobuf:"bytes,4,opt,name=FlagTableBytes,json=flagTableBytes,proto3" json:"FlagTableBytes,omitempty"`
	RootTableBytes   []byte        `protobuf:"bytes,5,opt,name=RootTableBytes,json=rootTableBytes,proto3" json:"RootTableBytes,omitempty"`
	Root             bool          `protobuf:"varint,6,opt,name=Root,json=root" json:"Root,omitempty"`
	Clotho           bool          `protobuf:"varint,7,opt,name=Clotho,json=clotho" json:"Clotho,omitempty"`
	Atropos          bool          `protobuf:"varint,8,opt,name=Atropos,json=atropos" json:"Atropos,omitempty"`
	AtroposTimestamp int64         `protobuf:"varint,9,opt,name=AtroposTimestamp,json=atroposTimestamp" json:"AtroposTimestamp,omitempty"`
	AtTimes          []int64       `protobuf:"varint,10,rep,packed,name=AtTimes,json=atTimes" json:"AtTimes,omitempty"`
	AtVisited        int64         `protobuf:"varint,11,opt,name=AtVisited,json=atVisited" json:"AtVisited,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Event) GetMessage() *EventMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Event) GetLamportTimestamp() int64 {
	if m != nil {
		return m.LamportTimestamp
	}
	return 0
}

func (m *Event) GetFrame() int64 {
	if m != nil {
		return m.Frame
	}
	return 0
}

func (m *Event) GetFlagTableBytes() []byte {
	if m != nil {
		return m.FlagTableBytes
	}
	return nil
}

func (m *Event) GetRootTableBytes() []byte {
	if m != nil {
		return m.RootTableBytes
	}
	return nil
}

func (m *Event) GetRoot() bool {
	if m != nil {
		return m.Root
	}
	return false
}

func (m *Event) GetClotho() bool {
	if m != nil {
		return m.Clotho
	}
	return false
}

func (m *Event) GetAtropos() bool {
	if m != nil {
		return m.Atropos
	}
	return false
}

func (m *Event) GetAtroposTimestamp() int64 {
	if m != nil {
		return m.AtroposTimestamp
	}
	return 0
}

func (m *Event) GetAtTimes() []int64 {
	if m != nil {
		return m.AtTimes
	}
	return nil
}

func (m *Event) GetAtVisited() int64 {
	if m != nil {
		return m.AtVisited
	}
	return 0
}

func init() {
	proto.RegisterType((*InternalTransaction)(nil), "poset.InternalTransaction")
	proto.RegisterType((*BlockSignature)(nil), "poset.BlockSignature")
	proto.RegisterType((*EventBody)(nil), "poset.EventBody")
	proto.RegisterType((*EventMessage)(nil), "poset.EventMessage")
	proto.RegisterType((*Event)(nil), "poset.Event")
	proto.RegisterEnum("poset.TransactionType", TransactionType_name, TransactionType_value)
}

func init() { proto.RegisterFile("event.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdf, 0x6a, 0x1a, 0x41,
	0x14, 0xc6, 0xab, 0xbb, 0xfe, 0xd9, 0xa3, 0xe8, 0x32, 0xb1, 0x61, 0x09, 0xbd, 0x10, 0x29, 0x41,
	0x84, 0x28, 0xd8, 0xeb, 0x52, 0x34, 0x51, 0x1a, 0x68, 0x12, 0x19, 0x25, 0xb7, 0x61, 0x5c, 0x47,
	0x5d, 0xba, 0xbb, 0xb3, 0xcc, 0x8c, 0xa1, 0x5e, 0xf6, 0xb6, 0x8f, 0xd4, 0xc7, 0xea, 0x13, 0x94,
	0x39, 0xbb, 0x26, 0xab, 0x78, 0x23, 0x9c, 0xef, 0x9c, 0xf9, 0x66, 0xbe, 0xdf, 0x8c, 0x0b, 0x35,
	0xfe, 0xca, 0x63, 0xdd, 0x4f, 0xa4, 0xd0, 0x82, 0x94, 0x12, 0xa1, 0xb8, 0xbe, 0xfa, 0xba, 0x09,
	0xf4, 0x76, 0xb7, 0xec, 0xfb, 0x22, 0x1a, 0x4c, 0x59, 0xac, 0x45, 0x74, 0xb3, 0x16, 0xbb, 0x78,
	0xc5, 0x74, 0x20, 0xe2, 0xc1, 0x46, 0xdc, 0x84, 0xcc, 0xdf, 0x72, 0x15, 0xa8, 0x81, 0x92, 0xfe,
	0x20, 0xe1, 0x5c, 0x2a, 0xfc, 0x4d, 0x5d, 0x3a, 0xbf, 0x0b, 0x70, 0x71, 0x1f, 0x6b, 0x2e, 0x63,
	0x16, 0x2e, 0x24, 0x8b, 0x15, 0xf3, 0xcd, 0x42, 0xd2, 0x03, 0x7b, 0xb1, 0x4f, 0xb8, 0x57, 0x68,
	0x17, 0xba, 0x8d, 0xe1, 0x65, 0x1f, 0x37, 0xeb, 0xe7, 0x26, 0x4c, 0x97, 0xda, 0x7a, 0x9f, 0x70,
	0x72, 0x0d, 0xb6, 0x71, 0xf4, 0x8a, 0xed, 0x42, 0xb7, 0x36, 0x24, 0x7d, 0xdc, 0xa4, 0x3f, 0xe3,
	0x5c, 0x3e, 0x70, 0xa5, 0xd8, 0x86, 0x53, 0xec, 0x93, 0x4b, 0x28, 0x8f, 0x22, 0xb1, 0x8b, 0xb5,
	0x67, 0xb5, 0x0b, 0x5d, 0x9b, 0x96, 0x19, 0x56, 0x9d, 0x25, 0x34, 0xc6, 0xa1, 0xf0, 0x7f, 0xce,
	0x83, 0x4d, 0xcc, 0xf4, 0x4e, 0x72, 0xf2, 0x09, 0x9c, 0x67, 0x16, 0x06, 0x2b, 0xa6, 0x85, 0xc4,
	0x23, 0xd4, 0xa9, 0xf3, 0x7a, 0x10, 0x48, 0x0b, 0x4a, 0xf7, 0xf1, 0x8a, 0xff, 0xc2, 0x0d, 0x2d,
	0x5a, 0x0a, 0x4c, 0x61, 0xd6, 0xbc, 0x19, 0xe0, 0x06, 0x0e, 0x75, 0xd4, 0x41, 0xe8, 0xfc, 0x29,
	0x82, 0x33, 0x31, 0xf4, 0xc6, 0x62, 0xb5, 0x27, 0x1d, 0xa8, 0xe7, 0xa2, 0x28, 0xaf, 0xd0, 0xb6,
	0xba, 0x75, 0x5a, 0xd7, 0x39, 0x8d, 0x3c, 0x42, 0xeb, 0x0c, 0x18, 0xe5, 0x15, 0xdb, 0x56, 0xb7,
	0x36, 0xbc, 0xca, 0x88, 0x9c, 0x19, 0xa1, 0xad, 0xe0, 0xcc, 0x3a, 0xe2, 0x41, 0x65, 0xc6, 0x24,
	0x8f, 0xb5, 0xf2, 0x2c, 0xdc, 0xae, 0x92, 0xa4, 0xa5, 0xe9, 0xdc, 0x4a, 0x8e, 0x59, 0x6d, 0xcc,
	0x5a, 0xf1, 0xd3, 0xf2, 0x3d, 0x69, 0x29, 0x9f, 0xf4, 0x1b, 0x34, 0x8f, 0x79, 0x29, 0xaf, 0x8c,
	0x87, 0xfa, 0x98, 0x1d, 0xea, 0xb8, 0x4b, 0x9b, 0xcb, 0xe3, 0xe9, 0xce, 0xdf, 0x22, 0xd4, 0x11,
	0x46, 0x76, 0x3f, 0xe4, 0x33, 0xd8, 0x86, 0x0b, 0xa2, 0xae, 0x0d, 0xdd, 0xcc, 0xe6, 0x8d, 0x17,
	0xb5, 0x97, 0x86, 0xda, 0x11, 0xe1, 0xe2, 0x09, 0x61, 0xd2, 0x85, 0xe6, 0x9c, 0x87, 0xeb, 0x34,
	0x63, 0x7a, 0x6a, 0x0b, 0x4f, 0xdd, 0x54, 0xc7, 0x32, 0x19, 0x42, 0xeb, 0x49, 0x6f, 0xb9, 0x4c,
	0xb5, 0x2c, 0xfa, 0xfd, 0x1d, 0x86, 0xb7, 0x69, 0x4b, 0x9c, 0xe9, 0x91, 0x1e, 0xb8, 0xb9, 0x35,
	0x79, 0x28, 0xae, 0x38, 0xd1, 0xcd, 0x39, 0xdf, 0x4d, 0xcb, 0x68, 0xea, 0xf8, 0x79, 0xa7, 0x85,
	0x48, 0x44, 0x28, 0x36, 0x81, 0xcf, 0xc2, 0xd4, 0xa9, 0x92, 0x3a, 0xe9, 0x13, 0x9d, 0x10, 0xb0,
	0xbf, 0x33, 0xb5, 0xf5, 0xaa, 0x78, 0x2d, 0xf6, 0x96, 0xa9, 0x6d, 0xe7, 0x5f, 0x11, 0x4a, 0x48,
	0x86, 0xdc, 0x40, 0x25, 0x03, 0x98, 0x81, 0xbb, 0xc8, 0x83, 0x3b, 0xbc, 0xfd, 0x4a, 0x94, 0x41,
	0xee, 0x81, 0xfb, 0x83, 0x45, 0x89, 0x90, 0x7a, 0x11, 0x44, 0x5c, 0x69, 0x16, 0x25, 0xd9, 0x0b,
	0x76, 0xc3, 0x13, 0xdd, 0x5c, 0xfc, 0x54, 0xb2, 0x88, 0x67, 0x08, 0x4b, 0x6b, 0x53, 0x90, 0x6b,
	0x68, 0x4c, 0x43, 0xb6, 0x59, 0xb0, 0x65, 0xc8, 0xc7, 0x7b, 0xcd, 0x55, 0xf6, 0x5e, 0x1a, 0xeb,
	0x23, 0xd5, 0xcc, 0x51, 0x21, 0x74, 0x6e, 0xae, 0x94, 0xce, 0xc9, 0x23, 0xd5, 0xc4, 0x33, 0x73,
	0xc8, 0xa8, 0x4a, 0x6d, 0xd3, 0x35, 0x7f, 0xd2, 0xdb, 0x50, 0xe8, 0xad, 0x40, 0x28, 0x55, 0x5a,
	0xf6, 0xb1, 0x32, 0x8f, 0x74, 0xa4, 0xa5, 0x48, 0x84, 0x42, 0x1a, 0x55, 0x5a, 0x61, 0x69, 0x69,
	0x72, 0x65, 0x9d, 0xf7, 0x5c, 0x4e, 0x9a, 0x8b, 0x9d, 0xe8, 0xa9, 0x0b, 0x96, 0x1e, 0xb4, 0xad,
	0xae, 0x65, 0x5c, 0xb0, 0x34, 0x97, 0x36, 0xd2, 0xcf, 0x81, 0x0a, 0x34, 0x5f, 0x79, 0x35, 0x5c,
	0xee, 0xb0, 0x83, 0xd0, 0x1b, 0x43, 0xf3, 0xe4, 0xdb, 0x43, 0xea, 0x50, 0x9d, 0x4d, 0x26, 0xf4,
	0x65, 0x74, 0x77, 0xe7, 0x7e, 0x20, 0x4d, 0xa8, 0x61, 0x45, 0x27, 0x0f, 0x4f, 0xcf, 0x13, 0xb7,
	0x40, 0x5c, 0xa8, 0xcf, 0x9e, 0xe6, 0x2f, 0x0b, 0x3a, 0x7a, 0x9c, 0x4f, 0x27, 0xd4, 0x2d, 0x2e,
	0xcb, 0xf8, 0xc5, 0xfb, 0xf2, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xae, 0x1f, 0x94, 0x46, 0x05,
	0x00, 0x00,
}
