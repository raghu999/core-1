// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ask_plan.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	ask_plan.proto
	benchmarks.proto
	bid.proto
	bigint.proto
	capabilities.proto
	container.proto
	deal.proto
	dwh.proto
	hub.proto
	insonmnia.proto
	marketplace.proto
	miner.proto
	net.proto
	node.proto
	relay.proto
	rendezvous.proto
	timestamp.proto
	volume.proto

It has these top-level messages:
	AskPlanCPU
	AskPlanGPU
	AskPlanRAM
	AskPlanStorage
	AskPlanNetwork
	AskPlanResources
	AskPlan
	Benchmark
	BidNetwork
	BidResources
	BidOrder
	BigInt
	CPUDevice
	CPU
	RAMDevice
	RAM
	GPUDevice
	GPU
	NetworkDevice
	Network
	StorageDevice
	Storage
	NetworkSpec
	Container
	DeprecatedDeal
	DealsRequest
	DWHDealsReply
	DWHDeal
	DealConditionsRequest
	DealConditionsReply
	OrdersRequest
	MatchingOrdersRequest
	DWHOrdersReply
	DWHOrder
	DealCondition
	DWHWorker
	ProfilesRequest
	ProfilesReply
	Profile
	BlacklistReply
	ValidatorsRequest
	ValidatorsReply
	Validator
	DealChangeRequestsReply
	WorkersRequest
	WorkersReply
	Certificate
	DWHBenchmarkConditions
	MaxMinUint64
	MaxMinBig
	CmpUint64
	DealChangeRequest
	StartTaskRequest
	HubJoinNetworkRequest
	StartTaskReply
	HubStatusReply
	AskPlansReply
	TaskListReply
	DevicesReply
	PullTaskRequest
	DealInfoReply
	Empty
	ID
	TaskID
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	TaskStatusReply
	AvailableResources
	StatusMapReply
	ContainerRestartPolicy
	TaskLogsRequest
	TaskLogsChunk
	TaskResourceRequirements
	Chunk
	Progress
	Duration
	EthAddress
	DataSize
	DataSizeRate
	Price
	GetOrdersRequest
	GetOrdersReply
	Benchmarks
	Deal
	Order
	MinerStartRequest
	MinerStartReply
	TaskInfo
	Endpoints
	SaveRequest
	Addr
	SocketAddr
	JoinNetworkRequest
	DealListRequest
	DealsReply
	Worker
	WorkerListReply
	HandshakeRequest
	DiscoverResponse
	HandshakeResponse
	RelayClusterReply
	RelayMetrics
	NetMetrics
	ConnectRequest
	PublishRequest
	RendezvousReply
	RendezvousState
	RendezvousMeeting
	ResolveMetaReply
	Timestamp
	Volume
*/
package sonm

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

type AskPlanCPU struct {
	CorePercents uint64 `protobuf:"varint,1,opt,name=core_percents,json=corePercents" json:"core_percents,omitempty"`
}

func (m *AskPlanCPU) Reset()                    { *m = AskPlanCPU{} }
func (m *AskPlanCPU) String() string            { return proto.CompactTextString(m) }
func (*AskPlanCPU) ProtoMessage()               {}
func (*AskPlanCPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AskPlanCPU) GetCorePercents() uint64 {
	if m != nil {
		return m.CorePercents
	}
	return 0
}

type AskPlanGPU struct {
	Indexes []uint64 `protobuf:"varint,1,rep,packed,name=indexes" json:"indexes,omitempty"`
	Hashes  []string `protobuf:"bytes,2,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *AskPlanGPU) Reset()                    { *m = AskPlanGPU{} }
func (m *AskPlanGPU) String() string            { return proto.CompactTextString(m) }
func (*AskPlanGPU) ProtoMessage()               {}
func (*AskPlanGPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AskPlanGPU) GetIndexes() []uint64 {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *AskPlanGPU) GetHashes() []string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type AskPlanRAM struct {
	Size *DataSize `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
}

