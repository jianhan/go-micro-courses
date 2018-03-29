package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jianhan/go-micro-courses/db"
	pcategory "github.com/jianhan/go-micro-courses/proto/category"
)

// Categories handles all incomming request related to categories.
type Categories struct {
	DB db.Categories
}

func (c *Categories) InsertCategories(ctx context.Context, req *pcategory.CategorySlice, rsp *pcategory.InsertCategoriesResponse) (err error) {
	req.GenerateIDs().GenerateCreatedUpdated().GenerateSlugs()
	if rsp.Inserted, err = c.DB.InsertCategories(req); err != nil {
		return
	}
	return
}

func (c *Categories) UpdateCategories(ctx context.Context, req *pcategory.CategorySlice, rsp *pcategory.UpdateCategoriesResponse) (err error) {
	req.GenerateSlugs()
	if err = req.Validate(); err != nil {
		return
	}
	req.GenerateCreatedUpdated()
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
	return nil
}
