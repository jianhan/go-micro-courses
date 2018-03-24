package handlers

import (
	"context"

	pcourses "github.com/jianhan/go-micro-courses/proto/courses"
	pmysql "github.com/jianhan/pkg/proto/mysql"
)

// API is a constant to define the name of API.
const API = "go_micro_srv_courses"

// Courses handles all incomming request related to course.
type Courses struct {
}

// UpsertCourses upsert multiply courses.
func (c *Courses) UpsertCourses(ctx context.Context, req *pcourses.UpsertCoursesRequest, rsp *pmysql.UpsertResult) (err error) {
	return
}

// GetCoursesByFilters retrieves courses by filters.
func (c *Courses) GetCoursesByFilters(ctx context.Context, req *pcourses.GetCoursesByFiltersRequest, rsp *pcourses.GetCoursesByFiltersResponse) (err error) {
	return
}

// DeleteCoursesByFilters remove courses according to filter set.
func (c *Courses) DeleteCoursesByFilters(ctx context.Context, req *pcourses.DeleteCoursesByFiltersRequest, rsp *pcourses.DeleteCoursesByFiltersResponse) (err error) {
	return
}
