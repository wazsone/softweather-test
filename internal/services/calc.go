package services

import (
	"strconv"
	"strings"
)

func EvaluateExpression(expr string) (int, error) {
	tokens := strings.Split(expr, "+")
	result := 0

	for _, token := range tokens {
		if strings.Contains(token, "-") {
			subTokens := strings.Split(token, "-")
			subResult, err := strconv.Atoi(subTokens[0])
			if err != nil {
				return 0, err
			}

			for _, subToken := range subTokens[1:] {
				num, err := strconv.Atoi(subToken)
				if err != nil {
					return 0, err
				}
				subResult -= num
			}

			result += subResult
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			result += num
		}
	}

	return result, nil
}
