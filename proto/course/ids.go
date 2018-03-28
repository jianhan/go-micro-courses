package course

import (
	"github.com/asaskevich/govalidator"
	"fmt"
	"errors"
)

func (i *IDs) Validate() error {
	if len(i.Ids) == 0 {
		return errors.New("empty ids")
	}
	for _, v := range i.Ids {
		if !govalidator.IsUUID(v) {
			return fmt.Errorf("ID %s is not a valid UUID", v)
		}
	}
	return nil
}
