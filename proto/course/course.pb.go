// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course/course.proto

/*
Package course is a generated protocol buffer package.

It is generated from these files:
	proto/course/course.proto

It has these top-level messages:
	UpsertCourseReq
	UpsertCourseRsp
	InsertCoursesRsp
	DeleteCoursesRsp
	IDs
	UpdateCoursesRsp
	FindCoursesRsp
	Courses
	FindCoursesReq
	CourseAndCategories
	SyncCategoriesReq
	AddCategoriesReq
	DeleteCategoriesReq
	PurgeCategoriesReq
	Course
*/
package course

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/empty"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpsertCourseReq struct {
	Course *Course `protobuf:"bytes,1,opt,name=course" json:"course,omitempty"`
	IsNew  bool    `protobuf:"varint,2,opt,name=is_new,json=isNew" json:"is_new,omitempty"`
}

func (m *UpsertCourseReq) Reset()                    { *m = UpsertCourseReq{} }
func (m *UpsertCourseReq) String() string            { return proto.CompactTextString(m) }
func (*UpsertCourseReq) ProtoMessage()               {}
func (*UpsertCourseReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpsertCourseReq) GetCourse() *Course {
	if m != nil {
		return m.Course
	}
	return nil
}

func (m *UpsertCourseReq) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

type UpsertCourseRsp struct {
	Course   *Course `protobuf:"bytes,1,opt,name=course" json:"course,omitempty"`
	Inserted int64   `protobuf:"varint,2,opt,name=inserted" json:"inserted,omitempty"`
	Updated  int64   `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpsertCourseRsp) Reset()                    { *m = UpsertCourseRsp{} }
func (m *UpsertCourseRsp) String() string            { return proto.CompactTextString(m) }
func (*UpsertCourseRsp) ProtoMessage()               {}
func (*UpsertCourseRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpsertCourseRsp) GetCourse() *Course {
	if m != nil {
		return m.Course
	}
	return nil
}

func (m *UpsertCourseRsp) GetInserted() int64 {
	if m != nil {
		return m.Inserted
	}
	return 0
}

func (m *UpsertCourseRsp) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type InsertCoursesRsp struct {
	Inserted int64 `protobuf:"varint,1,opt,name=inserted" json:"inserted,omitempty"`
}

func (m *InsertCoursesRsp) Reset()                    { *m = InsertCoursesRsp{} }
func (m *InsertCoursesRsp) String() string            { return proto.CompactTextString(m) }
func (*InsertCoursesRsp) ProtoMessage()               {}
func (*InsertCoursesRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InsertCoursesRsp) GetInserted() int64 {
	if m != nil {
		return m.Inserted
	}
	return 0
}

type DeleteCoursesRsp struct {
	Deleted int64 `protobuf:"varint,1,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteCoursesRsp) Reset()                    { *m = DeleteCoursesRsp{} }
func (m *DeleteCoursesRsp) String() string            { return proto.CompactTextString(m) }
func (*DeleteCoursesRsp) ProtoMessage()               {}
func (*DeleteCoursesRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteCoursesRsp) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

type IDs struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *IDs) Reset()                    { *m = IDs{} }
func (m *IDs) String() string            { return proto.CompactTextString(m) }
func (*IDs) ProtoMessage()               {}
func (*IDs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IDs) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type UpdateCoursesRsp struct {
	Updated int64 `protobuf:"varint,1,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateCoursesRsp) Reset()                    { *m = UpdateCoursesRsp{} }
func (m *UpdateCoursesRsp) String() string            { return proto.CompactTextString(m) }
func (*UpdateCoursesRsp) ProtoMessage()               {}
func (*UpdateCoursesRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateCoursesRsp) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type FindCoursesRsp struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (m *FindCoursesRsp) Reset()                    { *m = FindCoursesRsp{} }
func (m *FindCoursesRsp) String() string            { return proto.CompactTextString(m) }
func (*FindCoursesRsp) ProtoMessage()               {}
func (*FindCoursesRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FindCoursesRsp) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type Courses struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (m *Courses) Reset()                    { *m = Courses{} }
func (m *Courses) String() string            { return proto.CompactTextString(m) }
func (*Courses) ProtoMessage()               {}
func (*Courses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Courses) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type FindCoursesReq struct {
	Ids         []string                   `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Query       string                     `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Start       *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=start" json:"start,omitempty"`
	End         *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
	Inclusive   bool                       `protobuf:"varint,5,opt,name=inclusive" json:"inclusive,omitempty"`
	Hidden      bool                       `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
	Sort        []string                   `protobuf:"bytes,7,rep,name=sort" json:"sort,omitempty"`
	PerPage     int64                      `protobuf:"varint,8,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
	CurrentPage int64                      `protobuf:"varint,9,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
}

func (m *FindCoursesReq) Reset()                    { *m = FindCoursesReq{} }
func (m *FindCoursesReq) String() string            { return proto.CompactTextString(m) }
func (*FindCoursesReq) ProtoMessage()               {}
func (*FindCoursesReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FindCoursesReq) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FindCoursesReq) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *FindCoursesReq) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *FindCoursesReq) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *FindCoursesReq) GetInclusive() bool {
	if m != nil {
		return m.Inclusive
	}
	return false
}

func (m *FindCoursesReq) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *FindCoursesReq) GetSort() []string {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *FindCoursesReq) GetPerPage() int64 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *FindCoursesReq) GetCurrentPage() int64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

type CourseAndCategories struct {
	// @inject_tag: valid:"uuid,required~ID is required"
	CourseId    string   `protobuf:"bytes,1,opt,name=course_id,json=courseId" json:"course_id,omitempty" valid:"uuid,required~ID is required"`
	CategoryIds []string `protobuf:"bytes,2,rep,name=category_ids,json=categoryIds" json:"category_ids,omitempty"`
}

func (m *CourseAndCategories) Reset()                    { *m = CourseAndCategories{} }
func (m *CourseAndCategories) String() string            { return proto.CompactTextString(m) }
func (*CourseAndCategories) ProtoMessage()               {}
func (*CourseAndCategories) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CourseAndCategories) GetCourseId() string {
	if m != nil {
		return m.CourseId
	}
	return ""
}

func (m *CourseAndCategories) GetCategoryIds() []string {
	if m != nil {
		return m.CategoryIds
	}
	return nil
}

type SyncCategoriesReq struct {
	CourseAndCategories []*CourseAndCategories `protobuf:"bytes,1,rep,name=course_and_categories,json=courseAndCategories" json:"course_and_categories,omitempty"`
}

func (m *SyncCategoriesReq) Reset()                    { *m = SyncCategoriesReq{} }
func (m *SyncCategoriesReq) String() string            { return proto.CompactTextString(m) }
func (*SyncCategoriesReq) ProtoMessage()               {}
func (*SyncCategoriesReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SyncCategoriesReq) GetCourseAndCategories() []*CourseAndCategories {
	if m != nil {
		return m.CourseAndCategories
	}
	return nil
}

type AddCategoriesReq struct {
	CourseAndCategories []*CourseAndCategories `protobuf:"bytes,1,rep,name=course_and_categories,json=courseAndCategories" json:"course_and_categories,omitempty"`
}

func (m *AddCategoriesReq) Reset()                    { *m = AddCategoriesReq{} }
func (m *AddCategoriesReq) String() string            { return proto.CompactTextString(m) }
func (*AddCategoriesReq) ProtoMessage()               {}
func (*AddCategoriesReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AddCategoriesReq) GetCourseAndCategories() []*CourseAndCategories {
	if m != nil {
		return m.CourseAndCategories
	}
	return nil
}

type DeleteCategoriesReq struct {
	CourseAndCategories []*CourseAndCategories `protobuf:"bytes,1,rep,name=course_and_categories,json=courseAndCategories" json:"course_and_categories,omitempty"`
}

func (m *DeleteCategoriesReq) Reset()                    { *m = DeleteCategoriesReq{} }
func (m *DeleteCategoriesReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoriesReq) ProtoMessage()               {}
func (*DeleteCategoriesReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteCategoriesReq) GetCourseAndCategories() []*CourseAndCategories {
	if m != nil {
		return m.CourseAndCategories
	}
	return nil
}

type PurgeCategoriesReq struct {
	CourseAndCategories []*CourseAndCategories `protobuf:"bytes,1,rep,name=course_and_categories,json=courseAndCategories" json:"course_and_categories,omitempty"`
}

func (m *PurgeCategoriesReq) Reset()                    { *m = PurgeCategoriesReq{} }
func (m *PurgeCategoriesReq) String() string            { return proto.CompactTextString(m) }
func (*PurgeCategoriesReq) ProtoMessage()               {}
func (*PurgeCategoriesReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PurgeCategoriesReq) GetCourseAndCategories() []*CourseAndCategories {
	if m != nil {
		return m.CourseAndCategories
	}
	return nil
}

// Course defines data structure of Course.
type Course struct {
	// @inject_tag: bson:"_id" valid:"uuid,required~ID is required"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"uuid,required~ID is required"`
	// @inject_tag: bson:"name" valid:"required~name is required,length(1|256)~name must be max 256 characters"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~name is required,length(1|256)~name must be max 256 characters"`
	// @inject_tag: valid:"required~Slug is required"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" valid:"required~Slug is required"`
	// @inject_tag: bson:"display_order"
	DisplayOrder uint64 `protobuf:"varint,4,opt,name=display_order,json=displayOrder" json:"display_order,omitempty" bson:"display_order"`
	// @inject_tag: valid:"required~Description is required"
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty" valid:"required~Description is required"`
	Hidden      bool   `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
	// @inject_tag: valid:"required~Start date time is required"
	Start *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=start" json:"start,omitempty" valid:"required~Start date time is required"`
	// @inject_tag: valid:"required~End date time is required"
	End *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=end" json:"end,omitempty" valid:"required~End date time is required"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
	// @inject_tag: bson:"category_ids"
	CategoryIds []string `protobuf:"bytes,11,rep,name=category_ids,json=categoryIds" json:"category_ids,omitempty" bson:"category_ids"`
}

