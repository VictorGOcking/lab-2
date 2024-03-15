package lab2

import (
	"errors"
	"regexp"
	"strings"
)

func PostfixToPrefix(postfix string) (string, error) {

	input := strings.TrimSpace(postfix)

	// validation start
	if len(input) == 0 {
		return "", errors.New("empty expression")
	}

	if matched, _ := regexp.MatchString(`[+\-*/\^]{2,}`, input); matched {
		return "", errors.New("incorrect operators separation")
	}

	if matched, _ := regexp.MatchString(`[^0-9\s+\-*/\^]+`, input); matched {
		return "", errors.New("incorrect syntax")
	}

	numOperators := len(regexp.MustCompile(`[+\-*/\^]`).FindAllString(input, -1))
	numOperands := len(regexp.MustCompile(`\d+`).FindAllString(input, -1))

	if numOperands != numOperators+1 {
		return "", errors.New("incorrect number of operators")
	}
	// validation end

	// converting start
	stack := []string{}

	tokens := strings.Fields(input)

	for _, token := range tokens {
		if strings.ContainsAny(token, "+-*/^") {
			operand2Index := len(stack) - 1
			if operand2Index < 0 {
				return "", errors.New("incorrect syntax")
			}
			operand2 := stack[operand2Index]
			stack = stack[:operand2Index]

			operand1Index := len(stack) - 1
			if operand1Index < 0 {
				return "", errors.New("incorrect syntax")
			}
			operand1 := stack[operand1Index]
			stack = stack[:operand1Index]

			prefixExpr := token + " " + operand1 + " " + operand2
			stack = append(stack, prefixExpr)
		} else {
			stack = append(stack, token)
		}
	}
	// converting end

	return stack[len(stack)-1], nil
}
