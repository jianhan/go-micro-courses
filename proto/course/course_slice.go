package course

import (
	"github.com/asaskevich/govalidator"
	"github.com/satori/go.uuid"
)

func (r *CourseSlice) Validate() error {
	for k := range r.Courses {
		if _, err := govalidator.ValidateStruct(r.Courses[k]); err != nil {
			return err
		}
	}
	return nil
}

func (r *CourseSlice) GenerateIDs() {
	for k := range r.Courses {
		r.Courses[k].ID = uuid.Must(uuid.NewV4()).String()
	}
}
