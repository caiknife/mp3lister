package fjson

import (
	"io"

	"github.com/pkg/errors"

	"github.com/json-iterator/go"
)

var (
	j = jsoniter.ConfigCompatibleWithStandardLibrary
)

var (
	ErrStructToJson   = errors.New("parse struct to json string error")
	ErrStringToStruct = errors.New("parse json string to struct error")
)

func Marshal(v any) ([]byte, error) {
	return j.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return j.Unmarshal(data, v)
}

func MarshalToString(v any) (string, error) {
	return j.MarshalToString(v)
}

func UnmarshalFromString(data string, v any) error {
	return j.UnmarshalFromString(data, v)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return j.MarshalIndent(v, prefix, indent)
}

func NewDecoder(reader io.Reader) *jsoniter.Decoder {
	return j.NewDecoder(reader)
}

func NewEncoder(writer io.Writer) *jsoniter.Encoder {
	return j.NewEncoder(writer)
}

func Get(data []byte, path ...any) jsoniter.Any {
	return j.Get(data, path...)
}

func Valid(data []byte) bool {
	return j.Valid(data)
}

func ValidFromString(s string) bool {
	data := []byte(s)
	return Valid(data)
}
