package reverse_string

func ReverseString(input string) (output string) {

	toRune := []rune(input)
	x := len(toRune) - 1

	for x >= 0 {
		output = output + string(toRune[x])
		x--
	}
	return output

	return output

}
