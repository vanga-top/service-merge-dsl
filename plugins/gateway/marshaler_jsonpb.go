package gateway

import (
	"bytes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
)

type JSONPb struct {
	protojson.MarshalOptions
	protojson.UnmarshalOptions
}

func (j *JSONPb) Marshal(v interface{}) ([]byte, error) {
	if _, ok := v.(proto.Message); !ok {
		return j.marshalNonProtoField(v)
	}

	var buf bytes.Buffer
	if err := j.marshalTo(&buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (j *JSONPb) Unmarshal(data []byte, v interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (j *JSONPb) NewDecoder(r io.Reader) Decoder {
	//TODO implement me
	panic("implement me")
}

func (j *JSONPb) NewEncoder(w io.Writer) Encoder {
	//TODO implement me
	panic("implement me")
}

func (j *JSONPb) ContentType(_ interface{}) string {
	return "application/json"
}

func (j *JSONPb) marshalNonProtoField(v interface{}) ([]byte, error) {

	return nil, nil
}

func (j *JSONPb) marshalTo(b *bytes.Buffer, v interface{}) error {

	return nil
}
