package roman

type OutputRoman struct {
	returnValue int
	result      string
}

func IntToRoman(i int) string {
	outputRoman := OutputRoman{returnValue: i}
	var result string
	if i >= 1000 {
		outputRoman = IntToRomanProcess(outputRoman.returnValue, 1000, "M", "", "")
		result += outputRoman.result
	}
	if i >= 100 {
		outputRoman = IntToRomanProcess(outputRoman.returnValue, 100, "C", "D", "M")
		result += outputRoman.result
	}
	if i >= 10 {
		outputRoman = IntToRomanProcess(outputRoman.returnValue, 10, "X", "L", "C")
		result += outputRoman.result
	}
	outputRoman = IntToRomanProcess(outputRoman.returnValue, 1, "I", "V", "X")
	result += outputRoman.result
	//fmt.Println(outputRoman)

	return result
}

func IntToRomanProcess(i, mod int, romanString, romanStringHalf, romanStringOver string) OutputRoman {
	var result string
	if i >= mod {
		temp := (i - (i % mod)) / mod
		if temp < 4 {
			for j := 1; j <= temp; j++ {
				result += romanString
			}
		} else if temp == 4 {
			result += romanString + romanStringHalf
		} else if temp < 9 {
			result += romanStringHalf
			for j := 6; j <= temp; j++ {
				result += romanString
			}
		} else {
			result += romanString + romanStringOver
		}
	}
	returnValue := i % mod
	return OutputRoman{returnValue: returnValue, result: result}
}

//func main() {
//	for i := 1; i <= 20; i++ {
//		fmt.Printf("%v => %v\n", i, IntToRoman(i))
//
//	}
//}
