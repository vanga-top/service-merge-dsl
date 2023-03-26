package gateway

import (
	"context"
	"dsl/plugins/gateway/utils"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"net/http"
	"net/textproto"
	"regexp"
	"strings"
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

func (s *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s1 := "muxServer_v1"
	s2 := "now"
	fmt.Fprintf(w, "mux: "+s1+"\r\n register time:"+s2+"\r\n")
}

type ServeMuxOption func(mux *ServeMux)

// todo
func NewServeMux(opts ...ServeMuxOption) *ServeMux {
	serveMux := &ServeMux{
		handlers:               make(map[string][]handler),
		forwardResponseOptions: make([]func(context.Context, http.ResponseWriter, proto.Message) error, 0),
		marshalers:             makeMarshalerMIMERegistry(),
		errorHandler:           nil,
		streamErrorHandler:     nil,
		routingErrorHandler:    nil,
		unescapingMode:         UnescapingModeDefault,
	}

	for _, opt := range opts {
		opt(serveMux)
	}

	if serveMux.incomingHeaderMatcher == nil {

	}

	return serveMux
}

func DefaultHeaderMatcher(key string) (string, bool) {
	switch key = textproto.CanonicalMIMEHeaderKey(key); {
	case isPermanentHTTPHeader(key):
		return MetadataPrefix + key, true
	case strings.HasPrefix(key, MetadataHeaderPrefix):
		return key[len(MetadataHeaderPrefix):], true
	}
	return "", false
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
