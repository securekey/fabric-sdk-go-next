/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/collection.proto
package common // import "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
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
// CollectionType enumerates the various types of private data collections.
type CollectionType int32
const (
	CollectionType_COL_UNKNOWN   CollectionType = 0
	CollectionType_COL_PRIVATE   CollectionType = 1
	CollectionType_COL_TRANSIENT CollectionType = 2
	CollectionType_COL_OFFLEDGER CollectionType = 3
	CollectionType_COL_DCAS      CollectionType = 4
)
var CollectionType_name = map[int32]string{
	0: "COL_UNKNOWN",
	1: "COL_PRIVATE",
	2: "COL_TRANSIENT",
	3: "COL_OFFLEDGER",
	4: "COL_DCAS",
}
var CollectionType_value = map[string]int32{
	"COL_UNKNOWN":   0,
	"COL_PRIVATE":   1,
	"COL_TRANSIENT": 2,
	"COL_OFFLEDGER": 3,
	"COL_DCAS":      4,
}
func (x CollectionType) String() string {
	return proto.EnumName(CollectionType_name, int32(x))
}
func (CollectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_collection_59b2e02e8b8b23b5, []int{0}
}
// CollectionConfigPackage represents an array of CollectionConfig
// messages; the extra struct is required because repeated oneof is
// forbidden by the protobuf syntax
type CollectionConfigPackage struct {
	Config               []*CollectionConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
func (m *CollectionConfigPackage) Reset()         { *m = CollectionConfigPackage{} }
func (m *CollectionConfigPackage) String() string { return proto.CompactTextString(m) }
func (*CollectionConfigPackage) ProtoMessage()    {}
func (*CollectionConfigPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_collection_59b2e02e8b8b23b5, []int{0}
}
func (m *CollectionConfigPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfigPackage.Unmarshal(m, b)
}
func (m *CollectionConfigPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfigPackage.Marshal(b, m, deterministic)
}
func (dst *CollectionConfigPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfigPackage.Merge(dst, src)
}
func (m *CollectionConfigPackage) XXX_Size() int {
	return xxx_messageInfo_CollectionConfigPackage.Size(m)
}
func (m *CollectionConfigPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfigPackage.DiscardUnknown(m)
}
var xxx_messageInfo_CollectionConfigPackage proto.InternalMessageInfo
func (m *CollectionConfigPackage) GetConfig() []*CollectionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}
// CollectionConfig defines the configuration of a collection object;
// it currently contains a single, static type.
// Dynamic collections are deferred.
type CollectionConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionConfig_StaticCollectionConfig
	Payload              isCollectionConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}
func (m *CollectionConfig) Reset()         { *m = CollectionConfig{} }
func (m *CollectionConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionConfig) ProtoMessage()    {}
func (*CollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_collection_59b2e02e8b8b23b5, []int{1}
}
func (m *CollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfig.Unmarshal(m, b)
}
func (m *CollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfig.Marshal(b, m, deterministic)
}
func (dst *CollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfig.Merge(dst, src)
}
func (m *CollectionConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionConfig.Size(m)
}
func (m *CollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfig.DiscardUnknown(m)
}
var xxx_messageInfo_CollectionConfig proto.InternalMessageInfo
type isCollectionConfig_Payload interface {
	isCollectionConfig_Payload()
}
type CollectionConfig_StaticCollectionConfig struct {
	StaticCollectionConfig *StaticCollectionConfig `protobuf:"bytes,1,opt,name=static_collection_config,json=staticCollectionConfig,proto3,oneof"`
}
func (*CollectionConfig_StaticCollectionConfig) isCollectionConfig_Payload() {}
func (m *CollectionConfig) GetPayload() isCollectionConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}
func (m *CollectionConfig) GetStaticCollectionConfig() *StaticCollectionConfig {
	if x, ok := m.GetPayload().(*CollectionConfig_StaticCollectionConfig); ok {
		return x.StaticCollectionConfig
	}
	return nil
}
// XXX_OneofFuncs is for the internal use of the proto package.
func (*CollectionConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CollectionConfig_OneofMarshaler, _CollectionConfig_OneofUnmarshaler, _CollectionConfig_OneofSizer, []interface{}{
		(*CollectionConfig_StaticCollectionConfig)(nil),
	}
}
func _CollectionConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CollectionConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionConfig_StaticCollectionConfig:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StaticCollectionConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CollectionConfig.Payload has unexpected type %T", x)
	}
	return nil
}
func _CollectionConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CollectionConfig)
	switch tag {
	case 1: // payload.static_collection_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StaticCollectionConfig)
		err := b.DecodeMessage(msg)
		m.Payload = &CollectionConfig_StaticCollectionConfig{msg}
		return true, err
	default:
		return false, nil
	}
}
func _CollectionConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CollectionConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionConfig_StaticCollectionConfig:
		s := proto.Size(x.StaticCollectionConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}
