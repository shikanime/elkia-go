// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.12
// source: pkg/api/world/v1alpha1/combat.proto

package v1alpha1

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

type MultiTargetCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId  string `protobuf:"bytes,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	TargetId uint32 `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
}

func (x *MultiTargetCommand) Reset() {
	*x = MultiTargetCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiTargetCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiTargetCommand) ProtoMessage() {}

func (x *MultiTargetCommand) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiTargetCommand.ProtoReflect.Descriptor instead.
func (*MultiTargetCommand) Descriptor() ([]byte, []int) {
	return file_pkg_api_world_v1alpha1_combat_proto_rawDescGZIP(), []int{0}
}

func (x *MultiTargetCommand) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

func (x *MultiTargetCommand) GetTargetId() uint32 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

type MultiTargetListCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAmount uint32                `protobuf:"varint,1,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	Targets      []*MultiTargetCommand `protobuf:"bytes,2,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *MultiTargetListCommand) Reset() {
	*x = MultiTargetListCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiTargetListCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiTargetListCommand) ProtoMessage() {}

func (x *MultiTargetListCommand) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiTargetListCommand.ProtoReflect.Descriptor instead.
func (*MultiTargetListCommand) Descriptor() ([]byte, []int) {
	return file_pkg_api_world_v1alpha1_combat_proto_rawDescGZIP(), []int{1}
}

func (x *MultiTargetListCommand) GetTargetAmount() uint32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *MultiTargetListCommand) GetTargets() []*MultiTargetCommand {
	if x != nil {
		return x.Targets
	}
	return nil
}

type CombatCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//
	//	*CombatCommand_MultiTargetListCommand
	Command isCombatCommand_Command `protobuf_oneof:"command"`
}

func (x *CombatCommand) Reset() {
	*x = CombatCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombatCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombatCommand) ProtoMessage() {}

func (x *CombatCommand) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombatCommand.ProtoReflect.Descriptor instead.
func (*CombatCommand) Descriptor() ([]byte, []int) {
	return file_pkg_api_world_v1alpha1_combat_proto_rawDescGZIP(), []int{2}
}

func (m *CombatCommand) GetCommand() isCombatCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *CombatCommand) GetMultiTargetListCommand() *MultiTargetListCommand {
	if x, ok := x.GetCommand().(*CombatCommand_MultiTargetListCommand); ok {
		return x.MultiTargetListCommand
	}
	return nil
}

type isCombatCommand_Command interface {
	isCombatCommand_Command()
}

type CombatCommand_MultiTargetListCommand struct {
	MultiTargetListCommand *MultiTargetListCommand `protobuf:"bytes,2,opt,name=multi_target_list_command,json=multiTargetListCommand,proto3,oneof"`
}

func (*CombatCommand_MultiTargetListCommand) isCombatCommand_Command() {}

type CombatInteractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence uint32         `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Command  *CombatCommand `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *CombatInteractRequest) Reset() {
	*x = CombatInteractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombatInteractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombatInteractRequest) ProtoMessage() {}

func (x *CombatInteractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombatInteractRequest.ProtoReflect.Descriptor instead.
func (*CombatInteractRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_world_v1alpha1_combat_proto_rawDescGZIP(), []int{3}
}

func (x *CombatInteractRequest) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *CombatInteractRequest) GetCommand() *CombatCommand {
	if x != nil {
		return x.Command
	}
	return nil
}

type CombatEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CombatEvent) Reset() {
	*x = CombatEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombatEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombatEvent) ProtoMessage() {}

