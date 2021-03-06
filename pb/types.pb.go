// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

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

type Status int32

const (
	Status_activated   Status = 0
	Status_deactivated Status = 1
)

var Status_name = map[int32]string{
	0: "activated",
	1: "deactivated",
}

var Status_value = map[string]int32{
	"activated":   0,
	"deactivated": 1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

type Availability int32

const (
	Availability_offline Availability = 0
	Availability_online  Availability = 1
)

var Availability_name = map[int32]string{
	0: "offline",
	1: "online",
}

var Availability_value = map[string]int32{
	"offline": 0,
	"online":  1,
}

func (x Availability) String() string {
	return proto.EnumName(Availability_name, int32(x))
}

func (Availability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

type CampaignStatus int32

const (
	CampaignStatus_empty  CampaignStatus = 0
	CampaignStatus_cancel CampaignStatus = 1
	CampaignStatus_draft  CampaignStatus = 2
	CampaignStatus_live   CampaignStatus = 3
	CampaignStatus_pause  CampaignStatus = 4
)

var CampaignStatus_name = map[int32]string{
	0: "empty",
	1: "cancel",
	2: "draft",
	3: "live",
	4: "pause",
}

var CampaignStatus_value = map[string]int32{
	"empty":  0,
	"cancel": 1,
	"draft":  2,
	"live":   3,
	"pause":  4,
}

func (x CampaignStatus) String() string {
	return proto.EnumName(CampaignStatus_name, int32(x))
}

func (CampaignStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

type CampaignTimeRule int32

const (
	CampaignTimeRule_free_in_this_time CampaignTimeRule = 0
)

var CampaignTimeRule_name = map[int32]string{
	0: "free_in_this_time",
}

var CampaignTimeRule_value = map[string]int32{
	"free_in_this_time": 0,
}

func (x CampaignTimeRule) String() string {
	return proto.EnumName(CampaignTimeRule_name, int32(x))
}

func (CampaignTimeRule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

type EmptyMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyMessage) Reset()         { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()    {}
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}
func (m *EmptyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyMessage.Unmarshal(m, b)
}
func (m *EmptyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyMessage.Marshal(b, m, deterministic)
}
func (m *EmptyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyMessage.Merge(m, src)
}
func (m *EmptyMessage) XXX_Size() int {
	return xxx_messageInfo_EmptyMessage.Size(m)
}
func (m *EmptyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyMessage proto.InternalMessageInfo

type UserProfile struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Company              string   `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	Role                 string   `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	Status               Status   `protobuf:"varint,7,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	CreatedAt            int64    `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}
func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (m *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(m, src)
}
func (m *UserProfile) XXX_Size() int {
	return xxx_messageInfo_UserProfile.Size(m)
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserProfile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserProfile) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserProfile) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *UserProfile) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *UserProfile) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_activated
}

func (m *UserProfile) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Area struct {
	AreaId               int32    `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LocationId           int32    `protobuf:"varint,3,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Status               Status   `protobuf:"varint,4,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}
func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (m *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(m, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetAreaId() int32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *Area) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Area) GetLocationId() int32 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *Area) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_activated
}

type ScreenActivity struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Activity             string   `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScreenActivity) Reset()         { *m = ScreenActivity{} }
func (m *ScreenActivity) String() string { return proto.CompactTextString(m) }
func (*ScreenActivity) ProtoMessage()    {}
func (*ScreenActivity) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}
func (m *ScreenActivity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreenActivity.Unmarshal(m, b)
}
func (m *ScreenActivity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreenActivity.Marshal(b, m, deterministic)
}
func (m *ScreenActivity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreenActivity.Merge(m, src)
}
func (m *ScreenActivity) XXX_Size() int {
	return xxx_messageInfo_ScreenActivity.Size(m)
}
func (m *ScreenActivity) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreenActivity.DiscardUnknown(m)
}

var xxx_messageInfo_ScreenActivity proto.InternalMessageInfo

func (m *ScreenActivity) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ScreenActivity) GetActivity() string {
	if m != nil {
		return m.Activity
	}
	return ""
}

type ScreenActivities struct {
	Data                 []*ScreenActivity `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ScreenActivities) Reset()         { *m = ScreenActivities{} }
