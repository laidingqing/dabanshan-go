package config

import (
	"log"

	"github.com/laidingqing/dabanshan/common/util"
)

var (
	//ServiceAccountName 账号微服务名称
	ServiceAccountName = "account_service"
	//ServiceEpisodeName 买卖需求微服务名称
	ServiceEpisodeName = "episode_service"
	//ServiceNotificationName 通知服务
	ServiceNotificationName = "notification_service"
	//ServiceProductsName 商品服务
	ServiceProductsName = "products_service"
)

//APIVersion api prefix
var APIVersion = "v1"

//Service 微服务Registry配置
var Service struct {
	DomainName       string
	ServiceName      string
	Port             int
	RegistryLocation string
}

//Logger 微服务日志配置
var Logger struct {
	LogFile    string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

//Database 数据库配置
var Database struct {
	HostURI      string
	DatabaseName string
}

// LoadDefaults Initialize Default values
func LoadDefaults() {
	_, err := util.GetExecDirectory()
	if err != nil {
		log.Fatal(err)
	}

	Logger.LogFile = "service.log"
	Logger.MaxSize = 10
	Logger.MaxBackups = 3
	Logger.MaxAge = 30

	Database.HostURI = "localhost:27017"
	Database.DatabaseName = "dabanshan"
}
