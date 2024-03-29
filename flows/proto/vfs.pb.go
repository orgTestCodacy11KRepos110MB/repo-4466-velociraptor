// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vfs.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is stored in the VFS datastores to indicate where the actual
// download exist (usually in the flow filestore).
type VFSDownloadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated but leave it here so we can read older files
	// downloaded within the filestore.
	VfsPath string `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Size    uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Mtime   uint64 `protobuf:"varint,3,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Sparse  bool   `protobuf:"varint,4,opt,name=sparse,proto3" json:"sparse,omitempty"`
	MD5     string `protobuf:"bytes,5,opt,name=MD5,proto3" json:"MD5,omitempty"`
	SHA256  string `protobuf:"bytes,6,opt,name=SHA256,proto3" json:"SHA256,omitempty"`
	// The VFS path is now expressed as a list of components
	Components []string `protobuf:"bytes,7,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *VFSDownloadInfo) Reset() {
	*x = VFSDownloadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VFSDownloadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VFSDownloadInfo) ProtoMessage() {}

func (x *VFSDownloadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vfs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VFSDownloadInfo.ProtoReflect.Descriptor instead.
func (*VFSDownloadInfo) Descriptor() ([]byte, []int) {
	return file_vfs_proto_rawDescGZIP(), []int{0}
}

func (x *VFSDownloadInfo) GetVfsPath() string {
	if x != nil {
		return x.VfsPath
	}
	return ""
}

func (x *VFSDownloadInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *VFSDownloadInfo) GetMtime() uint64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

func (x *VFSDownloadInfo) GetSparse() bool {
	if x != nil {
		return x.Sparse
	}
	return false
}

func (x *VFSDownloadInfo) GetMD5() string {
	if x != nil {
		return x.MD5
	}
	return ""
}

func (x *VFSDownloadInfo) GetSHA256() string {
	if x != nil {
		return x.SHA256
	}
	return ""
}

func (x *VFSDownloadInfo) GetComponents() []string {
	if x != nil {
		return x.Components
	}
	return nil
}

type ClientMonitoringState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version             uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	LastUpdateTimestamp uint64 `protobuf:"varint,2,opt,name=last_update_timestamp,json=lastUpdateTimestamp,proto3" json:"last_update_timestamp,omitempty"`
}

func (x *ClientMonitoringState) Reset() {
	*x = ClientMonitoringState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMonitoringState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMonitoringState) ProtoMessage() {}

func (x *ClientMonitoringState) ProtoReflect() protoreflect.Message {
	mi := &file_vfs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMonitoringState.ProtoReflect.Descriptor instead.
func (*ClientMonitoringState) Descriptor() ([]byte, []int) {
	return file_vfs_proto_rawDescGZIP(), []int{1}
}

func (x *ClientMonitoringState) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ClientMonitoringState) GetLastUpdateTimestamp() uint64 {
	if x != nil {
		return x.LastUpdateTimestamp
	}
	return 0
}

var File_vfs_proto protoreflect.FileDescriptor

var file_vfs_proto_rawDesc = []byte{
	0x0a, 0x09, 0x76, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0f, 0x56, 0x46, 0x53,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x76, 0x66, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x66, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x61, 0x72, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x44, 0x35,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x44, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x48, 0x41, 0x32, 0x35, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x48, 0x41,
	0x32, 0x35, 0x36, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x26,
	0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x20, 0x12, 0x1e, 0x54, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x27, 0x73, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x60, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2c,
	0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x26, 0x12, 0x24, 0x54, 0x68, 0x65, 0x20, 0x6c, 0x61, 0x73, 0x74,
	0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x77, 0x65, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x13, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x33, 0x5a, 0x31, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vfs_proto_rawDescOnce sync.Once
	file_vfs_proto_rawDescData = file_vfs_proto_rawDesc
)

func file_vfs_proto_rawDescGZIP() []byte {
	file_vfs_proto_rawDescOnce.Do(func() {
		file_vfs_proto_rawDescData = protoimpl.X.CompressGZIP(file_vfs_proto_rawDescData)
	})
	return file_vfs_proto_rawDescData
}

var file_vfs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vfs_proto_goTypes = []interface{}{
	(*VFSDownloadInfo)(nil),       // 0: proto.VFSDownloadInfo
	(*ClientMonitoringState)(nil), // 1: proto.ClientMonitoringState
}
var file_vfs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vfs_proto_init() }
func file_vfs_proto_init() {
	if File_vfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VFSDownloadInfo); i {
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
		file_vfs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMonitoringState); i {
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
			RawDescriptor: file_vfs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vfs_proto_goTypes,
		DependencyIndexes: file_vfs_proto_depIdxs,
		MessageInfos:      file_vfs_proto_msgTypes,
	}.Build()
	File_vfs_proto = out.File
	file_vfs_proto_rawDesc = nil
	file_vfs_proto_goTypes = nil
	file_vfs_proto_depIdxs = nil
}
