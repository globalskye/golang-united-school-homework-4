package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return "", fmt.Errorf("Input went wrong: %w", errorEmptyInput)
	}
	b := 0
	a := make([]int, 0, 2)
	for i := 0; i < len(input); i++ {

		if i == 0 {

			continue
		}

		if input[i:i+1] == "-" || input[i:i+1] == "+" {

			value, err := strconv.Atoi(input[b:i])
			if err != nil {
				return "", fmt.Errorf("Input is wrong: %w", err)
			}
			b = i
			a = append(a, value)
		}
		if i == len(input)-1 {
			value, err := strconv.Atoi(input[b:i])
			if err != nil {
				return "", fmt.Errorf("Input is wrong: %w", err)
			}
			a = append(a, value)
		}

	}
	if len(a) != 2 {
		return "", fmt.Errorf("Input is wrong: %w", errorNotTwoOperands)
	}
	output = strconv.Itoa(a[0] + a[1])
	return output, nil

}
