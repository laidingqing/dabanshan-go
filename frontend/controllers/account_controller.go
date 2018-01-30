package controllers

import (
	"context"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	model "github.com/laidingqing/dabanshan/accounts/model"
	"github.com/laidingqing/dabanshan/common/auth"
	"github.com/laidingqing/dabanshan/common/clients"
	. "github.com/laidingqing/dabanshan/common/controller"
	"github.com/laidingqing/dabanshan/common/util"
	"github.com/laidingqing/dabanshan/pb"
)

// AccountsController user api struct
type AccountsController struct{}

//AccountResponse user api response
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
		Doc("session a User").
		Operation("session").
		Reads(model.Account{}).
		Writes(auth.AccountCredentials{}))

	usersWebService.Route(usersWebService.POST("").To(uc.create).
		Doc("Create a User").
		Operation("create").
		Reads(model.Account{}).
		Writes(AccountResponse{}))

	usersWebService.Route(usersWebService.GET("/{user-id}").To(uc.read).
		Filter(AuthUser).
		Doc("Gets a User").
		Operation("read").
		Param(usersWebService.PathParameter("user-id", "User Name").DataType("string")).
		Writes(AccountResponse{}))

	usersWebService.Route(usersWebService.POST("/{user-id}/interests").To(uc.createInterests).
		Filter(AuthUser).
		Doc("Post a User's Tags").
		Operation("createTags").
		Param(usersWebService.PathParameter("user-id", "User Name").DataType("string")).
		Writes(AccountResponse{}))

	usersWebService.Route(usersWebService.POST("/{user-id}/follows").To(uc.createFollows).
		Filter(AuthUser).
		Doc("Post a User's Follow").
		Operation("createFollows").
		Param(usersWebService.PathParameter("user-id", "User Name").DataType("string")).
		Writes(AccountResponse{}))

	usersWebService.Route(usersWebService.GET("/{user-id}/follows").To(uc.getFollows).
		Filter(AuthUser).
		Doc("Get a User's Follow").
		Operation("getFollows").
		Param(usersWebService.PathParameter("user-id", "User Name").DataType("string")).
		Writes(AccountResponse{}))

	container.Add(usersWebService)
}

//Create a User
func (uc AccountsController) create(request *restful.Request, response *restful.Response) {
	newUser := new(model.Account)
	err := request.ReadEntity(newUser)
	log.Printf("username: %s", newUser.UserName)
	if err != nil {
		log.Printf("err: %s", err.Error())
		WriteBadRequestError(response)
		return
	}
	rev, err := clients.GetAccountClient().CreateAccount(context.Background(), &pb.CreateAccountRequest{Username: newUser.UserName, Password: newUser.Password})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", rev.Account.Id)
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(AccountResponse{
		Account: model.Account{
			AccountID: rev.Account.Id,
		},
	})
}

// read user info
func (uc AccountsController) read(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("user-id")
	log.Printf("user-id is %s", userID)
	res, err := clients.GetAccountClient().GetByUsername(context.Background(), &pb.GetByUsernameRequest{Username: "asone"})
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
	res, err := clients.GetAccountClient().GetByUsername(context.Background(), &pb.GetByUsernameRequest{Username: acct.UserName})
	if err != nil {
		WriteError(err, response)
		return
	}
	diffPassword := util.CalculatePassHash(acct.Password, acct.UserName)
	log.Printf("db password: %s, login password: %s", res.Account.Password, diffPassword)
	if res.Account.Password != diffPassword {
		WriteError(auth.UnauthenticatedError(), response)
		return
	}
	jwt, err := auth.CreateJWT()
	if err != nil {
		WriteError(err, response)
		return
	}
	//Save current account token.

	response.WriteEntity(auth.AccountCredentials{
		Username:    res.Account.Username,
		AccessToken: jwt,
	})
}

//createTags, update a tags by buyer.
func (uc AccountsController) createInterests(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("user-id")
	var tags []model.Tag
	err := request.ReadEntity(&tags)
	if err != nil {
		WriteBadRequestErrorInfo(response, err)
		return
	}
	res, err := clients.GetAccountClient().CreateInterests(context.Background(), &pb.CreateTagsRequest{
		Accountid: userID,
		Tags:      DecodeTags(tags),
	})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(res.Created)
}

//createFollows create a follow by account
func (uc AccountsController) createFollows(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("user-id")
	followID := request.QueryParameter("followId")
	_, err := clients.GetAccountClient().FollowUser(context.Background(), &pb.FollowUserRequest{Accountid: userID, Followid: followID})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.WriteHeader(http.StatusCreated)
}

//getFollows get follows by account
func (uc AccountsController) getFollows(request *restful.Request, response *restful.Response) {
	userID := request.PathParameter("user-id")
	res, err := clients.GetAccountClient().GetFollows(context.Background(), &pb.GetFollowsRequest{Accountid: userID})
	if err != nil {
		WriteError(err, response)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.WriteEntity(EncodeAccounts(res.Accounts))
}
