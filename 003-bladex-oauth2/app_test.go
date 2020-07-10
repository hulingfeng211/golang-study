package main

import (
	"testing"
)

func Test_encode(t *testing.T) {
	type args struct {
		appid  string
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "abc",
			args: args{appid: "abc", secret: "def"},
			want: "YWJjZGVm",
		},
		{name: "abc",
			args: args{appid: "sword", secret: "sword_secret"},
			want: "YWJjZGVm",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.args.appid, tt.args.secret); got != tt.want {
				//t.Error(got)
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passwordEncode(t *testing.T) {
	type args struct {
		orgin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abc",
			args: args{orgin: "admin"},
			want: "YWJjZGVm"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := passwordEncode(tt.args.orgin); got != tt.want {
				t.Errorf("passwordEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
