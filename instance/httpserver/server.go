package httpserver

import (
	"dsl/instance"
	"dsl/plugins"
	"errors"
	"sync"
)

// Server implements for Instance
type Server struct {
	Ctx     *instance.InstanceCtx
	Name    string
	Port    int
	Plugins []plugins.Plugin
	sync.Mutex
}

func (s *Server) ListPlugins() []plugins.Plugin {
	return s.Plugins
}

func (s *Server) LoadPlugin(plugin plugins.Plugin, initImmediately bool) error {
	if plugin == nil || s.contains(plugin) {
		return errors.New("plugin is null or plugin has been registered")
	}
	s.Plugins = append(s.Plugins, plugin)
	if initImmediately {
		plugin.Init()
	}
	return nil
}

func (s *Server) RemovePlugin(pluginID string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Start(ctx *instance.InstanceCtx) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Stop(ctx *instance.InstanceCtx) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Restart(netCtx *instance.InstanceCtx) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) contains(plugin plugins.Plugin) bool {
	size := len(s.Plugins)
	for i := 0; i < size; i++ {
		p := s.Plugins[i]
		if p.ID() == plugin.ID() {
			return true
		}
	}
	return false
}
