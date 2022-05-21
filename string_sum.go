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

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	
	input = strings.Join(strings.Fields(input), "")
	if input == "" {
		return "", fmt.Errorf("my err: %w", errorEmptyInput)
	}
	if strings.Count(input, "+")+strings.Count(input, "-") > 2 {
		return "", fmt.Errorf("my err: %w", errorNotTwoOperands)
	}
	k := len(input)
	for i := 1; i < k-1; i++ {
		if input[i] == '+' || input[i] == '-' {
			a, err1 := strconv.Atoi(input[0:i])
			b, err2 := strconv.Atoi(input[i:k])
			if err1 != nil {
				return "", fmt.Errorf("my err: %w", err1)
			}
			if err2 != nil {
				return "", fmt.Errorf("my err: %w", err2)
			}
			return strconv.Itoa(a + b), nil
		}
	}
	return "", nil
}
