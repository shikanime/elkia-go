// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.12
// source: pkg/api/eventing/v1alpha1/auth.proto

package v1alpha1

import (
	v1alpha1 "go.shikanime.studio/elkia/pkg/api/fleet/v1alpha1"
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

type AuthCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//
	//	*AuthCommand_Presence
	Command isAuthCommand_Command `protobuf_oneof:"command"`
}

func (x *AuthCommand) Reset() {
	*x = AuthCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCommand) ProtoMessage() {}

func (x *AuthCommand) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCommand.ProtoReflect.Descriptor instead.
func (*AuthCommand) Descriptor() ([]byte, []int) {
	return file_pkg_api_eventing_v1alpha1_auth_proto_rawDescGZIP(), []int{0}
}

func (m *AuthCommand) GetCommand() isAuthCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *AuthCommand) GetPresence() *v1alpha1.PresenceCommand {
	if x, ok := x.GetCommand().(*AuthCommand_Presence); ok {
		return x.Presence
	}
	return nil
}

type isAuthCommand_Command interface {
	isAuthCommand_Command()
}

type AuthCommand_Presence struct {
	Presence *v1alpha1.PresenceCommand `protobuf:"bytes,1,opt,name=presence,proto3,oneof"`
}

func (*AuthCommand_Presence) isAuthCommand_Command() {}

type AuthEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*AuthEvent_Presence
	//	*AuthEvent_Client
	Event isAuthEvent_Event `protobuf_oneof:"event"`
}

func (x *AuthEvent) Reset() {
	*x = AuthEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthEvent) ProtoMessage() {}

func (x *AuthEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthEvent.ProtoReflect.Descriptor instead.
func (*AuthEvent) Descriptor() ([]byte, []int) {
	return file_pkg_api_eventing_v1alpha1_auth_proto_rawDescGZIP(), []int{1}
}

func (m *AuthEvent) GetEvent() isAuthEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *AuthEvent) GetPresence() *v1alpha1.PresenceEvent {
	if x, ok := x.GetEvent().(*AuthEvent_Presence); ok {
		return x.Presence
	}
	return nil
}

func (x *AuthEvent) GetClient() *ClientEvent {
	if x, ok := x.GetEvent().(*AuthEvent_Client); ok {
		return x.Client
	}
	return nil
}

type isAuthEvent_Event interface {
	isAuthEvent_Event()
}

type AuthEvent_Presence struct {
	Presence *v1alpha1.PresenceEvent `protobuf:"bytes,1,opt,name=presence,proto3,oneof"`
}

type AuthEvent_Client struct {
	Client *ClientEvent `protobuf:"bytes,2,opt,name=client,proto3,oneof"`
}

func (*AuthEvent_Presence) isAuthEvent_Event() {}

func (*AuthEvent_Client) isAuthEvent_Event() {}

var File_pkg_api_eventing_v1alpha1_auth_proto protoreflect.FileDescriptor

var file_pkg_api_eventing_v1alpha1_auth_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x4d, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x68, 0x69, 0x6b,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x22, 0xab, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x4b, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6c,
	0x6b, 0x69, 0x61, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73,
	0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32,
	0x78, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x70, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x12, 0x2e, 0x2e, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x2c, 0x2e, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e,
	0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x2e,
	0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2f, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_eventing_v1alpha1_auth_proto_rawDescOnce sync.Once
	file_pkg_api_eventing_v1alpha1_auth_proto_rawDescData = file_pkg_api_eventing_v1alpha1_auth_proto_rawDesc
)

func file_pkg_api_eventing_v1alpha1_auth_proto_rawDescGZIP() []byte {
	file_pkg_api_eventing_v1alpha1_auth_proto_rawDescOnce.Do(func() {
		file_pkg_api_eventing_v1alpha1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_eventing_v1alpha1_auth_proto_rawDescData)
	})
	return file_pkg_api_eventing_v1alpha1_auth_proto_rawDescData
}

var file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_api_eventing_v1alpha1_auth_proto_goTypes = []interface{}{
	(*AuthCommand)(nil),              // 0: shikanime.elkia.eventing.v1alpha1.AuthCommand
	(*AuthEvent)(nil),                // 1: shikanime.elkia.eventing.v1alpha1.AuthEvent
	(*v1alpha1.PresenceCommand)(nil), // 2: shikanime.elkia.fleet.v1alpha1.PresenceCommand
	(*v1alpha1.PresenceEvent)(nil),   // 3: shikanime.elkia.fleet.v1alpha1.PresenceEvent
	(*ClientEvent)(nil),              // 4: shikanime.elkia.eventing.v1alpha1.ClientEvent
}
var file_pkg_api_eventing_v1alpha1_auth_proto_depIdxs = []int32{
	2, // 0: shikanime.elkia.eventing.v1alpha1.AuthCommand.presence:type_name -> shikanime.elkia.fleet.v1alpha1.PresenceCommand
	3, // 1: shikanime.elkia.eventing.v1alpha1.AuthEvent.presence:type_name -> shikanime.elkia.fleet.v1alpha1.PresenceEvent
	4, // 2: shikanime.elkia.eventing.v1alpha1.AuthEvent.client:type_name -> shikanime.elkia.eventing.v1alpha1.ClientEvent
	0, // 3: shikanime.elkia.eventing.v1alpha1.Auth.AuthInteract:input_type -> shikanime.elkia.eventing.v1alpha1.AuthCommand
	1, // 4: shikanime.elkia.eventing.v1alpha1.Auth.AuthInteract:output_type -> shikanime.elkia.eventing.v1alpha1.AuthEvent
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_api_eventing_v1alpha1_auth_proto_init() }
func file_pkg_api_eventing_v1alpha1_auth_proto_init() {
	if File_pkg_api_eventing_v1alpha1_auth_proto != nil {
		return
	}
	file_pkg_api_eventing_v1alpha1_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCommand); i {
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
		file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthEvent); i {
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
	file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AuthCommand_Presence)(nil),
	}
	file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AuthEvent_Presence)(nil),
		(*AuthEvent_Client)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_api_eventing_v1alpha1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_eventing_v1alpha1_auth_proto_goTypes,
		DependencyIndexes: file_pkg_api_eventing_v1alpha1_auth_proto_depIdxs,
		MessageInfos:      file_pkg_api_eventing_v1alpha1_auth_proto_msgTypes,
	}.Build()
	File_pkg_api_eventing_v1alpha1_auth_proto = out.File
	file_pkg_api_eventing_v1alpha1_auth_proto_rawDesc = nil
	file_pkg_api_eventing_v1alpha1_auth_proto_goTypes = nil
	file_pkg_api_eventing_v1alpha1_auth_proto_depIdxs = nil
}
