package remote

import "dsl/plugins"

// remote caller

type RemoteClient interface {
	plugins.Plugin
	Call(serviceName, version, method string, params *RemoteParams) *RemoteResult
}
