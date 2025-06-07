package main

import (
	"bytes"
	"slices"
	"testing"
)

func TestSplit(t *testing.T) {
	testCases := []struct {
		name     string
		input    rune
		expected bool
	}{
		{
			name:     "dash should split",
			input:    '—',
			expected: true,
		}, {
			name:     "hyphen should not split",
			input:    '-',
			expected: false,
		}, {
			name:     "whitespace should split",
			input:    ' ',
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := split(tc.input)

			if output != tc.expected {
				t.Errorf("split(%c) returned %v; expected %v", tc.input, output, tc.expected)
			}
		})
	}
}

func TestParseWords(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []string
	}{
		{
			name:     "parse it easy",
			input:    []byte("parse this—okay?\nALRIGHT\the’s The commander-in-chief"),
			expected: []string{"parse", "this", "okay", "alright", "he’s", "the", "commander-in-chief"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := parseWords(tc.input)

			if !slices.Equal(output, tc.expected) {
				t.Errorf("parseWords returned %v;\n\texpected %v", output, tc.expected)
			}
		})
	}
}

func TestParseText(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []string
	}{
		{
			name:     "parse it easy",
			input:    []byte("parse this—okay?\nALRIGHT\the’s The commander-in-chief"),
			expected: []string{"parse", "this", "okay", "alright", "he’s", "the", "commander-in-chief"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parseText(bytes.NewReader(tc.input))
		})
	}
}
