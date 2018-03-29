package db

import pcategory "github.com/jianhan/go-micro-courses/proto/category"

type Categories interface {
	InsertCategories(categories *pcategory.CategorySlice) (int64, error)
	UpdateCategories(categories *pcategory.CategorySlice) (int64, error)
	FindCategories(req *pcategory.FindCategoriesRequest) ([]*pcategory.Category, error)
	DeleteCategoriesByIDs(ids []string) error
}
