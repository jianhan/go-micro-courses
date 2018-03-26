package handler

import (
	"context"

	"github.com/asaskevich/govalidator"
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
		if _, err := govalidator.ValidateStruct(req.Courses[k]); err != nil {
			return err
		}
	}
	c.DB.InsertCourses(req)
	return
}

func (c *Courses) UpdateCourses(ctx context.Context, req *pcourse.CourseSlice, rsp *pcourse.UpdateCoursesRsp) (err error) {
	updated, err := c.DB.UpdateCourses(req)
	if err != nil {
		return err
	}
	rsp.Updated = int64(updated)
	return
}

func (c *Courses) FindCourses(ctx context.Context, req *pcourse.FindCoursesRequest, rsp *pcourse.CourseSlice) (err error) {
	r, err := c.DB.FindCourses(req)
	if err != nil {
		return
	}
	rsp.Courses = r.Courses
	return
}
