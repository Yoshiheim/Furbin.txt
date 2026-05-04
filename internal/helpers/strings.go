package helpers

import (
	"unicode/utf8"
)

// Its cut string by maxBytes size
func TruncateByte(s string, maxBytes int) string {
	if len(s) <= maxBytes { // len() returns byte count
		return s
	}
	for !utf8.ValidString(s[:maxBytes]) {
		maxBytes--
	}
	return string(s[:maxBytes])
}

func CheckSizeString(s string, maxBytes int) bool {
	if utf8.RuneCountInString(s) >= maxBytes { // len() returns byte count
		return true
	}
	return false
}

/*
func DestroySpaces(text string) string {
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\b", "")
	text = strings.ReplaceAll(text, "⠀", "")
	text = strings.ReplaceAll(text, " ", "")
	return text
}
*/

func DestroySpaces(s string) string {
	var b []rune

	for _, r := range s {
		switch r {
		case '\n', '\t', '\b', ' ', '⠀':
			continue
		default:
			b = append(b, r)
		}
	}
	return string(b)
}

func OnlyASCII(s string) string {
	var b []rune

	for _, r := range s {
		switch r {
		case '\n', '\t':
			b = append(b, r)
		default:
			if r >= 32 && r <= 126 { // printable ASCII
				b = append(b, r)
			}
		}
	}

	return string(b)
}
