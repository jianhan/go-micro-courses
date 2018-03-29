package category

import (
	fmt "fmt"

	"github.com/asaskevich/govalidator"
)

func (f *FindCategoriesRequest) Validate() error {
	for _, v := range f.Ids {
		if !govalidator.IsUUID(v) {
			return fmt.Errorf("ID %s is not a valid UUID", v)
		}
	}
	return nil
}
