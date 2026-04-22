package helpers

import (
	"strings"
	"unicode"
)

func ToASCII(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] <= unicode.MaxASCII {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

func CleanForASCIIArt(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]

		if (char >= 32 && char <= 126) || char == '\n' || char == '\r' {
			b.WriteByte(char)
		}
	}
	return b.String()
}