func (m *ScreenActivities) String() string { return proto.CompactTextString(m) }
func (*ScreenActivities) ProtoMessage()    {}
func (*ScreenActivities) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}
func (m *ScreenActivities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreenActivities.Unmarshal(m, b)
}
func (m *ScreenActivities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreenActivities.Marshal(b, m, deterministic)
}
func (m *ScreenActivities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreenActivities.Merge(m, src)
}
func (m *ScreenActivities) XXX_Size() int {
	return xxx_messageInfo_ScreenActivities.Size(m)
}
func (m *ScreenActivities) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreenActivities.DiscardUnknown(m)
}

var xxx_messageInfo_ScreenActivities proto.InternalMessageInfo

func (m *ScreenActivities) GetData() []*ScreenActivity {
	if m != nil {
		return m.Data
	}
	return nil
}

type ScreenInfo struct {
	ScreenId             int32        `protobuf:"varint,1,opt,name=screen_id,json=screenId,proto3" json:"screen_id,omitempty"`
	ScreenName           string       `protobuf:"bytes,2,opt,name=screen_name,json=screenName,proto3" json:"screen_name,omitempty"`
	LocationId           int32        `protobuf:"varint,3,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	LocationName         string       `protobuf:"bytes,4,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
	LocationAddress      string       `protobuf:"bytes,5,opt,name=location_address,json=locationAddress,proto3" json:"location_address,omitempty"`
	AreaId               int32        `protobuf:"varint,6,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	AreaName             string       `protobuf:"bytes,7,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	Status               Status       `protobuf:"varint,8,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Availability         Availability `protobuf:"varint,9,opt,name=availability,proto3,enum=pb.Availability" json:"availability,omitempty"`
	TypeOfDevice         string       `protobuf:"bytes,10,opt,name=type_of_device,json=typeOfDevice,proto3" json:"type_of_device,omitempty"`
	MacAddress           string       `protobuf:"bytes,11,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Os                   string       `protobuf:"bytes,12,opt,name=os,proto3" json:"os,omitempty"`
	AppVersion           string       `protobuf:"bytes,13,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	IpAddress            string       `protobuf:"bytes,14,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScreenInfo) Reset()         { *m = ScreenInfo{} }
func (m *ScreenInfo) String() string { return proto.CompactTextString(m) }
func (*ScreenInfo) ProtoMessage()    {}
func (*ScreenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{5}
}
func (m *ScreenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScreenInfo.Unmarshal(m, b)
}
func (m *ScreenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScreenInfo.Marshal(b, m, deterministic)
}
func (m *ScreenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScreenInfo.Merge(m, src)
}
func (m *ScreenInfo) XXX_Size() int {
	return xxx_messageInfo_ScreenInfo.Size(m)
}
func (m *ScreenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ScreenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ScreenInfo proto.InternalMessageInfo

func (m *ScreenInfo) GetScreenId() int32 {
	if m != nil {
		return m.ScreenId
	}
	return 0
}

func (m *ScreenInfo) GetScreenName() string {
	if m != nil {
		return m.ScreenName
	}
	return ""
}

func (m *ScreenInfo) GetLocationId() int32 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *ScreenInfo) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *ScreenInfo) GetLocationAddress() string {
	if m != nil {
		return m.LocationAddress
	}
	return ""
}

func (m *ScreenInfo) GetAreaId() int32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *ScreenInfo) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *ScreenInfo) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_activated
}

func (m *ScreenInfo) GetAvailability() Availability {
	if m != nil {
		return m.Availability
	}
	return Availability_offline
}

func (m *ScreenInfo) GetTypeOfDevice() string {
	if m != nil {
		return m.TypeOfDevice
	}
	return ""
}

func (m *ScreenInfo) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *ScreenInfo) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *ScreenInfo) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *ScreenInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type Video struct {
	VideoId              string   `protobuf:"bytes,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	FileName             string   `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Format               string   `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	Url                  string   `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title                string   `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Tag                  string   `protobuf:"bytes,8,opt,name=tag,proto3" json:"tag,omitempty"`
	Picture              string   `protobuf:"bytes,9,opt,name=picture,proto3" json:"picture,omitempty"`
	Status               Status   `protobuf:"varint,10,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{6}
}
func (m *Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Video.Unmarshal(m, b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Video.Marshal(b, m, deterministic)
}
func (m *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(m, src)
}
func (m *Video) XXX_Size() int {
	return xxx_messageInfo_Video.Size(m)
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

func (m *Video) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Video) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Video) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *Video) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Video) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Video) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Video) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Video) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *Video) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_activated
}

