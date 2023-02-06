package remote

import "dsl/plugins"

// RemoteServer for service producer which is linked by phone or web
type RemoteServer interface {
	plugins.Plugin
	Provider(serviceName, version, method string, params *RemoteParams) *RemoteResult
}
