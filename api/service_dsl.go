package api

type ServiceDSL interface {
	// Call remote rpc call
	Merge([]interface{}) interface{}
}
