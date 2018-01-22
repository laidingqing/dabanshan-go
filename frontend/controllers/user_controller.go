package controllers

import (
	"context"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/Dabanshan/common/controller"
	client "github.com/laidingqing/Dabanshan/frontend/clients"
	model "github.com/laidingqing/Dabanshan/users/model"
	"github.com/laidingqing/dabanshan/pb"
)

// UsersController user api struct
type UsersController struct{}

//UserResponse user api response
type UserResponse struct {
	User model.User `json:"user"`
}

var usersWebService *restful.WebService

func (uc UsersController) userURI() string {
	return APIPrefix() + "/users"
}

// Service ..
func (uc UsersController) Service() *restful.WebService {
	return usersWebService
}

//Register Define routes
func (uc UsersController) Register(container *restful.Container) {
	usersWebService = new(restful.WebService)
	// usersWebService.Filter(LogRequest)
	usersWebService.
		Path(uc.userURI()).
		Doc("Manage Users").
		ApiVersion(APIVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	usersWebService.Route(usersWebService.POST("").To(uc.create).
		// Filter(AuthUser).
		Doc("Create a User").
		Operation("create").
		Reads(model.User{}).
		Writes(UserResponse{}))

	usersWebService.Route(usersWebService.GET("/{user-id}").To(uc.read).
		// Filter(AuthUser).
		Doc("Gets a User").
		Operation("read").
		Param(usersWebService.PathParameter("user-id", "User Name").DataType("string")).
		Writes(UserResponse{}))

	container.Add(usersWebService)
}

//Create a User
func (uc UsersController) create(request *restful.Request, response *restful.Response) {
	newUser := new(model.User)
	err := request.ReadEntity(newUser)
	if err != nil {
		WriteBadRequestError(response)
		return
	}
	rev, err := client.GetClient().CreateUser(context.Background(), &pb.CreateUserRequest{Username: newUser.UserName, Password: newUser.Password})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", rev.User.Id)
	response.WriteHeader(http.StatusCreated)
}

// read user info
func (uc UsersController) read(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("user-id")
	log.Printf("user-id is %s", userID)
	res, err := client.GetClient().GetByUsername(context.Background(), &pb.GetByUsernameRequest{Username: "asone"})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", "")
	response.WriteEntity(UserResponse{User: model.User{
		UserName: res.User.Username,
	}})
}