func (m *AskPlanRAM) Reset()                    { *m = AskPlanRAM{} }
func (m *AskPlanRAM) String() string            { return proto.CompactTextString(m) }
func (*AskPlanRAM) ProtoMessage()               {}
func (*AskPlanRAM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AskPlanRAM) GetSize() *DataSize {
	if m != nil {
		return m.Size
	}
	return nil
}

type AskPlanStorage struct {
	Size *DataSize `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
}

func (m *AskPlanStorage) Reset()                    { *m = AskPlanStorage{} }
func (m *AskPlanStorage) String() string            { return proto.CompactTextString(m) }
func (*AskPlanStorage) ProtoMessage()               {}
func (*AskPlanStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AskPlanStorage) GetSize() *DataSize {
	if m != nil {
		return m.Size
	}
	return nil
}

type AskPlanNetwork struct {
	ThroughputIn  *DataSizeRate `protobuf:"bytes,1,opt,name=throughputIn" json:"throughputIn,omitempty"`
	ThroughputOut *DataSizeRate `protobuf:"bytes,2,opt,name=throughputOut" json:"throughputOut,omitempty"`
	Overlay       bool          `protobuf:"varint,3,opt,name=overlay" json:"overlay,omitempty"`
	Outbound      bool          `protobuf:"varint,4,opt,name=outbound" json:"outbound,omitempty"`
	Incoming      bool          `protobuf:"varint,5,opt,name=incoming" json:"incoming,omitempty"`
}

func (m *AskPlanNetwork) Reset()                    { *m = AskPlanNetwork{} }
func (m *AskPlanNetwork) String() string            { return proto.CompactTextString(m) }
func (*AskPlanNetwork) ProtoMessage()               {}
func (*AskPlanNetwork) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AskPlanNetwork) GetThroughputIn() *DataSizeRate {
	if m != nil {
		return m.ThroughputIn
	}
	return nil
}

func (m *AskPlanNetwork) GetThroughputOut() *DataSizeRate {
	if m != nil {
		return m.ThroughputOut
	}
	return nil
}

func (m *AskPlanNetwork) GetOverlay() bool {
	if m != nil {
		return m.Overlay
	}
	return false
}

func (m *AskPlanNetwork) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

func (m *AskPlanNetwork) GetIncoming() bool {
	if m != nil {
		return m.Incoming
	}
	return false
}

type AskPlanResources struct {
	CPU     *AskPlanCPU     `protobuf:"bytes,1,opt,name=CPU" json:"CPU,omitempty"`
	RAM     *AskPlanRAM     `protobuf:"bytes,2,opt,name=RAM" json:"RAM,omitempty"`
	Storage *AskPlanStorage `protobuf:"bytes,3,opt,name=storage" json:"storage,omitempty"`
	GPU     *AskPlanGPU     `protobuf:"bytes,4,opt,name=GPU" json:"GPU,omitempty"`
	Network *AskPlanNetwork `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
}

func (m *AskPlanResources) Reset()                    { *m = AskPlanResources{} }
func (m *AskPlanResources) String() string            { return proto.CompactTextString(m) }
func (*AskPlanResources) ProtoMessage()               {}
func (*AskPlanResources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AskPlanResources) GetCPU() *AskPlanCPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *AskPlanResources) GetRAM() *AskPlanRAM {
	if m != nil {
		return m.RAM
	}
	return nil
}

func (m *AskPlanResources) GetStorage() *AskPlanStorage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *AskPlanResources) GetGPU() *AskPlanGPU {
	if m != nil {
		return m.GPU
	}
	return nil
}

func (m *AskPlanResources) GetNetwork() *AskPlanNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

