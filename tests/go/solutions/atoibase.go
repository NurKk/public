package correct

import "strings"

var m = map[rune]struct{}{}

func uniqueChar(s string) bool {
	for _, r := range s {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func validBase(base string) bool {
	return len(base) >= 2 && !strings.ContainsAny(base, "+-") && uniqueChar(base)
}

func power(nbr int, pwr int) int {
	if pwr == 0 {
		return 1
	}
	if pwr == 1 {
		return nbr
	}
	return nbr * power(nbr, pwr-1)
}

func AtoiBase(s string, base string) int {
	var result int
	var i int
	sign := 1
	lengthBase := len(base)
	lengths := len(s)

	if !validBase(base) {
		return 0
	}
	if s[i] == '-' {
		sign = -1
	}
	if s[i] == '-' || s[i] == '+' {
		i++
	}
	for i < len(s) {
		result += (Index(base, string(s[i])) * power(lengthBase, lengths-1))
		i++
		lengths--
	}
	return result * sign
}
