package instance

import (
	"context"
	"dsl/plugins"
)

type InstanceStatus int

const (
	CREATED InstanceStatus = iota
	INITED
	RUNNING
	STOP
	DESTROY
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
	Status() InstanceStatus
}
