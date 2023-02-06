package remote

// Router for http to remote call
type Router interface {
	Route(requestURI string, params []interface{}) chan interface{}
}
