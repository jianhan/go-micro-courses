package category

import (
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang/protobuf/ptypes"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
)

func (r *UpsertCategoryReq) Validate() error {
	// resolve ID
	if strings.TrimSpace(r.Catetory.ID) == "" {
		r.Catetory.ID = uuid.Must(uuid.NewV4()).String()
	}
	// resolve created at
	if r.IsNew {
		n, err := ptypes.TimestampProto(time.Now())
		if err != nil {
			return err
		}
		r.Catetory.CreatedAt = n
	}
	// resolve updated at
	n, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return err
	}
	r.Catetory.UpdatedAt = n
	// resolve slug
	if r.Catetory.Slug == "" {
		r.Catetory.Name = slug.Make(r.Catetory.Name)
	}
	if _, err := govalidator.ValidateStruct(r.Catetory); err != nil {
		return err
	}
	return nil
}
