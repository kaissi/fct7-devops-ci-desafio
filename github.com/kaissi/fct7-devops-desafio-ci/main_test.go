package main

import "testing"

func TestSoma(t *testing.T) {

	// Positive values
	result := soma(2, 3)
	if result != 5 {
		t.Errorf("soma(2, 3) failed, expected %v, got %v", 5, result)
	}

	// Positive and negative values
	result = soma(-3, 1)
	if result != -2 {
		t.Errorf("soma(-3, 1) failed, expected %v, got %v", -2, result)
	}

	// Negative values
	result = soma(-1, -3)
	if result != -4 {
		t.Errorf("soma(-1, -3) failed, expected %v, got %v", -4, result)
	}
}