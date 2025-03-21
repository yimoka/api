// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: common/common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Num struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         int32                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Num) Reset() {
	*x = Num{}
	mi := &file_common_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Num) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Num) ProtoMessage() {}

func (x *Num) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Num.ProtoReflect.Descriptor instead.
func (*Num) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *Num) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BigNum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         int64                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BigNum) Reset() {
	*x = BigNum{}
	mi := &file_common_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BigNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigNum) ProtoMessage() {}

func (x *BigNum) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigNum.ProtoReflect.Descriptor instead.
func (*BigNum) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *BigNum) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Bool struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         bool                   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bool) Reset() {
	*x = Bool{}
	mi := &file_common_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool) ProtoMessage() {}

func (x *Bool) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bool.ProtoReflect.Descriptor instead.
func (*Bool) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *Bool) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type Str struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Str) Reset() {
	*x = Str{}
	mi := &file_common_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Str) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Str) ProtoMessage() {}

func (x *Str) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Str.ProtoReflect.Descriptor instead.
func (*Str) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *Str) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type StrMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         map[string]string      `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StrMap) Reset() {
	*x = StrMap{}
	mi := &file_common_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StrMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrMap) ProtoMessage() {}

func (x *StrMap) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrMap.ProtoReflect.Descriptor instead.
func (*StrMap) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *StrMap) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

type NumMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         map[string]int32       `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NumMap) Reset() {
	*x = NumMap{}
	mi := &file_common_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumMap) ProtoMessage() {}

func (x *NumMap) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumMap.ProtoReflect.Descriptor instead.
func (*NumMap) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *NumMap) GetValue() map[string]int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type BigNumMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         map[string]int64       `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BigNumMap) Reset() {
	*x = BigNumMap{}
	mi := &file_common_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BigNumMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigNumMap) ProtoMessage() {}

func (x *BigNumMap) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigNumMap.ProtoReflect.Descriptor instead.
func (*BigNumMap) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *BigNumMap) GetValue() map[string]int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type BoolMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         map[string]bool        `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BoolMap) Reset() {
	*x = BoolMap{}
	mi := &file_common_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoolMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolMap) ProtoMessage() {}

func (x *BoolMap) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolMap.ProtoReflect.Descriptor instead.
func (*BoolMap) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *BoolMap) GetValue() map[string]bool {
	if x != nil {
		return x.Value
	}
	return nil
}

type AnyMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         *structpb.Struct       `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnyMap) Reset() {
	*x = AnyMap{}
	mi := &file_common_common_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyMap) ProtoMessage() {}

func (x *AnyMap) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyMap.ProtoReflect.Descriptor instead.
func (*AnyMap) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{8}
}

func (x *AnyMap) GetValue() *structpb.Struct {
	if x != nil {
		return x.Value
	}
	return nil
}

type StrArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []string               `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StrArray) Reset() {
	*x = StrArray{}
	mi := &file_common_common_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StrArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrArray) ProtoMessage() {}

func (x *StrArray) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrArray.ProtoReflect.Descriptor instead.
func (*StrArray) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{9}
}

func (x *StrArray) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type NumArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []int32                `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NumArray) Reset() {
	*x = NumArray{}
	mi := &file_common_common_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumArray) ProtoMessage() {}

func (x *NumArray) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumArray.ProtoReflect.Descriptor instead.
func (*NumArray) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{10}
}

func (x *NumArray) GetValue() []int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type BigNumArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []int64                `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BigNumArray) Reset() {
	*x = BigNumArray{}
	mi := &file_common_common_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BigNumArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigNumArray) ProtoMessage() {}

func (x *BigNumArray) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigNumArray.ProtoReflect.Descriptor instead.
func (*BigNumArray) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{11}
}

func (x *BigNumArray) GetValue() []int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type BoolArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []bool                 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BoolArray) Reset() {
	*x = BoolArray{}
	mi := &file_common_common_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoolArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolArray) ProtoMessage() {}

func (x *BoolArray) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolArray.ProtoReflect.Descriptor instead.
func (*BoolArray) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{12}
}

func (x *BoolArray) GetValue() []bool {
	if x != nil {
		return x.Value
	}
	return nil
}

// 选项
type Option struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Value      string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Label      string                 `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Desc       *string                `protobuf:"bytes,3,opt,name=desc,proto3,oneof" json:"desc,omitempty"`
	Url        *string                `protobuf:"bytes,4,opt,name=url,proto3,oneof" json:"url,omitempty"`
	Img        *string                `protobuf:"bytes,5,opt,name=img,proto3,oneof" json:"img,omitempty"`
	Icon       *string                `protobuf:"bytes,6,opt,name=icon,proto3,oneof" json:"icon,omitempty"`
	Color      *string                `protobuf:"bytes,7,opt,name=color,proto3,oneof" json:"color,omitempty"`
	Background *string                `protobuf:"bytes,8,opt,name=background,proto3,oneof" json:"background,omitempty"`
	// 激活态
	ActiveImg        *string `protobuf:"bytes,9,opt,name=activeImg,proto3,oneof" json:"activeImg,omitempty"`
	ActiveIcon       *string `protobuf:"bytes,10,opt,name=activeIcon,proto3,oneof" json:"activeIcon,omitempty"`
	ActiveColor      *string `protobuf:"bytes,11,opt,name=activeColor,proto3,oneof" json:"activeColor,omitempty"`
	ActiveBackground *string `protobuf:"bytes,12,opt,name=activeBackground,proto3,oneof" json:"activeBackground,omitempty"`
	Disabled         *bool   `protobuf:"varint,13,opt,name=disabled,proto3,oneof" json:"disabled,omitempty"`
	// 存放一些额外的数据
	Extra         *structpb.Struct `protobuf:"bytes,14,opt,name=extra,proto3" json:"extra,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Option) Reset() {
	*x = Option{}
	mi := &file_common_common_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{13}
}

