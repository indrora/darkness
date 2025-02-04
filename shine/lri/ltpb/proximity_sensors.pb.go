// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: proximity_sensors.proto

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

type ProximitySensors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor_1 *bool `protobuf:"varint,1,req,name=sensor_1,json=sensor1" json:"sensor_1,omitempty"`
	Sensor_2 *bool `protobuf:"varint,2,req,name=sensor_2,json=sensor2" json:"sensor_2,omitempty"`
	Sensor_3 *bool `protobuf:"varint,3,req,name=sensor_3,json=sensor3" json:"sensor_3,omitempty"`
	Sensor_4 *bool `protobuf:"varint,4,req,name=sensor_4,json=sensor4" json:"sensor_4,omitempty"`
	Sensor_5 *bool `protobuf:"varint,5,req,name=sensor_5,json=sensor5" json:"sensor_5,omitempty"`
}

func (x *ProximitySensors) Reset() {
	*x = ProximitySensors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proximity_sensors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProximitySensors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProximitySensors) ProtoMessage() {}

func (x *ProximitySensors) ProtoReflect() protoreflect.Message {
	mi := &file_proximity_sensors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProximitySensors.ProtoReflect.Descriptor instead.
func (*ProximitySensors) Descriptor() ([]byte, []int) {
	return file_proximity_sensors_proto_rawDescGZIP(), []int{0}
}

func (x *ProximitySensors) GetSensor_1() bool {
	if x != nil && x.Sensor_1 != nil {
		return *x.Sensor_1
	}
	return false
}

func (x *ProximitySensors) GetSensor_2() bool {
	if x != nil && x.Sensor_2 != nil {
		return *x.Sensor_2
	}
	return false
}

func (x *ProximitySensors) GetSensor_3() bool {
	if x != nil && x.Sensor_3 != nil {
		return *x.Sensor_3
	}
	return false
}

func (x *ProximitySensors) GetSensor_4() bool {
	if x != nil && x.Sensor_4 != nil {
		return *x.Sensor_4
	}
	return false
}

func (x *ProximitySensors) GetSensor_5() bool {
	if x != nil && x.Sensor_5 != nil {
		return *x.Sensor_5
	}
	return false
}

var File_proximity_sensors_proto protoreflect.FileDescriptor

var file_proximity_sensors_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x74, 0x70, 0x62, 0x22,
	0x99, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x31,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x31, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x33, 0x18, 0x03, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f,
	0x34, 0x18, 0x04, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x34,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x35, 0x18, 0x05, 0x20, 0x02,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x35, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x6c, 0x74, 0x70, 0x62, 0x3b, 0x6c, 0x74, 0x70, 0x62,
}

var (
	file_proximity_sensors_proto_rawDescOnce sync.Once
	file_proximity_sensors_proto_rawDescData = file_proximity_sensors_proto_rawDesc
)

func file_proximity_sensors_proto_rawDescGZIP() []byte {
	file_proximity_sensors_proto_rawDescOnce.Do(func() {
		file_proximity_sensors_proto_rawDescData = protoimpl.X.CompressGZIP(file_proximity_sensors_proto_rawDescData)
	})
	return file_proximity_sensors_proto_rawDescData
}

var file_proximity_sensors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proximity_sensors_proto_goTypes = []any{
	(*ProximitySensors)(nil), // 0: ltpb.ProximitySensors
}
var file_proximity_sensors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proximity_sensors_proto_init() }
func file_proximity_sensors_proto_init() {
	if File_proximity_sensors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proximity_sensors_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProximitySensors); i {
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
			RawDescriptor: file_proximity_sensors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proximity_sensors_proto_goTypes,
		DependencyIndexes: file_proximity_sensors_proto_depIdxs,
		MessageInfos:      file_proximity_sensors_proto_msgTypes,
	}.Build()
	File_proximity_sensors_proto = out.File
	file_proximity_sensors_proto_rawDesc = nil
	file_proximity_sensors_proto_goTypes = nil
	file_proximity_sensors_proto_depIdxs = nil
}
