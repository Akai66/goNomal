package word

import (
	"testing"
)

type testCase struct {
	input string
	want  bool
}

var tests = []testCase{
	{"", true},
	{"aba", true},
	{"cdegedc", true},
	{"我t我", true},
	{"Et se resservir, ivresse reste.", false},
	{"我爱静静，静静爱我", true},
	{"我爱中国，中国爱我", false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%s) = %v", test.input, got)
		}
	}
}
