package helpers

import "unicode"

// Its cut string by maxBytes size
func TruncateByte(s string, maxBytes int) string {
	if len(s) <= maxBytes { // len() returns byte count
		return s
	}
	return s[:maxBytes]
}

// this function doesn't work like I want.
func SanitizeString(s string) string {
	result := make([]rune, 0, len(s))

	for _, r := range s {
		if unicode.IsPrint(r) || r == '\n' || r == '\t' {
			result = append(result, r)
		}
	}
	return string(result)
}
