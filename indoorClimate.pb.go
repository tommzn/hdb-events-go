// HomeDashboard Indoor Climate Event Schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: indoorClimate.proto

package events

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Measurement types.
type MeasurementType int32

const (
	MeasurementType_TEMPERATURE MeasurementType = 0
	MeasurementType_HUMIDITY    MeasurementType = 1
	MeasurementType_BATTERY     MeasurementType = 2
)

// Enum value maps for MeasurementType.
var (
	MeasurementType_name = map[int32]string{
		0: "TEMPERATURE",
		1: "HUMIDITY",
		2: "BATTERY",
	}
	MeasurementType_value = map[string]int32{
		"TEMPERATURE": 0,
		"HUMIDITY":    1,
		"BATTERY":     2,
	}
)

func (x MeasurementType) Enum() *MeasurementType {
	p := new(MeasurementType)
	*p = x
	return p
}

func (x MeasurementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MeasurementType) Descriptor() protoreflect.EnumDescriptor {
	return file_indoorClimate_proto_enumTypes[0].Descriptor()
}

func (MeasurementType) Type() protoreflect.EnumType {
	return &file_indoorClimate_proto_enumTypes[0]
}

func (x MeasurementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MeasurementType.Descriptor instead.
func (MeasurementType) EnumDescriptor() ([]byte, []int) {
	return file_indoorClimate_proto_rawDescGZIP(), []int{0}
}

// Indoor Climate data contains a measurement and it's value given by a device at a specific time.
type IndoorClimate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp an indoor climate information is related to.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	// DEvide which sends this measurement.
	DeviceId string `protobuf:"bytes,2,opt,name=DeviceId,proto3" json:"DeviceId,omitempty"`
	// Measurement type.
	Type MeasurementType `protobuf:"varint,3,opt,name=Type,proto3,enum=events.MeasurementType" json:"Type,omitempty"`
	// Measurement value.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IndoorClimate) Reset() {
	*x = IndoorClimate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indoorClimate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndoorClimate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndoorClimate) ProtoMessage() {}

func (x *IndoorClimate) ProtoReflect() protoreflect.Message {
	mi := &file_indoorClimate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndoorClimate.ProtoReflect.Descriptor instead.
func (*IndoorClimate) Descriptor() ([]byte, []int) {
	return file_indoorClimate_proto_rawDescGZIP(), []int{0}
}

func (x *IndoorClimate) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *IndoorClimate) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *IndoorClimate) GetType() MeasurementType {
	if x != nil {
		return x.Type
	}
	return MeasurementType_TEMPERATURE
}

func (x *IndoorClimate) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_indoorClimate_proto protoreflect.FileDescriptor

var file_indoorClimate_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x64, 0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8,
	0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x64, 0x6f, 0x6f, 0x72, 0x43, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x3d, 0x0a, 0x0f, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b,
	0x54, 0x45, 0x4d, 0x50, 0x45, 0x52, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x48, 0x55, 0x4d, 0x49, 0x44, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42,
	0x41, 0x54, 0x54, 0x45, 0x52, 0x59, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indoorClimate_proto_rawDescOnce sync.Once
	file_indoorClimate_proto_rawDescData = file_indoorClimate_proto_rawDesc
)

func file_indoorClimate_proto_rawDescGZIP() []byte {
	file_indoorClimate_proto_rawDescOnce.Do(func() {
		file_indoorClimate_proto_rawDescData = protoimpl.X.CompressGZIP(file_indoorClimate_proto_rawDescData)
	})
	return file_indoorClimate_proto_rawDescData
}

var file_indoorClimate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_indoorClimate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_indoorClimate_proto_goTypes = []interface{}{
	(MeasurementType)(0),        // 0: events.MeasurementType
	(*IndoorClimate)(nil),       // 1: events.IndoorClimate
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_indoorClimate_proto_depIdxs = []int32{
	2, // 0: events.IndoorClimate.Timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: events.IndoorClimate.Type:type_name -> events.MeasurementType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_indoorClimate_proto_init() }
func file_indoorClimate_proto_init() {
	if File_indoorClimate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indoorClimate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndoorClimate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_indoorClimate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indoorClimate_proto_goTypes,
		DependencyIndexes: file_indoorClimate_proto_depIdxs,
		EnumInfos:         file_indoorClimate_proto_enumTypes,
		MessageInfos:      file_indoorClimate_proto_msgTypes,
	}.Build()
	File_indoorClimate_proto = out.File
	file_indoorClimate_proto_rawDesc = nil
	file_indoorClimate_proto_goTypes = nil
	file_indoorClimate_proto_depIdxs = nil
}