func (x *Option) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Option) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Option) GetDesc() string {
	if x != nil && x.Desc != nil {
		return *x.Desc
	}
	return ""
}

func (x *Option) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *Option) GetImg() string {
	if x != nil && x.Img != nil {
		return *x.Img
	}
	return ""
}

func (x *Option) GetIcon() string {
	if x != nil && x.Icon != nil {
		return *x.Icon
	}
	return ""
}

func (x *Option) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *Option) GetBackground() string {
	if x != nil && x.Background != nil {
		return *x.Background
	}
	return ""
}

func (x *Option) GetActiveImg() string {
	if x != nil && x.ActiveImg != nil {
		return *x.ActiveImg
	}
	return ""
}

func (x *Option) GetActiveIcon() string {
	if x != nil && x.ActiveIcon != nil {
		return *x.ActiveIcon
	}
	return ""
}

func (x *Option) GetActiveColor() string {
	if x != nil && x.ActiveColor != nil {
		return *x.ActiveColor
	}
	return ""
}

func (x *Option) GetActiveBackground() string {
	if x != nil && x.ActiveBackground != nil {
		return *x.ActiveBackground
	}
	return ""
}

func (x *Option) GetDisabled() bool {
	if x != nil && x.Disabled != nil {
		return *x.Disabled
	}
	return false
}

func (x *Option) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

// 选项数据
type Options struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []*Option              `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Options) Reset() {
	*x = Options{}
	mi := &file_common_common_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{14}
}

func (x *Options) GetData() []*Option {
	if x != nil {
		return x.Data
	}
	return nil
}

// 查询排序
type SortOrder struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Order         string                 `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SortOrder) Reset() {
	*x = SortOrder{}
	mi := &file_common_common_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SortOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortOrder) ProtoMessage() {}

func (x *SortOrder) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortOrder.ProtoReflect.Descriptor instead.
func (*SortOrder) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{15}
}

func (x *SortOrder) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SortOrder) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x42, 0x69, 0x67, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x1b, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7e,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x4d,
	0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7e,
	0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6d, 0x4d,
	0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x84,
	0x01, 0x0a, 0x09, 0x42, 0x69, 0x67, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x12, 0x3d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79, 0x69,
	0x6d, 0x6f, 0x6b, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x42, 0x69, 0x67, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x4d, 0x61,
	0x70, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38,
	0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x06, 0x41, 0x6e, 0x79, 0x4d,
	0x61, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x20, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x4e, 0x75, 0x6d, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x42, 0x69, 0x67, 0x4e, 0x75, 0x6d, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x42, 0x6f,
	0x6f, 0x6c, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xce, 0x04,
	0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6d,
	0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x49, 0x6d, 0x67, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x08, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52,
	0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0a, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x69, 0x6d, 0x67, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6d, 0x67, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x13, 0x0a, 0x11,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x38,
	0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x3e, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x69, 0x6d, 0x6f, 0x6b, 0x61, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData []byte
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_common_proto_rawDesc), len(file_common_common_proto_rawDesc)))
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_common_common_proto_goTypes = []any{
	(*Num)(nil),             // 0: yimoka.api.common.Num
	(*BigNum)(nil),          // 1: yimoka.api.common.BigNum
	(*Bool)(nil),            // 2: yimoka.api.common.Bool
	(*Str)(nil),             // 3: yimoka.api.common.Str
	(*StrMap)(nil),          // 4: yimoka.api.common.StrMap
	(*NumMap)(nil),          // 5: yimoka.api.common.NumMap
	(*BigNumMap)(nil),       // 6: yimoka.api.common.BigNumMap
	(*BoolMap)(nil),         // 7: yimoka.api.common.BoolMap
	(*AnyMap)(nil),          // 8: yimoka.api.common.AnyMap
	(*StrArray)(nil),        // 9: yimoka.api.common.StrArray
	(*NumArray)(nil),        // 10: yimoka.api.common.NumArray
	(*BigNumArray)(nil),     // 11: yimoka.api.common.BigNumArray
	(*BoolArray)(nil),       // 12: yimoka.api.common.BoolArray
	(*Option)(nil),          // 13: yimoka.api.common.Option
	(*Options)(nil),         // 14: yimoka.api.common.Options
	(*SortOrder)(nil),       // 15: yimoka.api.common.SortOrder
	nil,                     // 16: yimoka.api.common.StrMap.ValueEntry
	nil,                     // 17: yimoka.api.common.NumMap.ValueEntry
	nil,                     // 18: yimoka.api.common.BigNumMap.ValueEntry
	nil,                     // 19: yimoka.api.common.BoolMap.ValueEntry
	(*structpb.Struct)(nil), // 20: google.protobuf.Struct
}
var file_common_common_proto_depIdxs = []int32{
	16, // 0: yimoka.api.common.StrMap.value:type_name -> yimoka.api.common.StrMap.ValueEntry
	17, // 1: yimoka.api.common.NumMap.value:type_name -> yimoka.api.common.NumMap.ValueEntry
	18, // 2: yimoka.api.common.BigNumMap.value:type_name -> yimoka.api.common.BigNumMap.ValueEntry
	19, // 3: yimoka.api.common.BoolMap.value:type_name -> yimoka.api.common.BoolMap.ValueEntry
	20, // 4: yimoka.api.common.AnyMap.value:type_name -> google.protobuf.Struct
	20, // 5: yimoka.api.common.Option.extra:type_name -> google.protobuf.Struct
	13, // 6: yimoka.api.common.Options.data:type_name -> yimoka.api.common.Option
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	file_common_common_proto_msgTypes[13].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_common_proto_rawDesc), len(file_common_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