type Videos struct {
	Data                 []*Video `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Videos) Reset()         { *m = Videos{} }
func (m *Videos) String() string { return proto.CompactTextString(m) }
func (*Videos) ProtoMessage()    {}
func (*Videos) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{7}
}
func (m *Videos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Videos.Unmarshal(m, b)
}
func (m *Videos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Videos.Marshal(b, m, deterministic)
}
func (m *Videos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Videos.Merge(m, src)
}
func (m *Videos) XXX_Size() int {
	return xxx_messageInfo_Videos.Size(m)
}
func (m *Videos) XXX_DiscardUnknown() {
	xxx_messageInfo_Videos.DiscardUnknown(m)
}

var xxx_messageInfo_Videos proto.InternalMessageInfo

func (m *Videos) GetData() []*Video {
	if m != nil {
		return m.Data
	}
	return nil
}

type AdSetBasic struct {
	TotalVideo           int32    `protobuf:"varint,1,opt,name=total_video,json=totalVideo,proto3" json:"total_video,omitempty"`
	Played               int32    `protobuf:"varint,2,opt,name=played,proto3" json:"played,omitempty"`
	Videos               []string `protobuf:"bytes,3,rep,name=videos,proto3" json:"videos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdSetBasic) Reset()         { *m = AdSetBasic{} }
func (m *AdSetBasic) String() string { return proto.CompactTextString(m) }
func (*AdSetBasic) ProtoMessage()    {}
func (*AdSetBasic) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{8}
}
func (m *AdSetBasic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdSetBasic.Unmarshal(m, b)
}
func (m *AdSetBasic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdSetBasic.Marshal(b, m, deterministic)
}
func (m *AdSetBasic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdSetBasic.Merge(m, src)
}
func (m *AdSetBasic) XXX_Size() int {
	return xxx_messageInfo_AdSetBasic.Size(m)
}
func (m *AdSetBasic) XXX_DiscardUnknown() {
	xxx_messageInfo_AdSetBasic.DiscardUnknown(m)
}

var xxx_messageInfo_AdSetBasic proto.InternalMessageInfo

func (m *AdSetBasic) GetTotalVideo() int32 {
	if m != nil {
		return m.TotalVideo
	}
	return 0
}

func (m *AdSetBasic) GetPlayed() int32 {
	if m != nil {
		return m.Played
	}
	return 0
}

func (m *AdSetBasic) GetVideos() []string {
	if m != nil {
		return m.Videos
	}
	return nil
}

type AdSet struct {
	TotalVideo           int32          `protobuf:"varint,1,opt,name=total_video,json=totalVideo,proto3" json:"total_video,omitempty"`
	Played               int32          `protobuf:"varint,2,opt,name=played,proto3" json:"played,omitempty"`
	Data                 []*AdSet_Video `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AdSet) Reset()         { *m = AdSet{} }
func (m *AdSet) String() string { return proto.CompactTextString(m) }
func (*AdSet) ProtoMessage()    {}
func (*AdSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{9}
}
func (m *AdSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdSet.Unmarshal(m, b)
}
func (m *AdSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdSet.Marshal(b, m, deterministic)
}
func (m *AdSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdSet.Merge(m, src)
}
func (m *AdSet) XXX_Size() int {
	return xxx_messageInfo_AdSet.Size(m)
}
func (m *AdSet) XXX_DiscardUnknown() {
	xxx_messageInfo_AdSet.DiscardUnknown(m)
}

var xxx_messageInfo_AdSet proto.InternalMessageInfo

func (m *AdSet) GetTotalVideo() int32 {
	if m != nil {
		return m.TotalVideo
	}
	return 0
}

func (m *AdSet) GetPlayed() int32 {
	if m != nil {
		return m.Played
	}
	return 0
}

func (m *AdSet) GetData() []*AdSet_Video {
	if m != nil {
		return m.Data
	}
	return nil
}

type AdSet_Video struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Picture              string   `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdSet_Video) Reset()         { *m = AdSet_Video{} }
func (m *AdSet_Video) String() string { return proto.CompactTextString(m) }
func (*AdSet_Video) ProtoMessage()    {}
func (*AdSet_Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{9, 0}
}
func (m *AdSet_Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdSet_Video.Unmarshal(m, b)
}
func (m *AdSet_Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdSet_Video.Marshal(b, m, deterministic)
}
func (m *AdSet_Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdSet_Video.Merge(m, src)
}
func (m *AdSet_Video) XXX_Size() int {
	return xxx_messageInfo_AdSet_Video.Size(m)
}
func (m *AdSet_Video) XXX_DiscardUnknown() {
	xxx_messageInfo_AdSet_Video.DiscardUnknown(m)
}

var xxx_messageInfo_AdSet_Video proto.InternalMessageInfo

