// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.5.0
// source: esquema.proto

package pb

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

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Carnet int32  `protobuf:"varint,3,opt,name=carnet,proto3" json:"carnet,omitempty"`
	Course string `protobuf:"bytes,4,opt,name=course,proto3" json:"course,omitempty"`
	Body   string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Server string `protobuf:"bytes,6,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esquema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_esquema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_esquema_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Report) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Report) GetCarnet() int32 {
	if x != nil {
		return x.Carnet
	}
	return 0
}

func (x *Report) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Report) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Report) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

type Asistencia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Event  string `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Carnet int32  `protobuf:"varint,3,opt,name=carnet,proto3" json:"carnet,omitempty"`
	Img    string `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
	Id     int32  `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Server string `protobuf:"bytes,6,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *Asistencia) Reset() {
	*x = Asistencia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esquema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asistencia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asistencia) ProtoMessage() {}

func (x *Asistencia) ProtoReflect() protoreflect.Message {
	mi := &file_esquema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asistencia.ProtoReflect.Descriptor instead.
func (*Asistencia) Descriptor() ([]byte, []int) {
	return file_esquema_proto_rawDescGZIP(), []int{1}
}

func (x *Asistencia) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asistencia) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Asistencia) GetCarnet() int32 {
	if x != nil {
		return x.Carnet
	}
	return 0
}

func (x *Asistencia) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *Asistencia) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Asistencia) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

type ReportCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attended string    `protobuf:"bytes,1,opt,name=attended,proto3" json:"attended,omitempty"`
	Reports  []*Report `protobuf:"bytes,2,rep,name=reports,proto3" json:"reports,omitempty"`
}

func (x *ReportCollection) Reset() {
	*x = ReportCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esquema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportCollection) ProtoMessage() {}

func (x *ReportCollection) ProtoReflect() protoreflect.Message {
	mi := &file_esquema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportCollection.ProtoReflect.Descriptor instead.
func (*ReportCollection) Descriptor() ([]byte, []int) {
	return file_esquema_proto_rawDescGZIP(), []int{2}
}

func (x *ReportCollection) GetAttended() string {
	if x != nil {
		return x.Attended
	}
	return ""
}

func (x *ReportCollection) GetReports() []*Report {
	if x != nil {
		return x.Reports
	}
	return nil
}

type IndividualReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attended string  `protobuf:"bytes,1,opt,name=attended,proto3" json:"attended,omitempty"`
	Report   *Report `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *IndividualReport) Reset() {
	*x = IndividualReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esquema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndividualReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndividualReport) ProtoMessage() {}

func (x *IndividualReport) ProtoReflect() protoreflect.Message {
	mi := &file_esquema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndividualReport.ProtoReflect.Descriptor instead.
func (*IndividualReport) Descriptor() ([]byte, []int) {
	return file_esquema_proto_rawDescGZIP(), []int{3}
}

func (x *IndividualReport) GetAttended() string {
	if x != nil {
		return x.Attended
	}
	return ""
}

func (x *IndividualReport) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type AsistenciasCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attended    string        `protobuf:"bytes,1,opt,name=attended,proto3" json:"attended,omitempty"`
	Asistencias []*Asistencia `protobuf:"bytes,2,rep,name=asistencias,proto3" json:"asistencias,omitempty"`
}

func (x *AsistenciasCollection) Reset() {
	*x = AsistenciasCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esquema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsistenciasCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsistenciasCollection) ProtoMessage() {}

func (x *AsistenciasCollection) ProtoReflect() protoreflect.Message {
	mi := &file_esquema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsistenciasCollection.ProtoReflect.Descriptor instead.
func (*AsistenciasCollection) Descriptor() ([]byte, []int) {
	return file_esquema_proto_rawDescGZIP(), []int{4}
}

func (x *AsistenciasCollection) GetAttended() string {
	if x != nil {
		return x.Attended
	}
	return ""
}

func (x *AsistenciasCollection) GetAsistencias() []*Asistencia {
	if x != nil {
		return x.Asistencias
	}
	return nil
}

type IndividualAsistencia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attended   string      `protobuf:"bytes,1,opt,name=attended,proto3" json:"attended,omitempty"`
	Asistencia *Asistencia `protobuf:"bytes,2,opt,name=asistencia,proto3" json:"asistencia,omitempty"`
}

func (x *IndividualAsistencia) Reset() {
	*x = IndividualAsistencia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esquema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndividualAsistencia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndividualAsistencia) ProtoMessage() {}

func (x *IndividualAsistencia) ProtoReflect() protoreflect.Message {
	mi := &file_esquema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndividualAsistencia.ProtoReflect.Descriptor instead.
func (*IndividualAsistencia) Descriptor() ([]byte, []int) {
	return file_esquema_proto_rawDescGZIP(), []int{5}
}

func (x *IndividualAsistencia) GetAttended() string {
	if x != nil {
		return x.Attended
	}
	return ""
}

func (x *IndividualAsistencia) GetAsistencia() *Asistencia {
	if x != nil {
		return x.Asistencia
	}
	return nil
}

type Filtro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parametro int32 `protobuf:"varint,1,opt,name=parametro,proto3" json:"parametro,omitempty"`
}

func (x *Filtro) Reset() {
	*x = Filtro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esquema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filtro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filtro) ProtoMessage() {}

