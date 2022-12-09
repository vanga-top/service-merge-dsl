package eureka

import (
	"dsl/api"
	"strconv"
)
import "github.com/SimonWang00/goeureka"

type EurekaClient struct {
}

func (e *EurekaClient) Connect(url string, port int, namespace string, opts *api.SLBOptions) *api.SLBResult {
	config := make(map[string]string)
	config["username"] = opts.Username
	config["password"] = opts.Password
	goeureka.RegisterClient(url, "", namespace, strconv.Itoa(port), "43", config)
	return nil
}

func (e *EurekaClient) DisConnect() *api.SLBResult {
	//TODO implement me
	panic("implement me")
}
