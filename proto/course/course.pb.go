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
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

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
	Ids       []string                   `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Query     string                     `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Start     *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=start" json:"start,omitempty"`
	End       *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=end" json:"end,omitempty"`
	Inclusive bool                       `protobuf:"varint,5,opt,name=inclusive" json:"inclusive,omitempty"`
	Hidden    bool                       `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
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

// Course defines data structure of Course.
type Course struct {
	// @inject_tag: bson:"_id"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id"`
	// @inject_tag: bson:"name"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name"`
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	// @inject_tag: bson:"display_order"
	DisplayOrder uint64                     `protobuf:"varint,4,opt,name=display_order,json=displayOrder" json:"display_order,omitempty" bson:"display_order"`
	Description  string                     `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Hidden       bool                       `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
	Start        *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=start" json:"start,omitempty"`
	End          *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=end" json:"end,omitempty"`
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Courses service

type CoursesClient interface {
	InsertCourses(ctx context.Context, in *CourseSlice, opts ...client.CallOption) (*google_protobuf1.Empty, error)
	UpdateCourses(ctx context.Context, in *CourseSlice, opts ...client.CallOption) (*UpdateCoursesRsp, error)
	FindCourses(ctx context.Context, in *FindCoursesRequest, opts ...client.CallOption) (*CourseSlice, error)
}

type coursesClient struct {
	c           client.Client
	serviceName string
}

