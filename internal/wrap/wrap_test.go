package wrap

import (
	"reflect"
	"testing"

	"github.com/rRateLimit/linew/internal/config"
)

func TestWrapLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		config   *config.Config
		expected []string
	}{
		{
			name: "Short line",
			line: "This is a short line",
			config: &config.Config{
				Width:          80,
				PreserveIndent: true,
			},
			expected: []string{"This is a short line"},
		},
		{
			name: "Long line with wrap",
			line: "This is a very long line that needs to be wrapped because it exceeds the maximum width",
			config: &config.Config{
				Width:          30,
				PreserveIndent: true,
			},
			expected: []string{
				"This is a very long line that",
				"needs to be wrapped because it",
				"exceeds the maximum width",
			},
		},
		{
			name: "Line with indentation",
			line: "    This is an indented line that should preserve its indentation when wrapped",
			config: &config.Config{
				Width:          40,
				PreserveIndent: true,
			},
			expected: []string{
				"    This is an indented line that should",
				"    preserve its indentation when",
				"    wrapped",
			},
		},
		{
			name: "Line with indentation not preserved",
			line: "    This is an indented line that should not preserve indentation",
			config: &config.Config{
				Width:          40,
				PreserveIndent: false,
			},
			expected: []string{
				"This is an indented line that should not",
				"preserve indentation",
			},
		},
		{
			name: "Very long word",
			line: "This verylongwordthatexceedsthemaximumwidthandneedstobebrokenapart",
			config: &config.Config{
				Width:          20,
				PreserveIndent: false,
			},
			expected: []string{
				"This",
				"verylongwordthatexce",
				"edsthemaximumwidthan",
				"dneedstobebrokenapar",
				"t",
			},
		},
		{
			name: "Empty line",
			line: "",
			config: &config.Config{
				Width:          80,
				PreserveIndent: true,
			},
			expected: []string{""},
		},
		{
			name: "Only whitespace",
			line: "    ",
			config: &config.Config{
				Width:          80,
				PreserveIndent: true,
			},
			expected: []string{"    "},
		},
		{
			name: "Tab indentation",
			line: "\t\tThis line has tab indentation",
			config: &config.Config{
				Width:          40,
				PreserveIndent: true,
			},
			expected: []string{
				"\t\tThis line has tab indentation",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.config)
			result := w.WrapLine(tt.line)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WrapLine() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractIndent(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected string
	}{
		{
			name:     "No indentation",
			line:     "No indent here",
			expected: "",
		},
		{
			name:     "Space indentation",
			line:     "    Four spaces",
			expected: "    ",
		},
		{
			name:     "Tab indentation",
			line:     "\t\tTwo tabs",
			expected: "\t\t",
		},
		{
			name:     "Mixed indentation",
			line:     "  \t Mixed indent",
			expected: "  \t ",
		},
		{
			name:     "Only whitespace",
			line:     "    ",
			expected: "    ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractIndent(tt.line)
			if result != tt.expected {
				t.Errorf("extractIndent() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestBreakLongWord(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		maxWidth int
		expected []string
	}{
		{
			name:     "Simple break",
			word:     "verylongword",
			maxWidth: 4,
			expected: []string{"very", "long", "word"},
		},
		{
			name:     "Uneven break",
			word:     "verylongword",
			maxWidth: 5,
			expected: []string{"veryl", "ongwo", "rd"},
		},
		{
			name:     "Single character width",
			word:     "abc",
			maxWidth: 1,
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Word shorter than max",
			word:     "short",
			maxWidth: 10,
			expected: []string{"short"},
		},
		{
			name:     "Unicode characters",
			word:     "こんにちは",
			maxWidth: 2,
			expected: []string{"こん", "にち", "は"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := breakLongWord(tt.word, tt.maxWidth)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("breakLongWord() = %v, want %v", result, tt.expected)
			}
		})
	}
}