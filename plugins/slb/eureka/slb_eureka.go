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
	*slb.SLBOptions
	URL          string
	Port         int
	Namespace    string
	InstanceName string
}

func (e *EurekaClient) Destroy() error {
	fmt.Printf("destroy plugin Name:%s\n", e.ClientName)
	e.DisConnect()
	return nil
}

func (e *EurekaClient) ID() string {
	return e.ClientID
}

func (e *EurekaClient) Name() string {
	return e.ClientName
}

func (e *EurekaClient) Init() error {
	fmt.Println("init eureka client")
	e.Connect(e.URL, e.Port, e.InstanceName, nil)
	return nil
}

func (e *EurekaClient) Status() plugins.Status {
	return e.ClientStatus
}

func (e *EurekaClient) Connect(url string, port int, appName string, opts *slb.SLBOptions) *slb.SLBResult {
	config := make(map[string]string)
	if opts != nil {
		config["username"] = opts.Username
		config["password"] = opts.Password
	}
	go goeureka.RegisterClient(url, "", appName, strconv.Itoa(port), "43", config)
	e.ClientStatus = plugins.RUNNING //change status
	return nil
}

func (e *EurekaClient) DisConnect() *slb.SLBResult {
	return nil
}
