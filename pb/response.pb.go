// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: response.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CommonResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}
func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CommonResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CommonResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// user
type LoginResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiredIn            int64    `protobuf:"varint,4,opt,name=expired_in,json=expiredIn,proto3" json:"expired_in,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *LoginResponse) GetExpiredIn() int64 {
	if m != nil {
		return m.ExpiredIn
	}
	return 0
}

func (m *LoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetUserProfileResponse struct {
	Data                 *UserProfile `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code                 int32        `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserProfileResponse) Reset()         { *m = GetUserProfileResponse{} }
func (m *GetUserProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserProfileResponse) ProtoMessage()    {}
func (*GetUserProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{2}
}
func (m *GetUserProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserProfileResponse.Unmarshal(m, b)
}
func (m *GetUserProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetUserProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserProfileResponse.Merge(m, src)
}
func (m *GetUserProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserProfileResponse.Size(m)
}
func (m *GetUserProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserProfileResponse proto.InternalMessageInfo

func (m *GetUserProfileResponse) GetData() *UserProfile {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetUserProfileResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetUserProfileResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Location service
type AddLocationResponse struct {
	LocationId           int32    `protobuf:"varint,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int32    `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddLocationResponse) Reset()         { *m = AddLocationResponse{} }
