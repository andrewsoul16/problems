package str

var countMap = map[int32]func(string, int) int{
	'I': func(s string, i int) int { return 0 },
	'V': countV,
	'X': countX,
	'L': countL,
	'C': countC,
	'D': countD,
	'M': countM,
}

var romanMap = map[int32]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func toNumber(roman int32) int {
	return romanMap[roman]
}

func romanToInt(s string) int {
	result := 0
	//romanLetters := []rune{'M','D','C','L','X','V','I'}
	for i, v := range s {
		result += countMap[v](s, i)
	}
	// в этом случае у нас только единицы,
	// их количество равно количеству байт, поэтому берем len(s)
	if result == 0 {
		return len(s)
	}
	return result
}

func count(s string, iter int) int {
	res := romanMap[rune(s[iter])]
	if toNumber(rune(s[iter-1])) > toNumber(rune(s[iter])) {
		return res
	}
	return res
}

func countM(s string, iter int) int {
	res := 1000
	if iter == 0 {
		return res
	}
	if toNumber(rune(s[iter-1])) > toNumber(rune(s[iter])) {
		return res
	}
	for i := iter - 1; i >= 0; i-- {
		//r := s[i]
	}
	return res
}

func countV(s string, iter int) int {
	res := 5
	if iter == 0 {
		return res
	}
	var toCountBack = toNumber(rune(s[iter-1])) < toNumber(rune(s[iter]))
	if toCountBack {
		// идем в начало строки
		for i := iter; i >= 0; i-- {

		}
	}
	return res
}

func countX(s string, iter int) int {
	res := 10
	if iter == 0 {
		return res
	}
	return res
}

func countL(s string, iter int) int {
	res := 50
	if iter == 0 {
		return res
	}
	return res
}

func countC(s string, iter int) int {
	res := 100
	if iter == 0 {
		return res
	}
	return res
}

func countD(s string, iter int) int {
	res := 500
	if iter == 0 {
		return res
	}
	return res
}
