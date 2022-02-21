package solution

import "testing"

func TestGetMessage(t *testing.T) {
	var tests []struct {
		name string
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
