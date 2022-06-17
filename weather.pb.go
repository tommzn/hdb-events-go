// HomeDashboard Weather Event Schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: weather.proto

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

// Weather data contains current weather and a 7-days forecast for a given location.
type WeatherData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Location a weather information belongs to
	Location *Location `protobuf:"bytes,1,opt,name=Location,proto3" json:"Location,omitempty"`
	// Defines units for temperature and wind speed. Possible values are: default, metric, imperial
	Units string `protobuf:"bytes,2,opt,name=Units,proto3" json:"Units,omitempty"`
	// Current weather data
	Current *CurrentWeather `protobuf:"bytes,3,opt,name=Current,proto3" json:"Current,omitempty"`
	// 7-days forecast
	Forecast []*ForecastWeather `protobuf:"bytes,4,rep,name=Forecast,proto3" json:"Forecast,omitempty"`
}

func (x *WeatherData) Reset() {
	*x = WeatherData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData) ProtoMessage() {}

func (x *WeatherData) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData.ProtoReflect.Descriptor instead.
func (*WeatherData) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherData) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *WeatherData) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

func (x *WeatherData) GetCurrent() *CurrentWeather {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *WeatherData) GetForecast() []*ForecastWeather {
	if x != nil {
		return x.Forecast
	}
	return nil
}

// Current weather data.
type CurrentWeather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp a weather infirmation is related to.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	// Temperature. Units - default: kelvin, metric: Celsius, imperial: Fahrenheit
	Temperature float64 `protobuf:"fixed64,2,opt,name=Temperature,proto3" json:"Temperature,omitempty"`
	// Wind speed. Wind speed. Units – default: metre/sec, metric: metre/sec, imperial: miles/hour
	WindSpeed float64 `protobuf:"fixed64,3,opt,name=WindSpeed,proto3" json:"WindSpeed,omitempty"`
	// Weather details
	Weather *WeatherDetails `protobuf:"bytes,4,opt,name=Weather,proto3" json:"Weather,omitempty"`
	// Wind direction, degrees (meteorological)
	WindDirection int64 `protobuf:"varint,5,opt,name=WindDirection,proto3" json:"WindDirection,omitempty"`
	// Wind gust. Wind gust. Units – default: metre/sec, metric: metre/sec, imperial: miles/hour.
	WindGust float64 `protobuf:"fixed64,6,opt,name=WindGust,proto3" json:"WindGust,omitempty"`
}

func (x *CurrentWeather) Reset() {
	*x = CurrentWeather{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentWeather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentWeather) ProtoMessage() {}

func (x *CurrentWeather) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentWeather.ProtoReflect.Descriptor instead.
func (*CurrentWeather) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentWeather) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CurrentWeather) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *CurrentWeather) GetWindSpeed() float64 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

func (x *CurrentWeather) GetWeather() *WeatherDetails {
	if x != nil {
		return x.Weather
	}
	return nil
}

func (x *CurrentWeather) GetWindDirection() int64 {
	if x != nil {
		return x.WindDirection
	}
	return 0
}

func (x *CurrentWeather) GetWindGust() float64 {
	if x != nil {
		return x.WindGust
	}
	return 0
}

// ForeCastWeather contains forecast weather data for a single day
type ForecastWeather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp a weather infirmation is related to.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	// Forecast temperatures for the whole day
	Temperatures *ForecastTemperatures `protobuf:"bytes,2,opt,name=Temperatures,proto3" json:"Temperatures,omitempty"`
	// Wind speed. Wind speed. Units – default: metre/sec, metric: metre/sec, imperial: miles/hour
	WindSpeed float64 `protobuf:"fixed64,3,opt,name=WindSpeed,proto3" json:"WindSpeed,omitempty"`
	// Weather details
	Weather *WeatherDetails `protobuf:"bytes,4,opt,name=Weather,proto3" json:"Weather,omitempty"`
}

func (x *ForecastWeather) Reset() {
	*x = ForecastWeather{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastWeather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastWeather) ProtoMessage() {}

func (x *ForecastWeather) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastWeather.ProtoReflect.Descriptor instead.
func (*ForecastWeather) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{2}
}

func (x *ForecastWeather) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ForecastWeather) GetTemperatures() *ForecastTemperatures {
	if x != nil {
		return x.Temperatures
	}
	return nil
}

func (x *ForecastWeather) GetWindSpeed() float64 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

func (x *ForecastWeather) GetWeather() *WeatherDetails {
	if x != nil {
		return x.Weather
	}
	return nil
}

