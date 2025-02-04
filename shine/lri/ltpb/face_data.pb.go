// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: face_data.proto

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

type FaceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *CameraID       `protobuf:"varint,1,req,name=id,enum=ltpb.CameraID" json:"id,omitempty"`
	FrameIndex *uint32         `protobuf:"varint,2,req,name=frame_index,json=frameIndex" json:"frame_index,omitempty"`
	Rois       []*FaceData_ROI `protobuf:"bytes,3,rep,name=rois" json:"rois,omitempty"`
}

func (x *FaceData) Reset() {
	*x = FaceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceData) ProtoMessage() {}

func (x *FaceData) ProtoReflect() protoreflect.Message {
	mi := &file_face_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceData.ProtoReflect.Descriptor instead.
func (*FaceData) Descriptor() ([]byte, []int) {
	return file_face_data_proto_rawDescGZIP(), []int{0}
}

func (x *FaceData) GetId() CameraID {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return CameraID_A1
}

func (x *FaceData) GetFrameIndex() uint32 {
	if x != nil && x.FrameIndex != nil {
		return *x.FrameIndex
	}
	return 0
}

func (x *FaceData) GetRois() []*FaceData_ROI {
	if x != nil {
		return x.Rois
	}
	return nil
}

type FaceData_ROI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Roi        *RectangleI `protobuf:"bytes,2,opt,name=roi" json:"roi,omitempty"`
	Confidence *float32    `protobuf:"fixed32,3,opt,name=confidence" json:"confidence,omitempty"`
}

func (x *FaceData_ROI) Reset() {
	*x = FaceData_ROI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceData_ROI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceData_ROI) ProtoMessage() {}

func (x *FaceData_ROI) ProtoReflect() protoreflect.Message {
	mi := &file_face_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceData_ROI.ProtoReflect.Descriptor instead.
func (*FaceData_ROI) Descriptor() ([]byte, []int) {
	return file_face_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FaceData_ROI) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *FaceData_ROI) GetRoi() *RectangleI {
	if x != nil {
		return x.Roi
	}
	return nil
}

func (x *FaceData_ROI) GetConfidence() float32 {
	if x != nil && x.Confidence != nil {
		return *x.Confidence
	}
	return 0
}

var File_face_data_proto protoreflect.FileDescriptor

var file_face_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6c, 0x74, 0x70, 0x62, 0x1a, 0x0f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f,
	0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x72, 0x65, 0x63, 0x74, 0x61, 0x6e,
	0x67, 0x6c, 0x65, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x46,
	0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6c, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x6f, 0x69, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x74, 0x70, 0x62, 0x2e, 0x46, 0x61,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x4f, 0x49, 0x52, 0x04, 0x72, 0x6f, 0x69, 0x73,
	0x1a, 0x59, 0x0a, 0x03, 0x52, 0x4f, 0x49, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x03, 0x72, 0x6f, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x74,
	0x61, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x52, 0x03, 0x72, 0x6f, 0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x6c, 0x74, 0x70, 0x62, 0x3b, 0x6c, 0x74, 0x70, 0x62,
}

var (
	file_face_data_proto_rawDescOnce sync.Once
	file_face_data_proto_rawDescData = file_face_data_proto_rawDesc
)

func file_face_data_proto_rawDescGZIP() []byte {
	file_face_data_proto_rawDescOnce.Do(func() {
		file_face_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_face_data_proto_rawDescData)
	})
	return file_face_data_proto_rawDescData
}

var file_face_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_face_data_proto_goTypes = []any{
	(*FaceData)(nil),     // 0: ltpb.FaceData
	(*FaceData_ROI)(nil), // 1: ltpb.FaceData.ROI
	(CameraID)(0),        // 2: ltpb.CameraID
	(*RectangleI)(nil),   // 3: ltpb.RectangleI
}
var file_face_data_proto_depIdxs = []int32{
	2, // 0: ltpb.FaceData.id:type_name -> ltpb.CameraID
	1, // 1: ltpb.FaceData.rois:type_name -> ltpb.FaceData.ROI
	3, // 2: ltpb.FaceData.ROI.roi:type_name -> ltpb.RectangleI
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_face_data_proto_init() }
func file_face_data_proto_init() {
	if File_face_data_proto != nil {
		return
	}
	file_camera_id_proto_init()
	file_rectanglei_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_face_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FaceData); i {
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
		file_face_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FaceData_ROI); i {
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
			RawDescriptor: file_face_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_face_data_proto_goTypes,
		DependencyIndexes: file_face_data_proto_depIdxs,
		MessageInfos:      file_face_data_proto_msgTypes,
	}.Build()
	File_face_data_proto = out.File
	file_face_data_proto_rawDesc = nil
	file_face_data_proto_goTypes = nil
	file_face_data_proto_depIdxs = nil
}
