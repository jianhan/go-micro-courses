package db

import (
	pcourse "github.com/jianhan/go-micro-courses/proto/course"
)

type Courses interface {
	InsertCourses(courses *pcourse.CourseSlice) error
}
