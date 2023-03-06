package gateway

import "io"

type Marshaler interface {
	// Marshal marshals "v" into byte sequence.
	Marshal(v interface{}) ([]byte, error)

	Unmarshal(data []byte, v interface{}) error

	NewDecoder(r io.Reader) Decoder

	NewEncoder(w io.Writer) Encoder

	ContentType(v interface{}) string
}

type Decoder interface {
	Decode(v interface{}) error
}

type Encoder interface {
	Encode(v interface{}) error
}

type DecoderFunc func(v interface{}) error

func (f DecoderFunc) Decode(v interface{}) error {
	return f(v)
}

type EncoderFunc func(v interface{}) error

func (e EncoderFunc) Encode(v interface{}) error {
	return e(v)
}

type Delimited interface {
	// Delimiter returns the record separator for the stream.
	Delimiter() []byte
}
