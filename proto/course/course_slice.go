package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/satori/go.uuid"
	"github.com/golang/protobuf/ptypes"
)

func (r *CourseSlice) Validate() error {
	for k := range r.Courses {
		if _, err := govalidator.ValidateStruct(r.Courses[k]); err != nil {
			return err
		}
	}
	return nil
}

func (r *CourseSlice) GenerateIDs() *CourseSlice {
	for k := range r.Courses {
		r.Courses[k].ID = uuid.Must(uuid.NewV4()).String()
	}
	return r
}

func (r *CourseSlice) GenerateCreatedUpdated() *CourseSlice {
	for k := range r.Courses {
		r.Courses[k].UpdatedAt = ptypes.TimestampNow()
		if r.Courses[k].ID == "" {
			r.Courses[k].CreatedAt = ptypes.TimestampNow()
		}
	}
	return r
}

func (r *CourseSlice) GetCourseIDs() []string {
	courseIDs := []string{}
	for _,v := range r.Courses {
		courseIDs = append(courseIDs, v.ID)
	}
	return courseIDs
}