func (m *Course) Reset()                    { *m = Course{} }
func (m *Course) String() string            { return proto.CompactTextString(m) }
func (*Course) ProtoMessage()               {}
func (*Course) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Course) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Course) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Course) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Course) GetDisplayOrder() uint64 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

func (m *Course) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Course) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Course) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Course) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Course) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Course) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Course) GetCategoryIds() []string {
	if m != nil {
		return m.CategoryIds
	}
	return nil
}

func init() {
	proto.RegisterType((*UpsertCourseReq)(nil), "go.micro.srv.courses.UpsertCourseReq")
	proto.RegisterType((*UpsertCourseRsp)(nil), "go.micro.srv.courses.UpsertCourseRsp")
	proto.RegisterType((*InsertCoursesRsp)(nil), "go.micro.srv.courses.InsertCoursesRsp")
	proto.RegisterType((*DeleteCoursesRsp)(nil), "go.micro.srv.courses.DeleteCoursesRsp")
	proto.RegisterType((*IDs)(nil), "go.micro.srv.courses.IDs")
	proto.RegisterType((*UpdateCoursesRsp)(nil), "go.micro.srv.courses.UpdateCoursesRsp")
	proto.RegisterType((*FindCoursesRsp)(nil), "go.micro.srv.courses.FindCoursesRsp")
	proto.RegisterType((*Courses)(nil), "go.micro.srv.courses.Courses")
	proto.RegisterType((*FindCoursesReq)(nil), "go.micro.srv.courses.FindCoursesReq")
	proto.RegisterType((*CourseAndCategories)(nil), "go.micro.srv.courses.CourseAndCategories")
	proto.RegisterType((*SyncCategoriesReq)(nil), "go.micro.srv.courses.SyncCategoriesReq")
	proto.RegisterType((*AddCategoriesReq)(nil), "go.micro.srv.courses.AddCategoriesReq")
	proto.RegisterType((*DeleteCategoriesReq)(nil), "go.micro.srv.courses.DeleteCategoriesReq")
	proto.RegisterType((*PurgeCategoriesReq)(nil), "go.micro.srv.courses.PurgeCategoriesReq")
	proto.RegisterType((*Course)(nil), "go.micro.srv.courses.Course")
}

