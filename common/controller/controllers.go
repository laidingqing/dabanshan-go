package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/common/config"
	couchdb "github.com/rhinoman/couchdb-go"
)

//Controller interface
type Controller interface {
	Service() *restful.WebService
}

//APIVersion api version
func APIVersion() string {
	return config.ApiVersion
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
