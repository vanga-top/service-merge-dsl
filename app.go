package main

import (
	"dsl/api/config"
	"dsl/instance"
	"dsl/instance/httpserver"
	"fmt"
)

var serverInstance *httpserver.Server

func init() {
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
