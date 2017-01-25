package main

import (
	"testing"
)

func TestCentre(t *testing.T) {
	tests := []struct {
		str      string
		padding  string
		width    uint
		expected string
	}{
		{
			str:      "HELLO",
			padding:  "=",
			width:    5,
			expected: "HELLO",
		}, {
			str:      "HELLO",
			padding:  "=",
			width:    30,
			expected: "=========== HELLO ============",
		}, {
			str:      "SIXSIX",
			padding:  "=",
			width:    30,
			expected: "=========== SIXSIX ===========",
		}, /*{
			str:      "TEST",
			padding:  "<>",
			width:    15,
			expected: "<><> TEST <><>",
		}, {
			str:      "TESTS",
			padding:  "<>",
			width:    15,
			expected: "<><> TESTS <><>",
		},*/
	}
	for _, test := range tests {
		result := centre(test.str, test.padding, test.width)
		if result != test.expected {
			t.Errorf("'%s' '%s' %d failed.\r\nExpected %s\r\n     got %s", test.str, test.padding, test.width, test.expected, result)
		}
		if len(result) != int(test.width) {
			t.Errorf("'%s' '%s' %d incorrect length (%d)", test.str, test.padding, test.width, len(result))
		}
	}
}
