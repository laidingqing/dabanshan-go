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

//EncodeProducts ...
func EncodeProducts(request []*pb.Product) []model.ProductItem {
	var products []model.ProductItem
	for i := range request {
		products = append(products, model.ProductItem{
			Name: request[i].Name,
			SKU:  request[i].Sku,
		})
	}
	return products
}

//DencodeProduct ...
func DencodeProduct(request model.ProductItem) *pb.Product {
	return &pb.Product{
		Name: request.Name,
		Sku:  request.SKU,
	}
}
