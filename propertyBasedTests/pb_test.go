package propertyBasedTests

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"testing/quick"
)

type TestCase struct {
	Arabic int
	Roman  string
}

var cases = []TestCase{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%d converts to %s", testCase.Arabic, testCase.Roman), func(t *testing.T) {
			got := ConvertToRoman(testCase.Arabic)
			want := testCase.Roman
			assertEqual(t, want, got)
		})
	}
}

func TestRomanToArabic(t *testing.T) {
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%s converts to %d", testCase.Roman, testCase.Arabic), func(t *testing.T) {
			got := ConvertToArabic(testCase.Roman)
			want := testCase.Arabic
			assertEqual(t, want, got)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) (result bool) {
		if arabic > 3999 {
			log.Println(arabic)
			result = true
			return
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertToArabic(roman)
		result = fromRoman == int(arabic)
		return
	}
	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

func assertEqual(t *testing.T, expected interface{}, got interface{}) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("want \"%v\", got \"%v\"", expected, got)
	}
}
