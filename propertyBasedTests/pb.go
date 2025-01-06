package propertyBasedTests

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, num := range allRomanNumerals {
		for arabic >= num.Value {
			result.WriteString(num.Symbol)
			arabic -= num.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) int {
	arabic := 0
	for _, num := range allRomanNumerals {
		for strings.HasPrefix(roman, num.Symbol) {
			arabic += num.Value
			roman = strings.TrimPrefix(roman, num.Symbol)
		}
	}
	return arabic
}
