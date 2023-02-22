package main

import (
	"dsl/api/config"
	"dsl/instance"
	"dsl/instance/httpserver"
	"flag"
	"fmt"
)

var serverInstance *httpserver.Server

var (
	addr = flag.String("addr", ":9090", "endPoint of service ports")
)

func init() {
	flag.Parse()
	TrapSignals()
}

func main() {
	appConfig, err := config.ApplicationYamlParser("/Users/chenhui/code/service-merge-dsl/application-dev.yaml")
	if err != nil {
		panic(err)
	}
	serverInstance, err := httpserver.NewServer(appConfig)
	ctx := &instance.InstanceCtx{Config: &instance.Config{Env: "DEV"}}
	serverInstance.Start(ctx)
	fmt.Println("start server...")
	serverInstance.ListPlugins()
	//todo add signal discovery

	//
	serverInstance.Wait()
}
