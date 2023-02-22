package gateway

type Endpoint struct {
	Network, Addr string
}

type Options struct {
	Addr       string
	RPCServer  Endpoint
	OpenAPIDir string
	//Mux []
}