func (m *AdSet_Video) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AdSet_Video) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *AdSet_Video) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

type Campaign struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FromDate             int64          `protobuf:"varint,3,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate               int64          `protobuf:"varint,4,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
	Status               CampaignStatus `protobuf:"varint,5,opt,name=status,proto3,enum=pb.CampaignStatus" json:"status,omitempty"`
	AdSet                *AdSetBasic    `protobuf:"bytes,6,opt,name=ad_set,json=adSet,proto3" json:"ad_set,omitempty"`
	TotalLocation        int32          `protobuf:"varint,8,opt,name=total_location,json=totalLocation,proto3" json:"total_location,omitempty"`
	TotalScreen          int32          `protobuf:"varint,7,opt,name=total_screen,json=totalScreen,proto3" json:"total_screen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Campaign) Reset()         { *m = Campaign{} }
func (m *Campaign) String() string { return proto.CompactTextString(m) }
func (*Campaign) ProtoMessage()    {}
func (*Campaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{10}
}
func (m *Campaign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Campaign.Unmarshal(m, b)
}
func (m *Campaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Campaign.Marshal(b, m, deterministic)
}
func (m *Campaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Campaign.Merge(m, src)
}
func (m *Campaign) XXX_Size() int {
	return xxx_messageInfo_Campaign.Size(m)
}
func (m *Campaign) XXX_DiscardUnknown() {
	xxx_messageInfo_Campaign.DiscardUnknown(m)
}

var xxx_messageInfo_Campaign proto.InternalMessageInfo

func (m *Campaign) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Campaign) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Campaign) GetFromDate() int64 {
	if m != nil {
		return m.FromDate
	}
	return 0
}

func (m *Campaign) GetToDate() int64 {
	if m != nil {
		return m.ToDate
	}
	return 0
}

func (m *Campaign) GetStatus() CampaignStatus {
	if m != nil {
		return m.Status
	}
	return CampaignStatus_empty
}

func (m *Campaign) GetAdSet() *AdSetBasic {
	if m != nil {
		return m.AdSet
	}
	return nil
}

func (m *Campaign) GetTotalLocation() int32 {
	if m != nil {
		return m.TotalLocation
	}
	return 0
}

func (m *Campaign) GetTotalScreen() int32 {
	if m != nil {
		return m.TotalScreen
	}
	return 0
}

type CampaignRule struct {
	Locations            []int32          `protobuf:"varint,2,rep,packed,name=locations,proto3" json:"locations,omitempty"`
	Areas                []int32          `protobuf:"varint,3,rep,packed,name=areas,proto3" json:"areas,omitempty"`
	TimeRule             CampaignTimeRule `protobuf:"varint,4,opt,name=time_rule,json=timeRule,proto3,enum=pb.CampaignTimeRule" json:"time_rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CampaignRule) Reset()         { *m = CampaignRule{} }
func (m *CampaignRule) String() string { return proto.CompactTextString(m) }
func (*CampaignRule) ProtoMessage()    {}
func (*CampaignRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{11}
}
func (m *CampaignRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignRule.Unmarshal(m, b)
}
func (m *CampaignRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignRule.Marshal(b, m, deterministic)
}
func (m *CampaignRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignRule.Merge(m, src)
}
func (m *CampaignRule) XXX_Size() int {
	return xxx_messageInfo_CampaignRule.Size(m)
}
func (m *CampaignRule) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignRule.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignRule proto.InternalMessageInfo

func (m *CampaignRule) GetLocations() []int32 {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *CampaignRule) GetAreas() []int32 {
	if m != nil {
		return m.Areas
	}
	return nil
}

func (m *CampaignRule) GetTimeRule() CampaignTimeRule {
	if m != nil {
		return m.TimeRule
	}
	return CampaignTimeRule_free_in_this_time
}

