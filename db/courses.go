package db

import (
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type Courses interface {
	UpsertCourse(course *pcourse.Course) (*pcourse.Course, error)
	InsertCourses(courses *pcourse.CourseSlice) error
	UpdateCourses(courses *pcourse.CourseSlice) (int, error)
	FindCourses(request *pcourse.FindCoursesRequest) (*pcourse.CourseSlice, error)
	DeleteCoursesByIDs(ids []string) error
}
