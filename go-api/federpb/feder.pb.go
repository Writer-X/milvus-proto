// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feder.proto

package federpb

import (
	fmt "fmt"
	commonpb "github.com/Writer-X/milvus-proto/go-api/v2/commonpb"
	proto "github.com/golang/protobuf/proto"
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

type SegmentIndexData struct {
	SegmentID            int64    `protobuf:"varint,1,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	IndexData            string   `protobuf:"bytes,2,opt,name=index_data,json=indexData,proto3" json:"index_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentIndexData) Reset()         { *m = SegmentIndexData{} }
func (m *SegmentIndexData) String() string { return proto.CompactTextString(m) }
func (*SegmentIndexData) ProtoMessage()    {}
func (*SegmentIndexData) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{0}
}

func (m *SegmentIndexData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentIndexData.Unmarshal(m, b)
}
func (m *SegmentIndexData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentIndexData.Marshal(b, m, deterministic)
}
func (m *SegmentIndexData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentIndexData.Merge(m, src)
}
func (m *SegmentIndexData) XXX_Size() int {
	return xxx_messageInfo_SegmentIndexData.Size(m)
}
func (m *SegmentIndexData) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentIndexData.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentIndexData proto.InternalMessageInfo

func (m *SegmentIndexData) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *SegmentIndexData) GetIndexData() string {
	if m != nil {
		return m.IndexData
	}
	return ""
}

type FederSegmentSearchResult struct {
	SegmentID            int64    `protobuf:"varint,1,opt,name=segmentID,proto3" json:"segmentID,omitempty"`
	VisitInfo            string   `protobuf:"bytes,2,opt,name=visit_info,json=visitInfo,proto3" json:"visit_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FederSegmentSearchResult) Reset()         { *m = FederSegmentSearchResult{} }
func (m *FederSegmentSearchResult) String() string { return proto.CompactTextString(m) }
func (*FederSegmentSearchResult) ProtoMessage()    {}
func (*FederSegmentSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{1}
}

func (m *FederSegmentSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederSegmentSearchResult.Unmarshal(m, b)
}
func (m *FederSegmentSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederSegmentSearchResult.Marshal(b, m, deterministic)
}
func (m *FederSegmentSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederSegmentSearchResult.Merge(m, src)
}
func (m *FederSegmentSearchResult) XXX_Size() int {
	return xxx_messageInfo_FederSegmentSearchResult.Size(m)
}
func (m *FederSegmentSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FederSegmentSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_FederSegmentSearchResult proto.InternalMessageInfo

func (m *FederSegmentSearchResult) GetSegmentID() int64 {
	if m != nil {
		return m.SegmentID
	}
	return 0
}

func (m *FederSegmentSearchResult) GetVisitInfo() string {
	if m != nil {
		return m.VisitInfo
	}
	return ""
}

type ListIndexedSegmentRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	CollectionName       string            `protobuf:"bytes,2,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	IndexName            string            `protobuf:"bytes,3,opt,name=index_name,json=indexName,proto3" json:"index_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListIndexedSegmentRequest) Reset()         { *m = ListIndexedSegmentRequest{} }
func (m *ListIndexedSegmentRequest) String() string { return proto.CompactTextString(m) }
func (*ListIndexedSegmentRequest) ProtoMessage()    {}
func (*ListIndexedSegmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{2}
}

