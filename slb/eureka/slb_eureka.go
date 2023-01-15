package eureka

import (
	"dsl/slb"
	"strconv"
)
import "github.com/SimonWang00/goeureka"

type EurekaClient struct {
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
