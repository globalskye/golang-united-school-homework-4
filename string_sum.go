package string_sum

import (
	"errors"
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
	strings.TrimSpace(input)
	if len(input) == 0 {
		return "", errorEmptyInput
	}

	for i, v := range input {

		if v == '-' && i == 0 {
			continue
		}
		if v == '+' || v == '-' {
			a, err := strconv.Atoi(input[0:i])
			if err != nil {
				return "", errorEmptyInput
			}
			if a > 100 || a < 100-200 {
				return "", errorNotTwoOperands
			}

			f

			b, err1 := strconv.Atoi(input[i:len(input)])
			if err1 != nil {
				return "", errorNotTwoOperands
			}
			if b > 100 || b < 100-200 {
				return "", errorNotTwoOperands
			}

			output = strconv.Itoa(a + b)

		}

	}

	return output, nil

}
