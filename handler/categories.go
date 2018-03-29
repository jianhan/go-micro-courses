package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jianhan/go-micro-courses/db"
	pcategory "github.com/jianhan/go-micro-courses/proto/category"
)

// Categories handles all incomming request related to categories.
type Categories struct {
	DB db.Categories
}

func (c *Categories) InsertCategories(ctx context.Context, req *pcategory.CategorySlice, rsp *pcategory.InsertCategoriesResponse) (err error) {
	if err = req.GenerateIDs().GenerateCreatedUpdated().GenerateSlugs().Validate(); err != nil {
		return
	}
	if rsp.Inserted, err = c.DB.InsertCategories(req); err != nil {
		return
	}
	return
}

func (c *Categories) UpdateCategories(ctx context.Context, req *pcategory.CategorySlice, rsp *pcategory.UpdateCategoriesResponse) (err error) {
	if err = req.GenerateSlugs().GenerateCreatedUpdated().Validate(); err != nil {
		return err
	}
	rsp.Updated, err = c.DB.UpdateCategories(req)
	if err != nil {
		return
	}
	return
}

func (c *Categories) FindCategories(ctx context.Context, req *pcategory.FindCategoriesRequest, rsp *pcategory.CategorySlice) (err error) {
	if err = req.Validate(); err != nil {
		return err
	}
	rsp.Categories, err = c.DB.FindCategories(req)
	if err != nil {
		return
	}
	return
}

func (c *Categories) DeleteCategoriesByIDs(ctx context.Context, req *pcategory.DeleteCategoriesByIDsRequest, rsp *empty.Empty) (err error) {
	if len(req.Ids) == 0 {
		return errors.New("empty ids")
	}
	for _, v := range req.Ids {
		if !govalidator.IsUUID(v) {
			return fmt.Errorf("ID %s is not a valid UUID", v)
		}
	}
	if err = c.DB.DeleteCategoriesByIDs(req.Ids); err != nil {
		return
	}
	return
}
