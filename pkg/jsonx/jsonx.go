package jsonx

import (
	"encoding/json"
	"io"

	"github.com/spf13/cast"
)

func String(body any) string {
	data, _ := StringE(body)

	return data
}

func StringE(body any) (string, error) {
	data, err := json.Marshal(body)

	return string(data), err
}

func Bytes(body any) []byte {
	data, _ := BytesE(body)

	return data
}

func BytesE(body any) ([]byte, error) {
	data, err := json.Marshal(body)

	return data, err
}

func Pretty(body any) string {
	data, _ := PrettyE(body)

	return data
}

func PrettyE(body any) (string, error) {
	bytes, err := json.MarshalIndent(body, "", "\t")

	return string(bytes), err
}

// ---------------------------------------------------------------- Decode

func DecodeStruct(reader io.Reader, target any) error {
	if err := json.NewDecoder(reader).Decode(target); err != nil {
		return err
	}

	return nil
}

func UnmarshalStruct(data []byte, structy any) error {
	if err := json.Unmarshal(data, structy); err != nil {
		return err
	}

	return nil
}

func UnmarshalStructWithoutError(data []byte, structy any) {
	_ = json.Unmarshal(data, structy)
}

func ByteToMap(body []byte) map[string]any {
	ctx, _ := ByteToMapE(body)

	return ctx
}

func ByteToMapE(body []byte) (map[string]any, error) {
	return ToMapE(string(body))
}

func ToMap(body string) map[string]any {
	ctx, _ := ToMapE(body)

	return ctx
}

func ToMapE(body string) (map[string]any, error) {
	ctx := make(map[string]any)
	if err := json.Unmarshal([]byte(body), &ctx); err != nil {
		return ctx, err
	}

	return ctx, nil
}

func ByteToStringMap(body []byte) map[string]string {
	ctx, _ := ByteToStringMapE(body)

	return ctx
}

func ByteToStringMapE(body []byte) (map[string]string, error) {
	ctx, err := ToStringMapE(string(body))

	return ctx, err
}

func ToStringMap(body string) map[string]string {
	ctx, _ := ToStringMapE(body)

	return ctx
}

func ToStringMapE(body string) (map[string]string, error) {
	ctx := make(map[string]string)
	temp := make(map[string]any)
	if err := json.Unmarshal([]byte(body), &temp); err != nil {
		return ctx, err
	}
	ctx, err := cast.ToStringMapStringE(temp)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}
