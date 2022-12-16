package romans

func RomanToInt(s string) int {
	Romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result int
	max := 1
	for _, v := range s {
		letter := string(v)
		if Romans[letter] > max {
			result -= Romans[letter]
		} else if Romans[letter] == max {
			result += Romans[letter]
		} else {
			result += Romans[letter]
		}
		max = Romans[letter]
	}
	if result < 0 {
		return result * (-1)
	}
	return result
}
