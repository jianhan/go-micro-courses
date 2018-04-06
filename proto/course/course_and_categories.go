package course

import (
	"fmt"

	"errors"

	"github.com/asaskevich/govalidator"
)

func (c *CourseAndCategories) Validate() (err error) {
	if isValid := govalidator.IsUUID(c.CourseId); !isValid {
		return fmt.Errorf("course ID %s is not a valid UUID", c.CourseId)
	}
	for _, v := range c.CategoryIds {
		if isValid := govalidator.IsUUID(v); !isValid {
			return fmt.Errorf("category ID %s is not a valid UUID", v)
		}
	}
	return nil
}

func ValidateCourseAndCategories(f func() []*CourseAndCategories) (err error) {
	if c := f(); len(c) == 0 {
		return errors.New("empty input")
	}
	for _, v := range f() {
		if err = v.Validate(); err != nil {
			return
		}
	}
	return
}
