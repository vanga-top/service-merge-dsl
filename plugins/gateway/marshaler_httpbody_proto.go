package gateway

import (
	"google.golang.org/genproto/googleapis/api/httpbody"
)

type HTTPBodyMarshaler struct {
	Marshaler
}

func (h *HTTPBodyMarshaler) ContentType(v interface{}) string {
	if httpBody, ok := v.(*httpbody.HttpBody); ok {
		return httpBody.GetContentType()
	}
	return h.Marshaler.ContentType(v)
}

func (h *HTTPBodyMarshaler) Marshal(v interface{}) ([]byte, error) {
	if httpBody, ok := v.(*httpbody.HttpBody); ok {
		return httpBody.Data, nil
	}
	return h.Marshaler.Marshal(v)
}
