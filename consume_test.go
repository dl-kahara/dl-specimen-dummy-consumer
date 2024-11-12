package main

import "testing"

func Test_consume(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consume(tt.args.config)
		})
	}
}
