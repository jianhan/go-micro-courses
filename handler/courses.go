package handler

import (
	"context"

	"github.com/jianhan/go-micro-courses/db"
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

// Courses handles all incomming request related to course.
type Courses struct {
	DB db.Courses
}

func (c *Courses) UpsertCourse(ctx context.Context, req *pcourse.UpsertCourseReq, rsp *pcourse.UpsertCourseRsp) (err error) {
	if err = req.Validate(); err != nil {
		return
	}
	if rsp.Course, err = c.DB.UpsertCourse(req.Course); err != nil {
		return
	}
	return
}

func (c *Courses) InsertCourses(ctx context.Context, req *pcourse.Courses, rsp *pcourse.InsertCoursesRsp) (err error) {
	// TODO: fill rsp
	if err = req.GenerateIDs().GenerateCreatedUpdated().GenerateSlugs().Validate(); err != nil {
		return
	}
	if err = c.DB.InsertCourses(req); err != nil {
		return
	}
	return
}

func (c *Courses) UpdateCourses(ctx context.Context, req *pcourse.Courses, rsp *pcourse.UpdateCoursesRsp) (err error) {
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

func (c *Courses) FindCourses(ctx context.Context, req *pcourse.FindCoursesReq, rsp *pcourse.FindCoursesRsp) (err error) {
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

func (c *Courses) SyncCategories(ctx context.Context, req *pcourse.SyncCategoriesReq, rsp *pcourse.Courses) (err error) {
	if rsp, err = c.DB.SyncCategories(req); err != nil {
		return err
	}
	return nil
}

func (c *Courses) AddCategories(ctx context.Context, req *pcourse.AddCategoriesReq, rsp *pcourse.Courses) (err error) {
	if rsp, err = c.DB.AddCategories(req); err != nil {
		return err
	}
	return nil
}

func (c *Courses) DeleteCategories(ctx context.Context, req *pcourse.DeleteCategoriesReq, rsp *pcourse.Courses) (err error) {
	if rsp, err = c.DB.DeleteCategories(req); err != nil {
		return err
	}
	return nil
}

func (c *Courses) PurgeCategories(ctx context.Context, req *pcourse.PurgeCategoriesReq, rsp *pcourse.Courses) (err error) {
	if rsp, err = c.DB.PurgeCategories(req); err != nil {
		return err
	}
	return nil
}
