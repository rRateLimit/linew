package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rRateLimit/linew/internal/config"
)

func TestProcessText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		config   *config.Config
		expected string
	}{
		{
			name:  "Simple wrap",
			input: "This is a test line that should be wrapped at twenty characters",
			config: &config.Config{
				Width:          20,
				PreserveIndent: false,
			},
			expected: "This is a test line\nthat should be\nwrapped at twenty\ncharacters\n",
		},
		{
			name:  "Multiple lines",
			input: "First line\nSecond line that is very long and needs wrapping\nThird line",
			config: &config.Config{
				Width:          25,
				PreserveIndent: false,
			},
			expected: "First line\nSecond line that is very\nlong and needs wrapping\nThird line\n",
		},
		{
			name:  "Preserve indentation",
			input: "    Indented line that needs to be wrapped with preserved indentation",
			config: &config.Config{
				Width:          30,
				PreserveIndent: true,
			},
			expected: "    Indented line that needs\n    to be wrapped with\n    preserved indentation\n",
		},
		{
			name:  "Empty lines",
			input: "Line one\n\nLine three",
			config: &config.Config{
				Width:          80,
				PreserveIndent: true,
			},
			expected: "Line one\n\nLine three\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			var writer bytes.Buffer

			err := processText(reader, &writer, tt.config)
			if err != nil {
				t.Fatalf("processText() error = %v", err)
			}

			result := writer.String()
			if result != tt.expected {
				t.Errorf("processText() output = %q, want %q", result, tt.expected)
			}
		})
	}
}