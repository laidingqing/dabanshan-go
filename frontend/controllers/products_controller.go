package controllers

import (
	"context"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/dabanshan/common/clients"
	. "github.com/laidingqing/dabanshan/common/controller"
	"github.com/laidingqing/dabanshan/pb"
	"github.com/laidingqing/dabanshan/products/model"
)

// ProductsController user api struct
type ProductsController struct{}

type CategoriesResponse struct {
	Categories []model.Category `json:"categories"`
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

	container.Add(categoryWebService)
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