func init() { proto.RegisterFile("proto/course/course.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0x56, 0x92, 0xe6, 0xe2, 0x93, 0xa6, 0x1b, 0xa6, 0xbb, 0xe0, 0x7a, 0x17, 0x11, 0xcc, 0x2d,
	0x2b, 0xad, 0x5c, 0x54, 0x10, 0x12, 0x8f, 0xd9, 0x8d, 0x10, 0x79, 0x81, 0x95, 0x97, 0x16, 0xa8,
	0x0a, 0x91, 0xeb, 0x19, 0xcc, 0x48, 0x89, 0xed, 0xcc, 0x8c, 0x5b, 0xe5, 0x81, 0xff, 0xc1, 0x23,
	0x7f, 0x84, 0xff, 0xb6, 0x9a, 0x8b, 0x93, 0x38, 0xce, 0xa5, 0xed, 0x43, 0x9f, 0xe2, 0x39, 0xe7,
	0x3b, 0xdf, 0x9c, 0x39, 0xd7, 0xc0, 0x49, 0xca, 0x12, 0x91, 0x9c, 0x86, 0x49, 0xc6, 0x38, 0x31,
	0x3f, 0x9e, 0x92, 0xa1, 0xa7, 0x51, 0xe2, 0x4d, 0x69, 0xc8, 0x12, 0x8f, 0xb3, 0x1b, 0x4f, 0xab,
	0xb8, 0xf3, 0x49, 0x94, 0x24, 0xd1, 0x84, 0x9c, 0x2a, 0xcc, 0x75, 0xf6, 0xd7, 0xa9, 0xa0, 0x53,
	0xc2, 0x45, 0x30, 0x4d, 0xb5, 0x99, 0xf3, 0x7c, 0x1d, 0x40, 0xa6, 0xa9, 0x98, 0x6b, 0xa5, 0xfb,
	0x27, 0x3c, 0x39, 0x4f, 0x39, 0x61, 0xe2, 0x8d, 0xa2, 0xf3, 0xc9, 0x0c, 0x7d, 0x0b, 0x0d, 0xcd,
	0x6d, 0x57, 0x7a, 0x95, 0x7e, 0xfb, 0xec, 0x85, 0xb7, 0xe9, 0x5e, 0xcf, 0x18, 0x18, 0x2c, 0x7a,
	0x06, 0x0d, 0xca, 0xc7, 0x31, 0xb9, 0xb5, 0xab, 0xbd, 0x4a, 0xbf, 0xe5, 0xd7, 0x29, 0xff, 0x89,
	0xdc, 0xba, 0xff, 0xac, 0xf1, 0xf3, 0xf4, 0x81, 0xfc, 0x0e, 0xb4, 0x68, 0x2c, 0x89, 0x08, 0x56,
	0x37, 0xd4, 0xfc, 0xc5, 0x19, 0xd9, 0xd0, 0xcc, 0x52, 0x1c, 0x48, 0x55, 0x4d, 0xa9, 0xf2, 0xa3,
	0xeb, 0x41, 0x77, 0x14, 0x2f, 0xaf, 0xe7, 0xf2, 0xfe, 0x55, 0xa6, 0x4a, 0x91, 0xc9, 0x7d, 0x05,
	0xdd, 0x21, 0x99, 0x10, 0x41, 0x56, 0xf0, 0x36, 0x34, 0xb1, 0x92, 0xe5, 0xf0, 0xfc, 0xe8, 0x7e,
	0x04, 0xb5, 0xd1, 0x90, 0xa3, 0x2e, 0xd4, 0x28, 0xe6, 0x76, 0xa5, 0x57, 0xeb, 0x5b, 0xbe, 0xfc,
	0x94, 0x34, 0xe7, 0xca, 0x83, 0x22, 0x4d, 0xee, 0x64, 0xa5, 0xe8, 0xe4, 0x8f, 0x70, 0xf4, 0x03,
	0x8d, 0xf1, 0x0a, 0xf6, 0x3b, 0x68, 0x9a, 0x30, 0x28, 0xd6, 0x7d, 0x31, 0xca, 0xc1, 0xee, 0x00,
	0x9a, 0x86, 0xe5, 0xc1, 0x14, 0xff, 0x55, 0x8b, 0xde, 0x90, 0x59, 0xf9, 0x7d, 0xe8, 0x29, 0xd4,
	0x67, 0x19, 0x61, 0x73, 0x95, 0x09, 0xcb, 0xd7, 0x07, 0xf4, 0x35, 0xd4, 0xb9, 0x08, 0x98, 0x50,
	0x49, 0x68, 0x9f, 0x39, 0x9e, 0x2e, 0x3c, 0x2f, 0x2f, 0x3c, 0xef, 0x97, 0xbc, 0x32, 0x7d, 0x0d,
	0x44, 0xaf, 0xa0, 0x46, 0x62, 0x6c, 0x1f, 0xec, 0xc5, 0x4b, 0x18, 0x7a, 0x01, 0x16, 0x8d, 0xc3,
	0x49, 0xc6, 0xe9, 0x0d, 0xb1, 0xeb, 0xaa, 0xca, 0x96, 0x02, 0xf4, 0x21, 0x34, 0xfe, 0xa6, 0x18,
	0x93, 0xd8, 0x6e, 0x28, 0x95, 0x39, 0x21, 0x04, 0x07, 0x3c, 0x61, 0xc2, 0x6e, 0x2a, 0xf7, 0xd5,
	0x37, 0x3a, 0x81, 0x56, 0x4a, 0xd8, 0x38, 0x0d, 0x22, 0x62, 0xb7, 0x74, 0x32, 0x52, 0xc2, 0xde,
	0x06, 0x11, 0x41, 0x9f, 0xc2, 0x61, 0x98, 0x31, 0x46, 0x62, 0xa1, 0xd5, 0x96, 0x52, 0xb7, 0x8d,
	0x4c, 0x42, 0xdc, 0x73, 0x38, 0xd6, 0xd1, 0x19, 0xc4, 0xf8, 0x4d, 0x20, 0x48, 0x94, 0x30, 0x4a,
	0x38, 0x7a, 0x0e, 0x96, 0x0e, 0xe2, 0x98, 0xea, 0x14, 0x5b, 0x7e, 0x4b, 0x0b, 0x46, 0x58, 0xd1,
	0x6a, 0xe8, 0x7c, 0x2c, 0x83, 0x59, 0x55, 0xde, 0xb4, 0x73, 0xd9, 0x08, 0x73, 0x97, 0xc1, 0x07,
	0xef, 0xe6, 0x71, 0xb8, 0x64, 0x94, 0xb1, 0xff, 0x03, 0x9e, 0x19, 0xd2, 0x20, 0xc6, 0xe3, 0x70,
	0xa1, 0x33, 0x49, 0x7d, 0xb9, 0x2b, 0xa9, 0x05, 0xf7, 0xfc, 0xe3, 0xb0, 0x2c, 0x74, 0x67, 0xd0,
	0x1d, 0x60, 0xfc, 0xa8, 0x57, 0x0a, 0x38, 0x36, 0x2d, 0xf6, 0x98, 0xb7, 0x72, 0x40, 0x6f, 0x33,
	0x16, 0x3d, 0xee, 0xa5, 0xff, 0xd6, 0xa0, 0xa1, 0xc1, 0xe8, 0x08, 0xaa, 0xa3, 0xa1, 0xa9, 0x8a,
	0xea, 0x68, 0x28, 0xab, 0x32, 0x0e, 0xa6, 0xc4, 0x34, 0x90, 0xfa, 0x56, 0x95, 0x3a, 0xc9, 0x22,
	0xd5, 0x3e, 0xb2, 0x52, 0x27, 0x59, 0x84, 0x3e, 0x83, 0x0e, 0xa6, 0x3c, 0x9d, 0x04, 0xf3, 0x71,
	0xc2, 0x30, 0x61, 0xaa, 0x57, 0x0e, 0xfc, 0x43, 0x23, 0xfc, 0x59, 0xca, 0x50, 0x0f, 0xda, 0x98,
	0xf0, 0x90, 0xd1, 0x54, 0xd0, 0x24, 0x56, 0xad, 0x61, 0xf9, 0xab, 0xa2, 0xad, 0xcd, 0xb1, 0x68,
	0xd9, 0xe6, 0x3d, 0x5b, 0xb6, 0x75, 0xb7, 0x96, 0xfd, 0x1e, 0x20, 0x64, 0x44, 0x4e, 0xb9, 0x71,
	0x20, 0x54, 0x2f, 0xed, 0x36, 0xb2, 0x0c, 0x7a, 0x20, 0xa4, 0xa9, 0x19, 0x90, 0xd2, 0x14, 0xf6,
	0x9b, 0x1a, 0xf4, 0x40, 0x94, 0x9a, 0xad, 0x5d, 0x6a, 0xb6, 0xb3, 0xff, 0x1b, 0xd0, 0xd1, 0xa9,
	0x79, 0x47, 0xd8, 0x0d, 0x0d, 0x09, 0xba, 0x82, 0xc3, 0xd5, 0x4d, 0x85, 0xbe, 0xd8, 0x9c, 0xfc,
	0xb5, 0x6d, 0xe9, 0xdc, 0x05, 0xc6, 0x53, 0x74, 0x01, 0x9d, 0xc2, 0x22, 0x42, 0x1f, 0xef, 0xaa,
	0x2d, 0xee, 0x7c, 0xb9, 0x59, 0x5d, 0x5a, 0x66, 0x17, 0xd0, 0x29, 0x6c, 0x9a, 0x07, 0xf2, 0x96,
	0xb6, 0xd5, 0xef, 0xd0, 0x5e, 0xd9, 0x02, 0xe8, 0xf3, 0xcd, 0x66, 0xc5, 0x45, 0xe1, 0xdc, 0x01,
	0xc5, 0x53, 0xf4, 0x2b, 0xa0, 0xc2, 0x8e, 0x7d, 0x3d, 0x97, 0x4b, 0xf4, 0x64, 0xcb, 0x83, 0x87,
	0x5b, 0x7d, 0x2e, 0x2d, 0xea, 0xdf, 0xe0, 0xa8, 0x38, 0x40, 0xd1, 0x57, 0x9b, 0x2d, 0x4b, 0x63,
	0xd6, 0xd9, 0x1d, 0x35, 0x19, 0xe5, 0xc2, 0x98, 0x44, 0x5b, 0x5c, 0x5a, 0x9f, 0xa5, 0xfb, 0x78,
	0xaf, 0x16, 0x7f, 0x37, 0x96, 0xd4, 0x2f, 0x77, 0xbe, 0xf6, 0x3e, 0xec, 0x97, 0xf0, 0x64, 0x6d,
	0xe6, 0xa1, 0xfe, 0x66, 0x8b, 0xf2, 0x68, 0xdc, 0xc3, 0xfd, 0xba, 0x75, 0x69, 0xfe, 0x98, 0x5d,
	0x37, 0x54, 0x2f, 0x7e, 0xf3, 0x3e, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x4f, 0x91, 0x80, 0xb9, 0x0a,
	0x00, 0x00,
}
