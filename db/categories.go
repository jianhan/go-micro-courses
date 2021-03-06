package db

import pcategory "github.com/jianhan/go-micro-courses/proto/category"

type Categories interface {
	UpsertCategory(category *pcategory.Category) (*pcategory.Category, error)
	InsertCategories(categories *pcategory.Categories) (int64, error)
	UpdateCategories(categories *pcategory.Categories) (int64, error)
	FindCategories(req *pcategory.FindCategoriesRequest) ([]*pcategory.Category, error)
	DeleteCategoriesByIDs(ids []string) error
}
