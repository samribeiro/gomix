package unlockbug

import (
	"testing"
)

func TestUnlockBug(t *testing.T) {
	tests := []struct {
		s Struct
	}{
		{
			s: Struct{},
		},
		{
			s: Struct{},
		},
		{
			s: Struct{},
		},
		{
			s: Struct{},
		},
	}
	for _, test := range tests {
		go test.s.CallNotify()
	}
}
