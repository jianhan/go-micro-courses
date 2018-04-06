package course

import (
	"fmt"

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
