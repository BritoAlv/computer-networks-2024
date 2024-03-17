package main

import "testing"

func Test_is_valid_ip(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 1", args{"192.168.1.1"}, true},
		{"Test 2", args{"129.200.00.34"}, true},
		{"Test 3", args{"256.200.00.34"}, false},
		{"Test 4", args{"est.200.00.34"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_valid_ip(tt.args.ip); got != tt.want {
				t.Errorf("is_valid_ip() = %v, want %v", got, tt.want)
			}
		})
	}
}
