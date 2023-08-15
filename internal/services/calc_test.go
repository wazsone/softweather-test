package services

import (
	"testing"
)

func TestCalc_EvaluateExpression(t *testing.T) {

	// Testing expression with only addition
	result, err := EvaluateExpression("1+2+3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 6 {
		t.Errorf("Expected result to be 6, but got %d", result)
	}

	// Testing expression with only subtraction
	result, err = EvaluateExpression("10-5-3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("Expected result to be 2, but got %d", result)
	}

	// Testing expression with both addition and subtraction
	result, err = EvaluateExpression("5+3-2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 6 {
		t.Errorf("Expected result to be 6, but got %d", result)
	}

	// Testing expression with invalid token
	_, err = EvaluateExpression("10+abc")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
