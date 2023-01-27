package main

import (
	"dsl/api/config"
	"dsl/instance"
	"dsl/instance/httpserver"
	"fmt"
)

func main() {

	appConfig, err := config.ApplicationYamlParser("/Users/chenhui/code/service-merge-dsl/application-dev.yaml")
	if err != nil {
		panic(err)
	}
	server, err := httpserver.NewServer(appConfig)
	ctx := &instance.InstanceCtx{Config: &instance.Config{Env: "DEV"}}
	server.Start(ctx)
	fmt.Println("start server...")
	server.ListPlugins()
	//todo add signal discovery
	server.Wait()
}
