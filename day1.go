package main

func captcha (input string) int {
	var sum, nextIndex int
	for index := 0; index < len(input); index++  {
		char := input[index] - 48
		nextIndex = index + 1
		if int32(index) == int32(len(input) - 1) {
			nextIndex = 0
		}

		nextChar := input[nextIndex] - 48

		if int(char) == int(nextChar) {
			sum += int(char)
		}
	}

	return sum
}

func captcha2 (input string) int {
	var sum, nextIndex int
	indexIncrement := int(len(input)) / 2
	for index := 0; index < len(input); index++  {
		char := input[index] - 48
		nextIndex = index + indexIncrement
		if nextIndex >= len(input) {
			nextIndex = nextIndex - len(input)
		}

		nextChar := input[nextIndex] - 48

		if int(char) == int(nextChar) {
			sum += int(char)
		}
	}

	return sum
}