type AskPlan struct {
	ID        string            `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	MarketID  string            `protobuf:"bytes,2,opt,name=marketID" json:"marketID,omitempty"`
	Duration  *Duration         `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	Price     *Price            `protobuf:"bytes,4,opt,name=price" json:"price,omitempty"`
	Blacklist *EthAddress       `protobuf:"bytes,5,opt,name=blacklist" json:"blacklist,omitempty"`
	Resources *AskPlanResources `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
}

func (m *AskPlan) Reset()                    { *m = AskPlan{} }
func (m *AskPlan) String() string            { return proto.CompactTextString(m) }
func (*AskPlan) ProtoMessage()               {}
func (*AskPlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AskPlan) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AskPlan) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func (m *AskPlan) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AskPlan) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *AskPlan) GetBlacklist() *EthAddress {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *AskPlan) GetResources() *AskPlanResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*AskPlanCPU)(nil), "sonm.AskPlanCPU")
	proto.RegisterType((*AskPlanGPU)(nil), "sonm.AskPlanGPU")
	proto.RegisterType((*AskPlanRAM)(nil), "sonm.AskPlanRAM")
	proto.RegisterType((*AskPlanStorage)(nil), "sonm.AskPlanStorage")
	proto.RegisterType((*AskPlanNetwork)(nil), "sonm.AskPlanNetwork")
	proto.RegisterType((*AskPlanResources)(nil), "sonm.AskPlanResources")
	proto.RegisterType((*AskPlan)(nil), "sonm.AskPlan")
}

func init() { proto.RegisterFile("ask_plan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x8a, 0xdb, 0x30,
	0x18, 0xc4, 0x49, 0x36, 0x89, 0xbf, 0xec, 0xa6, 0x8b, 0x28, 0x8b, 0xd9, 0x53, 0xea, 0x5e, 0x42,
	0x0f, 0xa6, 0xdd, 0x2e, 0xa5, 0xa7, 0x82, 0xd9, 0x94, 0x90, 0x43, 0x5a, 0xa3, 0x25, 0xe7, 0x45,
	0xb1, 0x45, 0x2c, 0xec, 0x48, 0x46, 0x92, 0xfb, 0xb3, 0xcf, 0xd9, 0x73, 0x1f, 0xa0, 0x4f, 0x51,
	0x24, 0xcb, 0x0e, 0x0e, 0x14, 0x7a, 0x0a, 0xf3, 0xcd, 0x8c, 0x3c, 0xdf, 0x48, 0x81, 0x39, 0x51,
	0xc5, 0x53, 0x55, 0x12, 0x1e, 0x55, 0x52, 0x68, 0x81, 0x46, 0x4a, 0xf0, 0xe3, 0xed, 0x0b, 0xc6,
	0xcd, 0x2f, 0x67, 0xa4, 0x19, 0x87, 0xef, 0x00, 0x62, 0x55, 0x24, 0x25, 0xe1, 0x0f, 0xc9, 0x0e,
	0xbd, 0x86, 0xab, 0x54, 0x48, 0xfa, 0x54, 0x51, 0x99, 0x52, 0xae, 0x55, 0xe0, 0x2d, 0xbc, 0xe5,
	0x08, 0x5f, 0x9a, 0x61, 0xe2, 0x66, 0xe1, 0xa7, 0xce, 0xb2, 0x4e, 0x76, 0x28, 0x80, 0x09, 0xe3,
	0x19, 0xfd, 0x41, 0x8d, 0x78, 0xb8, 0x1c, 0xe1, 0x16, 0xa2, 0x1b, 0x18, 0xe7, 0x44, 0xe5, 0x54,
	0x05, 0x83, 0xc5, 0x70, 0xe9, 0x63, 0x87, 0xc2, 0xb7, 0x9d, 0x1f, 0xc7, 0x5b, 0x14, 0xc2, 0x48,
	0xb1, 0x67, 0x6a, 0xbf, 0x34, 0xbb, 0x9b, 0x47, 0x26, 0x5e, 0xb4, 0x22, 0x9a, 0x3c, 0xb2, 0x67,
	0x8a, 0x2d, 0x17, 0xde, 0xc3, 0xdc, 0x39, 0x1e, 0xb5, 0x90, 0xe4, 0x40, 0xff, 0xcb, 0xf5, 0xcb,
	0xeb, 0x6c, 0x5f, 0xa8, 0xfe, 0x2e, 0x64, 0x81, 0x3e, 0xc0, 0xa5, 0xce, 0xa5, 0xa8, 0x0f, 0x79,
	0x55, 0xeb, 0x0d, 0x77, 0x76, 0x74, 0x66, 0x27, 0x9a, 0xe2, 0x9e, 0x0e, 0x7d, 0x84, 0xab, 0x13,
	0xfe, 0x5a, 0xeb, 0x60, 0xf0, 0x4f, 0x63, 0x5f, 0x68, 0xea, 0x11, 0xdf, 0xa8, 0x2c, 0xc9, 0xcf,
	0x60, 0xb8, 0xf0, 0x96, 0x53, 0xdc, 0x42, 0x74, 0x0b, 0x53, 0x51, 0xeb, 0xbd, 0xa8, 0x79, 0x16,
	0x8c, 0x2c, 0xd5, 0x61, 0xc3, 0x31, 0x9e, 0x8a, 0x23, 0xe3, 0x87, 0xe0, 0xa2, 0xe1, 0x5a, 0x1c,
	0xfe, 0xf6, 0xe0, 0xba, 0xed, 0x8f, 0x2a, 0x51, 0xcb, 0x94, 0x2a, 0x14, 0xc2, 0xf0, 0x21, 0xd9,
	0xb9, 0x7d, 0xae, 0x9b, 0x58, 0xa7, 0x7b, 0xc5, 0x86, 0x34, 0x1a, 0x1c, 0x6f, 0x5d, 0xf4, 0xbe,
	0x06, 0xc7, 0x5b, 0x6c, 0x48, 0x14, 0xc1, 0x44, 0x35, 0x15, 0xdb, 0xb8, 0xb3, 0xbb, 0x97, 0x3d,
	0x9d, 0xab, 0x1f, 0xb7, 0x22, 0x73, 0xe6, 0x3a, 0xd9, 0xd9, 0xfc, 0xe7, 0x67, 0xae, 0xcd, 0x77,
	0xcd, 0x0b, 0x89, 0x60, 0xc2, 0x9b, 0xfe, 0xed, 0x2e, 0xe7, 0x67, 0xba, 0xbb, 0xc1, 0xad, 0x28,
	0xfc, 0xe3, 0xc1, 0xc4, 0x71, 0x68, 0x0e, 0x83, 0xcd, 0xca, 0xae, 0xe5, 0xe3, 0xc1, 0x66, 0x65,
	0x8a, 0x39, 0x12, 0x59, 0x50, 0xbd, 0x59, 0xd9, 0x45, 0x7c, 0xdc, 0x61, 0xf4, 0x06, 0xa6, 0x59,
	0x2d, 0x89, 0x66, 0x82, 0xbb, 0xf0, 0xed, 0xbb, 0x70, 0x53, 0xdc, 0xf1, 0xe8, 0x15, 0x5c, 0x54,
	0x92, 0xa5, 0xd4, 0x25, 0x9f, 0x35, 0xc2, 0xc4, 0x8c, 0x70, 0xc3, 0xa0, 0x08, 0xfc, 0x7d, 0x49,
	0xd2, 0xa2, 0x64, 0x4a, 0xbb, 0xe0, 0x6e, 0xc1, 0xcf, 0x3a, 0x8f, 0xb3, 0x4c, 0x52, 0xa5, 0xf0,
	0x49, 0x82, 0xee, 0xc1, 0x97, 0xed, 0x7d, 0x04, 0x63, 0xab, 0xbf, 0xe9, 0x97, 0xdc, 0xb2, 0xf8,
	0x24, 0xdc, 0x8f, 0xed, 0xdf, 0xf0, 0xfd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x05, 0x76,
	0xf0, 0xaf, 0x03, 0x00, 0x00,
}