// WeatherDetails contains information of current weather or forecast
// Full list of condition id, group and description and Icons is available at: https://openweathermap.org/weather-conditions#Weather-Condition-Codes-2
type WeatherDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Weather condition id
	ConditionId int64 `protobuf:"varint,1,opt,name=ConditionId,proto3" json:"ConditionId,omitempty"`
	// Group of weather parameters (Rain, Snow, Extreme etc.)
	Group string `protobuf:"bytes,2,opt,name=Group,proto3" json:"Group,omitempty"`
	// Weather condition within the group
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	// Weather icon from Open Weather Map
	Icon string `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
}

func (x *WeatherDetails) Reset() {
	*x = WeatherDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherDetails) ProtoMessage() {}

func (x *WeatherDetails) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherDetails.ProtoReflect.Descriptor instead.
func (*WeatherDetails) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{3}
}

func (x *WeatherDetails) GetConditionId() int64 {
	if x != nil {
		return x.ConditionId
	}
	return 0
}

func (x *WeatherDetails) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *WeatherDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WeatherDetails) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

// forecastTemperature contains forecast temperature data for a single day
// Units – default: kelvin, metric: Celsius, imperial: Fahrenheit
type ForecastTemperatures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Morning temperature
	Morning float64 `protobuf:"fixed64,1,opt,name=Morning,proto3" json:"Morning,omitempty"`
	// Day temperature
	Day float64 `protobuf:"fixed64,2,opt,name=Day,proto3" json:"Day,omitempty"`
	// Evening temperature
	Evening float64 `protobuf:"fixed64,3,opt,name=Evening,proto3" json:"Evening,omitempty"`
	// Night temperature
	Night float64 `protobuf:"fixed64,4,opt,name=Night,proto3" json:"Night,omitempty"`
	// Min daily temperature.
	DayMin float64 `protobuf:"fixed64,5,opt,name=DayMin,proto3" json:"DayMin,omitempty"`
	// Max daily temperature.
	DayMax float64 `protobuf:"fixed64,6,opt,name=DayMax,proto3" json:"DayMax,omitempty"`
}

func (x *ForecastTemperatures) Reset() {
	*x = ForecastTemperatures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForecastTemperatures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForecastTemperatures) ProtoMessage() {}

func (x *ForecastTemperatures) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForecastTemperatures.ProtoReflect.Descriptor instead.
func (*ForecastTemperatures) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{4}
}

func (x *ForecastTemperatures) GetMorning() float64 {
	if x != nil {
		return x.Morning
	}
	return 0
}

func (x *ForecastTemperatures) GetDay() float64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *ForecastTemperatures) GetEvening() float64 {
	if x != nil {
		return x.Evening
	}
	return 0
}

func (x *ForecastTemperatures) GetNight() float64 {
	if x != nil {
		return x.Night
	}
	return 0
}

func (x *ForecastTemperatures) GetDayMin() float64 {
	if x != nil {
		return x.DayMin
	}
	return 0
}

func (x *ForecastTemperatures) GetDayMax() float64 {
	if x != nil {
		return x.DayMax
	}
	return 0
}

// Geographical coordinates weather information are related to
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Longitude coordinate.
	Longitude float64 `protobuf:"fixed64,1,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	// Latitude coordinate.
	Latitude float64 `protobuf:"fixed64,2,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{5}
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

var File_weather_proto protoreflect.FileDescriptor

var file_weather_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0b, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x07,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x33,
	0x0a, 0x08, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x08, 0x46, 0x6f, 0x72, 0x65, 0x63,
	0x61, 0x73, 0x74, 0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x30, 0x0a, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x57, 0x69, 0x6e, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x57, 0x69, 0x6e, 0x64, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x69, 0x6e, 0x64,
	0x47, 0x75, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x57, 0x69, 0x6e, 0x64,
	0x47, 0x75, 0x73, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x40, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x22, 0x7e, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x49, 0x63, 0x6f, 0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x14, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x6f, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x4d, 0x6f, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x44, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x76, 0x65,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x45, 0x76, 0x65, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x4e, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x61, 0x79,
	0x4d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x44, 0x61, 0x79, 0x4d, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x61, 0x79, 0x4d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x44, 0x61, 0x79, 0x4d, 0x61, 0x78, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_weather_proto_rawDescOnce sync.Once
	file_weather_proto_rawDescData = file_weather_proto_rawDesc
)

func file_weather_proto_rawDescGZIP() []byte {
	file_weather_proto_rawDescOnce.Do(func() {
		file_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_proto_rawDescData)
	})
	return file_weather_proto_rawDescData
}

var file_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_weather_proto_goTypes = []interface{}{
	(*WeatherData)(nil),          // 0: events.WeatherData
	(*CurrentWeather)(nil),       // 1: events.CurrentWeather
	(*ForecastWeather)(nil),      // 2: events.ForecastWeather
	(*WeatherDetails)(nil),       // 3: events.WeatherDetails
	(*ForecastTemperatures)(nil), // 4: events.ForecastTemperatures
	(*Location)(nil),             // 5: events.Location
	(*timestamp.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_weather_proto_depIdxs = []int32{
	5, // 0: events.WeatherData.Location:type_name -> events.Location
	1, // 1: events.WeatherData.Current:type_name -> events.CurrentWeather
	2, // 2: events.WeatherData.Forecast:type_name -> events.ForecastWeather
	6, // 3: events.CurrentWeather.Timestamp:type_name -> google.protobuf.Timestamp
	3, // 4: events.CurrentWeather.Weather:type_name -> events.WeatherDetails
	6, // 5: events.ForecastWeather.Timestamp:type_name -> google.protobuf.Timestamp
	4, // 6: events.ForecastWeather.Temperatures:type_name -> events.ForecastTemperatures
	3, // 7: events.ForecastWeather.Weather:type_name -> events.WeatherDetails
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_weather_proto_init() }
func file_weather_proto_init() {
	if File_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherData); i {
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
		file_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentWeather); i {
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
		file_weather_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastWeather); i {
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
		file_weather_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherDetails); i {
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
		file_weather_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForecastTemperatures); i {
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
		file_weather_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
			RawDescriptor: file_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_weather_proto_goTypes,
		DependencyIndexes: file_weather_proto_depIdxs,
		MessageInfos:      file_weather_proto_msgTypes,
	}.Build()
	File_weather_proto = out.File
	file_weather_proto_rawDesc = nil
	file_weather_proto_goTypes = nil
	file_weather_proto_depIdxs = nil
}
