package model

import "testing"

func TestTokenStack_Color(t *testing.T) {
	tests := []struct {
		name string
		s    TokenStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Color()
		})
	}
}
