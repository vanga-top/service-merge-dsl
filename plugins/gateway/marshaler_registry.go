package gateway

import (
	"errors"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

const MIMEWildcard = "*"

var (
	acceptHeader      = http.CanonicalHeaderKey("Accept")
	contentTypeHeader = http.CanonicalHeaderKey("Content-Type")

	defaultMarshaler = &HTTPBodyMarshaler{
		Marshaler: &JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}
)

// marshalerRegistry is a mapping from MIME types to Marshalers.
type marshalerRegistry struct {
	mimeMap map[string]Marshaler
}

func (m marshalerRegistry) add(mime string, marshaler Marshaler) error {
	if len(mime) == 0 {
		return errors.New("empty MIME type")
	}
	m.mimeMap[mime] = marshaler
	return nil
}

func makeMarshalerMIMERegistry() marshalerRegistry {
	return marshalerRegistry{
		mimeMap: map[string]Marshaler{
			MIMEWildcard: defaultMarshaler,
		},
	}
}