func (x *Filtro) ProtoReflect() protoreflect.Message {
	mi := &file_esquema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filtro.ProtoReflect.Descriptor instead.
func (*Filtro) Descriptor() ([]byte, []int) {
	return file_esquema_proto_rawDescGZIP(), []int{6}
}

func (x *Filtro) GetParametro() int32 {
	if x != nil {
		return x.Parametro
	}
	return 0
}

var File_esquema_proto protoreflect.FileDescriptor

var file_esquema_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x73, 0x71, 0x75, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x6e, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x63, 0x61, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x88,
	0x01, 0x0a, 0x0a, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x6e, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x61, 0x72, 0x6e, 0x65, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x10, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x07, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22,
	0x52, 0x0a, 0x10, 0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12,
	0x22, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x65, 0x0a, 0x15, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69,
	0x61, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x61, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x52, 0x0b, 0x61,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x73, 0x22, 0x62, 0x0a, 0x14, 0x49, 0x6e,
	0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x69, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x2e,
	0x0a, 0x0a, 0x61, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x69, 0x61, 0x52, 0x0a, 0x61, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x22, 0x26,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x32, 0xbf, 0x01, 0x0a, 0x0e, 0x41, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x12, 0x33, 0x0a, 0x05, 0x43, 0x72, 0x65,
	0x61, 0x72, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x69, 0x61, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75,
	0x61, 0x6c, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0e, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x5f, 0x49, 0x44,
	0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x6f, 0x1a, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x73, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x12, 0x4f, 0x62, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x5f, 0x43, 0x41, 0x52, 0x4e, 0x45, 0x54, 0x12,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x6f, 0x1a, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x73, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x32, 0xaa, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x12, 0x2b, 0x0a, 0x05, 0x43, 0x72, 0x65, 0x61,
	0x72, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x5f, 0x70, 0x6f, 0x72, 0x5f, 0x49, 0x44, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x72, 0x6f, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x12, 0x4f,
	0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x5f, 0x43, 0x41, 0x52, 0x4e, 0x45,
	0x54, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x72, 0x6f, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6e, 0x61, 0x74, 0x6f, 0x37, 0x37, 0x36, 0x2f, 0x70, 0x72,
	0x6f, 0x79, 0x65, 0x63, 0x74, 0x6f, 0x2d, 0x67, 0x31, 0x36, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_esquema_proto_rawDescOnce sync.Once
	file_esquema_proto_rawDescData = file_esquema_proto_rawDesc
)

func file_esquema_proto_rawDescGZIP() []byte {
	file_esquema_proto_rawDescOnce.Do(func() {
		file_esquema_proto_rawDescData = protoimpl.X.CompressGZIP(file_esquema_proto_rawDescData)
	})
	return file_esquema_proto_rawDescData
}

var file_esquema_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_esquema_proto_goTypes = []interface{}{
	(*Report)(nil),                // 0: pb.Report
	(*Asistencia)(nil),            // 1: pb.Asistencia
	(*ReportCollection)(nil),      // 2: pb.ReportCollection
	(*IndividualReport)(nil),      // 3: pb.IndividualReport
	(*AsistenciasCollection)(nil), // 4: pb.AsistenciasCollection
	(*IndividualAsistencia)(nil),  // 5: pb.IndividualAsistencia
	(*Filtro)(nil),                // 6: pb.Filtro
}
var file_esquema_proto_depIdxs = []int32{
	0,  // 0: pb.ReportCollection.reports:type_name -> pb.Report
	0,  // 1: pb.IndividualReport.report:type_name -> pb.Report
	1,  // 2: pb.AsistenciasCollection.asistencias:type_name -> pb.Asistencia
	1,  // 3: pb.IndividualAsistencia.asistencia:type_name -> pb.Asistencia
	1,  // 4: pb.AsistenciaServ.Crear:input_type -> pb.Asistencia
	6,  // 5: pb.AsistenciaServ.Obtener_por_ID:input_type -> pb.Filtro
	6,  // 6: pb.AsistenciaServ.Obtener_por_CARNET:input_type -> pb.Filtro
	0,  // 7: pb.ReporteServ.Crear:input_type -> pb.Report
	6,  // 8: pb.ReporteServ.Obtener_por_ID:input_type -> pb.Filtro
	6,  // 9: pb.ReporteServ.Obtener_por_CARNET:input_type -> pb.Filtro
	5,  // 10: pb.AsistenciaServ.Crear:output_type -> pb.IndividualAsistencia
	4,  // 11: pb.AsistenciaServ.Obtener_por_ID:output_type -> pb.AsistenciasCollection
	4,  // 12: pb.AsistenciaServ.Obtener_por_CARNET:output_type -> pb.AsistenciasCollection
	3,  // 13: pb.ReporteServ.Crear:output_type -> pb.IndividualReport
	2,  // 14: pb.ReporteServ.Obtener_por_ID:output_type -> pb.ReportCollection
	2,  // 15: pb.ReporteServ.Obtener_por_CARNET:output_type -> pb.ReportCollection
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_esquema_proto_init() }
func file_esquema_proto_init() {
	if File_esquema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_esquema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_esquema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asistencia); i {
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
		file_esquema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportCollection); i {
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
		file_esquema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndividualReport); i {
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
		file_esquema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsistenciasCollection); i {
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
		file_esquema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndividualAsistencia); i {
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
		file_esquema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filtro); i {
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
			RawDescriptor: file_esquema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_esquema_proto_goTypes,
		DependencyIndexes: file_esquema_proto_depIdxs,
		MessageInfos:      file_esquema_proto_msgTypes,
	}.Build()
	File_esquema_proto = out.File
	file_esquema_proto_rawDesc = nil
	file_esquema_proto_goTypes = nil
	file_esquema_proto_depIdxs = nil
}