func (m *ListIndexedSegmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIndexedSegmentRequest.Unmarshal(m, b)
}
func (m *ListIndexedSegmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIndexedSegmentRequest.Marshal(b, m, deterministic)
}
func (m *ListIndexedSegmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIndexedSegmentRequest.Merge(m, src)
}
func (m *ListIndexedSegmentRequest) XXX_Size() int {
	return xxx_messageInfo_ListIndexedSegmentRequest.Size(m)
}
func (m *ListIndexedSegmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIndexedSegmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListIndexedSegmentRequest proto.InternalMessageInfo

func (m *ListIndexedSegmentRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *ListIndexedSegmentRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *ListIndexedSegmentRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type ListIndexedSegmentResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SegmentIDs           []int64          `protobuf:"varint,2,rep,packed,name=segmentIDs,proto3" json:"segmentIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListIndexedSegmentResponse) Reset()         { *m = ListIndexedSegmentResponse{} }
func (m *ListIndexedSegmentResponse) String() string { return proto.CompactTextString(m) }
func (*ListIndexedSegmentResponse) ProtoMessage()    {}
func (*ListIndexedSegmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{3}
}

func (m *ListIndexedSegmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIndexedSegmentResponse.Unmarshal(m, b)
}
func (m *ListIndexedSegmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIndexedSegmentResponse.Marshal(b, m, deterministic)
}
func (m *ListIndexedSegmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIndexedSegmentResponse.Merge(m, src)
}
func (m *ListIndexedSegmentResponse) XXX_Size() int {
	return xxx_messageInfo_ListIndexedSegmentResponse.Size(m)
}
func (m *ListIndexedSegmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIndexedSegmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListIndexedSegmentResponse proto.InternalMessageInfo

func (m *ListIndexedSegmentResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListIndexedSegmentResponse) GetSegmentIDs() []int64 {
	if m != nil {
		return m.SegmentIDs
	}
	return nil
}

type DescribeSegmentIndexDataRequest struct {
	Base                 *commonpb.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	CollectionName       string            `protobuf:"bytes,2,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	IndexName            string            `protobuf:"bytes,3,opt,name=index_name,json=indexName,proto3" json:"index_name,omitempty"`
	SegmentsIDs          []int64           `protobuf:"varint,4,rep,packed,name=segmentsIDs,proto3" json:"segmentsIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DescribeSegmentIndexDataRequest) Reset()         { *m = DescribeSegmentIndexDataRequest{} }
func (m *DescribeSegmentIndexDataRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeSegmentIndexDataRequest) ProtoMessage()    {}
func (*DescribeSegmentIndexDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{4}
}

func (m *DescribeSegmentIndexDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeSegmentIndexDataRequest.Unmarshal(m, b)
}
func (m *DescribeSegmentIndexDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeSegmentIndexDataRequest.Marshal(b, m, deterministic)
}
func (m *DescribeSegmentIndexDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeSegmentIndexDataRequest.Merge(m, src)
}
func (m *DescribeSegmentIndexDataRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeSegmentIndexDataRequest.Size(m)
}
func (m *DescribeSegmentIndexDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeSegmentIndexDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeSegmentIndexDataRequest proto.InternalMessageInfo

func (m *DescribeSegmentIndexDataRequest) GetBase() *commonpb.MsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *DescribeSegmentIndexDataRequest) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *DescribeSegmentIndexDataRequest) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *DescribeSegmentIndexDataRequest) GetSegmentsIDs() []int64 {
	if m != nil {
		return m.SegmentsIDs
	}
	return nil
}

type DescribeSegmentIndexDataResponse struct {
	Status *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// segmentID => segmentIndexData
	IndexData            map[int64]*SegmentIndexData `protobuf:"bytes,2,rep,name=index_data,json=indexData,proto3" json:"index_data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IndexParams          []*commonpb.KeyValuePair    `protobuf:"bytes,3,rep,name=index_params,json=indexParams,proto3" json:"index_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DescribeSegmentIndexDataResponse) Reset()         { *m = DescribeSegmentIndexDataResponse{} }
func (m *DescribeSegmentIndexDataResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeSegmentIndexDataResponse) ProtoMessage()    {}
func (*DescribeSegmentIndexDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d670fd126c7825, []int{5}
}

func (m *DescribeSegmentIndexDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeSegmentIndexDataResponse.Unmarshal(m, b)
}
func (m *DescribeSegmentIndexDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeSegmentIndexDataResponse.Marshal(b, m, deterministic)
}
func (m *DescribeSegmentIndexDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeSegmentIndexDataResponse.Merge(m, src)
}
func (m *DescribeSegmentIndexDataResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeSegmentIndexDataResponse.Size(m)
}
func (m *DescribeSegmentIndexDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeSegmentIndexDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeSegmentIndexDataResponse proto.InternalMessageInfo

func (m *DescribeSegmentIndexDataResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeSegmentIndexDataResponse) GetIndexData() map[int64]*SegmentIndexData {
	if m != nil {
		return m.IndexData
	}
	return nil
}

func (m *DescribeSegmentIndexDataResponse) GetIndexParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.IndexParams
	}
	return nil
}

