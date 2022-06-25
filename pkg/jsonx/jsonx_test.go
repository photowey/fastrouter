package jsonx

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

// Book
type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

var jsonData = `{
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "authors": [
    "Alan A.A.Donovan",
    "Brian W. Kergnighan"
  ],
  "press": "Pearson Education"
}`

func TestToStruct(t *testing.T) {
	type args struct {
		data   []byte
		target any
	}
	book := &Book{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test json string to struct(Unmarshal)",
			args: args{
				data:   []byte(jsonData),
				target: book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnmarshalStruct(tt.args.data, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToStructd(t *testing.T) {
	type args struct {
		reader io.Reader
		target any
	}
	book := &Book{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test json string to struct(Decode)",
			args: args{
				reader: strings.NewReader(jsonData),
				target: book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecodeStruct(tt.args.reader, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("DecodeStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var (
	apiErrorBody = `{
  "code": "9787111558422",
  "message": "I'm full message"
}`
	apiErrorShortMessage = `{
  "code": "9787111558422",
  "msg": "I'm short message"
}`
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message" json:"msg"` // 解析失败
}

func TestToStructWithoutError(t *testing.T) {
	type args struct {
		data   []byte
		target any
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test deserialize json without error",
			args: args{
				data:   []byte(apiErrorBody),
				target: &APIError{},
			},
		},
		{
			name: "Test deserialize json without error",
			args: args{
				data:   []byte(apiErrorShortMessage),
				target: &APIError{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnmarshalStructWithoutError(tt.args.data, tt.args.target)
		})
	}
}

var mapData = `{
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "press": "Pearson Education"
}`

var badMapData = `
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "press": "Pearson Education"
}`

var mapInt64Data = `{
  "id": 9787111558422,
  "name": "The Go Programming Language",
  "press": "Pearson Education"
}`

func TestToMap(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "Test unmarshal json to map",
			args: args{
				body: mapData,
			},
			want: map[string]any{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToMap(tt.args.body)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapE(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "Test unmarshal json-byte to map<string,any>",
			args: args{
				body: mapData,
			},
			want: map[string]any{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
			wantErr: false,
		},
		{
			name: "Test unmarshal json-byte to map<string,any>",
			args: args{
				body: badMapData,
			},
			want:    map[string]any{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMapE(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMapE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToMap(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "Test unmarshal json to map",
			args: args{
				body: []byte(mapData),
			},
			want: map[string]any{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToMap(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToStringMap(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test unmarshal json-byte to map<string,string>",
			args: args{
				body: []byte(mapData),
			},
			want: map[string]string{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToStringMap(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteToStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToStringMapE(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Test unmarshal json-byte to map<string,string>",
			args: args{
				body: []byte(mapData),
			},
			want: map[string]string{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
			wantErr: false,
		},
		{
			name: "Test unmarshal json-byte to map<string,string>",
			args: args{
				body: []byte(badMapData),
			},
			want:    map[string]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ByteToStringMapE(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteToStringMapE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteToStringMapE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringMap(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test unmarshal json to map<string,string>",
			args: args{
				body: mapData,
			},
			want: map[string]string{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
		},
		{
			name: "Test unmarshal json to map<string,string>",
			args: args{
				body: mapInt64Data,
			},
			want: map[string]string{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringMap(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringMapE(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Test unmarshal json to map<string,string>",
			args: args{
				body: mapData,
			},
			want: map[string]string{
				"id":    "9787111558422",
				"name":  "The Go Programming Language",
				"press": "Pearson Education",
			},
			wantErr: false,
		},
		{
			name: "Test unmarshal json to map<string,string>",
			args: args{
				body: badMapData,
			},
			want:    map[string]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringMapE(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringMapE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
