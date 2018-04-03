package category

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/golang/protobuf/ptypes"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
)

func (c *Categories) Validate() error {
	if len(c.Categories) == 0 {
		return errors.New("Categories are empty")
	}
	for k := range c.Categories {
		if _, err := govalidator.ValidateStruct(c.Categories[k]); err != nil {
			return err
		}
	}
	return nil
}

func (c *Categories) GenerateIDs() *Categories {
	for k := range c.Categories {
		c.Categories[k].ID = uuid.Must(uuid.NewV4()).String()
	}
	return c
}

func (c *Categories) GenerateCreatedUpdated() *Categories {
	for k := range c.Categories {
		c.Categories[k].UpdatedAt = ptypes.TimestampNow()
		if c.Categories[k].ID == "" {
			c.Categories[k].CreatedAt = ptypes.TimestampNow()
		}
	}
	return c
}

func (c *Categories) GenerateSlugs() *Categories {
	for k := range c.Categories {
		if c.Categories[k].Slug == "" {
			c.Categories[k].Slug = slug.Make(c.Categories[k].Name)
		} else if !slug.IsSlug(c.Categories[k].Slug) {
			c.Categories[k].Slug = slug.Make(c.Categories[k].Slug)
		}
	}
	return c
}

func (c *Categories) GetCourseIDs() []string {
	courseIDs := []string{}
	for _, v := range c.Categories {
		courseIDs = append(courseIDs, v.ID)
	}
	return courseIDs
}
