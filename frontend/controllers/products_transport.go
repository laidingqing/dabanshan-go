package controllers

import (
	"github.com/laidingqing/dabanshan/pb"
	model "github.com/laidingqing/dabanshan/products/model"
)

//EncodeCategories ...
func EncodeCategories(categories []*pb.Category) []model.Category {
	var categoryList []model.Category
	for i := range categories {
		categoryList = append(categoryList, model.Category{
			CategoryID: categories[i].Id,
			Name:       categories[i].Name,
			Parent:     categories[i].Parent,
		})
	}
	return categoryList
}
