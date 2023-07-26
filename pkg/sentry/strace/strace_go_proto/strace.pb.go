// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.7
// source: pkg/sentry/strace/strace.proto

package strace_go_proto

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

type Strace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Process  string   `protobuf:"bytes,1,opt,name=process,proto3" json:"process,omitempty"`
	Function string   `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	Args     []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// Types that are assignable to Info:
	//
	//	*Strace_Enter
	//	*Strace_Exit
	Info isStrace_Info `protobuf_oneof:"info"`
}

func (x *Strace) Reset() {
	*x = Strace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_strace_strace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strace) ProtoMessage() {}

func (x *Strace) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_strace_strace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strace.ProtoReflect.Descriptor instead.
func (*Strace) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_strace_strace_proto_rawDescGZIP(), []int{0}
}

func (x *Strace) GetProcess() string {
	if x != nil {
		return x.Process
	}
	return ""
}

func (x *Strace) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *Strace) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (m *Strace) GetInfo() isStrace_Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (x *Strace) GetEnter() *StraceEnter {
	if x, ok := x.GetInfo().(*Strace_Enter); ok {
		return x.Enter
	}
	return nil
}

func (x *Strace) GetExit() *StraceExit {
	if x, ok := x.GetInfo().(*Strace_Exit); ok {
		return x.Exit
	}
	return nil
}

type isStrace_Info interface {
	isStrace_Info()
}

type Strace_Enter struct {
	Enter *StraceEnter `protobuf:"bytes,4,opt,name=enter,proto3,oneof"`
}

type Strace_Exit struct {
	Exit *StraceExit `protobuf:"bytes,5,opt,name=exit,proto3,oneof"`
}

func (*Strace_Enter) isStrace_Info() {}

func (*Strace_Exit) isStrace_Info() {}

type StraceEnter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StraceEnter) Reset() {
	*x = StraceEnter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_strace_strace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StraceEnter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StraceEnter) ProtoMessage() {}

func (x *StraceEnter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_strace_strace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StraceEnter.ProtoReflect.Descriptor instead.
func (*StraceEnter) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_strace_strace_proto_rawDescGZIP(), []int{1}
}

type StraceExit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Return    string `protobuf:"bytes,1,opt,name=return,proto3" json:"return,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ErrNo     int64  `protobuf:"varint,3,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ElapsedNs int64  `protobuf:"varint,4,opt,name=elapsed_ns,json=elapsedNs,proto3" json:"elapsed_ns,omitempty"`
}

func (x *StraceExit) Reset() {
	*x = StraceExit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_strace_strace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StraceExit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StraceExit) ProtoMessage() {}

func (x *StraceExit) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_strace_strace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StraceExit.ProtoReflect.Descriptor instead.
func (*StraceExit) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_strace_strace_proto_rawDescGZIP(), []int{2}
}

func (x *StraceExit) GetReturn() string {
	if x != nil {
		return x.Return
	}
	return ""
}

func (x *StraceExit) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *StraceExit) GetErrNo() int64 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *StraceExit) GetElapsedNs() int64 {
	if x != nil {
		return x.ElapsedNs
	}
	return 0
}

var File_pkg_sentry_strace_strace_proto protoreflect.FileDescriptor

var file_pkg_sentry_strace_strace_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x2b, 0x0a,
	0x05, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x65, 0x78,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x63, 0x65, 0x45, 0x78, 0x69, 0x74, 0x48, 0x00, 0x52, 0x04,
	0x65, 0x78, 0x69, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x0d, 0x0a, 0x0b,
	0x53, 0x74, 0x72, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x45, 0x78, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x4e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sentry_strace_strace_proto_rawDescOnce sync.Once
	file_pkg_sentry_strace_strace_proto_rawDescData = file_pkg_sentry_strace_strace_proto_rawDesc
)

func file_pkg_sentry_strace_strace_proto_rawDescGZIP() []byte {
	file_pkg_sentry_strace_strace_proto_rawDescOnce.Do(func() {
		file_pkg_sentry_strace_strace_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sentry_strace_strace_proto_rawDescData)
	})
	return file_pkg_sentry_strace_strace_proto_rawDescData
}

var file_pkg_sentry_strace_strace_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_sentry_strace_strace_proto_goTypes = []interface{}{
	(*Strace)(nil),      // 0: gvisor.Strace
	(*StraceEnter)(nil), // 1: gvisor.StraceEnter
	(*StraceExit)(nil),  // 2: gvisor.StraceExit
}
var file_pkg_sentry_strace_strace_proto_depIdxs = []int32{
	1, // 0: gvisor.Strace.enter:type_name -> gvisor.StraceEnter
	2, // 1: gvisor.Strace.exit:type_name -> gvisor.StraceExit
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_sentry_strace_strace_proto_init() }
func file_pkg_sentry_strace_strace_proto_init() {
	if File_pkg_sentry_strace_strace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sentry_strace_strace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strace); i {
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
		file_pkg_sentry_strace_strace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StraceEnter); i {
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
		file_pkg_sentry_strace_strace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StraceExit); i {
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
	file_pkg_sentry_strace_strace_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Strace_Enter)(nil),
		(*Strace_Exit)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_sentry_strace_strace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sentry_strace_strace_proto_goTypes,
		DependencyIndexes: file_pkg_sentry_strace_strace_proto_depIdxs,
		MessageInfos:      file_pkg_sentry_strace_strace_proto_msgTypes,
	}.Build()
	File_pkg_sentry_strace_strace_proto = out.File
	file_pkg_sentry_strace_strace_proto_rawDesc = nil
	file_pkg_sentry_strace_strace_proto_goTypes = nil
	file_pkg_sentry_strace_strace_proto_depIdxs = nil
}