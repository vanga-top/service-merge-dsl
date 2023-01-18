package httpserver

import (
	"dsl/api/config"
	"dsl/instance"
	"dsl/plugins"
	"errors"
	"sync"
)

// Server implements for Instance
type Server struct {
	Ctx     *instance.InstanceCtx
	Name    string
	Env     string
	Port    int
	Plugins []plugins.Plugin
	sync.Mutex
}

func NewServer(appConfig *config.ApplicationConfig) (*Server, error) {
	if appConfig == nil {
		return nil, errors.New("application config is nil")
	}
	//struct server
	server := &Server{
		Name: appConfig.Name,
		Port: appConfig.Port,
		Env:  appConfig.Env,
	}
	//load plugin
	if appConfig.SLBFragments != nil {

	}

	//server.LoadPlugin()
	return server, nil
}

func (s *Server) ListPlugins() []plugins.Plugin {
	return s.Plugins
}

func (s *Server) LoadPlugin(plugin plugins.Plugin, initImmediately bool) error {
	if plugin == nil || s.contains(plugin) {
		return errors.New("plugin is null or plugin has been registered")
	}
	s.Plugins = append(s.Plugins, plugin)
	if initImmediately && plugin.Status() == plugins.CREATED { //check status
		return plugin.Init()
	}
	return nil
}

func (s *Server) RemovePlugin(pluginID string) error {
	if pluginID == "" {
		return errors.New("id cannot be null")
	}
	if len(s.Plugins) == 0 {
		return nil
	}

	for i, v := range s.Plugins {
		if v.ID() == pluginID {
			s.Plugins = append(s.Plugins[:i], s.Plugins[i+1:]...)
			break
		}
	}
	return nil
}

func (s *Server) Start(ctx *instance.InstanceCtx) error {
	if ctx == nil || ctx.Config == nil {
		return errors.New("ctx or ctx.config is nil")
	}
	return nil
}

func (s *Server) Stop(ctx *instance.InstanceCtx) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Restart(netCtx *instance.InstanceCtx) error {
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