func NewCoursesClient(serviceName string, c client.Client) CoursesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.courses"
	}
	return &coursesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *coursesClient) InsertCourses(ctx context.Context, in *CourseSlice, opts ...client.CallOption) (*google_protobuf1.Empty, error) {
	req := c.c.NewRequest(c.serviceName, "Courses.InsertCourses", in)
	out := new(google_protobuf1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesClient) UpdateCourses(ctx context.Context, in *CourseSlice, opts ...client.CallOption) (*UpdateCoursesRsp, error) {
	req := c.c.NewRequest(c.serviceName, "Courses.UpdateCourses", in)
	out := new(UpdateCoursesRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesClient) FindCourses(ctx context.Context, in *FindCoursesRequest, opts ...client.CallOption) (*CourseSlice, error) {
	req := c.c.NewRequest(c.serviceName, "Courses.FindCourses", in)
	out := new(CourseSlice)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Courses service

type CoursesHandler interface {
	InsertCourses(context.Context, *CourseSlice, *google_protobuf1.Empty) error
	UpdateCourses(context.Context, *CourseSlice, *UpdateCoursesRsp) error
	FindCourses(context.Context, *FindCoursesRequest, *CourseSlice) error
}

func RegisterCoursesHandler(s server.Server, hdlr CoursesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Courses{hdlr}, opts...))
}

type Courses struct {
	CoursesHandler
}

func (h *Courses) InsertCourses(ctx context.Context, in *CourseSlice, out *google_protobuf1.Empty) error {
	return h.CoursesHandler.InsertCourses(ctx, in, out)
}

func (h *Courses) UpdateCourses(ctx context.Context, in *CourseSlice, out *UpdateCoursesRsp) error {
	return h.CoursesHandler.UpdateCourses(ctx, in, out)
}

func (h *Courses) FindCourses(ctx context.Context, in *FindCoursesRequest, out *CourseSlice) error {
	return h.CoursesHandler.FindCourses(ctx, in, out)
}

func init() { proto.RegisterFile("proto/course/course.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0x9a, 0xf5, 0x23, 0x37, 0x2b, 0x9a, 0xac, 0x69, 0x0a, 0x65, 0x12, 0x59, 0x91, 0x50,
	0x1e, 0x26, 0x07, 0x15, 0x09, 0x89, 0xc7, 0xc1, 0x86, 0x54, 0x5e, 0x90, 0x0c, 0xbc, 0xf4, 0xa5,
	0x4a, 0x63, 0x93, 0x1a, 0x92, 0x38, 0xb5, 0x9d, 0x49, 0xfd, 0x11, 0xbc, 0xf3, 0xbb, 0xf8, 0x45,
	0x28, 0x76, 0xa2, 0x0d, 0x5a, 0xe8, 0x78, 0x49, 0xec, 0xe3, 0x73, 0xec, 0x7b, 0xae, 0xce, 0x85,
	0xc7, 0x95, 0x14, 0x5a, 0xc4, 0xa9, 0xa8, 0xa5, 0x62, 0xed, 0x0f, 0x1b, 0x0c, 0x9d, 0x66, 0x02,
	0x17, 0x3c, 0x95, 0x02, 0x2b, 0x79, 0x8b, 0xed, 0x91, 0x9a, 0x3c, 0xcd, 0x84, 0xc8, 0x72, 0x16,
	0x1b, 0xce, 0xaa, 0xfe, 0x12, 0x6b, 0x5e, 0x30, 0xa5, 0x93, 0xa2, 0xb2, 0xb2, 0x09, 0xce, 0xb8,
	0x5e, 0xd7, 0x2b, 0x9c, 0x8a, 0x22, 0xfe, 0xca, 0x93, 0x72, 0x9d, 0x94, 0x71, 0xf5, 0x2d, 0xb3,
	0x82, 0xb8, 0xd8, 0xaa, 0x4d, 0x6e, 0xbf, 0x2d, 0xff, 0xc9, 0x9f, 0x17, 0xb2, 0xa2, 0xd2, 0x5b,
	0x7b, 0x38, 0xbd, 0x84, 0x93, 0xcf, 0x15, 0x4d, 0x34, 0x7b, 0x6b, 0x9f, 0x27, 0xaa, 0x42, 0x01,
	0x0c, 0x6b, 0x83, 0xd1, 0xc0, 0x09, 0x9d, 0xc8, 0x25, 0xdd, 0x76, 0x7a, 0x03, 0xbe, 0xe5, 0x7d,
	0xcc, 0x79, 0xca, 0xd0, 0x2b, 0x18, 0xb6, 0x55, 0x07, 0x4e, 0xe8, 0x46, 0xfe, 0xec, 0x1c, 0xef,
	0xb3, 0x84, 0xad, 0x86, 0x74, 0xe4, 0xe9, 0x4f, 0x07, 0xd0, 0x3b, 0x5e, 0xd2, 0xee, 0x4d, 0xb6,
	0xa9, 0x99, 0xd2, 0xe8, 0x04, 0x5c, 0x4e, 0xed, 0x55, 0x1e, 0x69, 0x96, 0xe8, 0x14, 0xfa, 0x9b,
	0x9a, 0xc9, 0x6d, 0xd0, 0x0b, 0x9d, 0xc8, 0x23, 0x76, 0x83, 0x5e, 0x40, 0x5f, 0xe9, 0x44, 0xea,
	0xc0, 0x0d, 0x9d, 0xc8, 0x9f, 0x4d, 0xb0, 0x35, 0x88, 0x3b, 0x83, 0xf8, 0x53, 0xd7, 0x31, 0x62,
	0x89, 0xe8, 0x12, 0x5c, 0x56, 0xd2, 0xe0, 0xe8, 0x20, 0xbf, 0xa1, 0xa1, 0x73, 0xf0, 0x78, 0x99,
	0xe6, 0xb5, 0xe2, 0xb7, 0x2c, 0xe8, 0x87, 0x4e, 0x34, 0x22, 0x77, 0x00, 0x3a, 0x83, 0xc1, 0x9a,
	0x53, 0xca, 0xca, 0x60, 0x60, 0x8e, 0xda, 0xdd, 0xf4, 0x87, 0x0b, 0x03, 0x6b, 0x08, 0x3d, 0x82,
	0xde, 0xfc, 0xda, 0xf4, 0xce, 0x23, 0xbd, 0xf9, 0x35, 0x42, 0x70, 0x54, 0x26, 0x05, 0x6b, 0x5d,
	0x98, 0x75, 0x83, 0xa9, 0xbc, 0xce, 0x8c, 0x07, 0x8f, 0x98, 0x35, 0x7a, 0x06, 0x63, 0xca, 0x55,
	0x95, 0x27, 0xdb, 0xa5, 0x90, 0x94, 0x49, 0x53, 0xf0, 0x11, 0x39, 0x6e, 0xc1, 0x0f, 0x0d, 0x86,
	0x42, 0xf0, 0x29, 0x53, 0xa9, 0xe4, 0x95, 0xe6, 0xa2, 0x34, 0xf5, 0x79, 0xe4, 0x3e, 0xf4, 0xb7,
	0x0a, 0xef, 0xfa, 0x36, 0xfc, 0xcf, 0xbe, 0x8d, 0x1e, 0xd6, 0xb7, 0xd7, 0x00, 0xa9, 0x64, 0x4d,
	0x50, 0x96, 0x89, 0x0e, 0xbc, 0x83, 0x22, 0xaf, 0x65, 0x5f, 0xe9, 0x46, 0xda, 0x66, 0xac, 0x91,
	0xc2, 0x61, 0x69, 0xcb, 0xbe, 0xd2, 0xe8, 0x02, 0x8e, 0xd3, 0x44, 0xb3, 0x4c, 0xc8, 0xed, 0xb2,
	0x89, 0x8f, 0x6f, 0xe2, 0xe3, 0x77, 0xd8, 0x9c, 0xaa, 0xd9, 0xf7, 0x1e, 0x0c, 0xdb, 0xac, 0xa1,
	0xf7, 0x30, 0x9e, 0x97, 0x8a, 0x49, 0xdd, 0x01, 0x17, 0xff, 0xca, 0xac, 0xc9, 0xf9, 0xe4, 0x6c,
	0xa7, 0x92, 0x9b, 0x66, 0x84, 0xd0, 0x02, 0xc6, 0xbf, 0x0d, 0xcf, 0x43, 0xee, 0x7a, 0xbe, 0x9f,
	0xb2, 0x33, 0x84, 0x0b, 0xf0, 0xef, 0x8d, 0x08, 0x8a, 0xf6, 0xcb, 0x76, 0xa7, 0x68, 0x72, 0xb8,
	0x86, 0x37, 0xa3, 0xc5, 0xc0, 0xc2, 0xab, 0x81, 0x71, 0xf4, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0d, 0xc0, 0x5d, 0x3d, 0xa6, 0x04, 0x00, 0x00,
}
