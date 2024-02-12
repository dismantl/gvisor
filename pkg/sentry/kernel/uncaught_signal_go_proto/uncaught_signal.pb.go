// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: pkg/sentry/kernel/uncaught_signal.proto

package uncaught_signal_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	registers_go_proto "github.com/dismantl/gvisor/pkg/sentry/arch/registers_go_proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UncaughtSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid          int32                         `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Pid          int32                         `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Registers    *registers_go_proto.Registers `protobuf:"bytes,3,opt,name=registers,proto3" json:"registers,omitempty"`
	SignalNumber int32                         `protobuf:"varint,4,opt,name=signal_number,json=signalNumber,proto3" json:"signal_number,omitempty"`
	FaultAddr    uint64                        `protobuf:"varint,5,opt,name=fault_addr,json=faultAddr,proto3" json:"fault_addr,omitempty"`
}

func (x *UncaughtSignal) Reset() {
	*x = UncaughtSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_kernel_uncaught_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UncaughtSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncaughtSignal) ProtoMessage() {}

func (x *UncaughtSignal) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_kernel_uncaught_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncaughtSignal.ProtoReflect.Descriptor instead.
func (*UncaughtSignal) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_kernel_uncaught_signal_proto_rawDescGZIP(), []int{0}
}

func (x *UncaughtSignal) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *UncaughtSignal) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *UncaughtSignal) GetRegisters() *registers_go_proto.Registers {
	if x != nil {
		return x.Registers
	}
	return nil
}

func (x *UncaughtSignal) GetSignalNumber() int32 {
	if x != nil {
		return x.SignalNumber
	}
	return 0
}

func (x *UncaughtSignal) GetFaultAddr() uint64 {
	if x != nil {
		return x.FaultAddr
	}
	return 0
}

var File_pkg_sentry_kernel_uncaught_signal_proto protoreflect.FileDescriptor

var file_pkg_sentry_kernel_uncaught_signal_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x2f, 0x75, 0x6e, 0x63, 0x61, 0x75, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x72,
	0x63, 0x68, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0e, 0x55, 0x6e, 0x63, 0x61, 0x75, 0x67, 0x68, 0x74, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x09, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x64, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sentry_kernel_uncaught_signal_proto_rawDescOnce sync.Once
	file_pkg_sentry_kernel_uncaught_signal_proto_rawDescData = file_pkg_sentry_kernel_uncaught_signal_proto_rawDesc
)

func file_pkg_sentry_kernel_uncaught_signal_proto_rawDescGZIP() []byte {
	file_pkg_sentry_kernel_uncaught_signal_proto_rawDescOnce.Do(func() {
		file_pkg_sentry_kernel_uncaught_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sentry_kernel_uncaught_signal_proto_rawDescData)
	})
	return file_pkg_sentry_kernel_uncaught_signal_proto_rawDescData
}

var file_pkg_sentry_kernel_uncaught_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_sentry_kernel_uncaught_signal_proto_goTypes = []interface{}{
	(*UncaughtSignal)(nil),               // 0: gvisor.UncaughtSignal
	(*registers_go_proto.Registers)(nil), // 1: gvisor.Registers
}
var file_pkg_sentry_kernel_uncaught_signal_proto_depIdxs = []int32{
	1, // 0: gvisor.UncaughtSignal.registers:type_name -> gvisor.Registers
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_sentry_kernel_uncaught_signal_proto_init() }
func file_pkg_sentry_kernel_uncaught_signal_proto_init() {
	if File_pkg_sentry_kernel_uncaught_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sentry_kernel_uncaught_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UncaughtSignal); i {
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
			RawDescriptor: file_pkg_sentry_kernel_uncaught_signal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sentry_kernel_uncaught_signal_proto_goTypes,
		DependencyIndexes: file_pkg_sentry_kernel_uncaught_signal_proto_depIdxs,
		MessageInfos:      file_pkg_sentry_kernel_uncaught_signal_proto_msgTypes,
	}.Build()
	File_pkg_sentry_kernel_uncaught_signal_proto = out.File
	file_pkg_sentry_kernel_uncaught_signal_proto_rawDesc = nil
	file_pkg_sentry_kernel_uncaught_signal_proto_goTypes = nil
	file_pkg_sentry_kernel_uncaught_signal_proto_depIdxs = nil
}
