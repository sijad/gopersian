package stringutils

import "testing"

func TestPersianToEnglishDigit(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			"123",
			"123",
		},
		{
			"۱۲۳",
			"123",
		},
		{
			"۱۲۳۴۵۶۷۸۹۰",
			"1234567890",
		},
		{
			"123۱۲۳123۱۲۳",
			"123123123123",
		},
		{
			"سلام 123",
			"سلام 123",
		},
		{
			"سلام ۱۲۳",
			"سلام 123",
		},
	}

	for _, test := range tests {
		if result := PersianToEnglishDigits(test.in); result != test.out {
			t.Errorf("PersianToEnglishDigits(%s) != %s, got: %s", test.in, test.out, result)
		}
	}
}

func TestEnglishTopPersianDigits(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{
			"123",
			"۱۲۳",
		},
		{
			"۱۲۳",
			"۱۲۳",
		},
		{
			"۱۲۳۴۵۶۷۸۹۰",
			"۱۲۳۴۵۶۷۸۹۰",
		},
		{
			"123۱۲۳123۱۲۳",
			"۱۲۳۱۲۳۱۲۳۱۲۳",
		},
		{
			"سلام 123",
			"سلام ۱۲۳",
		},
		{
			"سلام ۱۲۳",
			"سلام ۱۲۳",
		},
	}

	for _, test := range tests {
		if result := EnglishToPersianDigits(test.in); result != test.out {
			t.Errorf("PersianToEnglishDigits(%s) != %s, got: %s", test.in, test.out, result)
		}
	}
}
