package service

import "unicode/utf8"

func ValidateTitle(s string) bool {
	// Заголовок не пустой, короче 200 символов и валиден UTF-8
	return utf8.ValidString(s) && len([]rune(s)) > 0 && len([]rune(s)) <= 200
}
