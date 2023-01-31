package judgment

import "testing"

func TestIsUrl(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"../movie.mp4", false},
		{"movies.com/movie/0.mp4", false},
		{"http://127.0.0.1:11470/78b3743b34d37f73cd643da5027e590713ae6b19/0", true},
	}

	for _, test := range tests {
		if got := IsUrl(test.input); got != test.expected {
			t.Errorf("IsUrl(%v) = %v", test.input, got)
		}
	}
}
