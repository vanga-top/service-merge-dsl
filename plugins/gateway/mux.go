package gateway

import (
	"context"
	"dsl/plugins/gateway/utils"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"net/http"
	"regexp"
)

type UnescapingMode int

const (
	UnescapingModeLegacy UnescapingMode = iota
	UnescapingModeAllExceptReserved
	UnescapingModeAllExceptSlash
	UnescapingModeAllCharacters
	UnescapingModeDefault = UnescapingModeLegacy
)

var encodedPathSplitter = regexp.MustCompile("(/|%2F)")

type ServeMux struct {
	// handlers maps HTTP method to a list of handlers.
	handlers                  map[string][]handler
	forwardResponseOptions    []func(context.Context, http.ResponseWriter, proto.Message) error
	marshalers                marshalerRegistry
	incomingHeaderMatcher     HeaderMatcherFunc
	outgoingHeaderMatcher     HeaderMatcherFunc
	metadataAnnotators        []func(ctx context.Context, r *http.Request) metadata.MD
	errorHandler              ErrorHandlerFunc
	streamErrorHandler        StreamErrorHandlerFunc
	routingErrorHandler       RoutingErrorHandlerFunc
	disablePathLengthFallback bool
	unescapingMode            UnescapingMode
}

func (s *ServeMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

type ServeMuxOption func(mux *ServeMux)

func NewServeMux(opts ...ServeMuxOption) *ServeMux {
	
	return nil
}

type handler struct {
	pat Pattern
	h   HandlerFunc
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, pathParams map[string]string)

type HeaderMatcherFunc func(string) (string, bool)

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
