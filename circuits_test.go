package main

import (
	"fmt"
	"testing"
)

func TestFinalTwoDigitMultiplier(t *testing.T) {
	for a := int64(0); a <= 9; a++ {
		for b := int64(0); b <= 9; b++ {
			result := finalTwoDigitMultiplier(DecimalToBCD(a)[0], DecimalToBCD(b)[0])
			r := BCDToDecimal(BinaryToBCDConverter(result))
			expected := a * b
			fmt.Printf("%d * %d = %08b (%d)\n", a, b, result, result)
			if r != expected {
				t.Errorf("finalTwoDigitMultiplier(%d,%d): expected %d, got %d", a, b, expected, r)
			}
		}
	}
}
