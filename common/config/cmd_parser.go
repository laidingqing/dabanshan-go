package config

import (
	"flag"
)

//DefaultCmdLine default cmd flag struct
type DefaultCmdLine struct {
	HostName         string
	ServiceName      string
	Port             int
	RegistryLocation string
}

//ParseCmdParams parse from flag.
func ParseCmdParams(defaults DefaultCmdLine) {
	hostName := flag.String("hostName", defaults.HostName, "The host name for this instance")
	port := flag.Int("port", defaults.Port, "The port number for this instance")
	registryLocation := flag.String("registryLocation", defaults.RegistryLocation, "URL for etcd")
	ServiceName := flag.String("serviceName", defaults.ServiceName, "The node Id for this instance")
	flag.Parse()
	Service.DomainName = *hostName
	Service.Port = *port
	Service.RegistryLocation = *registryLocation
	Service.ServiceName = *ServiceName
}
