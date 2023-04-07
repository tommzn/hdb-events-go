// HomeDashboard Strava Athlete Stats Event Schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: strava.proto

package events

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// AthleteStats contains stats, current year and total, together with latest activities of an athlete.
type AthleteStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ActivityStats contains current year and total stats for an athlete.
	ActivityStats *ActivityStats `protobuf:"bytes,1,opt,name=ActivityStats,proto3" json:"ActivityStats,omitempty"`
	// List of latest activities of an athlete.
	Activities []*Activity `protobuf:"bytes,2,rep,name=Activities,proto3" json:"Activities,omitempty"`
	// Timestamp,  point in time this stats has been retrieved from Strava.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *AthleteStats) Reset() {
	*x = AthleteStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strava_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AthleteStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AthleteStats) ProtoMessage() {}

func (x *AthleteStats) ProtoReflect() protoreflect.Message {
	mi := &file_strava_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AthleteStats.ProtoReflect.Descriptor instead.
func (*AthleteStats) Descriptor() ([]byte, []int) {
	return file_strava_proto_rawDescGZIP(), []int{0}
}

func (x *AthleteStats) GetActivityStats() *ActivityStats {
	if x != nil {
		return x.ActivityStats
	}
	return nil
}

func (x *AthleteStats) GetActivities() []*Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

func (x *AthleteStats) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// ActivityTotal contains current year and total stats for an athlete.
type ActivityStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The longest distance ridden by the athlete.
	BiggestRideDistance float64 `protobuf:"fixed64,1,opt,name=BiggestRideDistance,proto3" json:"BiggestRideDistance,omitempty"`
	// The recent (last 4 weeks) ride stats for the athlete.
	RecentRideTotals *ActivityTotal `protobuf:"bytes,2,opt,name=RecentRideTotals,proto3" json:"RecentRideTotals,omitempty"`
	// The recent (last 4 weeks) run stats for the athlete.
	RecentRunTotals *ActivityTotal `protobuf:"bytes,3,opt,name=RecentRunTotals,proto3" json:"RecentRunTotals,omitempty"`
	// The year to date ride stats for the athlete.
	YearToDateRideTotals *ActivityTotal `protobuf:"bytes,4,opt,name=YearToDateRideTotals,proto3" json:"YearToDateRideTotals,omitempty"`
	// The year to date run stats for the athlete.
	YearToDateRunTotals *ActivityTotal `protobuf:"bytes,5,opt,name=YearToDateRunTotals,proto3" json:"YearToDateRunTotals,omitempty"`
	// The all time ride stats for the aTthlete.
	AllRideotals *ActivityTotal `protobuf:"bytes,6,opt,name=AllRideotals,proto3" json:"AllRideotals,omitempty"`
	// The all time run stats for the athlete.
	AllRunTotals *ActivityTotal `protobuf:"bytes,7,opt,name=AllRunTotals,proto3" json:"AllRunTotals,omitempty"`
}

func (x *ActivityStats) Reset() {
	*x = ActivityStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strava_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityStats) ProtoMessage() {}

func (x *ActivityStats) ProtoReflect() protoreflect.Message {
	mi := &file_strava_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityStats.ProtoReflect.Descriptor instead.
func (*ActivityStats) Descriptor() ([]byte, []int) {
	return file_strava_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityStats) GetBiggestRideDistance() float64 {
	if x != nil {
		return x.BiggestRideDistance
	}
	return 0
}

func (x *ActivityStats) GetRecentRideTotals() *ActivityTotal {
	if x != nil {
		return x.RecentRideTotals
	}
	return nil
}

func (x *ActivityStats) GetRecentRunTotals() *ActivityTotal {
	if x != nil {
		return x.RecentRunTotals
	}
	return nil
}

func (x *ActivityStats) GetYearToDateRideTotals() *ActivityTotal {
	if x != nil {
		return x.YearToDateRideTotals
	}
	return nil
}

func (x *ActivityStats) GetYearToDateRunTotals() *ActivityTotal {
	if x != nil {
		return x.YearToDateRunTotals
	}
	return nil
}

func (x *ActivityStats) GetAllRideotals() *ActivityTotal {
	if x != nil {
		return x.AllRideotals
	}
	return nil
}

func (x *ActivityStats) GetAllRunTotals() *ActivityTotal {
	if x != nil {
		return x.AllRunTotals
	}
	return nil
}

type ActivityTotal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of activities considered in this total.
	Count int64 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	// The total distance covered by the considered activities.
	Distance float64 `protobuf:"fixed64,2,opt,name=Distance,proto3" json:"Distance,omitempty"`
	// The total moving time of the considered activities in seconds.
	MovingTime int64 `protobuf:"varint,3,opt,name=MovingTime,proto3" json:"MovingTime,omitempty"`
}

func (x *ActivityTotal) Reset() {
	*x = ActivityTotal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strava_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityTotal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityTotal) ProtoMessage() {}

