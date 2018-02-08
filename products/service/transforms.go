package service

import (
	"github.com/laidingqing/dabanshan/pb"
	"github.com/laidingqing/dabanshan/products/model"
)

func encodeCategoriesInfo(request []model.Category) []*pb.Category {
	var categories []*pb.Category
	for i := range request {
		categories = append(categories, &pb.Category{
			Id:     request[i].ID.Hex(),
			Name:   request[i].Name,
			Parent: request[i].Parent,
		})
	}
	return categories
}

func encodeProductLib(request []model.ProductLib) []*pb.Product {
	var products []*pb.Product
	for i := range request {
		products = append(products, &pb.Product{
			Id:   request[i].ID.Hex(),
			Name: request[i].Name,
		})
	}
	return products
}

func decodeProductItem(request *pb.Product) model.ProductItem {
	return model.ProductItem{
		Name: request.Name,
	}
}

func encodeProductItem(request []model.ProductItem) []*pb.Product {
	var products []*pb.Product
	for i := range request {
		products = append(products, &pb.Product{
			Id:   request[i].ID.Hex(),
			Name: request[i].Name,
		})
	}
	return products
}
