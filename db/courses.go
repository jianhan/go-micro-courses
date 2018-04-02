package db

import (
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type Courses interface {
	UpsertCourse(course *pcourse.Course) (*pcourse.Course, error)
	InsertCourses(courses *pcourse.courses) error
	UpdateCourses(courses *pcourse.courses) (int, error)
	FindCourses(request *pcourse.FindCoursesRequest) (*pcourse.courses, error)
	DeleteCoursesByIDs(ids []string) error
}
