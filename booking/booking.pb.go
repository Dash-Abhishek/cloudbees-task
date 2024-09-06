// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: booking/booking.proto

package booking

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

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *BookingRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *BookingRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From  string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	User  *User   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Price float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Seat  *Seat   `protobuf:"bytes,6,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{1}
}

func (x *Receipt) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Receipt) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Receipt) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Receipt) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Receipt) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Receipt) GetSeat() *Seat {
	if x != nil {
		return x.Seat
	}
	return nil
}

type Seat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatNumber int32  `protobuf:"varint,1,opt,name=seatNumber,proto3" json:"seatNumber,omitempty"`
	Section    string `protobuf:"bytes,2,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *Seat) Reset() {
	*x = Seat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seat) ProtoMessage() {}

func (x *Seat) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seat.ProtoReflect.Descriptor instead.
func (*Seat) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{2}
}

func (x *Seat) GetSeatNumber() int32 {
	if x != nil {
		return x.SeatNumber
	}
	return 0
}

func (x *Seat) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type BookingId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BookingId) Reset() {
	*x = BookingId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingId) ProtoMessage() {}

func (x *BookingId) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingId.ProtoReflect.Descriptor instead.
func (*BookingId) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{4}
}

func (x *BookingId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{5}
}

func (x *Section) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Allocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId string `protobuf:"bytes,1,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	User      *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Seat      *Seat  `protobuf:"bytes,3,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *Allocation) Reset() {
	*x = Allocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Allocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Allocation) ProtoMessage() {}

func (x *Allocation) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Allocation.ProtoReflect.Descriptor instead.
func (*Allocation) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{6}
}

func (x *Allocation) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *Allocation) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Allocation) GetSeat() *Seat {
	if x != nil {
		return x.Seat
	}
	return nil
}

type AllocationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allocations []*Allocation `protobuf:"bytes,1,rep,name=allocations,proto3" json:"allocations,omitempty"`
}

func (x *AllocationList) Reset() {
	*x = AllocationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationList) ProtoMessage() {}

func (x *AllocationList) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationList.ProtoReflect.Descriptor instead.
func (*AllocationList) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{7}
}

func (x *AllocationList) GetAllocations() []*Allocation {
	if x != nil {
		return x.Allocations
	}
	return nil
}

type UpdateSeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId *BookingId `protobuf:"bytes,1,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	Seat      *Seat      `protobuf:"bytes,2,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *UpdateSeatRequest) Reset() {
	*x = UpdateSeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeatRequest) ProtoMessage() {}

func (x *UpdateSeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeatRequest.ProtoReflect.Descriptor instead.
func (*UpdateSeatRequest) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateSeatRequest) GetBookingId() *BookingId {
	if x != nil {
		return x.BookingId
	}
	return nil
}

func (x *UpdateSeatRequest) GetSeat() *Seat {
	if x != nil {
		return x.Seat
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_booking_booking_proto_rawDescGZIP(), []int{9}
}

var File_booking_booking_proto protoreflect.FileDescriptor

var file_booking_booking_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x19, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x04,
	0x73, 0x65, 0x61, 0x74, 0x22, 0x40, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x21,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x1d, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x60, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x04, 0x73, 0x65,
	0x61, 0x74, 0x22, 0x3f, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd7, 0x01, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x0a, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x12, 0x12, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x08, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x1e, 0x0a,
	0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x0a, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x08, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44,
	0x61, 0x73, 0x68, 0x2d, 0x41, 0x62, 0x68, 0x69, 0x73, 0x68, 0x65, 0x6b, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x62, 0x65, 0x65, 0x73, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_booking_proto_rawDescOnce sync.Once
	file_booking_booking_proto_rawDescData = file_booking_booking_proto_rawDesc
)

func file_booking_booking_proto_rawDescGZIP() []byte {
	file_booking_booking_proto_rawDescOnce.Do(func() {
		file_booking_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_booking_proto_rawDescData)
	})
	return file_booking_booking_proto_rawDescData
}

var file_booking_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_booking_booking_proto_goTypes = []interface{}{
	(*BookingRequest)(nil),    // 0: BookingRequest
	(*Receipt)(nil),           // 1: Receipt
	(*Seat)(nil),              // 2: Seat
	(*User)(nil),              // 3: User
	(*BookingId)(nil),         // 4: BookingId
	(*Section)(nil),           // 5: Section
	(*Allocation)(nil),        // 6: Allocation
	(*AllocationList)(nil),    // 7: AllocationList
	(*UpdateSeatRequest)(nil), // 8: UpdateSeatRequest
	(*Empty)(nil),             // 9: Empty
}
var file_booking_booking_proto_depIdxs = []int32{
	3,  // 0: BookingRequest.user:type_name -> User
	3,  // 1: Receipt.user:type_name -> User
	2,  // 2: Receipt.seat:type_name -> Seat
	3,  // 3: Allocation.user:type_name -> User
	2,  // 4: Allocation.seat:type_name -> Seat
	6,  // 5: AllocationList.allocations:type_name -> Allocation
	4,  // 6: UpdateSeatRequest.bookingId:type_name -> BookingId
	2,  // 7: UpdateSeatRequest.seat:type_name -> Seat
	0,  // 8: Booking.Create:input_type -> BookingRequest
	4,  // 9: Booking.GetReceipt:input_type -> BookingId
	8,  // 10: Booking.UpdateSeat:input_type -> UpdateSeatRequest
	4,  // 11: Booking.Cancel:input_type -> BookingId
	5,  // 12: Booking.GetSeatAllocations:input_type -> Section
	1,  // 13: Booking.Create:output_type -> Receipt
	1,  // 14: Booking.GetReceipt:output_type -> Receipt
	1,  // 15: Booking.UpdateSeat:output_type -> Receipt
	9,  // 16: Booking.Cancel:output_type -> Empty
	7,  // 17: Booking.GetSeatAllocations:output_type -> AllocationList
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_booking_booking_proto_init() }
func file_booking_booking_proto_init() {
	if File_booking_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequest); i {
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
		file_booking_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
		file_booking_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seat); i {
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
		file_booking_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_booking_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingId); i {
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
		file_booking_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Section); i {
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
		file_booking_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Allocation); i {
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
		file_booking_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationList); i {
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
		file_booking_booking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSeatRequest); i {
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
		file_booking_booking_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_booking_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_booking_proto_goTypes,
		DependencyIndexes: file_booking_booking_proto_depIdxs,
		MessageInfos:      file_booking_booking_proto_msgTypes,
	}.Build()
	File_booking_booking_proto = out.File
	file_booking_booking_proto_rawDesc = nil
	file_booking_booking_proto_goTypes = nil
	file_booking_booking_proto_depIdxs = nil
}
