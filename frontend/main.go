package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful"
	"github.com/laidingqing/dabanshan/common/config"
	controllers "github.com/laidingqing/dabanshan/frontend/controllers"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	config.LoadDefaults()
	config.ParseCmdParams(config.DefaultCmdLine{
		HostName:         "localhost",
		Port:             7701,
		ServiceName:      "front_end",
		RegistryLocation: "http://127.0.0.1:2379",
	})
	// Set up the core logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/frontend.log",
		MaxSize:    config.Logger.MaxSize,
		MaxBackups: config.Logger.MaxBackups,
		MaxAge:     config.Logger.MaxAge,
	})
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	wsContainer.EnableContentEncoding(true)
	uc := controllers.AccountsController{}
	uc.Register(wsContainer)
	httpAddr := ":" + strconv.Itoa(config.Service.Port)
	log.Printf("starting frontend at %s", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, wsContainer))

}
