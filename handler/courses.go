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
	if err = req.GenerateIDs().GenerateCreatedUpdated().GenerateSlugs().Validate(); err != nil {
		return
	}
	if err := c.DB.InsertCourses(req); err != nil {
		return
	}
	return
}

func (c *Courses) UpdateCourses(ctx context.Context, req *pcourse.CourseSlice, rsp *pcourse.UpdateCoursesRsp) (err error) {
	if err = req.GenerateSlugs().GenerateCreatedUpdated().Validate(); err != nil {
		return err
	}
	updated, err := c.DB.UpdateCourses(req)
	if err != nil {
		return
	}
	rsp.Updated = int64(updated)
	return
}

func (c *Courses) FindCourses(ctx context.Context, req *pcourse.FindCoursesRequest, rsp *pcourse.CourseSlice) (err error) {
	if err := req.Validate(); err != nil {
		return err
	}
	r, err := c.DB.FindCourses(req)
	if err != nil {
		return
	}
	rsp.Courses = r.Courses
	return
}

func (c *Courses) DeleteCoursesByIDs(ctx context.Context, req *pcourse.IDs, rsp *pcourse.DeleteCoursesRsp) (err error) {
	if err = req.Validate(); err != nil {
		return err
	}
	if err = c.DB.DeleteCoursesByIDs(req.Ids); err != nil {
		return err
	}
	return nil
}
