package levenshtein

import "testing"

func TestDistance(t *testing.T) {
	testCases := []struct {
		a, b string
		dist int
	}{
		// empty strings
		{"", "", 0},
		{"a", "", 1},
		{"", "a", 1},
		{"abc", "", 3},
		{"", "abc", 3},

		// equal strings
		{"a", "a", 0},
		{"abc", "abc", 0},

		// only inserts
		{"a", "ab", 1},
		{"b", "ab", 1},
		{"ac", "abc", 1},
		{"abcdefg", "xabxcdxxefxgx", 6},

		// only deletes
		{"ab", "a", 1},
		{"ab", "b", 1},
		{"abc", "ac", 1},
		{"xabxcdxxefxgx", "abcdefg", 6},

		// only substitions
		{"a", "b", 1},
		{"ab", "ac", 1},
		{"abc", "axc", 1},
		{"xabxcdxxefxgx", "1ab2cd34ef5g6", 6},

		// many oprations
		{"example", "samples", 3},
		{"sturgeon", "urgently", 6},
		{"levenshtein", "frankenstein", 6},
		{"distance", "difference", 5},
		{"java was neat", "scala is great", 7},

		// utf-8
		{"тест", "Тест", 1},
		{"gross", "groß", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.a+"·"+tc.b, func(t *testing.T) {
			dist := Distance(tc.a, tc.b)
			if dist != tc.dist {
				t.Errorf("got %v, want %v", dist, tc.dist)
			}
		})
	}
}
