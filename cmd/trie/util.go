package main

import (
	"strings"

	"unicode/utf8"
)

func GetIndex(c rune) int {
	return int(c) - 97
}

func Lower(str string) string {
	return strings.ToLower(str)
}
func isSpace(r rune) bool {
	switch r {
	case ' ', '\n', '\t':
		return true
	}
	return false
}

func SplitEnglishWords(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	start, i := 0, 0
	for i < len(data) {
		isEnglishWords := true

		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if !isSpace(r) {
				break
			}
		}
		i = start
		// Scan until space, marking end of word.
		for width := 0; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if isSpace(r) {
				if isEnglishWords {
					return i, data[start:i], nil
				} else {
					start = i + width
					break
				}
			}
			// A-Z  [65:90] [0x41:0x5A]
			// a-z  [97:122] [0x61:0x7a]
			if !(r >= 0x41 && r <= 0x5a) && !(r >= 0x61 && r <= 0x7a) {
				isEnglishWords = false
			}
		}
	}
	return start, nil, nil
}
