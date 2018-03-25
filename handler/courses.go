package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jianhan/go-micro-courses/db"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

// API is a constant to define the name of API.
const API = "go_micro_srv_courses"

// Courses handles all incomming request related to course.
type Courses struct {
	DB db.Courses
}

func (c *Courses) InsertCourses(ctx context.Context, req *pcourse.CourseSlice, rsp *empty.Empty) (err error) {
	c.DB.InsertCourses(req)
	return
}
