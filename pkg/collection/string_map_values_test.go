package collection

import (
	"reflect"
	"testing"
)

func TestStringValues(t *testing.T) {
	type args struct {
		ctx map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test map<string,string> values",
			args: args{
				ctx: map[string]string{
					"hello": "world",
					"tom":   "cat",
					"lilei": "hanmeimei",
				},
			},
			want: []string{"world", "cat", "hanmeimei"},
		},
		{
			name: "Test map<string,string> values",
			args: args{
				ctx: map[string]string{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringValues(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
