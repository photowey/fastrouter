package collection

import (
	"testing"
)

func TestStringKeys(t *testing.T) {
	type args struct {
		ctx map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test map<string,string> keys",
			args: args{
				ctx: map[string]string{
					"hello": "world",
					"tom":   "cat",
					"lilei": "hanmeimei",
				},
			},
			want: []string{"hello", "tom", "lilei"},
		},
		{
			name: "Test map<string,string> keys",
			args: args{
				ctx: map[string]string{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringKeys(tt.args.ctx); !DeepEqual(got, tt.want) {
				t.Errorf("StringKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func DeepEqual(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}

OUTER:
	for _, item := range got {
		for _, wantItem := range want {
			if item == wantItem {
				continue OUTER
			}
		}

		return false
	}

	return true
}
