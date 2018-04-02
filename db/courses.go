package db

import (
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type Courses interface {
	UpsertCourse(course *pcourse.Course) (*pcourse.Course, error)
	InsertCourses(courses *pcourse.Courses) error
	UpdateCourses(courses *pcourse.Courses) (int, error)
	FindCourses(request *pcourse.FindCoursesReq) (*pcourse.Courses, error)
	DeleteCoursesByIDs(ids []string) error
}
