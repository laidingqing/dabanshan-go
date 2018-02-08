package controllers

import (
	"context"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/dabanshan/common/clients"
	. "github.com/laidingqing/dabanshan/common/controller"
	"github.com/laidingqing/dabanshan/pb"
	"github.com/laidingqing/dabanshan/products/model"
)

// ProductsController user api struct
type ProductsController struct{}

type CategoriesResponse struct {
	Categories []model.Category `json:"categories,omitempty"`
}

//ProductsResponse ..
type ProductsResponse struct {
	Products []model.ProductItem `json:"products,omitempty"`
}

//ProductResponse ...
type ProductResponse struct {
	Product model.ProductItem `json:"product,omitempty"`
	Error   error             `json:"error,omitempty"`
}

var categoryWebService *restful.WebService
var productsWebService *restful.WebService

func (pc ProductsController) categoryURI() string {
	return APIPrefix() + "/categories"
}

func (pc ProductsController) productsURI() string {
	return APIPrefix() + "/products"
}

//Register Define routes
func (pc ProductsController) Register(container *restful.Container) {

	categoryWebService = new(restful.WebService)
	categoryWebService.Filter(LogRequest)
	categoryWebService.
		Path(pc.categoryURI()).
		Doc("Manage Category").
		ApiVersion(APIVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	categoryWebService.Route(categoryWebService.GET("").To(pc.findCategories).
		// Filter(AuthUser).
		Doc("Get All Category").
		Operation("findCategories").
		Reads(model.Category{}).
		Writes(CategoriesResponse{}))

	productsWebService = new(restful.WebService)
	productsWebService.Filter(LogRequest)
	productsWebService.
		Path(pc.productsURI()).
		Doc("Manage Product Item").
		ApiVersion(APIVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	productsWebService.Route(productsWebService.POST("").To(pc.createProductItem).
		// Filter(AuthUser).
		Doc("Post Product Item").
		Operation("createProductItem").
		Reads(model.ProductItem{}).
		Writes(ProductResponse{}))

	productsWebService.Route(productsWebService.GET("").To(pc.findProducts).
		// Filter(AuthUser).
		Doc("Get All Product Items").
		Operation("findProducts").
		Reads(model.ProductItem{}).
		Writes(ProductsResponse{}))

	container.Add(categoryWebService)
	container.Add(productsWebService)
}

func (pc ProductsController) findCategories(request *restful.Request, response *restful.Response) {
	parentID := request.QueryParameter("parent")
	res, err := clients.GetProductsClient().FindCategories(context.Background(), &pb.FindCategoryRequest{Parent: parentID})
	if err != nil {
		WriteError(err, response)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.WriteEntity(CategoriesResponse{
		Categories: EncodeCategories(res.Categories),
	})

}

// product api ::on

func (pc ProductsController) findProducts(request *restful.Request, response *restful.Response) {
	// TODO pagination
	storeID := request.QueryParameter("storeId")
	categoryID := request.QueryParameter("categoryId")
	keyword := request.QueryParameter("keyword")
	offset, _ := strconv.ParseInt(request.QueryParameter("offset"), 10, 64)
	limit, _ := strconv.ParseInt(request.QueryParameter("limit"), 10, 64)
	res, err := clients.GetProductsClient().FindProductItems(context.Background(), &pb.FindProductItemRequest{
		Storeid:    storeID,
		Categoryid: categoryID,
		Keyword:    keyword,
		Offset:     offset,
		Limit:      limit,
	})
	if err != nil {
		WriteError(err, response)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.WriteEntity(ProductsResponse{
		Products: EncodeProducts(res.Products),
	})
}

func (pc ProductsController) createProductItem(request *restful.Request, response *restful.Response) {
	newProductItem := new(model.ProductItem)
	err := request.ReadEntity(newProductItem)
	if err != nil {
		WriteBadRequestError(response)
		return
	}

	err = newProductItem.ValidateProductItemRequired()
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.WriteEntity(ProductResponse{
			Error: err,
		})
		return
	}

	res, err := clients.GetProductsClient().CreateProductItem(context.Background(), &pb.CreateProductItemRequest{
		Product: DencodeProduct(*newProductItem),
	})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(ProductResponse{
		Product: model.ProductItem{ProductID: res.Id},
	})
}
