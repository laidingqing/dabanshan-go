package service

import (
	"context"

	"github.com/laidingqing/dabanshan/pb"
	"github.com/laidingqing/dabanshan/products/mongo"
)

// RPCProductsServer is used to implement product_service.RPCProductsServer.
type RPCProductsServer struct{}

var (
	categoryManager    = mongo.NewCagegoryManager()
	productLibManager  = mongo.NewProductLibManager()
	productItemManager = mongo.NewProductItemManager()
	activityManager    = mongo.NewActivityManager()
)

//FindCategories find categories by parent.
func (s *RPCProductsServer) FindCategories(ctx context.Context, in *pb.FindCategoryRequest) (*pb.FindCategoryResponse, error) {
	enties, err := categoryManager.FindAll(in.Parent)
	if err != nil {
		return &pb.FindCategoryResponse{}, err
	}
	return &pb.FindCategoryResponse{
		Categories: encodeCategoriesInfo(enties),
	}, nil
}

//FindProductsByLib find match product libs
func (s *RPCProductsServer) FindProductsByLib(ctx context.Context, in *pb.FindProductLibRequest) (*pb.FindProductLibResponse, error) {
	enties, err := productLibManager.Find(in.Word)
	if err != nil {
		return &pb.FindProductLibResponse{}, err
	}
	return &pb.FindProductLibResponse{
		Products: encodeProductLib(enties),
	}, nil
}

//FindActivities find all activity list
func (s *RPCProductsServer) FindActivities(ctx context.Context, in *pb.FindActivitiesRequest) (*pb.FindActivitiesResponse, error) {

	return &pb.FindActivitiesResponse{}, nil
}

//CreateProductItem create a sale item
func (s *RPCProductsServer) CreateProductItem(ctx context.Context, in *pb.CreateProductItemRequest) (*pb.CreateProductItemResponse, error) {
	productItem := decodeProductItem(in.Product)
	rev, err := productItemManager.Insert(productItem)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductItemResponse{
		Id: rev,
	}, nil
}

//FindProductItems find all product list
func (s *RPCProductsServer) FindProductItems(ctx context.Context, in *pb.FindProductItemRequest) (*pb.FindProductItemResponse, error) {
	revs, err := productItemManager.FindAll(in.Keyword, in.Storeid, in.Categoryid, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}
	return &pb.FindProductItemResponse{
		Products: encodeProductItem(revs),
	}, nil
}
