package validator

import "unicode/utf8"

func CheckMaxChars(value string, number int) bool {
	return utf8.RuneCountInString(value) <= number
}

func CheckMinChars(value string, number int) bool {
	return utf8.RuneCountInString(value) >= number
}

func CheckMaxInt(value int, number int) bool {
	return value <= number
}

func CheckMinInt(value int, number int) bool {
	return value >= number
}
