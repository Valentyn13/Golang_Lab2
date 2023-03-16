package lab2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func checkInput(str string) bool {
	regex := regexp.MustCompile("[a-zA-Z]")
	containsLetters := regex.MatchString(str)
	return containsLetters
}

const ERROR_CODE float64 = 0

func PrefixSolver(exprsn string) (float64, error) {

	letterExist := checkInput(exprsn)

	if letterExist {
		return ERROR_CODE, fmt.Errorf("invalid expression")
	}

	if exprsn == "" {
		return ERROR_CODE, fmt.Errorf("EOF")
	}

	stack := make([]float64, 0)
	tokens := strings.Split(exprsn, " ")

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		num, err := strconv.ParseFloat(token, 64)
		if err == nil {
			stack = append(stack, num)
			continue
		}

		if len(stack) < 2 {
			return ERROR_CODE, fmt.Errorf("missing operand:stack index out of range")
		}
		o1, o2 := stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch token {
		case "+":
			stack = append(stack, o1+o2)
		case "-":
			stack = append(stack, o1-o2)
		case "*":
			stack = append(stack, o1*o2)
		case "/":
			if o2 == 0 {
				return ERROR_CODE, fmt.Errorf("division by zero")
			}
			stack = append(stack, o1/o2)
		default:
			return ERROR_CODE, fmt.Errorf("invalid input")
		}
	}

	if len(stack) != 1 {
		return ERROR_CODE, fmt.Errorf("missing operator:stack index out of range")
	}
	return stack[0], nil
}
