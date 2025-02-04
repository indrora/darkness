// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: camera_id.proto

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

type CameraID int32

const (
	CameraID_A1 CameraID = 0
	CameraID_A2 CameraID = 1
	CameraID_A3 CameraID = 2
	CameraID_A4 CameraID = 3
	CameraID_A5 CameraID = 4
	CameraID_B1 CameraID = 5
	CameraID_B2 CameraID = 6
	CameraID_B3 CameraID = 7
	CameraID_B4 CameraID = 8
	CameraID_B5 CameraID = 9
	CameraID_C1 CameraID = 10
	CameraID_C2 CameraID = 11
	CameraID_C3 CameraID = 12
	CameraID_C4 CameraID = 13
	CameraID_C5 CameraID = 14
	CameraID_C6 CameraID = 15
)

// Enum value maps for CameraID.
var (
	CameraID_name = map[int32]string{
		0:  "A1",
		1:  "A2",
		2:  "A3",
		3:  "A4",
		4:  "A5",
		5:  "B1",
		6:  "B2",
		7:  "B3",
		8:  "B4",
		9:  "B5",
		10: "C1",
		11: "C2",
		12: "C3",
		13: "C4",
		14: "C5",
		15: "C6",
	}
	CameraID_value = map[string]int32{
		"A1": 0,
		"A2": 1,
		"A3": 2,
		"A4": 3,
		"A5": 4,
		"B1": 5,
		"B2": 6,
		"B3": 7,
		"B4": 8,
		"B5": 9,
		"C1": 10,
		"C2": 11,
		"C3": 12,
		"C4": 13,
		"C5": 14,
		"C6": 15,
	}
)

func (x CameraID) Enum() *CameraID {
	p := new(CameraID)
	*p = x
	return p
}

func (x CameraID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CameraID) Descriptor() protoreflect.EnumDescriptor {
	return file_camera_id_proto_enumTypes[0].Descriptor()
}

func (CameraID) Type() protoreflect.EnumType {
	return &file_camera_id_proto_enumTypes[0]
}

func (x CameraID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CameraID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CameraID(num)
	return nil
}

// Deprecated: Use CameraID.Descriptor instead.
func (CameraID) EnumDescriptor() ([]byte, []int) {
	return file_camera_id_proto_rawDescGZIP(), []int{0}
}

var File_camera_id_proto protoreflect.FileDescriptor

var file_camera_id_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6c, 0x74, 0x70, 0x62, 0x2a, 0x8a, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x49, 0x44, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x31, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x41, 0x32, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x33, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02,
	0x41, 0x34, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x35, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02,
	0x42, 0x31, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x32, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02,
	0x42, 0x33, 0x10, 0x07, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x34, 0x10, 0x08, 0x12, 0x06, 0x0a, 0x02,
	0x42, 0x35, 0x10, 0x09, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x31, 0x10, 0x0a, 0x12, 0x06, 0x0a, 0x02,
	0x43, 0x32, 0x10, 0x0b, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x33, 0x10, 0x0c, 0x12, 0x06, 0x0a, 0x02,
	0x43, 0x34, 0x10, 0x0d, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x35, 0x10, 0x0e, 0x12, 0x06, 0x0a, 0x02,
	0x43, 0x36, 0x10, 0x0f, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6c, 0x74, 0x70, 0x62, 0x3b, 0x6c,
	0x74, 0x70, 0x62,
}

var (
	file_camera_id_proto_rawDescOnce sync.Once
	file_camera_id_proto_rawDescData = file_camera_id_proto_rawDesc
)

func file_camera_id_proto_rawDescGZIP() []byte {
	file_camera_id_proto_rawDescOnce.Do(func() {
		file_camera_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_camera_id_proto_rawDescData)
	})
	return file_camera_id_proto_rawDescData
}

var file_camera_id_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_camera_id_proto_goTypes = []any{
	(CameraID)(0), // 0: ltpb.CameraID
}
var file_camera_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_camera_id_proto_init() }
func file_camera_id_proto_init() {
	if File_camera_id_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_camera_id_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_camera_id_proto_goTypes,
		DependencyIndexes: file_camera_id_proto_depIdxs,
		EnumInfos:         file_camera_id_proto_enumTypes,
	}.Build()
	File_camera_id_proto = out.File
	file_camera_id_proto_rawDesc = nil
	file_camera_id_proto_goTypes = nil
	file_camera_id_proto_depIdxs = nil
}
