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
	if len(input) == 0 || strings.ContainsAny(input, "") {
		return "", fmt.Errorf("Input went wrong: %w", errorEmptyInput)
	}

	for i, v := range input {
		if i == 0 && (v == '+' || v == '-') {
			continue
		}
		if v == '+' || v == '-' {
			value1, err1 := strconv.Atoi(input[0:i])
			if err1 != nil {
				return "", fmt.Errorf("Input is wrong: %w", err1)
			}
			value2, err2 := strconv.Atoi(input[i:len(input)])
			if err2 != nil {
				return "", fmt.Errorf("Input is wrong: %w", err2)
			}
			output = strconv.Itoa(value1 + value2)
		}

	}
	return output, nil
}
