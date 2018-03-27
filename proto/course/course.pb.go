// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course/course.proto

/*
Package course is a generated protocol buffer package.

It is generated from these files:
	proto/course/course.proto

It has these top-level messages:
	UpdateCoursesRsp
	CourseSlice
	FindCoursesRequest
	Course
*/
package course

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/jianhan/pkg/proto/mysql"
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

type UpdateCoursesRsp struct {
	Updated int64 `protobuf:"varint,1,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateCoursesRsp) Reset()                    { *m = UpdateCoursesRsp{} }
func (m *UpdateCoursesRsp) String() string            { return proto.CompactTextString(m) }
func (*UpdateCoursesRsp) ProtoMessage()               {}
func (*UpdateCoursesRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateCoursesRsp) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type CourseSlice struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (m *CourseSlice) Reset()                    { *m = CourseSlice{} }
func (m *CourseSlice) String() string            { return proto.CompactTextString(m) }
func (*CourseSlice) ProtoMessage()               {}
func (*CourseSlice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CourseSlice) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type FindCoursesRequest struct {
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

func (m *FindCoursesRequest) Reset()                    { *m = FindCoursesRequest{} }
func (m *FindCoursesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindCoursesRequest) ProtoMessage()               {}
func (*FindCoursesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FindCoursesRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FindCoursesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *FindCoursesRequest) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *FindCoursesRequest) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *FindCoursesRequest) GetInclusive() bool {
	if m != nil {
		return m.Inclusive
	}
	return false
}

func (m *FindCoursesRequest) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *FindCoursesRequest) GetSort() []string {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *FindCoursesRequest) GetPerPage() int64 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *FindCoursesRequest) GetCurrentPage() int64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

// Course defines data structure of Course.
type Course struct {
	// @inject_tag: bson:"_id" valid:"uuid,required~ID is required"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"uuid,required~ID is required" bson:"_id" valid:"uuid,required~ID is required"`
	// @inject_tag: bson:"name" valid:"required~name is required,length(1|256)~name must be max 256 characters"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~name is required,length(1|256)~name must be max 256 characters" bson:"name" valid:"required~name is required,length(1|256)~name must be max 256 characters"`
	// @inject_tag: valid:"required~Slug is required"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" valid:"required~Slug is required" valid:"required~Slug is required"`
	// @inject_tag: bson:"display_order"
	DisplayOrder uint64 `protobuf:"varint,4,opt,name=display_order,json=displayOrder" json:"display_order,omitempty" bson:"display_order" bson:"display_order"`
	// @inject_tag: valid:"required~Description is required"
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty" valid:"required~Description is required" valid:"required~Description is required"`
	Hidden      bool   `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
	// @inject_tag: valid:"required~Start date time is required"
	Start *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=start" json:"start,omitempty" valid:"required~Start date time is required" valid:"required~Start date time is required"`
	// @inject_tag: valid:"required~End date time is required"
	End *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=end" json:"end,omitempty" valid:"required~End date time is required" valid:"required~End date time is required"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at" bson:"updated_at"`
	// @inject_tag: bson:"category_ids"
	CategoryIds []string `protobuf:"bytes,11,rep,name=category_ids,json=categoryIds" json:"category_ids,omitempty" bson:"category_ids" bson:"category_ids"`
}

func (m *Course) Reset()                    { *m = Course{} }
func (m *Course) String() string            { return proto.CompactTextString(m) }
func (*Course) ProtoMessage()               {}
func (*Course) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	proto.RegisterType((*UpdateCoursesRsp)(nil), "go.micro.srv.courses.UpdateCoursesRsp")
	proto.RegisterType((*CourseSlice)(nil), "go.micro.srv.courses.CourseSlice")
	proto.RegisterType((*FindCoursesRequest)(nil), "go.micro.srv.courses.FindCoursesRequest")
	proto.RegisterType((*Course)(nil), "go.micro.srv.courses.Course")
}

func init() { proto.RegisterFile("proto/course/course.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0x9a, 0xad, 0x6d, 0x6e, 0xd6, 0x9f, 0x26, 0x6b, 0x9a, 0xb2, 0xfe, 0x26, 0x91, 0x15,
	0x09, 0xe5, 0x61, 0x4a, 0x50, 0x91, 0x90, 0x78, 0x1c, 0x6c, 0x48, 0xe5, 0x05, 0x14, 0xe0, 0xa5,
	0x2f, 0x55, 0x9a, 0x5c, 0x52, 0x43, 0x12, 0xa7, 0xb6, 0x33, 0xa9, 0x1f, 0x82, 0x77, 0x3e, 0x00,
	0x1f, 0x14, 0xc5, 0x76, 0x60, 0xd0, 0xb2, 0x8e, 0x97, 0xc4, 0x3e, 0xf7, 0x1c, 0xfb, 0xfe, 0xf1,
	0x81, 0xb3, 0x9a, 0x33, 0xc9, 0xa2, 0x94, 0x35, 0x5c, 0xa0, 0xf9, 0x85, 0x0a, 0x23, 0x27, 0x39,
	0x0b, 0x4b, 0x9a, 0x72, 0x16, 0x0a, 0x7e, 0x1b, 0xea, 0x90, 0x18, 0x3f, 0xca, 0x19, 0xcb, 0x0b,
	0x8c, 0x14, 0x67, 0xd9, 0x7c, 0x8a, 0x24, 0x2d, 0x51, 0xc8, 0xa4, 0xac, 0xb5, 0x6c, 0x1c, 0xe6,
	0x54, 0xae, 0x9a, 0x65, 0x98, 0xb2, 0x32, 0xfa, 0x4c, 0x93, 0x6a, 0x95, 0x54, 0x51, 0xfd, 0x25,
	0xd7, 0x82, 0xa8, 0xdc, 0x88, 0x75, 0xa1, 0xbf, 0x86, 0xff, 0xff, 0x9f, 0x07, 0x62, 0x59, 0xcb,
	0x8d, 0x0e, 0x4e, 0x2e, 0xe1, 0xf8, 0x63, 0x9d, 0x25, 0x12, 0x5f, 0xe9, 0xeb, 0x63, 0x51, 0x13,
	0x0f, 0x06, 0x8d, 0xc2, 0x32, 0xcf, 0xf2, 0xad, 0xc0, 0x8e, 0xbb, 0xed, 0xe4, 0x06, 0x5c, 0xcd,
	0x7b, 0x5f, 0xd0, 0x14, 0xc9, 0x73, 0x18, 0x98, 0xac, 0x3d, 0xcb, 0xb7, 0x03, 0x77, 0x7a, 0x1e,
	0xee, 0x2a, 0x29, 0xd4, 0x9a, 0xb8, 0x23, 0x4f, 0xbe, 0xf7, 0x80, 0xbc, 0xa6, 0x55, 0xd6, 0xdd,
	0x89, 0xeb, 0x06, 0x85, 0x24, 0xc7, 0x60, 0xd3, 0x4c, 0x1f, 0xe5, 0xc4, 0xed, 0x92, 0x9c, 0xc0,
	0xe1, 0xba, 0x41, 0xbe, 0xf1, 0x7a, 0xbe, 0x15, 0x38, 0xb1, 0xde, 0x90, 0xa7, 0x70, 0x28, 0x64,
	0xc2, 0xa5, 0x67, 0xfb, 0x56, 0xe0, 0x4e, 0xc7, 0xa1, 0x2e, 0x30, 0xec, 0x0a, 0x0c, 0x3f, 0x74,
	0x1d, 0x8b, 0x35, 0x91, 0x5c, 0x82, 0x8d, 0x55, 0xe6, 0x1d, 0xec, 0xe5, 0xb7, 0x34, 0x72, 0x0e,
	0x0e, 0xad, 0xd2, 0xa2, 0x11, 0xf4, 0x16, 0xbd, 0x43, 0xdf, 0x0a, 0x86, 0xf1, 0x2f, 0x80, 0x9c,
	0x42, 0x7f, 0x45, 0xb3, 0x0c, 0x2b, 0xaf, 0xaf, 0x42, 0x66, 0x47, 0x08, 0x1c, 0x08, 0xc6, 0xa5,
	0x37, 0x50, 0xe9, 0xab, 0x35, 0x39, 0x83, 0x61, 0x8d, 0x7c, 0x51, 0x27, 0x39, 0x7a, 0x43, 0xdd,
	0xca, 0x1a, 0xf9, 0xbb, 0x24, 0x47, 0x72, 0x01, 0x47, 0x69, 0xc3, 0x39, 0x56, 0x52, 0x87, 0x1d,
	0x15, 0x76, 0x0d, 0xd6, 0x52, 0x26, 0xdf, 0x6c, 0xe8, 0xeb, 0x16, 0x91, 0xff, 0xa0, 0x37, 0xbb,
	0x56, 0xd3, 0x70, 0xe2, 0xde, 0xec, 0xba, 0xbd, 0xac, 0x4a, 0x4a, 0x34, 0x7d, 0x51, 0x6b, 0x95,
	0x40, 0xd1, 0xe4, 0xaa, 0x2b, 0x6d, 0x02, 0x45, 0x93, 0x93, 0xc7, 0x30, 0xca, 0xa8, 0xa8, 0x8b,
	0x64, 0xb3, 0x60, 0x3c, 0x43, 0xae, 0x5a, 0x70, 0x10, 0x1f, 0x19, 0xf0, 0x6d, 0x8b, 0x11, 0x1f,
	0xdc, 0x0c, 0x45, 0xca, 0x69, 0x2d, 0x29, 0xab, 0x54, 0xc5, 0x4e, 0x7c, 0x17, 0xfa, 0x6b, 0xcd,
	0x3f, 0x27, 0x31, 0xf8, 0xc7, 0x49, 0x0c, 0x1f, 0x36, 0x89, 0x17, 0x00, 0x29, 0xc7, 0xf6, 0xe9,
	0x2d, 0x12, 0xa9, 0x5a, 0x74, 0xbf, 0xc8, 0x31, 0xec, 0x2b, 0xd9, 0x4a, 0xcd, 0xab, 0x6d, 0xa5,
	0xb0, 0x5f, 0x6a, 0xd8, 0x57, 0x52, 0x8d, 0x26, 0x91, 0x98, 0x33, 0xbe, 0x59, 0xb4, 0x0f, 0xd2,
	0x55, 0x13, 0x75, 0x3b, 0x6c, 0x96, 0x89, 0xe9, 0xd7, 0x1e, 0x0c, 0xcc, 0xeb, 0x25, 0x6f, 0x60,
	0x34, 0xab, 0x04, 0x72, 0xd9, 0x01, 0x17, 0xf7, 0xb9, 0x40, 0x39, 0x67, 0x7c, 0xba, 0x95, 0xc9,
	0x4d, 0x6b, 0x4a, 0x32, 0x87, 0xd1, 0x6f, 0x76, 0x7c, 0xc8, 0x59, 0x4f, 0x76, 0x53, 0xb6, 0x6c,
	0x3d, 0x07, 0xf7, 0x8e, 0xe9, 0x48, 0xb0, 0x5b, 0xb6, 0xed, 0xcb, 0xf1, 0xfe, 0x1c, 0x5e, 0x0e,
	0xe7, 0x7d, 0x0d, 0x2f, 0xfb, 0xaa, 0xa2, 0x67, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x31, 0x7d,
	0x6a, 0xf9, 0xf8, 0x04, 0x00, 0x00,
}
