package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/dabanshan/common/auth"
	"github.com/laidingqing/dabanshan/common/config"
	couchdb "github.com/rhinoman/couchdb-go"
)

//Controller interface
type Controller interface {
	Service() *restful.WebService
}

//APIVersion api version
func APIVersion() string {
	return config.APIVersion
}

//APIPrefix .. api prefix
func APIPrefix() string {
	return "/api/" + APIVersion()
}

//LogError Log an error
func LogError(request *restful.Request, resp *restful.Response, err error) {
	method := request.Request.Method
	url := request.Request.URL.String()
	remoteAddr := request.Request.RemoteAddr
	log.Printf("[ERROR] %v : %v : %v %v", err, remoteAddr, method, url)
}

//WriteBadRequestError write bad request error
func WriteBadRequestError(response *restful.Response) {
	log.Printf("400: Bad Request")
	response.WriteErrorString(http.StatusBadRequest, "Bad Request")
}

//WriteBadRequestErrorInfo ...
func WriteBadRequestErrorInfo(response *restful.Response, err error) {
	log.Printf("400: Bad Request, %s", err.Error())
	response.WriteErrorString(http.StatusBadRequest, err.Error())
}

//WriteError Writes and logs errors from the couchdb driver
func WriteError(err error, response *restful.Response) {
	var statusCode int
	var reason string = "error"
	//Is this a couchdb error?
	cErr, ok := err.(*couchdb.Error)
	if ok { // Yes!
		statusCode = cErr.StatusCode
		reason = cErr.Reason
	} else { // No, try to parse :(
		str := err.Error()
		errStrings := strings.Split(str, ":")
		statusCode := 0
		var cErr error
		if len(errStrings) > 1 {
			statusCode, cErr = strconv.Atoi(errStrings[1])
			reason = http.StatusText(statusCode)
		}
		if cErr != nil || statusCode == 0 {
			statusCode = 500
		}
	}
	//Write the error to the response
	response.WriteErrorString(statusCode, reason)
	//Log the error
	log.Printf("%v", err)
}

//LogRequest Filter function.  Logs incoming requests
func LogRequest(request *restful.Request, resp *restful.Response,
	chain *restful.FilterChain) {
	method := request.Request.Method
	url := request.Request.URL.String()
	remoteAddr := request.Request.RemoteAddr
	log.Printf("[API] %v : %v %v", remoteAddr, method, url)
	chain.ProcessFilter(request, resp)
}

//Unauthenticated respone a unauthenticated error
func Unauthenticated(request *restful.Request, response *restful.Response) {
	LogError(request, response, errors.New("Unauthenticated"))
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(401, "Unauthenticated")
}

//AuthUser 受保护的请求
func AuthUser(request *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	//TODO ..
	err := auth.ValidateTokenMiddleware(request)
	if err != nil {
		Unauthenticated(request, resp)
		return
	}
	chain.ProcessFilter(request, resp)
}
