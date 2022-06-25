package headermap

import (
	"net/http"
	"reflect"
	"testing"
)

func TestHeaderMap_Add(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   HeaderMap
	}{
		{
			name: "Test add-0",
			fields: fields{
				keyMap: make(map[string]string),
				ctx:    make(http.Header),
			},
			args: args{
				key:   "hello",
				value: "world",
			},
			want: HeaderMap{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
		},
		{
			name: "Test add-1",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			args: args{
				key:   "hello",
				value: "world2",
			},
			want: HeaderMap{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world", "world2"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			if got := hm.Add(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderMap_Clean(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	tests := []struct {
		name   string
		fields fields
		want   HeaderMap
	}{
		{
			name: "Test Clean",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			want: HeaderMap{
				keyMap: map[string]string{},
				ctx:    http.Header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			if got := hm.Clean(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderMap_Get(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		ok     bool
	}{
		{
			name: "Test Get-true",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			args: args{
				key: "hello",
			},
			want: "world",
			ok:   true,
		},
		{
			name: "Test Get-false",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			args: args{
				key: "mut-hello",
			},
			want: "",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			got, got1 := hm.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestHeaderMap_Has(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test Get-true",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			args: args{
				key: "hello",
			},
			want: true,
		},
		{
			name: "Test Get-false",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			args: args{
				key: "mut-hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			if got := hm.Has(tt.args.key); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderMap_Length(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test Length-2",
			fields: fields{
				keyMap: map[string]string{"hello": "", "world": "hello"},
				ctx:    http.Header{"Hello": []string{"world"}, "World": []string{"hello"}},
			},
			want: 2,
		},
		{
			name: "Test Length-1",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			want: 1,
		},
		{
			name: "Test Length-0",
			fields: fields{
				keyMap: map[string]string{},
				ctx:    http.Header{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			if got := hm.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderMap_Put(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   HeaderMap
	}{
		{
			name: "Test Put",
			fields: fields{
				keyMap: make(map[string]string),
				ctx:    make(http.Header),
			},
			args: args{
				key:   "hello",
				value: "world",
			},
			want: HeaderMap{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			if got := hm.Put(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderMap_Remove(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   HeaderMap
	}{
		{
			name: "Test Remove",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			args: args{
				key: "hello",
			},
			want: HeaderMap{
				keyMap: map[string]string{},
				ctx:    http.Header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			if got := hm.Remove(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderMap_Values(t *testing.T) {
	type fields struct {
		keyMap map[string]string
		ctx    http.Header
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		{
			name: "Test Values",
			fields: fields{
				keyMap: map[string]string{"hello": ""},
				ctx:    http.Header{"Hello": []string{"world"}},
			},
			want: http.Header{"Hello": []string{"world"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HeaderMap{
				keyMap: tt.fields.keyMap,
				ctx:    tt.fields.ctx,
			}
			if got := hm.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHeaderMap(t *testing.T) {
	tests := []struct {
		name string
		want HeaderMap
	}{
		{
			name: "Test NewHeaderMap",
			want: HeaderMap{
				keyMap: map[string]string{},
				ctx:    http.Header{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeaderMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeaderMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
