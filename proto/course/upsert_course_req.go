package course

import (
	"strings"

	"time"

	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/golang/protobuf/ptypes"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
)

func (r *UpsertCourseReq) Validate() error {
	// resolve ID
	if strings.TrimSpace(r.Course.ID) == "" {
		r.Course.ID = uuid.Must(uuid.NewV4()).String()
	}
	// resolve created at
	if r.IsNew {
		n, err := ptypes.TimestampProto(time.Now())
		if err != nil {
			return err
		}
		r.Course.CreatedAt = n
	}
	// resolve updated at
	n, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return err
	}
	r.Course.UpdatedAt = n
	// resolve slug
	if r.Course.Slug == "" {
		r.Course.Name = slug.Make(r.Course.Name)
	}
	if _, err := govalidator.ValidateStruct(r.Course); err != nil {
		return err
	}
	// validate category IDs
	for _, v := range r.Course.CategoryIds {
		if !govalidator.IsUUID(v) {
			return fmt.Errorf("category ID %s is not a valid UUID", v)
		}
	}
	return nil
}