func (m *AddLocationResponse) String() string { return proto.CompactTextString(m) }
func (*AddLocationResponse) ProtoMessage()    {}
func (*AddLocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{3}
}
func (m *AddLocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddLocationResponse.Unmarshal(m, b)
}
func (m *AddLocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddLocationResponse.Marshal(b, m, deterministic)
}
func (m *AddLocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddLocationResponse.Merge(m, src)
}
func (m *AddLocationResponse) XXX_Size() int {
	return xxx_messageInfo_AddLocationResponse.Size(m)
}
func (m *AddLocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddLocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddLocationResponse proto.InternalMessageInfo

func (m *AddLocationResponse) GetLocationId() int32 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *AddLocationResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddLocationResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type Location struct {
	LocationId           int32    `protobuf:"varint,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Areas                []*Area  `protobuf:"bytes,5,rep,name=areas,proto3" json:"areas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{4}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLocationId() int32 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *Location) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Location) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Location) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Location) GetAreas() []*Area {
	if m != nil {
		return m.Areas
	}
	return nil
}

type ListLocationResponse struct {
	Data                 []*Location `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Code                 int32       `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListLocationResponse) Reset()         { *m = ListLocationResponse{} }
func (m *ListLocationResponse) String() string { return proto.CompactTextString(m) }
func (*ListLocationResponse) ProtoMessage()    {}
func (*ListLocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{5}
}
func (m *ListLocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLocationResponse.Unmarshal(m, b)
}
func (m *ListLocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLocationResponse.Marshal(b, m, deterministic)
}
func (m *ListLocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLocationResponse.Merge(m, src)
}
func (m *ListLocationResponse) XXX_Size() int {
	return xxx_messageInfo_ListLocationResponse.Size(m)
}
func (m *ListLocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLocationResponse proto.InternalMessageInfo

func (m *ListLocationResponse) GetData() []*Location {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListLocationResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListLocationResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Area service
type AddAreaResponse struct {
	AreaId               int32    `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAreaResponse) Reset()         { *m = AddAreaResponse{} }
func (m *AddAreaResponse) String() string { return proto.CompactTextString(m) }
func (*AddAreaResponse) ProtoMessage()    {}
func (*AddAreaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{6}
}
func (m *AddAreaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAreaResponse.Unmarshal(m, b)
}
func (m *AddAreaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAreaResponse.Marshal(b, m, deterministic)
}
func (m *AddAreaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAreaResponse.Merge(m, src)
}
func (m *AddAreaResponse) XXX_Size() int {
	return xxx_messageInfo_AddAreaResponse.Size(m)
}
func (m *AddAreaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAreaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddAreaResponse proto.InternalMessageInfo

func (m *AddAreaResponse) GetAreaId() int32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *AddAreaResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AddAreaResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListAreaResponse struct {
	Data                 []*Area  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAreaResponse) Reset()         { *m = ListAreaResponse{} }
func (m *ListAreaResponse) String() string { return proto.CompactTextString(m) }
func (*ListAreaResponse) ProtoMessage()    {}
func (*ListAreaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{7}
}
func (m *ListAreaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAreaResponse.Unmarshal(m, b)
}
func (m *ListAreaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAreaResponse.Marshal(b, m, deterministic)
}
func (m *ListAreaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAreaResponse.Merge(m, src)
}
func (m *ListAreaResponse) XXX_Size() int {
	return xxx_messageInfo_ListAreaResponse.Size(m)
}
func (m *ListAreaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAreaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAreaResponse proto.InternalMessageInfo

func (m *ListAreaResponse) GetData() []*Area {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListAreaResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListAreaResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// screen
type ListScreenResponse struct {
	Data                 []*ScreenInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Code                 int32         `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string        `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListScreenResponse) Reset()         { *m = ListScreenResponse{} }
func (m *ListScreenResponse) String() string { return proto.CompactTextString(m) }
func (*ListScreenResponse) ProtoMessage()    {}
func (*ListScreenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{8}
}
func (m *ListScreenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListScreenResponse.Unmarshal(m, b)
}
func (m *ListScreenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListScreenResponse.Marshal(b, m, deterministic)
}
func (m *ListScreenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListScreenResponse.Merge(m, src)
}
func (m *ListScreenResponse) XXX_Size() int {
	return xxx_messageInfo_ListScreenResponse.Size(m)
}
func (m *ListScreenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListScreenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListScreenResponse proto.InternalMessageInfo

func (m *ListScreenResponse) GetData() []*ScreenInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListScreenResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListScreenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddScreenResponse struct {
	Data                 *ScreenInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code                 int32       `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddScreenResponse) Reset()         { *m = AddScreenResponse{} }
func (m *AddScreenResponse) String() string { return proto.CompactTextString(m) }
func (*AddScreenResponse) ProtoMessage()    {}
func (*AddScreenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{9}
}
func (m *AddScreenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddScreenResponse.Unmarshal(m, b)
}
func (m *AddScreenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddScreenResponse.Marshal(b, m, deterministic)
}
func (m *AddScreenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddScreenResponse.Merge(m, src)
}
func (m *AddScreenResponse) XXX_Size() int {
	return xxx_messageInfo_AddScreenResponse.Size(m)
}
func (m *AddScreenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddScreenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddScreenResponse proto.InternalMessageInfo

func (m *AddScreenResponse) GetData() *ScreenInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AddScreenResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AddScreenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// activities
type ListActivitiesResponse struct {
	Activities           *ScreenActivities `protobuf:"bytes,1,opt,name=activities,proto3" json:"activities,omitempty"`
	Code                 int32             `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListActivitiesResponse) Reset()         { *m = ListActivitiesResponse{} }
func (m *ListActivitiesResponse) String() string { return proto.CompactTextString(m) }
func (*ListActivitiesResponse) ProtoMessage()    {}
func (*ListActivitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{10}
}
func (m *ListActivitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListActivitiesResponse.Unmarshal(m, b)
}
func (m *ListActivitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListActivitiesResponse.Marshal(b, m, deterministic)
}
func (m *ListActivitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListActivitiesResponse.Merge(m, src)
}
func (m *ListActivitiesResponse) XXX_Size() int {
	return xxx_messageInfo_ListActivitiesResponse.Size(m)
}
func (m *ListActivitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListActivitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListActivitiesResponse proto.InternalMessageInfo

func (m *ListActivitiesResponse) GetActivities() *ScreenActivities {
	if m != nil {
		return m.Activities
	}
	return nil
}

func (m *ListActivitiesResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListActivitiesResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UploadVideoResponse struct {
	VideoId              int32    `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadVideoResponse) Reset()         { *m = UploadVideoResponse{} }
func (m *UploadVideoResponse) String() string { return proto.CompactTextString(m) }
func (*UploadVideoResponse) ProtoMessage()    {}
func (*UploadVideoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{11}
}
func (m *UploadVideoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadVideoResponse.Unmarshal(m, b)
}
func (m *UploadVideoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadVideoResponse.Marshal(b, m, deterministic)
}
func (m *UploadVideoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadVideoResponse.Merge(m, src)
}
func (m *UploadVideoResponse) XXX_Size() int {
	return xxx_messageInfo_UploadVideoResponse.Size(m)
}
func (m *UploadVideoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadVideoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadVideoResponse proto.InternalMessageInfo

func (m *UploadVideoResponse) GetVideoId() int32 {
	if m != nil {
		return m.VideoId
	}
	return 0
}

func (m *UploadVideoResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UploadVideoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListVideoResponse struct {
	Data                 *Videos  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVideoResponse) Reset()         { *m = ListVideoResponse{} }
func (m *ListVideoResponse) String() string { return proto.CompactTextString(m) }
func (*ListVideoResponse) ProtoMessage()    {}
func (*ListVideoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{12}
}
func (m *ListVideoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVideoResponse.Unmarshal(m, b)
}
func (m *ListVideoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVideoResponse.Marshal(b, m, deterministic)
}
func (m *ListVideoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVideoResponse.Merge(m, src)
}
func (m *ListVideoResponse) XXX_Size() int {
	return xxx_messageInfo_ListVideoResponse.Size(m)
}
func (m *ListVideoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVideoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVideoResponse proto.InternalMessageInfo

func (m *ListVideoResponse) GetData() *Videos {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListVideoResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListVideoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetVideoResponse struct {
	Data                 *Video   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVideoResponse) Reset()         { *m = GetVideoResponse{} }
func (m *GetVideoResponse) String() string { return proto.CompactTextString(m) }
func (*GetVideoResponse) ProtoMessage()    {}
func (*GetVideoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{13}
}
func (m *GetVideoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVideoResponse.Unmarshal(m, b)
}
func (m *GetVideoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVideoResponse.Marshal(b, m, deterministic)
}
func (m *GetVideoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVideoResponse.Merge(m, src)
}
func (m *GetVideoResponse) XXX_Size() int {
	return xxx_messageInfo_GetVideoResponse.Size(m)
}
func (m *GetVideoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVideoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVideoResponse proto.InternalMessageInfo

func (m *GetVideoResponse) GetData() *Video {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetVideoResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetVideoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// campaign
type ListCampaignResponse struct {
	Data                 []*Campaign `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Code                 int32       `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListCampaignResponse) Reset()         { *m = ListCampaignResponse{} }
func (m *ListCampaignResponse) String() string { return proto.CompactTextString(m) }
func (*ListCampaignResponse) ProtoMessage()    {}
func (*ListCampaignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{14}
}
func (m *ListCampaignResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCampaignResponse.Unmarshal(m, b)
}
func (m *ListCampaignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCampaignResponse.Marshal(b, m, deterministic)
}
func (m *ListCampaignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCampaignResponse.Merge(m, src)
}
func (m *ListCampaignResponse) XXX_Size() int {
	return xxx_messageInfo_ListCampaignResponse.Size(m)
}
func (m *ListCampaignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCampaignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCampaignResponse proto.InternalMessageInfo

func (m *ListCampaignResponse) GetData() []*Campaign {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListCampaignResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListCampaignResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CampaignResponse struct {
	Data                 *Campaign `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code                 int32     `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CampaignResponse) Reset()         { *m = CampaignResponse{} }
func (m *CampaignResponse) String() string { return proto.CompactTextString(m) }
func (*CampaignResponse) ProtoMessage()    {}
func (*CampaignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{15}
}
func (m *CampaignResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignResponse.Unmarshal(m, b)
}
func (m *CampaignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignResponse.Marshal(b, m, deterministic)
}
func (m *CampaignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignResponse.Merge(m, src)
}
func (m *CampaignResponse) XXX_Size() int {
	return xxx_messageInfo_CampaignResponse.Size(m)
}
func (m *CampaignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignResponse proto.InternalMessageInfo

func (m *CampaignResponse) GetData() *Campaign {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CampaignResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CampaignResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonResponse)(nil), "pb.CommonResponse")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*GetUserProfileResponse)(nil), "pb.GetUserProfileResponse")
	proto.RegisterType((*AddLocationResponse)(nil), "pb.AddLocationResponse")
	proto.RegisterType((*Location)(nil), "pb.Location")
	proto.RegisterType((*ListLocationResponse)(nil), "pb.ListLocationResponse")
	proto.RegisterType((*AddAreaResponse)(nil), "pb.AddAreaResponse")
	proto.RegisterType((*ListAreaResponse)(nil), "pb.ListAreaResponse")
	proto.RegisterType((*ListScreenResponse)(nil), "pb.ListScreenResponse")
	proto.RegisterType((*AddScreenResponse)(nil), "pb.AddScreenResponse")
	proto.RegisterType((*ListActivitiesResponse)(nil), "pb.ListActivitiesResponse")
	proto.RegisterType((*UploadVideoResponse)(nil), "pb.UploadVideoResponse")
	proto.RegisterType((*ListVideoResponse)(nil), "pb.ListVideoResponse")
	proto.RegisterType((*GetVideoResponse)(nil), "pb.GetVideoResponse")
	proto.RegisterType((*ListCampaignResponse)(nil), "pb.ListCampaignResponse")
	proto.RegisterType((*CampaignResponse)(nil), "pb.CampaignResponse")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_0fbc901015fa5021) }

var fileDescriptor_0fbc901015fa5021 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xe5, 0xfc, 0x6b, 0x32, 0x49, 0xd3, 0x74, 0x5b, 0x05, 0x83, 0x68, 0x09, 0xee, 0x25,
	0xa7, 0x1c, 0x0a, 0x37, 0x4e, 0xa1, 0x87, 0x2a, 0x52, 0x84, 0x90, 0xa1, 0x08, 0x71, 0x20, 0xda,
	0x78, 0x27, 0x61, 0xd5, 0xc4, 0x6b, 0xed, 0x2e, 0x15, 0x48, 0x7c, 0x08, 0xf8, 0xc6, 0x68, 0xd7,
	0xff, 0x0b, 0x04, 0xc5, 0xea, 0xcd, 0xf3, 0xf2, 0x76, 0x7e, 0x33, 0x2f, 0x6b, 0x19, 0xfa, 0x12,
	0x55, 0x24, 0x42, 0x85, 0x93, 0x48, 0x0a, 0x2d, 0x48, 0x2d, 0x5a, 0x3e, 0xe9, 0xea, 0xef, 0x11,
	0xaa, 0x58, 0xf0, 0xde, 0x40, 0xff, 0x4a, 0x6c, 0xb7, 0x22, 0xf4, 0x13, 0x23, 0x21, 0xd0, 0x08,
	0x04, 0x43, 0xd7, 0x19, 0x39, 0xe3, 0xa6, 0x6f, 0x9f, 0x89, 0x0b, 0x07, 0x5b, 0x54, 0x8a, 0xae,
	0xd1, 0xad, 0x8d, 0x9c, 0x71, 0xc7, 0x4f, 0x4b, 0xd2, 0x87, 0x1a, 0x67, 0x6e, 0xdd, 0x7a, 0x6b,
	0x9c, 0x79, 0x3f, 0x1d, 0x38, 0x9c, 0x8b, 0x35, 0xcf, 0xfb, 0x3d, 0x87, 0x1e, 0x0d, 0x02, 0x54,
	0x6a, 0xa1, 0xc5, 0x2d, 0x86, 0xb6, 0x6f, 0xc7, 0xef, 0xc6, 0xda, 0x7b, 0x23, 0x91, 0x0b, 0x38,
	0x94, 0xb8, 0x92, 0xa8, 0xbe, 0x24, 0x9e, 0x18, 0xd2, 0x4b, 0xc4, 0xd8, 0x74, 0x06, 0x80, 0xdf,
	0x22, 0x2e, 0x91, 0x2d, 0x78, 0xe8, 0x36, 0x46, 0xce, 0xb8, 0xee, 0x77, 0x12, 0x65, 0x16, 0x16,
	0x47, 0xac, 0x97, 0x46, 0xf4, 0x6e, 0x61, 0x78, 0x8d, 0xfa, 0x46, 0xa1, 0x7c, 0x2b, 0xc5, 0x8a,
	0x6f, 0x30, 0x1b, 0xed, 0x02, 0x1a, 0x8c, 0x6a, 0x6a, 0x47, 0xea, 0x5e, 0x1e, 0x4d, 0xa2, 0xe5,
	0xa4, 0x68, 0xb3, 0x3f, 0x66, 0x79, 0xd4, 0xfe, 0x9e, 0xc7, 0x3d, 0x18, 0x83, 0x93, 0x29, 0x63,
	0x73, 0x11, 0x50, 0xcd, 0x0b, 0xa1, 0x3e, 0x83, 0xee, 0x26, 0xd1, 0x16, 0x9c, 0x25, 0xd9, 0x42,
	0x2a, 0xcd, 0xd8, 0x8e, 0x84, 0x53, 0x7e, 0x3d, 0xe7, 0x7b, 0xbf, 0x1c, 0x68, 0xa7, 0x8c, 0xff,
	0xf7, 0x26, 0xd0, 0x08, 0xe9, 0x36, 0x6d, 0x6c, 0x9f, 0x0d, 0x8f, 0x32, 0x26, 0x51, 0xa9, 0x74,
	0x83, 0xa4, 0x24, 0x43, 0x68, 0x29, 0x4d, 0xf5, 0x57, 0x65, 0x33, 0x6e, 0xfa, 0x49, 0x45, 0xce,
	0xa1, 0x49, 0x25, 0x52, 0xe5, 0x36, 0x47, 0xf5, 0x71, 0xf7, 0xb2, 0x6d, 0xd2, 0x9a, 0x4a, 0xa4,
	0x7e, 0x2c, 0x7b, 0x2b, 0x38, 0x9d, 0x73, 0xa5, 0xff, 0x58, 0x7d, 0x94, 0x85, 0x6c, 0x8e, 0xf5,
	0xcc, 0xb1, 0xcc, 0x53, 0x25, 0xe1, 0x8f, 0x70, 0x34, 0x65, 0xcc, 0x92, 0x53, 0xc4, 0x23, 0x38,
	0x30, 0x33, 0xe4, 0xdb, 0xb7, 0x4c, 0x19, 0x6f, 0xbe, 0x47, 0xe7, 0xcf, 0x30, 0x30, 0x1b, 0x94,
	0x5a, 0x3f, 0x2d, 0x4d, 0x9f, 0x2f, 0x5d, 0x65, 0xf2, 0x15, 0x10, 0xd3, 0xff, 0x5d, 0x20, 0x11,
	0xf3, 0x7c, 0xbc, 0x12, 0xa1, 0x6f, 0x08, 0xb1, 0x63, 0x16, 0xae, 0x44, 0x25, 0x0e, 0xc2, 0xf1,
	0x94, 0xb1, 0x7f, 0x62, 0x9c, 0x07, 0xc2, 0xfc, 0x80, 0xa1, 0x8d, 0x2b, 0xd0, 0xfc, 0x8e, 0x6b,
	0x8e, 0x2a, 0x63, 0xbd, 0x04, 0xa0, 0x99, 0x9a, 0x10, 0x4f, 0x73, 0x62, 0xe1, 0x44, 0xc1, 0xb7,
	0xf7, 0x9f, 0x75, 0x72, 0x13, 0x6d, 0x04, 0x65, 0x1f, 0x38, 0x43, 0x91, 0xa1, 0x1f, 0x43, 0xfb,
	0xce, 0x08, 0xf9, 0x5d, 0x38, 0xb0, 0xf5, 0xde, 0x97, 0x81, 0xc2, 0xb1, 0xd9, 0xae, 0xdc, 0xfd,
	0xbc, 0x14, 0x22, 0x98, 0x95, 0xac, 0x41, 0x55, 0x0a, 0x70, 0x01, 0x83, 0x6b, 0xbc, 0x47, 0x38,
	0x2b, 0x11, 0x3a, 0x19, 0xa1, 0xe2, 0x85, 0xb3, 0xaf, 0xe4, 0x15, 0xdd, 0x46, 0x94, 0xaf, 0x77,
	0xbe, 0x92, 0x99, 0xa7, 0x0a, 0x67, 0x09, 0x83, 0x1d, 0x0c, 0xe7, 0x21, 0x18, 0xaf, 0x5b, 0x9f,
	0x1a, 0x93, 0x57, 0xd1, 0x72, 0xd9, 0xb2, 0xdf, 0xad, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x12, 0x1e, 0x3b, 0x9f, 0xda, 0x06, 0x00, 0x00,
}