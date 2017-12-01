package day1

func Captcha (input string) int {
	var sum, nextIndex int
	for index := 0; index < len(input); index++  {
		nextIndex = index + 1
		if int32(index) == int32(len(input) - 1) {
			nextIndex = 0
		}

		if input[index] == input[nextIndex] {
			sum += int(input[index]) - 48
		}
	}

	return sum
}

func Captcha2 (input string) int {
	var sum, nextIndex int
	indexIncrement := int(len(input)) / 2
	for index := 0; index < len(input); index++  {
		nextIndex = index + indexIncrement
		if nextIndex >= len(input) {
			nextIndex = nextIndex - len(input)
		}

		if input[index] == input[nextIndex] {
			sum += int(input[index]) - 48

		}
	}

	return sum
}
