package controllers

import (
	"context"
	"net/http"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/dabanshan/common/controller"
	"github.com/laidingqing/dabanshan/episodes/model"
	client "github.com/laidingqing/dabanshan/frontend/clients"
)

// EpisodesController user api struct
type EpisodesController struct{}

//EpisodeResponse user api response
type EpisodeResponse struct {
	Episode model.Episode `json:"episode"`
}

var episodesWebService *restful.WebService

func (ec EpisodesController) episodesURI() string {
	return APIPrefix() + "/episodes"
}

// Service ..
func (ec EpisodesController) Service() *restful.WebService {
	return episodesWebService
}

//Register Define routes
func (ec EpisodesController) Register(container *restful.Container) {
	episodesWebService = new(restful.WebService)
	episodesWebService.Filter(LogRequest)
	episodesWebService.
		Path(ec.episodesURI()).
		Doc("Manage Users").
		ApiVersion(APIVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	episodesWebService.Route(episodesWebService.POST("").To(ec.create).
		Filter(AuthUser).
		Doc("Create a User").
		Operation("create").
		Reads(model.Episode{}).
		Writes(EpisodeResponse{}))

	episodesWebService.Route(episodesWebService.GET("/{episode-id}").To(ec.read).
		Filter(AuthUser).
		Doc("Gets a User").
		Operation("read").
		Param(episodesWebService.PathParameter("episode-id", "Episode ID").DataType("string")).
		Writes(EpisodeResponse{}))

	container.Add(episodesWebService)
}

//create 创建一个供应需求，角色为卖方
func (ec EpisodesController) create(request *restful.Request, response *restful.Response) {
	newEpisode := new(model.Episode)
	err := request.ReadEntity(newEpisode)
	if err != nil {
		WriteBadRequestError(response)
		return
	}
	rev, err := client.GetEpisodeClient().CreateEpisode(context.Background(), DecodeEpisode(*newEpisode))
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", rev.Episode.Id)
	response.WriteHeader(http.StatusCreated)
}

//read 获取已发布的供应需求信息, return a Episode entry
func (ec EpisodesController) read(request *restful.Request, response *restful.Response) {

	response.AddHeader("ETag", "")
	response.WriteHeader(http.StatusOK)
}
