package instance

import (
	"context"
	"dsl/plugins"
)

type Config struct {
	Env     string
	Name    string
	GroupID string
	Port    int
}

type InstanceCtx struct {
	Config *Config
	context.Context
}

type Instance interface {
	LoadPlugin(plugin plugins.Plugin, initImmediately bool) error
	ListPlugins() []plugins.Plugin
	RemovePlugin(pluginID string) error
	Start(ctx *InstanceCtx) error
	Stop(ctx *InstanceCtx) error
	Restart(netCtx *InstanceCtx) error
}
