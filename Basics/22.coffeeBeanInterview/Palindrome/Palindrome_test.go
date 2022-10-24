package main

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid Palindrome",
			input:    "madam",
			expected: "Palindrome",
		},
		{
			name:     "Invalid Palindrome",
			input:    "hemnath",
			expected: "Not Palindrome",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := palindrome(test.input)
			want := test.expected
			if got != want {
				t.Errorf("Unmatched values %q %q", got, want)
			}
		})
	}
}
