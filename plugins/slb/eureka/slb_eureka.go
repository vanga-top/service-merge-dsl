package eureka

import (
	"dsl/plugins"
	"dsl/plugins/slb"
	"fmt"
	"strconv"
)
import "github.com/SimonWang00/goeureka"

type EurekaClient struct {
	ClientID     string
	ClientName   string
	ClientStatus plugins.Status
}

func (e *EurekaClient) ID() string {
	return e.ClientID
}

func (e *EurekaClient) Name() string {
	return e.ClientName
}

func (e *EurekaClient) Init() error {
	fmt.Println("init eureka client")
	return nil
}

func (e *EurekaClient) Status() plugins.Status {
	return e.ClientStatus
}

func (e *EurekaClient) Connect(url string, port int, namespace string, opts *slb.SLBOptions) *slb.SLBResult {
	config := make(map[string]string)
	config["username"] = opts.Username
	config["password"] = opts.Password
	goeureka.RegisterClient(url, "", namespace, strconv.Itoa(port), "43", config)
	return nil
}

func (e *EurekaClient) DisConnect() *slb.SLBResult {
	//TODO implement me
	panic("implement me")
}
