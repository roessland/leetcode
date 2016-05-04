package main

func romanToInt(s string) int {
	value := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := value[s[len(s)-1]]
	for i := 0; i < len(s)-1; i++ {
		if i+1 < len(s) && value[s[i]] < value[s[i+1]] {
			sum -= value[s[i]]
		} else {
			sum += value[s[i]]
		}
	}

	return sum
}
