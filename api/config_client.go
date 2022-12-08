package api

type ConfigClient interface {
	Connect(host string, port int, namespace string, token string) *ConfigResult
	DisConnect() *ConfigResult
}

type ConfigResult struct {
}
