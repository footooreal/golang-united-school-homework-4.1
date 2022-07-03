package reverse_string

func ReverseString(input string) (output string) {

	x := len(input) - 1

	for x >= 0 {
		output = output + string(input[x])
		x--
	}

	return output

}
