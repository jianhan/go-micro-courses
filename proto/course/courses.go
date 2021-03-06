package course

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/golang/protobuf/ptypes"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
)

func (r *Courses) Validate() error {
	if len(r.Courses) == 0 {
		return errors.New("Courses are empty")
	}
	for k := range r.Courses {
		if _, err := govalidator.ValidateStruct(r.Courses[k]); err != nil {
			return err
		}
	}
	return nil
}

func (r *Courses) GenerateIDs() *Courses {
	for k := range r.Courses {
		r.Courses[k].ID = uuid.Must(uuid.NewV4()).String()
	}
	return r
}

func (r *Courses) GenerateCreatedUpdated() *Courses {
	for k := range r.Courses {
		r.Courses[k].UpdatedAt = ptypes.TimestampNow()
		if r.Courses[k].ID == "" {
			r.Courses[k].CreatedAt = ptypes.TimestampNow()
		}
	}
	return r
}

func (r *Courses) GenerateSlugs() *Courses {
	for k := range r.Courses {
		if r.Courses[k].Slug == "" {
			r.Courses[k].Slug = slug.Make(r.Courses[k].Name)
		} else if !slug.IsSlug(r.Courses[k].Slug) {
			r.Courses[k].Slug = slug.Make(r.Courses[k].Slug)
		}
	}
	return r
}

func (r *Courses) GetCourseIDs() []string {
	courseIDs := []string{}
	for _, v := range r.Courses {
		courseIDs = append(courseIDs, v.ID)
	}
	return courseIDs
}
