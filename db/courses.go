package db

import (
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type Courses interface {
	// Course related rpcs
	UpsertCourse(course *pcourse.Course) (*pcourse.Course, error)
	InsertCourses(courses *pcourse.Courses) error
	UpdateCourses(courses *pcourse.Courses) (int, error)
	FindCourses(request *pcourse.FindCoursesReq) (*pcourse.Courses, error)
	DeleteCoursesByIDs(ids []string) error
	// Category related rpcs
	SyncCategories(req *pcourse.SyncCategoriesReq) (*pcourse.Courses, error)
	AddCategories(req *pcourse.AddCategoriesReq) (*pcourse.Courses, error)
	DeleteCategories(req *pcourse.DeleteCategoriesReq) (*pcourse.Courses, error)
	PurgeCategories(req *pcourse.PurgeCategoriesReq) error
}
