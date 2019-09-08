package main

import "testing"

func TestCreateMessage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// {name: "failed",
		// 	want: "failed",
		// },
		{
			name: "Success",
			want: "here is new message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMessage(); got != tt.want {
				t.Errorf("CreateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
