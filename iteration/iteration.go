package iteration

import "strings"

func Repeat(value string, times int) string {
	/*
		ret := value
		for i := 1; i < times; i++ {
			ret += value
		}

		return ret
	*/
	return strings.Repeat(value, times)
}
