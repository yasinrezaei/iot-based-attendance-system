// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: service.proto

package attendance_system

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

type LedColorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *LedColorRequest) Reset() {
	*x = LedColorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedColorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedColorRequest) ProtoMessage() {}

func (x *LedColorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedColorRequest.ProtoReflect.Descriptor instead.
func (*LedColorRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *LedColorRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type LedColorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Red   int32 `protobuf:"varint,1,opt,name=red,proto3" json:"red,omitempty"`
	Green int32 `protobuf:"varint,2,opt,name=green,proto3" json:"green,omitempty"`
	Blue  int32 `protobuf:"varint,3,opt,name=blue,proto3" json:"blue,omitempty"`
}

func (x *LedColorResponse) Reset() {
	*x = LedColorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedColorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedColorResponse) ProtoMessage() {}

func (x *LedColorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedColorResponse.ProtoReflect.Descriptor instead.
func (*LedColorResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *LedColorResponse) GetRed() int32 {
	if x != nil {
		return x.Red
	}
	return 0
}

func (x *LedColorResponse) GetGreen() int32 {
	if x != nil {
		return x.Green
	}
	return 0
}

func (x *LedColorResponse) GetBlue() int32 {
	if x != nil {
		return x.Blue
	}
	return 0
}

type LockOpenedHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *LockOpenedHistoryRequest) Reset() {
	*x = LockOpenedHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockOpenedHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockOpenedHistoryRequest) ProtoMessage() {}

func (x *LockOpenedHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockOpenedHistoryRequest.ProtoReflect.Descriptor instead.
func (*LockOpenedHistoryRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *LockOpenedHistoryRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type LockOpenedHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Histories []*LockOpenedHistoryResponseHistory `protobuf:"bytes,1,rep,name=histories,proto3" json:"histories,omitempty"`
}

func (x *LockOpenedHistoryResponse) Reset() {
	*x = LockOpenedHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockOpenedHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockOpenedHistoryResponse) ProtoMessage() {}

func (x *LockOpenedHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockOpenedHistoryResponse.ProtoReflect.Descriptor instead.
func (*LockOpenedHistoryResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *LockOpenedHistoryResponse) GetHistories() []*LockOpenedHistoryResponseHistory {
	if x != nil {
		return x.Histories
	}
	return nil
}

type ChangeLedColorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Red      int32  `protobuf:"varint,2,opt,name=red,proto3" json:"red,omitempty"`
	Green    int32  `protobuf:"varint,3,opt,name=green,proto3" json:"green,omitempty"`
	Blue     int32  `protobuf:"varint,4,opt,name=blue,proto3" json:"blue,omitempty"`
}

func (x *ChangeLedColorRequest) Reset() {
	*x = ChangeLedColorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLedColorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLedColorRequest) ProtoMessage() {}

func (x *ChangeLedColorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLedColorRequest.ProtoReflect.Descriptor instead.
func (*ChangeLedColorRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *ChangeLedColorRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ChangeLedColorRequest) GetRed() int32 {
	if x != nil {
		return x.Red
	}
	return 0
}

func (x *ChangeLedColorRequest) GetGreen() int32 {
	if x != nil {
		return x.Green
	}
	return 0
}

func (x *ChangeLedColorRequest) GetBlue() int32 {
	if x != nil {
		return x.Blue
	}
	return 0
}

type ChangeLedColorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeLedColorResponse) Reset() {
	*x = ChangeLedColorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLedColorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLedColorResponse) ProtoMessage() {}

func (x *ChangeLedColorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLedColorResponse.ProtoReflect.Descriptor instead.
func (*ChangeLedColorResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

type OpenDoorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *OpenDoorRequest) Reset() {
	*x = OpenDoorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenDoorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenDoorRequest) ProtoMessage() {}

func (x *OpenDoorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenDoorRequest.ProtoReflect.Descriptor instead.
func (*OpenDoorRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *OpenDoorRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type OpenDoorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpenDoorResponse) Reset() {
	*x = OpenDoorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenDoorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenDoorResponse) ProtoMessage() {}

func (x *OpenDoorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenDoorResponse.ProtoReflect.Descriptor instead.
func (*OpenDoorResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

type GetDeviceIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetDeviceIdsRequest) Reset() {
	*x = GetDeviceIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceIdsRequest) ProtoMessage() {}

func (x *GetDeviceIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceIdsRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceIdsRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetDeviceIdsRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9}
}

func (x *Device) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetDeviceIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *GetDeviceIdsResponse) Reset() {
	*x = GetDeviceIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceIdsResponse) ProtoMessage() {}

func (x *GetDeviceIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceIdsResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceIdsResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetDeviceIdsResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type EmployeePresenceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *EmployeePresenceStatusRequest) Reset() {
	*x = EmployeePresenceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeePresenceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeePresenceStatusRequest) ProtoMessage() {}

func (x *EmployeePresenceStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeePresenceStatusRequest.ProtoReflect.Descriptor instead.
func (*EmployeePresenceStatusRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{11}
}

func (x *EmployeePresenceStatusRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type EmployeePresenceStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	IsPresent bool   `protobuf:"varint,3,opt,name=is_present,json=isPresent,proto3" json:"is_present,omitempty"`
}

func (x *EmployeePresenceStatusResponse) Reset() {
	*x = EmployeePresenceStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeePresenceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeePresenceStatusResponse) ProtoMessage() {}

func (x *EmployeePresenceStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeePresenceStatusResponse.ProtoReflect.Descriptor instead.
func (*EmployeePresenceStatusResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{12}
}

func (x *EmployeePresenceStatusResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *EmployeePresenceStatusResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *EmployeePresenceStatusResponse) GetIsPresent() bool {
	if x != nil {
		return x.IsPresent
	}
	return false
}

type LockOpenedHistoryResponseHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTimestamp     int64 `protobuf:"varint,1,opt,name=server_timestamp,json=serverTimestamp,proto3" json:"server_timestamp,omitempty"`
	TimeAfterStartupSec int32 `protobuf:"varint,2,opt,name=time_after_startup_sec,json=timeAfterStartupSec,proto3" json:"time_after_startup_sec,omitempty"`
}

func (x *LockOpenedHistoryResponseHistory) Reset() {
	*x = LockOpenedHistoryResponseHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockOpenedHistoryResponseHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockOpenedHistoryResponseHistory) ProtoMessage() {}

func (x *LockOpenedHistoryResponseHistory) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockOpenedHistoryResponseHistory.ProtoReflect.Descriptor instead.
func (*LockOpenedHistoryResponseHistory) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *LockOpenedHistoryResponseHistory) GetServerTimestamp() int64 {
	if x != nil {
		return x.ServerTimestamp
	}
	return 0
}

func (x *LockOpenedHistoryResponseHistory) GetTimeAfterStartupSec() int32 {
	if x != nil {
		return x.TimeAfterStartupSec
	}
	return 0
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f,
	0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x22, 0x2e, 0x0a, 0x0f, 0x4c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x6c, 0x75,
	0x65, 0x22, 0x37, 0x0a, 0x18, 0x4c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0xe9, 0x01, 0x0a, 0x19, 0x4c,
	0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x69, 0x72,
	0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c,
	0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x69, 0x0a, 0x07, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x33, 0x0a, 0x16, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x13, 0x74, 0x69, 0x6d, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x75, 0x70, 0x53, 0x65, 0x63, 0x22, 0x70, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67,
	0x72, 0x65, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x62, 0x6c, 0x75, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2e, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x06, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x5a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x72, 0x2e,
	0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65,
	0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x3c, 0x0a,
	0x1d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x1e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x32, 0xb5, 0x06, 0x0a, 0x10, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x73, 0x0a,
	0x08, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x31, 0x2e, 0x69, 0x72, 0x2e, 0x6d,
	0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x65, 0x64,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69,
	0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x4c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x8e, 0x01, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x65, 0x6e, 0x65,
	0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3a, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61,
	0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x63, 0x6b,
	0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75,
	0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x4f, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x83, 0x01, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x65,
	0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x37, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d,
	0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74,
	0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x6e, 0x44, 0x6f, 0x6f, 0x72, 0x12, 0x31, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f,
	0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x6f, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61,
	0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x44, 0x6f, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x80, 0x01, 0x0a,
	0x0f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73,
	0x12, 0x35, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68,
	0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x9e, 0x01, 0x0a, 0x17, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x2e, 0x69, 0x72,
	0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x69,
	0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x37, 0x0a, 0x1c, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x68, 0x6d, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x50, 0x01, 0x5a, 0x15, 0x69, 0x6f, 0x74, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_service_proto_goTypes = []interface{}{
	(*LedColorRequest)(nil),                  // 0: ir.mahmoud.iot_attendance_system.LedColorRequest
	(*LedColorResponse)(nil),                 // 1: ir.mahmoud.iot_attendance_system.LedColorResponse
	(*LockOpenedHistoryRequest)(nil),         // 2: ir.mahmoud.iot_attendance_system.LockOpenedHistoryRequest
	(*LockOpenedHistoryResponse)(nil),        // 3: ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse
	(*ChangeLedColorRequest)(nil),            // 4: ir.mahmoud.iot_attendance_system.ChangeLedColorRequest
	(*ChangeLedColorResponse)(nil),           // 5: ir.mahmoud.iot_attendance_system.ChangeLedColorResponse
	(*OpenDoorRequest)(nil),                  // 6: ir.mahmoud.iot_attendance_system.OpenDoorRequest
	(*OpenDoorResponse)(nil),                 // 7: ir.mahmoud.iot_attendance_system.OpenDoorResponse
	(*GetDeviceIdsRequest)(nil),              // 8: ir.mahmoud.iot_attendance_system.GetDeviceIdsRequest
	(*Device)(nil),                           // 9: ir.mahmoud.iot_attendance_system.Device
	(*GetDeviceIdsResponse)(nil),             // 10: ir.mahmoud.iot_attendance_system.GetDeviceIdsResponse
	(*EmployeePresenceStatusRequest)(nil),    // 11: ir.mahmoud.iot_attendance_system.EmployeePresenceStatusRequest
	(*EmployeePresenceStatusResponse)(nil),   // 12: ir.mahmoud.iot_attendance_system.EmployeePresenceStatusResponse
	(*LockOpenedHistoryResponseHistory)(nil), // 13: ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.history
}
var file_service_proto_depIdxs = []int32{
	13, // 0: ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.histories:type_name -> ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.history
	9,  // 1: ir.mahmoud.iot_attendance_system.GetDeviceIdsResponse.devices:type_name -> ir.mahmoud.iot_attendance_system.Device
	0,  // 2: ir.mahmoud.iot_attendance_system.attendanceSystem.ledColor:input_type -> ir.mahmoud.iot_attendance_system.LedColorRequest
	2,  // 3: ir.mahmoud.iot_attendance_system.attendanceSystem.lockOpenedHistory:input_type -> ir.mahmoud.iot_attendance_system.LockOpenedHistoryRequest
	4,  // 4: ir.mahmoud.iot_attendance_system.attendanceSystem.changeLedColor:input_type -> ir.mahmoud.iot_attendance_system.ChangeLedColorRequest
	6,  // 5: ir.mahmoud.iot_attendance_system.attendanceSystem.openDoor:input_type -> ir.mahmoud.iot_attendance_system.OpenDoorRequest
	8,  // 6: ir.mahmoud.iot_attendance_system.attendanceSystem.getAllDeviceIds:input_type -> ir.mahmoud.iot_attendance_system.GetDeviceIdsRequest
	11, // 7: ir.mahmoud.iot_attendance_system.attendanceSystem.employeesPresenceStatus:input_type -> ir.mahmoud.iot_attendance_system.EmployeePresenceStatusRequest
	1,  // 8: ir.mahmoud.iot_attendance_system.attendanceSystem.ledColor:output_type -> ir.mahmoud.iot_attendance_system.LedColorResponse
	3,  // 9: ir.mahmoud.iot_attendance_system.attendanceSystem.lockOpenedHistory:output_type -> ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse
	5,  // 10: ir.mahmoud.iot_attendance_system.attendanceSystem.changeLedColor:output_type -> ir.mahmoud.iot_attendance_system.ChangeLedColorResponse
	7,  // 11: ir.mahmoud.iot_attendance_system.attendanceSystem.openDoor:output_type -> ir.mahmoud.iot_attendance_system.OpenDoorResponse
	10, // 12: ir.mahmoud.iot_attendance_system.attendanceSystem.getAllDeviceIds:output_type -> ir.mahmoud.iot_attendance_system.GetDeviceIdsResponse
	12, // 13: ir.mahmoud.iot_attendance_system.attendanceSystem.employeesPresenceStatus:output_type -> ir.mahmoud.iot_attendance_system.EmployeePresenceStatusResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedColorRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedColorResponse); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockOpenedHistoryRequest); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockOpenedHistoryResponse); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeLedColorRequest); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeLedColorResponse); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenDoorRequest); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenDoorResponse); i {
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
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceIdsRequest); i {
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
		file_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceIdsResponse); i {
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
		file_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeePresenceStatusRequest); i {
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
		file_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeePresenceStatusResponse); i {
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
		file_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockOpenedHistoryResponseHistory); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