func init() {
	proto.RegisterType((*SegmentIndexData)(nil), "milvus.proto.feder.SegmentIndexData")
	proto.RegisterType((*FederSegmentSearchResult)(nil), "milvus.proto.feder.FederSegmentSearchResult")
	proto.RegisterType((*ListIndexedSegmentRequest)(nil), "milvus.proto.feder.ListIndexedSegmentRequest")
	proto.RegisterType((*ListIndexedSegmentResponse)(nil), "milvus.proto.feder.ListIndexedSegmentResponse")
	proto.RegisterType((*DescribeSegmentIndexDataRequest)(nil), "milvus.proto.feder.DescribeSegmentIndexDataRequest")
	proto.RegisterType((*DescribeSegmentIndexDataResponse)(nil), "milvus.proto.feder.DescribeSegmentIndexDataResponse")
	proto.RegisterMapType((map[int64]*SegmentIndexData)(nil), "milvus.proto.feder.DescribeSegmentIndexDataResponse.IndexDataEntry")
}

func init() { proto.RegisterFile("feder.proto", fileDescriptor_84d670fd126c7825) }

var fileDescriptor_84d670fd126c7825 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0x52, 0x29, 0xe3, 0xaa, 0x54, 0x7b, 0x32, 0xa1, 0x80, 0xb1, 0x90, 0xc8, 0x25,
	0x36, 0x72, 0x39, 0xa0, 0x1e, 0x4b, 0x40, 0xaa, 0xf8, 0xab, 0x1c, 0x89, 0x22, 0x2e, 0xd1, 0xda,
	0x99, 0xa4, 0x2b, 0xec, 0x5d, 0x77, 0x77, 0x1d, 0x91, 0x07, 0xe1, 0x71, 0x78, 0x1c, 0xde, 0x03,
	0x79, 0xbd, 0x6d, 0xd2, 0xd6, 0x28, 0x07, 0x0e, 0xbd, 0xd9, 0xdf, 0xcc, 0x7c, 0x3f, 0xb3, 0xbb,
	0xe0, 0xcd, 0x71, 0x86, 0x32, 0xaa, 0xa4, 0xd0, 0x82, 0x90, 0x92, 0x15, 0xcb, 0x5a, 0xb5, 0x7f,
	0x91, 0xa9, 0x0c, 0xf6, 0x72, 0x51, 0x96, 0x82, 0xb7, 0x58, 0xf8, 0x05, 0x0e, 0x26, 0xb8, 0x28,
	0x91, 0xeb, 0x53, 0x3e, 0xc3, 0x9f, 0x63, 0xaa, 0x29, 0x39, 0x84, 0xbe, 0xb2, 0xd8, 0xd8, 0x77,
	0x02, 0x67, 0xe8, 0xa6, 0x6b, 0x80, 0x3c, 0x01, 0x60, 0x4d, 0xeb, 0x74, 0x46, 0x35, 0xf5, 0x7b,
	0x81, 0x33, 0xec, 0xa7, 0x7d, 0x76, 0x35, 0x1c, 0x9e, 0x83, 0xff, 0xbe, 0xd1, 0xb1, 0xac, 0x13,
	0xa4, 0x32, 0xbf, 0x48, 0x51, 0xd5, 0x85, 0xde, 0x4e, 0xbc, 0x64, 0x8a, 0xe9, 0x29, 0xe3, 0x73,
	0x71, 0x45, 0x6c, 0x90, 0x53, 0x3e, 0x17, 0xe1, 0x2f, 0x07, 0x1e, 0x7d, 0x64, 0xaa, 0xf5, 0x89,
	0x33, 0xcb, 0x9f, 0xe2, 0x65, 0x8d, 0x4a, 0x93, 0x57, 0xb0, 0x93, 0x51, 0x85, 0x86, 0xd5, 0x4b,
	0x0e, 0xa3, 0x1b, 0xc1, 0x6d, 0xe2, 0x4f, 0x6a, 0x71, 0x42, 0x15, 0xa6, 0xa6, 0x93, 0xbc, 0x84,
	0x87, 0xb9, 0x28, 0x0a, 0xcc, 0x35, 0x13, 0x7c, 0xca, 0x69, 0x89, 0x56, 0x73, 0x7f, 0x0d, 0x7f,
	0xa6, 0x25, 0xae, 0x03, 0x9b, 0x1e, 0x77, 0x23, 0x70, 0x53, 0x0e, 0x2f, 0x61, 0xd0, 0x65, 0x4b,
	0x55, 0x82, 0x2b, 0x24, 0x47, 0xb0, 0xab, 0x34, 0xd5, 0xb5, 0xb2, 0xce, 0x1e, 0x77, 0x3a, 0x9b,
	0x98, 0x96, 0xd4, 0xb6, 0x92, 0xa7, 0x00, 0xd7, 0x6b, 0x51, 0x7e, 0x2f, 0x70, 0x87, 0x6e, 0xba,
	0x81, 0x84, 0xbf, 0x1d, 0x78, 0x36, 0x46, 0x95, 0x4b, 0x96, 0xe1, 0xed, 0xd3, 0xbb, 0xf7, 0x85,
	0x90, 0x00, 0x3c, 0xeb, 0x55, 0x35, 0xf6, 0x77, 0x8c, 0xfd, 0x4d, 0x28, 0xfc, 0xd3, 0x83, 0xe0,
	0xdf, 0xfe, 0xff, 0x67, 0x73, 0xd9, 0xad, 0xcb, 0xe9, 0x0e, 0xbd, 0xe4, 0x6d, 0x74, 0xf7, 0x15,
	0x44, 0xdb, 0xe4, 0xa3, 0x6b, 0xe4, 0x1d, 0xd7, 0x72, 0xb5, 0x71, 0xc3, 0xc9, 0x18, 0xf6, 0x5a,
	0x8d, 0x8a, 0x4a, 0x5a, 0x2a, 0xdf, 0x35, 0x2a, 0xcf, 0x3b, 0xed, 0x7d, 0xc0, 0xd5, 0x57, 0x5a,
	0xd4, 0x78, 0x46, 0x99, 0x4c, 0x3d, 0x33, 0x76, 0x66, 0xa6, 0x06, 0x19, 0xec, 0xdf, 0x94, 0x20,
	0x07, 0xe0, 0xfe, 0xc0, 0x95, 0x7d, 0x17, 0xcd, 0x27, 0x39, 0x86, 0x07, 0xcb, 0x66, 0xda, 0x9c,
	0x83, 0x97, 0xbc, 0xe8, 0x0a, 0x72, 0x27, 0x40, 0x3b, 0x72, 0xdc, 0x7b, 0xe3, 0x9c, 0xbc, 0xfe,
	0x9e, 0x2c, 0x98, 0xbe, 0xa8, 0xb3, 0xc6, 0x4e, 0x7c, 0x2e, 0x99, 0x46, 0x39, 0xfa, 0x16, 0xb7,
	0x2c, 0x23, 0xc3, 0x12, 0x2f, 0xc4, 0x88, 0x56, 0x2c, 0x5e, 0x26, 0xb1, 0xe1, 0xab, 0xb2, 0x6c,
	0xd7, 0x14, 0x8e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xe7, 0x3e, 0x88, 0x4a, 0x04, 0x00,
	0x00,
}