type CampaignDetail struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AdSet                *AdSet         `protobuf:"bytes,3,opt,name=ad_set,json=adSet,proto3" json:"ad_set,omitempty"`
	FromDate             int64          `protobuf:"varint,4,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate               int64          `protobuf:"varint,5,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
	CampaignRule         *CampaignRule  `protobuf:"bytes,6,opt,name=campaign_rule,json=campaignRule,proto3" json:"campaign_rule,omitempty"`
	SelectedScreens      []int32        `protobuf:"varint,7,rep,packed,name=selected_screens,json=selectedScreens,proto3" json:"selected_screens,omitempty"`
	Status               CampaignStatus `protobuf:"varint,8,opt,name=status,proto3,enum=pb.CampaignStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CampaignDetail) Reset()         { *m = CampaignDetail{} }
func (m *CampaignDetail) String() string { return proto.CompactTextString(m) }
func (*CampaignDetail) ProtoMessage()    {}
func (*CampaignDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{12}
}
func (m *CampaignDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignDetail.Unmarshal(m, b)
}
func (m *CampaignDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignDetail.Marshal(b, m, deterministic)
}
func (m *CampaignDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignDetail.Merge(m, src)
}
func (m *CampaignDetail) XXX_Size() int {
	return xxx_messageInfo_CampaignDetail.Size(m)
}
func (m *CampaignDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignDetail.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignDetail proto.InternalMessageInfo

func (m *CampaignDetail) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CampaignDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CampaignDetail) GetAdSet() *AdSet {
	if m != nil {
		return m.AdSet
	}
	return nil
}

func (m *CampaignDetail) GetFromDate() int64 {
	if m != nil {
		return m.FromDate
	}
	return 0
}

func (m *CampaignDetail) GetToDate() int64 {
	if m != nil {
		return m.ToDate
	}
	return 0
}

func (m *CampaignDetail) GetCampaignRule() *CampaignRule {
	if m != nil {
		return m.CampaignRule
	}
	return nil
}

func (m *CampaignDetail) GetSelectedScreens() []int32 {
	if m != nil {
		return m.SelectedScreens
	}
	return nil
}

func (m *CampaignDetail) GetStatus() CampaignStatus {
	if m != nil {
		return m.Status
	}
	return CampaignStatus_empty
}

func init() {
	proto.RegisterEnum("pb.Status", Status_name, Status_value)
	proto.RegisterEnum("pb.Availability", Availability_name, Availability_value)
	proto.RegisterEnum("pb.CampaignStatus", CampaignStatus_name, CampaignStatus_value)
	proto.RegisterEnum("pb.CampaignTimeRule", CampaignTimeRule_name, CampaignTimeRule_value)
	proto.RegisterType((*EmptyMessage)(nil), "pb.EmptyMessage")
	proto.RegisterType((*UserProfile)(nil), "pb.UserProfile")
	proto.RegisterType((*Area)(nil), "pb.Area")
	proto.RegisterType((*ScreenActivity)(nil), "pb.ScreenActivity")
	proto.RegisterType((*ScreenActivities)(nil), "pb.ScreenActivities")
	proto.RegisterType((*ScreenInfo)(nil), "pb.ScreenInfo")
	proto.RegisterType((*Video)(nil), "pb.Video")
	proto.RegisterType((*Videos)(nil), "pb.Videos")
	proto.RegisterType((*AdSetBasic)(nil), "pb.AdSetBasic")
	proto.RegisterType((*AdSet)(nil), "pb.AdSet")
	proto.RegisterType((*AdSet_Video)(nil), "pb.AdSet.Video")
	proto.RegisterType((*Campaign)(nil), "pb.Campaign")
	proto.RegisterType((*CampaignRule)(nil), "pb.CampaignRule")
	proto.RegisterType((*CampaignDetail)(nil), "pb.CampaignDetail")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 1076 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xde, 0xd8, 0xb1, 0x63, 0x57, 0x32, 0x19, 0xd3, 0x5a, 0x16, 0xb3, 0x0b, 0x22, 0x78, 0x59,
	0x36, 0x3b, 0x87, 0x91, 0x18, 0xe0, 0x02, 0x17, 0xb2, 0x0c, 0x42, 0x23, 0xc1, 0x82, 0x3c, 0xb0,
	0x07, 0x24, 0x64, 0xf5, 0xd8, 0x9d, 0xd9, 0x96, 0xfc, 0x27, 0x77, 0x27, 0x62, 0x8e, 0xdc, 0x78,
	0x06, 0xde, 0x83, 0x67, 0xe1, 0xce, 0x93, 0xa0, 0xaa, 0x6e, 0x27, 0xce, 0xc0, 0x88, 0x15, 0xb7,
	0xae, 0xaf, 0xda, 0xd5, 0x5d, 0xdf, 0x57, 0x55, 0x6d, 0x98, 0xea, 0x9b, 0x56, 0xa8, 0xd3, 0xb6,
	0x6b, 0x74, 0xc3, 0x9c, 0xf6, 0x2a, 0x99, 0xc3, 0xec, 0xab, 0xaa, 0xd5, 0x37, 0xdf, 0x0a, 0xa5,
	0xf8, 0xb5, 0x48, 0xfe, 0x1c, 0xc1, 0xf4, 0x47, 0x25, 0xba, 0xef, 0xbb, 0x66, 0x2d, 0x4b, 0xc1,
	0xde, 0x82, 0xc9, 0x46, 0x89, 0x2e, 0x93, 0x45, 0x3c, 0x5a, 0x8c, 0x96, 0x61, 0xea, 0xa3, 0x79,
	0x51, 0x30, 0x06, 0xe3, 0x9a, 0x57, 0x22, 0x76, 0x08, 0xa5, 0x35, 0xbb, 0x0f, 0x9e, 0xa8, 0xb8,
	0x2c, 0x63, 0x97, 0x40, 0x63, 0x20, 0xda, 0xbe, 0x6a, 0x6a, 0x11, 0x8f, 0x0d, 0x4a, 0x06, 0x8b,
	0x61, 0x92, 0x37, 0x55, 0xcb, 0xeb, 0x9b, 0xd8, 0x23, 0xbc, 0x37, 0x31, 0x72, 0xd7, 0x94, 0x22,
	0xf6, 0x4d, 0x64, 0x5c, 0xb3, 0x04, 0x7c, 0xa5, 0xb9, 0xde, 0xa8, 0x78, 0xb2, 0x18, 0x2d, 0xe7,
	0x67, 0x70, 0xda, 0x5e, 0x9d, 0x5e, 0x12, 0x92, 0x5a, 0x0f, 0x7b, 0x17, 0x20, 0xef, 0x04, 0xd7,
	0xa2, 0xc8, 0xb8, 0x8e, 0x83, 0xc5, 0x68, 0xe9, 0xa6, 0xa1, 0x45, 0x56, 0x3a, 0xf9, 0x05, 0xc6,
	0xab, 0x4e, 0x70, 0xcc, 0x88, 0x77, 0x82, 0xf7, 0x19, 0x79, 0xa9, 0x8f, 0xe6, 0x1d, 0x19, 0xbd,
	0x07, 0xd3, 0xb2, 0xc9, 0xb9, 0x96, 0x4d, 0x8d, 0x1f, 0xb8, 0xf4, 0x01, 0xf4, 0xd0, 0x45, 0x31,
	0xb8, 0xd8, 0xf8, 0xae, 0x8b, 0x25, 0x5f, 0xc0, 0xfc, 0x32, 0xef, 0x84, 0xa8, 0x57, 0xb9, 0x96,
	0x5b, 0xa9, 0x29, 0x45, 0x2d, 0x2b, 0x41, 0x17, 0x70, 0x53, 0x5a, 0xb3, 0x87, 0x10, 0x70, 0xeb,
	0xb7, 0x57, 0xd8, 0xd9, 0xc9, 0x67, 0x10, 0x1d, 0x44, 0x90, 0x42, 0xb1, 0x0f, 0x61, 0x5c, 0x70,
	0xcd, 0xe3, 0xd1, 0xc2, 0x5d, 0x4e, 0xcf, 0x18, 0x9d, 0x7b, 0x70, 0x4a, 0x4a, 0xfe, 0xe4, 0x2f,
	0x17, 0xc0, 0x38, 0x2e, 0xea, 0x75, 0xc3, 0x1e, 0x41, 0xa8, 0xc8, 0xda, 0x13, 0x10, 0x18, 0xe0,
	0xa2, 0xc0, 0x74, 0xad, 0x73, 0xc0, 0x04, 0x18, 0xe8, 0xc5, 0x6b, 0xf1, 0xf1, 0x18, 0x8e, 0x76,
	0x1b, 0x28, 0x86, 0x11, 0x7d, 0xd6, 0x83, 0x14, 0xe5, 0x19, 0x44, 0xbb, 0x4d, 0xbc, 0x28, 0x3a,
	0xa1, 0x94, 0x2d, 0x82, 0xe3, 0x1e, 0x5f, 0x19, 0x78, 0xa8, 0x96, 0x7f, 0xa0, 0xd6, 0x23, 0x08,
	0xc9, 0x41, 0x87, 0x4c, 0x2c, 0x5f, 0x9d, 0xe0, 0x74, 0xc0, 0x5e, 0x95, 0xe0, 0xce, 0x72, 0xf9,
	0x04, 0x66, 0x7c, 0xcb, 0x65, 0xc9, 0xaf, 0x64, 0x89, 0x9c, 0x87, 0xb4, 0x33, 0xc2, 0x9d, 0xab,
	0x01, 0x9e, 0x1e, 0xec, 0x62, 0x1f, 0xc0, 0x1c, 0x5b, 0x28, 0x6b, 0xd6, 0x59, 0x21, 0xb6, 0x32,
	0x17, 0x31, 0x98, 0x04, 0x11, 0xfd, 0x6e, 0x7d, 0x4e, 0x18, 0xd2, 0x54, 0xf1, 0x7c, 0x97, 0xdb,
	0xd4, 0xf0, 0x58, 0xf1, 0xbc, 0x4f, 0x6b, 0x0e, 0x4e, 0xa3, 0xe2, 0x19, 0xe1, 0x4e, 0xa3, 0xf0,
	0x03, 0xde, 0xb6, 0xd9, 0x56, 0x74, 0x4a, 0x36, 0x75, 0x7c, 0x64, 0x3e, 0xe0, 0x6d, 0xfb, 0xd2,
	0x20, 0x58, 0xdc, 0xb2, 0xdd, 0x05, 0x9c, 0x93, 0x3f, 0x94, 0xad, 0x8d, 0x97, 0xfc, 0xe6, 0x80,
	0xf7, 0x52, 0x16, 0xa2, 0x61, 0x6f, 0x43, 0xb0, 0xc5, 0xc5, 0xbe, 0x63, 0x27, 0x64, 0x5f, 0x14,
	0xa6, 0xe5, 0x6a, 0x2d, 0x6a, 0x4d, 0xca, 0xce, 0xd2, 0xde, 0x44, 0x32, 0xb1, 0xdb, 0x0d, 0x99,
	0xa6, 0x79, 0x03, 0x04, 0x88, 0xcc, 0x07, 0xe0, 0xaf, 0x9b, 0xae, 0xe2, 0xda, 0x6a, 0x69, 0x2d,
	0x16, 0x81, 0xbb, 0xe9, 0x4a, 0x2b, 0x1c, 0x2e, 0xf1, 0x92, 0x9b, 0xb6, 0xe8, 0x3b, 0xd0, 0x37,
	0x1d, 0x68, 0x91, 0x95, 0xc6, 0x41, 0xa0, 0xa5, 0x2e, 0x7b, 0xb9, 0x8c, 0x81, 0x61, 0x34, 0xbf,
	0x26, 0xa1, 0xc2, 0x14, 0x97, 0x78, 0xcf, 0x56, 0xe6, 0x7a, 0xd3, 0x09, 0x12, 0x25, 0x4c, 0x7b,
	0x73, 0xa0, 0x2b, 0xdc, 0xd9, 0x6d, 0x4f, 0xc1, 0x27, 0x26, 0x70, 0x20, 0x0c, 0x3b, 0x24, 0xc4,
	0xbd, 0xe4, 0xb1, 0x8d, 0xf1, 0x33, 0xc0, 0xaa, 0xb8, 0x14, 0xfa, 0x39, 0x57, 0x32, 0x47, 0x05,
	0x74, 0xa3, 0x79, 0x99, 0x11, 0x5b, 0xb6, 0x33, 0x80, 0x20, 0x43, 0xec, 0x03, 0xf0, 0xdb, 0x92,
	0xdf, 0x88, 0x82, 0xc8, 0xf3, 0x52, 0x6b, 0x21, 0x4e, 0x9f, 0xa8, 0xd8, 0x5d, 0xb8, 0x48, 0x8f,
	0xb1, 0x92, 0x3f, 0x46, 0xe0, 0x51, 0xfc, 0xff, 0x1f, 0xfa, 0xb1, 0x4d, 0xc0, 0xa5, 0x04, 0x8e,
	0xa9, 0x34, 0x31, 0xe2, 0x30, 0x8d, 0x87, 0x2f, 0x7a, 0xe5, 0xe7, 0xe0, 0xec, 0x5a, 0xda, 0x91,
	0xc5, 0xa1, 0xa8, 0xce, 0x2d, 0x51, 0x07, 0x1c, 0xbb, 0x07, 0x1c, 0x27, 0xbf, 0x3a, 0x10, 0x7c,
	0xc9, 0xab, 0x96, 0xcb, 0xeb, 0xfa, 0x1f, 0x31, 0xff, 0x6d, 0x46, 0xe2, 0x39, 0x5d, 0x53, 0x65,
	0x28, 0x33, 0x05, 0x73, 0xd3, 0x00, 0x81, 0x73, 0xae, 0xe9, 0xfd, 0xd0, 0x8d, 0x71, 0x8d, 0xc9,
	0xe5, 0xeb, 0x86, 0x1c, 0x27, 0x3b, 0x29, 0x3d, 0x92, 0x92, 0x06, 0x58, 0x7f, 0xee, 0xad, 0x56,
	0x7d, 0x02, 0x3e, 0x2f, 0x32, 0x25, 0x4c, 0x4d, 0x4d, 0xcf, 0xe6, 0x3b, 0x26, 0x48, 0xbb, 0xd4,
	0xe3, 0xc4, 0xf3, 0x13, 0x98, 0x1b, 0x9e, 0xfb, 0x21, 0x42, 0x45, 0xe5, 0xa5, 0x47, 0x84, 0x7e,
	0x63, 0x41, 0xf6, 0x3e, 0xcc, 0xcc, 0x36, 0x33, 0xd7, 0xa8, 0x1a, 0xbd, 0xd4, 0x48, 0x64, 0x06,
	0x65, 0xb2, 0x81, 0x59, 0x7f, 0x95, 0x74, 0x53, 0x0a, 0xf6, 0x0e, 0x84, 0x7d, 0x4c, 0x15, 0x3b,
	0x0b, 0x77, 0xe9, 0xa5, 0x7b, 0x00, 0xeb, 0x1a, 0x27, 0x8f, 0x29, 0x00, 0x2f, 0x35, 0x06, 0xfb,
	0x08, 0x42, 0x9c, 0xeb, 0x59, 0xb7, 0x29, 0x85, 0x7d, 0x1c, 0xee, 0x0f, 0x73, 0xfc, 0x41, 0x56,
	0x02, 0x83, 0xa7, 0x81, 0xb6, 0xab, 0xe4, 0x77, 0x07, 0xe6, 0xbd, 0xfb, 0x5c, 0x68, 0x7c, 0x3c,
	0x5f, 0x47, 0x80, 0xc5, 0x8e, 0x1e, 0x97, 0xe8, 0x09, 0x77, 0xf4, 0xf4, 0xcc, 0x1c, 0x48, 0x34,
	0xbe, 0x5b, 0x22, 0xef, 0x40, 0xa2, 0x4f, 0xe1, 0x28, 0xb7, 0xb7, 0x31, 0x59, 0x18, 0xf6, 0xa3,
	0x61, 0x16, 0x94, 0xc1, 0x2c, 0x1f, 0x92, 0xf5, 0x0c, 0x22, 0x25, 0x4a, 0x91, 0xe3, 0x18, 0x30,
	0x14, 0xe3, 0xab, 0x8d, 0xcc, 0x1c, 0xf7, 0xb8, 0xa1, 0x59, 0x0d, 0x8a, 0x20, 0xf8, 0xaf, 0x22,
	0x38, 0x59, 0x82, 0x6f, 0x10, 0x76, 0x04, 0x21, 0xbd, 0x8c, 0x38, 0x56, 0xa2, 0x7b, 0xec, 0x18,
	0xa6, 0x85, 0xd8, 0x03, 0xa3, 0x93, 0xa7, 0x30, 0x1b, 0x4e, 0x70, 0x36, 0x85, 0x49, 0xb3, 0x5e,
	0x97, 0xb2, 0x16, 0xd1, 0x3d, 0x06, 0xe0, 0x37, 0x35, 0xad, 0x47, 0x27, 0x5f, 0xef, 0xe9, 0xb6,
	0xa1, 0x43, 0xfc, 0x83, 0x69, 0xf5, 0x8d, 0xd9, 0x98, 0xf3, 0x3a, 0x17, 0x65, 0x34, 0x42, 0xb8,
	0xe8, 0xf8, 0x5a, 0x47, 0x0e, 0x0b, 0x60, 0x5c, 0xca, 0xad, 0x88, 0x5c, 0x04, 0x5b, 0xbe, 0x51,
	0x22, 0x1a, 0x9f, 0x3c, 0x83, 0xe8, 0xb6, 0xac, 0xec, 0x4d, 0x78, 0x63, 0xdd, 0x09, 0x91, 0xc9,
	0x3a, 0xd3, 0xaf, 0xa4, 0xca, 0x50, 0xe5, 0xe8, 0xde, 0x73, 0xff, 0xa7, 0xf1, 0xe9, 0xe7, 0xed,
	0xd5, 0x95, 0x4f, 0xff, 0x60, 0x1f, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x26, 0x5c, 0xef,
	0x92, 0x09, 0x00, 0x00,
}