// StaticCollectionConfig constitutes the configuration parameters of a
// static collection object. Static collections are collections that are
// known at chaincode instantiation time, and that cannot be changed.
// Dynamic collections are deferred.
type StaticCollectionConfig struct {
	// the name of the collection inside the denoted chaincode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define which orgs have access to this collection’s private data
	MemberOrgsPolicy *CollectionPolicyConfig `protobuf:"bytes,2,opt,name=member_orgs_policy,json=memberOrgsPolicy,proto3" json:"member_orgs_policy,omitempty"`
	// The minimum number of peers private data will be sent to upon
	// endorsement. The endorsement would fail if dissemination to at least
	// this number of peers is not achieved.
	RequiredPeerCount int32 `protobuf:"varint,3,opt,name=required_peer_count,json=requiredPeerCount,proto3" json:"required_peer_count,omitempty"`
	// The maximum number of peers that private data will be sent to
	// upon endorsement. This number has to be bigger than required_peer_count.
	MaximumPeerCount int32 `protobuf:"varint,4,opt,name=maximum_peer_count,json=maximumPeerCount,proto3" json:"maximum_peer_count,omitempty"`
	// The number of blocks after which the collection data expires.
	// For instance if the value is set to 10, a key last modified by block number 100
	// will be purged at block number 111. A zero value is treated same as MaxUint64
	BlockToLive uint64 `protobuf:"varint,5,opt,name=block_to_live,json=blockToLive,proto3" json:"block_to_live,omitempty"`
	// The member only read access denotes whether only collection member clients
	// can read the private data (if set to true), or even non members can
	// read the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyRead bool `protobuf:"varint,6,opt,name=member_only_read,json=memberOnlyRead,proto3" json:"member_only_read,omitempty"`
	// The type of collection.
	Type CollectionType `protobuf:"varint,9900,opt,name=type,proto3,enum=common.CollectionType" json:"type,omitempty"`
	// The time after which the collection data expires. For example,
	// if the value is set to "10m" then the data will be purged
	// 10 minutes after it was stored. An empty value indicates that
	// the data should never be purged.
	// The format of this string must be parseable by time.ParseDuration
	TimeToLive           string   `protobuf:"bytes,9901,opt,name=time_to_live,json=timeToLive,proto3" json:"time_to_live,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
func (m *StaticCollectionConfig) Reset()         { *m = StaticCollectionConfig{} }
func (m *StaticCollectionConfig) String() string { return proto.CompactTextString(m) }
func (*StaticCollectionConfig) ProtoMessage()    {}
func (*StaticCollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_collection_59b2e02e8b8b23b5, []int{2}
}
func (m *StaticCollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticCollectionConfig.Unmarshal(m, b)
}
func (m *StaticCollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticCollectionConfig.Marshal(b, m, deterministic)
}
func (dst *StaticCollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticCollectionConfig.Merge(dst, src)
}
func (m *StaticCollectionConfig) XXX_Size() int {
	return xxx_messageInfo_StaticCollectionConfig.Size(m)
}
func (m *StaticCollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticCollectionConfig.DiscardUnknown(m)
}
var xxx_messageInfo_StaticCollectionConfig proto.InternalMessageInfo
func (m *StaticCollectionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *StaticCollectionConfig) GetMemberOrgsPolicy() *CollectionPolicyConfig {
	if m != nil {
		return m.MemberOrgsPolicy
	}
	return nil
}
func (m *StaticCollectionConfig) GetRequiredPeerCount() int32 {
	if m != nil {
		return m.RequiredPeerCount
	}
	return 0
}
func (m *StaticCollectionConfig) GetMaximumPeerCount() int32 {
	if m != nil {
		return m.MaximumPeerCount
	}
	return 0
}
func (m *StaticCollectionConfig) GetBlockToLive() uint64 {
	if m != nil {
		return m.BlockToLive
	}
	return 0
}
func (m *StaticCollectionConfig) GetMemberOnlyRead() bool {
	if m != nil {
		return m.MemberOnlyRead
	}
	return false
}
func (m *StaticCollectionConfig) GetType() CollectionType {
	if m != nil {
		return m.Type
	}
	return CollectionType_COL_UNKNOWN
}
func (m *StaticCollectionConfig) GetTimeToLive() string {
	if m != nil {
		return m.TimeToLive
	}
	return ""
}
// Collection policy configuration. Initially, the configuration can only
// contain a SignaturePolicy. In the future, the SignaturePolicy may be a
// more general Policy. Instead of containing the actual policy, the
// configuration may in the future contain a string reference to a policy.
type CollectionPolicyConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionPolicyConfig_SignaturePolicy
	Payload              isCollectionPolicyConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}
func (m *CollectionPolicyConfig) Reset()         { *m = CollectionPolicyConfig{} }
func (m *CollectionPolicyConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionPolicyConfig) ProtoMessage()    {}
func (*CollectionPolicyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_collection_59b2e02e8b8b23b5, []int{3}
}
func (m *CollectionPolicyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionPolicyConfig.Unmarshal(m, b)
}
func (m *CollectionPolicyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionPolicyConfig.Marshal(b, m, deterministic)
}
func (dst *CollectionPolicyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionPolicyConfig.Merge(dst, src)
}
func (m *CollectionPolicyConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionPolicyConfig.Size(m)
}
func (m *CollectionPolicyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionPolicyConfig.DiscardUnknown(m)
}
var xxx_messageInfo_CollectionPolicyConfig proto.InternalMessageInfo
type isCollectionPolicyConfig_Payload interface {
	isCollectionPolicyConfig_Payload()
}
type CollectionPolicyConfig_SignaturePolicy struct {
	SignaturePolicy *SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}
func (*CollectionPolicyConfig_SignaturePolicy) isCollectionPolicyConfig_Payload() {}
func (m *CollectionPolicyConfig) GetPayload() isCollectionPolicyConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}
func (m *CollectionPolicyConfig) GetSignaturePolicy() *SignaturePolicyEnvelope {
	if x, ok := m.GetPayload().(*CollectionPolicyConfig_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}
// XXX_OneofFuncs is for the internal use of the proto package.
func (*CollectionPolicyConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CollectionPolicyConfig_OneofMarshaler, _CollectionPolicyConfig_OneofUnmarshaler, _CollectionPolicyConfig_OneofSizer, []interface{}{
		(*CollectionPolicyConfig_SignaturePolicy)(nil),
	}
}
func _CollectionPolicyConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CollectionPolicyConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionPolicyConfig_SignaturePolicy:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SignaturePolicy); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CollectionPolicyConfig.Payload has unexpected type %T", x)
	}
	return nil
}
func _CollectionPolicyConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CollectionPolicyConfig)
	switch tag {
	case 1: // payload.signature_policy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicyEnvelope)
		err := b.DecodeMessage(msg)
		m.Payload = &CollectionPolicyConfig_SignaturePolicy{msg}
		return true, err
	default:
		return false, nil
	}
}
func _CollectionPolicyConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CollectionPolicyConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionPolicyConfig_SignaturePolicy:
		s := proto.Size(x.SignaturePolicy)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}
// CollectionCriteria defines an element of a private data that corresponds
// to a certain transaction and collection
type CollectionCriteria struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	TxId                 string   `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Collection           string   `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	Namespace            string   `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
func (m *CollectionCriteria) Reset()         { *m = CollectionCriteria{} }
func (m *CollectionCriteria) String() string { return proto.CompactTextString(m) }
func (*CollectionCriteria) ProtoMessage()    {}
func (*CollectionCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_collection_59b2e02e8b8b23b5, []int{4}
}
func (m *CollectionCriteria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionCriteria.Unmarshal(m, b)
}
func (m *CollectionCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionCriteria.Marshal(b, m, deterministic)
}
func (dst *CollectionCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionCriteria.Merge(dst, src)
}
func (m *CollectionCriteria) XXX_Size() int {
	return xxx_messageInfo_CollectionCriteria.Size(m)
}
func (m *CollectionCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionCriteria.DiscardUnknown(m)
}
var xxx_messageInfo_CollectionCriteria proto.InternalMessageInfo
func (m *CollectionCriteria) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}
func (m *CollectionCriteria) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}
func (m *CollectionCriteria) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}
func (m *CollectionCriteria) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}
func init() {
	proto.RegisterType((*CollectionConfigPackage)(nil), "sdk.common.CollectionConfigPackage")
	proto.RegisterType((*CollectionConfig)(nil), "sdk.common.CollectionConfig")
	proto.RegisterType((*StaticCollectionConfig)(nil), "sdk.common.StaticCollectionConfig")
	proto.RegisterType((*CollectionPolicyConfig)(nil), "sdk.common.CollectionPolicyConfig")
	proto.RegisterType((*CollectionCriteria)(nil), "sdk.common.CollectionCriteria")
	proto.RegisterEnum("sdk.common.CollectionType", CollectionType_name, CollectionType_value)

}
func init() { proto.RegisterFile("common/collection.proto", fileDescriptor_collection_59b2e02e8b8b23b5) }
var fileDescriptor_collection_59b2e02e8b8b23b5 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5d, 0x4f, 0xdb, 0x30,
	0x14, 0x25, 0x50, 0x0a, 0xbd, 0xe5, 0x23, 0x18, 0xad, 0x44, 0xd3, 0xc4, 0xba, 0x6a, 0x0f, 0xd1,
	0x98, 0xda, 0x89, 0xfd, 0x02, 0x28, 0x65, 0x20, 0x4a, 0x5b, 0xb9, 0xdd, 0x26, 0xf1, 0x12, 0xb9,
	0xc9, 0x25, 0x58, 0x24, 0x71, 0x70, 0x5c, 0x44, 0x1e, 0xf7, 0x7f, 0xb6, 0xbf, 0xb7, 0xe7, 0x29,
	0x4e, 0xd2, 0x16, 0xc6, 0x5b, 0x7c, 0xce, 0xb9, 0xd7, 0xf7, 0x9e, 0xe3, 0xc0, 0x81, 0x2b, 0xc2,
	0x50, 0x44, 0x1d, 0x57, 0x04, 0x01, 0xba, 0x8a, 0x8b, 0xa8, 0x1d, 0x4b, 0xa1, 0x04, 0xa9, 0xe6,
	0xc4, 0xdb, 0x37, 0x85, 0x20, 0x16, 0x01, 0x77, 0x39, 0x26, 0x39, 0xdd, 0xba, 0x82, 0x83, 0xee,
	0xbc, 0xa4, 0x2b, 0xa2, 0x5b, 0xee, 0x8f, 0x98, 0x7b, 0xcf, 0x7c, 0x24, 0x5f, 0xa0, 0xea, 0x6a,
	0xc0, 0x32, 0x9a, 0x6b, 0x76, 0xfd, 0xd8, 0x6a, 0xe7, 0x2d, 0xda, 0x2f, 0x0b, 0x68, 0xa1, 0x6b,
	0xa5, 0x60, 0xbe, 0xe4, 0xc8, 0x0d, 0x58, 0x89, 0x62, 0x8a, 0xbb, 0xce, 0x62, 0x34, 0x67, 0xde,
	0xd7, 0xb0, 0xeb, 0xc7, 0x87, 0x65, 0xdf, 0xb1, 0xd6, 0xbd, 0xec, 0x70, 0xb1, 0x42, 0x1b, 0xc9,
	0xab, 0xcc, 0x69, 0x0d, 0x36, 0x62, 0x96, 0x06, 0x82, 0x79, 0xad, 0xbf, 0xab, 0xd0, 0x78, 0xbd,
	0x9e, 0x10, 0xa8, 0x44, 0x2c, 0x44, 0x7d, 0x5b, 0x8d, 0xea, 0x6f, 0xd2, 0x07, 0x12, 0x62, 0x38,
	0x45, 0xe9, 0x08, 0xe9, 0x27, 0x8e, 0x36, 0x25, 0xb5, 0x56, 0x9f, 0xcf, 0xb3, 0xe8, 0x34, 0xd2,
	0x7c, 0xb1, 0xad, 0x99, 0x57, 0x0e, 0xa5, 0x9f, 0xe4, 0x38, 0x69, 0xc3, 0xbe, 0xc4, 0x87, 0x19,
	0x97, 0xe8, 0x39, 0x31, 0xa2, 0x74, 0x5c, 0x31, 0x8b, 0x94, 0xb5, 0xd6, 0x34, 0xec, 0x75, 0xba,
	0x57, 0x52, 0x23, 0x44, 0xd9, 0xcd, 0x08, 0xf2, 0x19, 0x48, 0xc8, 0x9e, 0x78, 0x38, 0x0b, 0x97,
	0xe5, 0x15, 0x2d, 0x37, 0x0b, 0x66, 0xa1, 0x6e, 0xc1, 0xf6, 0x34, 0x10, 0xee, 0xbd, 0xa3, 0x84,
	0x13, 0xf0, 0x47, 0xb4, 0xd6, 0x9b, 0x86, 0x5d, 0xa1, 0x75, 0x0d, 0x4e, 0x44, 0x9f, 0x3f, 0x22,
	0xb1, 0xc1, 0x2c, 0xf7, 0x89, 0x82, 0xd4, 0x91, 0xc8, 0x3c, 0xab, 0xda, 0x34, 0xec, 0x4d, 0xba,
	0x53, 0x4c, 0x1b, 0x05, 0x29, 0x45, 0xe6, 0x91, 0x23, 0xa8, 0xa8, 0x34, 0x46, 0xeb, 0xf7, 0x75,
	0xd3, 0xb0, 0x77, 0x8e, 0x1b, 0xff, 0x2f, 0x3b, 0x49, 0x63, 0xa4, 0x5a, 0x44, 0x3e, 0xc0, 0x96,
	0xe2, 0x21, 0xce, 0x6f, 0xfe, 0x73, 0xad, 0x3d, 0x84, 0x0c, 0xcc, 0x6f, 0x6e, 0x3d, 0x40, 0xe3,
	0x75, 0x9f, 0x48, 0x1f, 0xcc, 0x84, 0xfb, 0x11, 0x53, 0x33, 0x89, 0xa5, 0xc3, 0x79, 0xe2, 0xef,
	0xe7, 0x89, 0x97, 0x7c, 0x5e, 0xd8, 0x8b, 0x1e, 0x31, 0x10, 0x31, 0x5e, 0xac, 0xd0, 0xdd, 0xe4,
	0x39, 0xb5, 0x9c, 0xf5, 0x2f, 0x03, 0xc8, 0x52, 0xca, 0x92, 0x2b, 0x94, 0x9c, 0x11, 0x0b, 0x36,
	0xdc, 0x3b, 0x16, 0x45, 0x18, 0x14, 0x51, 0x97, 0x47, 0xb2, 0x0f, 0xeb, 0xea, 0xc9, 0xe1, 0x9e,
	0x0e, 0xb8, 0x46, 0x2b, 0xea, 0xe9, 0xd2, 0x23, 0x87, 0x00, 0x8b, 0x17, 0xa9, 0xb3, 0xaa, 0xd1,
	0x25, 0x84, 0xbc, 0x83, 0x5a, 0xf6, 0x54, 0x92, 0x98, 0xb9, 0xa8, 0xb3, 0xa9, 0xd1, 0x05, 0xf0,
	0xe9, 0x16, 0x76, 0x9e, 0x3b, 0x46, 0x76, 0xa1, 0xde, 0x1d, 0xf6, 0x9d, 0xef, 0x83, 0xab, 0xc1,
	0xf0, 0xe7, 0xc0, 0x5c, 0x29, 0x81, 0x11, 0xbd, 0xfc, 0x71, 0x32, 0xe9, 0x99, 0x06, 0xd9, 0x83,
	0xed, 0x0c, 0x98, 0xd0, 0x93, 0xc1, 0xf8, 0xb2, 0x37, 0x98, 0x98, 0xab, 0x25, 0x34, 0x3c, 0x3f,
	0xef, 0xf7, 0xce, 0xbe, 0xf5, 0xa8, 0xb9, 0x46, 0xb6, 0x60, 0x33, 0x83, 0xce, 0xba, 0x27, 0x63,
	0xb3, 0x72, 0x3a, 0x86, 0x8f, 0x42, 0xfa, 0xed, 0xbb, 0x34, 0x46, 0x19, 0xa0, 0xe7, 0xa3, 0x6c,
	0xdf, 0xb2, 0xa9, 0xe4, 0x6e, 0xfe, 0xff, 0x26, 0x85, 0x93, 0x37, 0x47, 0x3e, 0x57, 0x77, 0xb3,
	0x69, 0x76, 0xec, 0x2c, 0x89, 0x3b, 0xb9, 0xb8, 0x93, 0x8b, 0x3b, 0xb9, 0x78, 0x5a, 0xd5, 0xc7,
	0xaf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x1a, 0x4f, 0xba, 0x35, 0x04, 0x00, 0x00,
}