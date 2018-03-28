// Code generated by protoc-gen-micro. DO NOT EDIT.
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
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/jianhan/pkg/proto/mysql"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = google_protobuf1.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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