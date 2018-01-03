package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/dabanshan/common/config"
	controllers "github.com/laidingqing/dabanshan/frontend/controllers"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	serv = flag.String("service", "user_service", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
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
	uc := controllers.UsersController{}
	uc.Register(wsContainer)
	httpAddr := ":" + strconv.Itoa(config.Service.Port)
	log.Fatal(http.ListenAndServe(httpAddr, wsContainer))
	// flag.Parse()
	// r := grpclb.NewResolver(*serv)
	// b := grpc.RoundRobin(r)
	//
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
	// if err != nil {
	// 	panic(err)
	// }
	//
	// ticker := time.NewTicker(1 * time.Second)
	// for t := range ticker.C {
	// 	client := pb.NewUserServiceClient(conn)
	// 	resp, err := client.GetUser(context.Background(), &pb.GetUserRequest{})
	// 	if err == nil {
	// 		fmt.Printf("Reply is %s, %s\n", resp.User.Username, t)
	// 	}
	// }
}
