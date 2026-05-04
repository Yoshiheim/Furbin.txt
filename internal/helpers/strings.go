package helpers

import (
	"strings"
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

func TrimLeft(s string) string {
	return strings.TrimLeft(s, " \n\t\r\v\f")
}

func DestroySpaces(s string) string {
	var b []rune

	for _, r := range s {
		switch r {
		case '\n', '\t', '\b', '\r', ' ', '⠀':
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
