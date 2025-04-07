package main

import "testing"

func TestMultiplySingleDigitBCD(t *testing.T) {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			result := BCDToDecimal(MultiplySingleDigitBCD(DecimalToBCD(a)[0], DecimalToBCD(b)[0]))
			expected := a * b
			if result != expected {
				t.Errorf("MultiplySingleDigitBCD(%d,%d): expected %d, got %d", a, b, expected, result)
			}
		}
	}
}

func TestMultiplyBCD(t *testing.T) {
	for a := 0; a <= 9999; a++ {
		for b := 0; b <= 9999; b++ {
			result := BCDToDecimal(MultiplyBCD(DecimalToBCD(a), DecimalToBCD(b)))
			expected := a * b
			if result != expected {
				t.Errorf("MultiplyBCD(%d,%d): expected %d, got %d", a, b, expected, result)
			}
		}
	}
}
