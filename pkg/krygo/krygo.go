package krygo

import (
	"strconv"
	"strings"
)



func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func Inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		return len(input), "char"
	}
	return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}

func Caesar(input string, shift int) {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	//lower_input := strings.ToLower(input)
	rune_input := []rune(input)

	for i := 0; i < len(rune_input); i++ {
		rune_input[i] = alphabet[]
	}

}
