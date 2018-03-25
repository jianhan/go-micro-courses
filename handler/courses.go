package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jianhan/go-micro-courses/db"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
	"github.com/satori/go.uuid"
)

// API is a constant to define the name of API.
const API = "go_micro_srv_courses"

// Courses handles all incomming request related to course.
type Courses struct {
	DB db.Courses
}

func (c *Courses) InsertCourses(ctx context.Context, req *pcourse.CourseSlice, rsp *empty.Empty) (err error) {
	for k := range req.Courses {
		req.Courses[k].ID = uuid.Must(uuid.NewV4()).String()
	}
	c.DB.InsertCourses(req)
	return
}
