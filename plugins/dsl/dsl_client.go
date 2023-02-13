package dsl

import (
	"dsl/api/dsl"
	"dsl/plugins"
)

type DSLClient struct {
	plugins.Plugin
}

func (client *DSLClient) Merge(request *dsl.Request) *dsl.Result {
	//TODO implement me
	panic("implement me")
}
