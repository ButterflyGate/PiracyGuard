package tests

import (
	"github.com/ButterflyGate/piracy_guard/main"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name:"test1"
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main.Check()
		})
	}
}
