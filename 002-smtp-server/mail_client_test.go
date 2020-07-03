package main

import "testing"

func Test_SendMail(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMail()
		})
	}
}
