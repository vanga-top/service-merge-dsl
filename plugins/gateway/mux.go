package gateway

import (
	"context"
	"dsl/plugins/gateway/utils"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type ServerMux struct {
	// handlers maps HTTP method to a list of handlers.
	handlers               map[string][]handler
	forwardResponseOptions []func(context.Context, http.ResponseWriter, proto.Message) error
}

type handler struct {
	pat Pattern
	h   HandlerFunc
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, pathParams map[string]string)

type Pattern struct {
	ops []op
	// pool is a constant pool indexed by the operands or vars.
	pool []string

	vars []string

	stackSize int

	tailLen int

	// verb is the VERB part of the path pattern. It is empty if the pattern does not have VERB part.
	verb string
}

type op struct {
	code    utils.OpCode
	operand int
}
