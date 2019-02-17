package leapyear

import (
	"testing"
)

func TestIsLeap(t *testing.T) {
	testCases := []struct {
		desc string
		in   int
		want bool
	}{
		{"Year 1900 is common", 1900, false},
		{"Year 2000 is leap", 2000, true},
		{"Year 2001 is common", 2001, false},
		{"Year 2004 is leap", 2004, true},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := isLeap(test.in); got != test.want {
				t.Errorf("isLeap(%v): got %v want %v", test.in, got, test.want)
			}
		})
	}
}
