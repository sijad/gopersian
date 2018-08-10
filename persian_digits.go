package gopersian

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
	newStr := make([]rune, 0, len(str))
	for _, r := range str {
		if r >= min && r <= max {
			r -= diff
		}
		newStr = append(newStr, r)
	}

	return string(newStr)
}
