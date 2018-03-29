package handler

import (
	pcategory "github.com/jianhan/go-micro-courses/proto/category"
	"github.com/jianhan/go-micro-courses/db"
	"context"
)

// Categories handles all incomming request related to categories.
type Categories struct {
	DB db.Courses
}

func (c *Categories) UpsertCategories (ctx context.Context,req *pcategory.CategorySlice, rsp *pcategory.UpsertCategoriesRsp) (err error) {
	return nil
}

func (c *Categories) FindCategories (ctx context.Context,req *pcategory.FindCategoriesRequest, rsp *pcategory.CategorySlice) (err error) {
	return nil
}

func (c *Categories) DeleteCategoriesByIDs (ctx context.Context,req *pcategory.DeleteCategoriesByIDsRequest, rsp *pcategory.DeleteCategoriesByIDsResponse) (err error) {
	return nil
}
