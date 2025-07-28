package wrap

import (
	"strings"
	"unicode"

	"github.com/rRateLimit/linew/internal/config"
)

type Wrapper struct {
	config *config.Config
}

func New(cfg *config.Config) *Wrapper {
	return &Wrapper{
		config: cfg,
	}
}

func (w *Wrapper) WrapLine(line string) []string {
	if len(line) <= w.config.Width {
		return []string{line}
	}

	var result []string
	indent := ""

	if w.config.PreserveIndent {
		indent = extractIndent(line)
	}

	trimmedLine := strings.TrimSpace(line)
	words := splitIntoWords(trimmedLine)

	if len(words) == 0 {
		return []string{line}
	}

	var currentLine strings.Builder
	if w.config.PreserveIndent {
		currentLine.WriteString(indent)
	}
	lineLength := len(currentLine.String())

	for i, word := range words {
		wordLen := len(word)
		
		// Check if we need to start a new line
		if i > 0 && lineLength+1+wordLen > w.config.Width {
			result = append(result, currentLine.String())
			currentLine.Reset()
			
			if w.config.PreserveIndent {
				currentLine.WriteString(indent)
				lineLength = len(indent)
			} else {
				lineLength = 0
			}
		} else if i > 0 {
			currentLine.WriteString(" ")
			lineLength++
		}

		// Handle very long words that need to be broken
		remainingWidth := w.config.Width - lineLength
		if wordLen > remainingWidth && remainingWidth > 0 {
			parts := breakLongWord(word, remainingWidth)
			for j, part := range parts {
				if j > 0 {
					result = append(result, currentLine.String())
					currentLine.Reset()
					if w.config.PreserveIndent {
						currentLine.WriteString(indent)
						lineLength = len(indent)
					} else {
						lineLength = 0
					}
					remainingWidth = w.config.Width - lineLength
				}
				currentLine.WriteString(part)
				lineLength += len(part)
			}
		} else {
			currentLine.WriteString(word)
			lineLength += wordLen
		}
	}

	if currentLine.Len() > 0 || (currentLine.Len() == 0 && w.config.PreserveIndent && len(indent) > 0) {
		result = append(result, currentLine.String())
	}

	return result
}

func extractIndent(line string) string {
	for i, r := range line {
		if !unicode.IsSpace(r) {
			return line[:i]
		}
	}
	return line
}

func splitIntoWords(text string) []string {
	return strings.Fields(text)
}

func breakLongWord(word string, maxWidth int) []string {
	if maxWidth <= 0 {
		maxWidth = 1
	}

	var result []string
	runes := []rune(word)
	
	for i := 0; i < len(runes); i += maxWidth {
		end := i + maxWidth
		if end > len(runes) {
			end = len(runes)
		}
		result = append(result, string(runes[i:end]))
	}
	
	return result
}