func (x *ActivityTotal) ProtoReflect() protoreflect.Message {
	mi := &file_strava_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityTotal.ProtoReflect.Descriptor instead.
func (*ActivityTotal) Descriptor() ([]byte, []int) {
	return file_strava_proto_rawDescGZIP(), []int{2}
}

func (x *ActivityTotal) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ActivityTotal) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *ActivityTotal) GetMovingTime() int64 {
	if x != nil {
		return x.MovingTime
	}
	return 0
}

// Description of a single activity.
type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the activity.
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// The activity's distance.
	Distance float64 `protobuf:"fixed64,2,opt,name=Distance,proto3" json:"Distance,omitempty"`
	// The activity's moving time.
	MovingTime int64 `protobuf:"varint,3,opt,name=MovingTime,proto3" json:"MovingTime,omitempty"`
	// Type of an activity.
	SportType string `protobuf:"bytes,4,opt,name=SportType,proto3" json:"SportType,omitempty"`
	// Timestamp an activity is related to.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strava_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_strava_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_strava_proto_rawDescGZIP(), []int{3}
}

func (x *Activity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Activity) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Activity) GetMovingTime() int64 {
	if x != nil {
		return x.MovingTime
	}
	return 0
}

func (x *Activity) GetSportType() string {
	if x != nil {
		return x.SportType
	}
	return ""
}

func (x *Activity) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_strava_proto protoreflect.FileDescriptor

var file_strava_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x41, 0x74, 0x68, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0xcf, 0x03, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x42, 0x69, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52, 0x69,
	0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x13, 0x42, 0x69, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52, 0x69, 0x64, 0x65, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x52,
	0x69, 0x64, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x10, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x69,
	0x64, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x52, 0x75, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x49, 0x0a, 0x14, 0x59, 0x65, 0x61,
	0x72, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x14,
	0x59, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x73, 0x12, 0x47, 0x0a, 0x13, 0x59, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x75, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x13, 0x59, 0x65, 0x61, 0x72, 0x54, 0x6f,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x39, 0x0a,
	0x0c, 0x41, 0x6c, 0x6c, 0x52, 0x69, 0x64, 0x65, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x0c, 0x41, 0x6c, 0x6c, 0x52,
	0x69, 0x64, 0x65, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x52,
	0x75, 0x6e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x0c, 0x41, 0x6c, 0x6c, 0x52, 0x75, 0x6e, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x73, 0x22, 0x61, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x44, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x6f, 0x76, 0x69,
	0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strava_proto_rawDescOnce sync.Once
	file_strava_proto_rawDescData = file_strava_proto_rawDesc
)

func file_strava_proto_rawDescGZIP() []byte {
	file_strava_proto_rawDescOnce.Do(func() {
		file_strava_proto_rawDescData = protoimpl.X.CompressGZIP(file_strava_proto_rawDescData)
	})
	return file_strava_proto_rawDescData
}

var file_strava_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_strava_proto_goTypes = []interface{}{
	(*AthleteStats)(nil),        // 0: events.AthleteStats
	(*ActivityStats)(nil),       // 1: events.ActivityStats
	(*ActivityTotal)(nil),       // 2: events.ActivityTotal
	(*Activity)(nil),            // 3: events.Activity
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_strava_proto_depIdxs = []int32{
	1,  // 0: events.AthleteStats.ActivityStats:type_name -> events.ActivityStats
	3,  // 1: events.AthleteStats.Activities:type_name -> events.Activity
	4,  // 2: events.AthleteStats.Timestamp:type_name -> google.protobuf.Timestamp
	2,  // 3: events.ActivityStats.RecentRideTotals:type_name -> events.ActivityTotal
	2,  // 4: events.ActivityStats.RecentRunTotals:type_name -> events.ActivityTotal
	2,  // 5: events.ActivityStats.YearToDateRideTotals:type_name -> events.ActivityTotal
	2,  // 6: events.ActivityStats.YearToDateRunTotals:type_name -> events.ActivityTotal
	2,  // 7: events.ActivityStats.AllRideotals:type_name -> events.ActivityTotal
	2,  // 8: events.ActivityStats.AllRunTotals:type_name -> events.ActivityTotal
	4,  // 9: events.Activity.Timestamp:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_strava_proto_init() }
func file_strava_proto_init() {
	if File_strava_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strava_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AthleteStats); i {
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
		file_strava_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityStats); i {
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
		file_strava_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityTotal); i {
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
		file_strava_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
			RawDescriptor: file_strava_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_strava_proto_goTypes,
		DependencyIndexes: file_strava_proto_depIdxs,
		MessageInfos:      file_strava_proto_msgTypes,
	}.Build()
	File_strava_proto = out.File
	file_strava_proto_rawDesc = nil
	file_strava_proto_goTypes = nil
	file_strava_proto_depIdxs = nil
}
