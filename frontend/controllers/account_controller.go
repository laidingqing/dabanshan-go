package controllers

import (
	"context"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	model "github.com/laidingqing/Dabanshan/accounts/model"
	. "github.com/laidingqing/Dabanshan/common/auth"
	. "github.com/laidingqing/Dabanshan/common/controller"
	client "github.com/laidingqing/Dabanshan/frontend/clients"
	"github.com/laidingqing/dabanshan/common/auth"
	"github.com/laidingqing/dabanshan/common/util"
	"github.com/laidingqing/dabanshan/pb"
)

// UsersController user api struct
type AccountsController struct{}

//UserResponse user api response
type AccountResponse struct {
	Account model.Account `json:"account"`
}

var usersWebService *restful.WebService

func (uc AccountsController) userURI() string {
	return APIPrefix() + "/accounts"
}

// Service ..
func (uc AccountsController) Service() *restful.WebService {
	return usersWebService
}

//Register Define routes
func (uc AccountsController) Register(container *restful.Container) {
	usersWebService = new(restful.WebService)
	//usersWebService.Filter(LogRequest)
	usersWebService.
		Path(uc.userURI()).
		Doc("Manage Users").
		ApiVersion(APIVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	usersWebService.Route(usersWebService.POST("/session").To(uc.session).
		// Filter(AuthUser).
		Doc("session a User").
		Operation("session").
		Reads(model.Account{}).
		Writes(auth.AccountCredentials{}))

	usersWebService.Route(usersWebService.POST("").To(uc.create).
		// Filter(AuthUser).
		Doc("Create a User").
		Operation("create").
		Reads(model.Account{}).
		Writes(AccountResponse{}))

	usersWebService.Route(usersWebService.GET("/{user-id}").To(uc.read).
		// Filter(AuthUser).
		Doc("Gets a User").
		Operation("read").
		Param(usersWebService.PathParameter("user-id", "User Name").DataType("string")).
		Writes(AccountResponse{}))

	container.Add(usersWebService)
}

//Create a User
func (uc AccountsController) create(request *restful.Request, response *restful.Response) {
	newUser := new(model.Account)
	err := request.ReadEntity(newUser)
	if err != nil {
		WriteBadRequestError(response)
		return
	}
	rev, err := client.GetAccountClient().CreateAccount(context.Background(), &pb.CreateAccountRequest{Username: newUser.UserName, Password: newUser.Password})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", rev.Account.Id)
	response.WriteHeader(http.StatusCreated)
}

// read user info
func (uc AccountsController) read(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("user-id")
	log.Printf("user-id is %s", userID)
	res, err := client.GetAccountClient().GetByUsername(context.Background(), &pb.GetByUsernameRequest{Username: "asone"})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", "")
	response.WriteEntity(AccountResponse{Account: model.Account{
		UserName: res.Account.Username,
	}})
}

//session user login and session
func (uc AccountsController) session(request *restful.Request, response *restful.Response) {
	acct := new(model.Account)
	err := request.ReadEntity(acct)
	if err != nil {
		WriteBadRequestError(response)
		return
	}
	res, err := client.GetAccountClient().GetByUsername(context.Background(), &pb.GetByUsernameRequest{Username: acct.UserName})
	if err != nil {
		WriteError(err, response)
		return
	}
	diffPassword := util.CalculatePassHash(acct.Password, request.Username)
	if res.Account.Password != diffPassword {
		WriteError(auth.UnauthenticatedError(), response)
		return
	}
}
