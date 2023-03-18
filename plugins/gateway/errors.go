package gateway

import (
	"context"
	"google.golang.org/grpc/status"
	"net/http"
)

type ErrorHandlerFunc func(ctx context.Context, mux *ServeMux, w http.ResponseWriter, r *http.Request, err error)

// StreamErrorHandlerFunc is the signature used to configure stream error handling.
type StreamErrorHandlerFunc func(context.Context, error) *status.Status

// RoutingErrorHandlerFunc is the signature used to configure error handling for routing errors.
type RoutingErrorHandlerFunc func(context.Context, *ServeMux, Marshaler, http.ResponseWriter, *http.Request, int)