func (x *CombatEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_world_v1alpha1_combat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombatEvent.ProtoReflect.Descriptor instead.
func (*CombatEvent) Descriptor() ([]byte, []int) {
	return file_pkg_api_world_v1alpha1_combat_proto_rawDescGZIP(), []int{4}
}

var File_pkg_api_world_v1alpha1_combat_proto protoreflect.FileDescriptor

var file_pkg_api_world_v1alpha1_combat_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x4c, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x16, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x73, 0x0a, 0x19, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00,
	0x52, 0x16, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x22, 0x7c, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x68, 0x69, 0x6b,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x32, 0x82, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x12, 0x78, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x12, 0x35, 0x2e,
	0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x6e, 0x69, 0x6d, 0x65,
	0x2e, 0x65, 0x6c, 0x6b, 0x69, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x73, 0x68, 0x69, 0x6b,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x65, 0x6c, 0x6b,
	0x69, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_api_world_v1alpha1_combat_proto_rawDescOnce sync.Once
	file_pkg_api_world_v1alpha1_combat_proto_rawDescData = file_pkg_api_world_v1alpha1_combat_proto_rawDesc
)

func file_pkg_api_world_v1alpha1_combat_proto_rawDescGZIP() []byte {
	file_pkg_api_world_v1alpha1_combat_proto_rawDescOnce.Do(func() {
		file_pkg_api_world_v1alpha1_combat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_world_v1alpha1_combat_proto_rawDescData)
	})
	return file_pkg_api_world_v1alpha1_combat_proto_rawDescData
}

var file_pkg_api_world_v1alpha1_combat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_api_world_v1alpha1_combat_proto_goTypes = []interface{}{
	(*MultiTargetCommand)(nil),     // 0: shikanime.elkia.world.v1alpha1.MultiTargetCommand
	(*MultiTargetListCommand)(nil), // 1: shikanime.elkia.world.v1alpha1.MultiTargetListCommand
	(*CombatCommand)(nil),          // 2: shikanime.elkia.world.v1alpha1.CombatCommand
	(*CombatInteractRequest)(nil),  // 3: shikanime.elkia.world.v1alpha1.CombatInteractRequest
	(*CombatEvent)(nil),            // 4: shikanime.elkia.world.v1alpha1.CombatEvent
}
var file_pkg_api_world_v1alpha1_combat_proto_depIdxs = []int32{
	0, // 0: shikanime.elkia.world.v1alpha1.MultiTargetListCommand.targets:type_name -> shikanime.elkia.world.v1alpha1.MultiTargetCommand
	1, // 1: shikanime.elkia.world.v1alpha1.CombatCommand.multi_target_list_command:type_name -> shikanime.elkia.world.v1alpha1.MultiTargetListCommand
	2, // 2: shikanime.elkia.world.v1alpha1.CombatInteractRequest.command:type_name -> shikanime.elkia.world.v1alpha1.CombatCommand
	3, // 3: shikanime.elkia.world.v1alpha1.Combat.CombatInteract:input_type -> shikanime.elkia.world.v1alpha1.CombatInteractRequest
	4, // 4: shikanime.elkia.world.v1alpha1.Combat.CombatInteract:output_type -> shikanime.elkia.world.v1alpha1.CombatEvent
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_api_world_v1alpha1_combat_proto_init() }
func file_pkg_api_world_v1alpha1_combat_proto_init() {
	if File_pkg_api_world_v1alpha1_combat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_world_v1alpha1_combat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiTargetCommand); i {
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
		file_pkg_api_world_v1alpha1_combat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiTargetListCommand); i {
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
		file_pkg_api_world_v1alpha1_combat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombatCommand); i {
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
		file_pkg_api_world_v1alpha1_combat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombatInteractRequest); i {
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
		file_pkg_api_world_v1alpha1_combat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombatEvent); i {
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
	file_pkg_api_world_v1alpha1_combat_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CombatCommand_MultiTargetListCommand)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_api_world_v1alpha1_combat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_world_v1alpha1_combat_proto_goTypes,
		DependencyIndexes: file_pkg_api_world_v1alpha1_combat_proto_depIdxs,
		MessageInfos:      file_pkg_api_world_v1alpha1_combat_proto_msgTypes,
	}.Build()
	File_pkg_api_world_v1alpha1_combat_proto = out.File
	file_pkg_api_world_v1alpha1_combat_proto_rawDesc = nil
	file_pkg_api_world_v1alpha1_combat_proto_goTypes = nil
	file_pkg_api_world_v1alpha1_combat_proto_depIdxs = nil
}
