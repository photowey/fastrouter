package nanoid

import (
	"fmt"
	"regexp"
	"testing"
	"unicode"
)

func TestNew(t *testing.T) {
	type args struct {
		sizes []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test nanoid in golang",
			args: args{
				sizes: []int{32},
			},
			want:    32,
			wantErr: false,
		},
		{
			name: "Test nanoid in golang-default",
			args: args{
				sizes: []int{},
			},
			want:    21,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.sizes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type test func(src string, len int) bool

func TestGenerate(t *testing.T) {
	type args struct {
		alphabet string
		sizes    []int
		test     test
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test nanoid in golang",
			args: args{
				alphabet: "0123456789",
				sizes:    []int{32},
				test:     isDigit,
			},
			want:    32,
			wantErr: false,
		},
		{
			name: "Test nanoid in golang-default",
			args: args{
				alphabet: "abcdefhgijklmnopqrstuvwxyz",
				sizes:    []int{},
				test:     alphabet,
			},
			want:    21,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.alphabet, tt.args.sizes...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want || !tt.args.test(got, tt.want) {
				t.Errorf("Generate() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func isDigit(src string, len int) bool {
	for _, x := range []rune(src) {
		if !unicode.IsDigit(x) {
			return false
		}
	}

	return true
}

func alphabet(src string, len int) bool {
	alpha := regexp.MustCompile(fmt.Sprintf("^[a-z]{%d}$", len))

	fmt.Println(alpha.MatchString(src))

	return true
}
