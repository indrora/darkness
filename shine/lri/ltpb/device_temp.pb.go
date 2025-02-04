// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: device_temp.proto

package ltpb

import (
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

type DeviceTemp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor_1     *int32 `protobuf:"zigzag32,1,opt,name=sensor_1,json=sensor1" json:"sensor_1,omitempty"`
	Sensor_2     *int32 `protobuf:"zigzag32,2,opt,name=sensor_2,json=sensor2" json:"sensor_2,omitempty"`
	Sensor_3     *int32 `protobuf:"zigzag32,3,opt,name=sensor_3,json=sensor3" json:"sensor_3,omitempty"`
	Sensor_4     *int32 `protobuf:"zigzag32,4,opt,name=sensor_4,json=sensor4" json:"sensor_4,omitempty"`
	FlexSensor_1 *int32 `protobuf:"zigzag32,5,opt,name=flex_sensor_1,json=flexSensor1" json:"flex_sensor_1,omitempty"`
}

func (x *DeviceTemp) Reset() {
	*x = DeviceTemp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_temp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceTemp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceTemp) ProtoMessage() {}

func (x *DeviceTemp) ProtoReflect() protoreflect.Message {
	mi := &file_device_temp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceTemp.ProtoReflect.Descriptor instead.
func (*DeviceTemp) Descriptor() ([]byte, []int) {
	return file_device_temp_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceTemp) GetSensor_1() int32 {
	if x != nil && x.Sensor_1 != nil {
		return *x.Sensor_1
	}
	return 0
}

func (x *DeviceTemp) GetSensor_2() int32 {
	if x != nil && x.Sensor_2 != nil {
		return *x.Sensor_2
	}
	return 0
}

func (x *DeviceTemp) GetSensor_3() int32 {
	if x != nil && x.Sensor_3 != nil {
		return *x.Sensor_3
	}
	return 0
}

func (x *DeviceTemp) GetSensor_4() int32 {
	if x != nil && x.Sensor_4 != nil {
		return *x.Sensor_4
	}
	return 0
}

func (x *DeviceTemp) GetFlexSensor_1() int32 {
	if x != nil && x.FlexSensor_1 != nil {
		return *x.FlexSensor_1
	}
	return 0
}

var File_device_temp_proto protoreflect.FileDescriptor

var file_device_temp_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x74, 0x70, 0x62, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x31, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x32, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x5f, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x34, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x5f, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0b, 0x66, 0x6c, 0x65,
	0x78, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x31, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6c, 0x74,
	0x70, 0x62, 0x3b, 0x6c, 0x74, 0x70, 0x62,
}

var (
	file_device_temp_proto_rawDescOnce sync.Once
	file_device_temp_proto_rawDescData = file_device_temp_proto_rawDesc
)

func file_device_temp_proto_rawDescGZIP() []byte {
	file_device_temp_proto_rawDescOnce.Do(func() {
		file_device_temp_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_temp_proto_rawDescData)
	})
	return file_device_temp_proto_rawDescData
}

var file_device_temp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_device_temp_proto_goTypes = []any{
	(*DeviceTemp)(nil), // 0: ltpb.DeviceTemp
}
var file_device_temp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_device_temp_proto_init() }
func file_device_temp_proto_init() {
	if File_device_temp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_temp_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceTemp); i {
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
			RawDescriptor: file_device_temp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_device_temp_proto_goTypes,
		DependencyIndexes: file_device_temp_proto_depIdxs,
		MessageInfos:      file_device_temp_proto_msgTypes,
	}.Build()
	File_device_temp_proto = out.File
	file_device_temp_proto_rawDesc = nil
	file_device_temp_proto_goTypes = nil
	file_device_temp_proto_depIdxs = nil
}
