package course

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func (f *FindCoursesReq) Validate() error {
	for _, v := range f.Ids {
		if !govalidator.IsUUID(v) {
			return fmt.Errorf("ID %s is not a valid UUID", v)
		}
	}
	return nil
}
