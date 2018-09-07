package gopersian

import (
	"strings"
)

const (
	persianToEnglishDiff = '۰' - '0' // 1728
)

// PersianToEnglishDigits converts all persian digits to english digits
func PersianToEnglishDigits(str string) string {
	return convertDigits(str, '۰', '۹', persianToEnglishDiff)
}

// EnglishToPersianDigits converts all english digits to persian digits
func EnglishToPersianDigits(str string) string {
	return convertDigits(str, '0', '9', -persianToEnglishDiff)
}

func convertDigits(str string, min, max, diff rune) string {
	if !containsAnyInRange(str, min, max) {
		return str
	}

	newStr := make([]rune, 0, len(str))
	for _, r := range str {
		if r >= min && r <= max {
			r -= diff
		}
		newStr = append(newStr, r)
	}

	return string(newStr)
}

func containsAnyInRange(s string, a, b rune) bool {
	for i := a; i <= b; i++ {
		if strings.ContainsRune(s, i) {
			return true
		}
	}

	return false
}
