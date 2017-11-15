package lavenshtein

import "testing"

func TestDistance(t *testing.T) {
	testCases := []struct {
		a, b string
		dist int
	}{
		{"", "", 0},
		{"aa", "aa", 0},
		{"aa", "ab", 1},
		{"aa", "ba", 1},
		{"aa", "bb", 2},
		{"aa", "", 2},
		{"", "bb", 2},
		{"test", "Test", 1},
		{"тест", "Тест", 1},
	}
	for _, tc := range testCases {
		t.Run(tc.a+" "+tc.b, func(t *testing.T) {
			dist := Distance(tc.a, tc.b)
			if dist != tc.dist {
				t.Errorf("got %v, want %v", dist, tc.dist)
			}
		})
	}
}
