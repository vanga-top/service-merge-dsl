package httpserver

import (
	"dsl/api/config"
	"dsl/instance"
	"dsl/plugins"
	"dsl/plugins/gateway"
	"dsl/plugins/slb/eureka"
	"errors"
	"fmt"
	"github.com/SimonWang00/goeureka"
	"net/http"
	"sync"
)

// Server implements for Instance
type Server struct {
	AppConfig *config.ApplicationConfig //init config application.yaml
	Ctx       *instance.InstanceCtx
	Name      string
	Env       string
	Port      int
	Plugins   []plugins.Plugin
	Lock      *sync.Mutex
	Stat      instance.InstanceStatus
	// wg is used to wait for all servers to shut down
	wg *sync.WaitGroup
}

func (s *Server) Wait() {
	s.wg.Wait()
}

func (s *Server) Status() instance.InstanceStatus {
	return s.Stat
}

func NewServer(appConfig *config.ApplicationConfig, ctx *instance.InstanceCtx, initImmediately bool) (*Server, error) {
	if appConfig == nil {
		return nil, errors.New("application config is nil")
	}
	//struct server
	server := &Server{
		Name:      appConfig.ApplicationFragment.Name,
		Port:      appConfig.ApplicationFragment.Port,
		Env:       appConfig.Env,
		wg:        new(sync.WaitGroup),
		AppConfig: appConfig,
		Ctx:       ctx,
	}
	server.parserAppConfig(initImmediately)
	return server, nil
}

func (s *Server) parserAppConfig(initImmediately bool) {
	//load plugin
	if s.AppConfig.SLBFragments != nil {
		slbClient := &eureka.EurekaClient{
			ClientName:   "eureka-dsl-client",
			ClientID:     "eureka-client",
			InstanceName: s.AppConfig.ApplicationFragment.Name,
			Port:         s.AppConfig.ApplicationFragment.Port,
			URL:          s.AppConfig.SLBFragments.Host,
		}
		s.LoadPlugin(slbClient, initImmediately)
	}

	//load gateway plugin
	if s.AppConfig.GatewayFragment != nil {
		// load gateway plugin
		opts := gateway.Options{
			Addr: ":" + s.AppConfig.GatewayFragment.Port,
			RPCServer: gateway.Endpoint{
				Network: s.AppConfig.GatewayFragment.RPCNetwork,
				Addr:    s.AppConfig.GatewayFragment.RPCAddr,
			},
			OpenAPIDir: "examples/hello",
			Handlers:   []gateway.ServiceHandler{},
		}
		// gateway
		g := &gateway.Gateway{
			Opts: opts,
			Ctx:  s.Ctx.Context,
		}
		s.LoadPlugin(g, initImmediately)
	}

}

func (s *Server) ListPlugins() []plugins.Plugin {
	fmt.Println("--------list all registered plugins--------")
	for i, v := range s.Plugins {
		fmt.Printf("%d. PluginID:%s  Plugin-Name:%s \n", i, v.ID(), v.Name())
	}
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

func (s *Server) Start() error {
	if s.Ctx == nil || s.Ctx.Config == nil {
		return errors.New("ctx or ctx.config is nil")
	}
	if s.Stat == instance.RUNNING {
		return errors.New("server has been run already")
	}

	// init plugin
	for i, v := range s.Plugins {
		fmt.Printf("start: %d. PluginID:%s  Plugin-Name:%s \n", i, v.ID(), v.Name())
		err := v.Init()
		if err != nil {
			return err
		}
	}

	// admin 入口
	http.HandleFunc("/", s.handler)
	go http.ListenAndServe(":8000", nil)

	//add instance
	s.wg.Add(1)
	return nil
}

// handle all request
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	s1, s2, err := goeureka.GetInfoWithappName(s.Name)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "local: "+s1+"\r\n register time:"+s2+"\r\n")
	applications, err2 := goeureka.GetServices()
	if err2 != nil {
		fmt.Fprintf(w, err2.Error())
		return
	}

	fmt.Fprintf(w, "register consumer:\r\n")
	for _, app := range applications {
		for _, ins := range app.Instance {
			fmt.Fprintf(w, "app:%s, ip:%s \r\n", ins.App, ins.IpAddr)
		}
	}

}

func (s *Server) Stop() error {
	defer s.wg.Done()
	fmt.Println("stop server:", s)
	for _, plugin := range s.Plugins {
		err := plugin.Destroy()
		if err != nil {
			return err
		}
	}
	return nil
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
