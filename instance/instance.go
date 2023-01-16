package instance

import (
	"context"
	"dsl/plugins"
)

type InstanceCtx struct {
	context.Context
}

type Instance interface {
	LoadPlugin(plugin *plugins.Plugin) error
	RemovePlugin(pluginID string) error
	Start(ctx *InstanceCtx)
	Stop(ctx *InstanceCtx)
	Restart(netCtx *InstanceCtx)
}
