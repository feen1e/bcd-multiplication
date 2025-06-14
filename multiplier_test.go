package main

import (
	"math/rand"
	"testing"
	"time"
)

func DecimalToBCD(decimal int64) []byte {
	var bcd []byte
	for {
		bcd = append([]byte{byte(decimal % 10)}, bcd...)
		decimal /= 10
		if decimal <= 0 {
			break
		}
	}
	return bcd
}

func BCDToDecimal(bcd []byte) int64 {
	decimal := int64(0)

	for _, digit := range bcd {
		decimal = decimal*10 + int64(digit)
	}

	return decimal
}

func TestMultiplySingleDigitBCD(t *testing.T) {
	for a := int64(0); a <= 9; a++ {
		for b := int64(0); b <= 9; b++ {
			result := BCDToDecimal(MultiplySingleDigitBCD(DecimalToBCD(a)[0], DecimalToBCD(b)[0]))
			expected := a * b
			if result != expected {
				t.Errorf("MultiplySingleDigitBCD(%d,%d): expected %d, got %d", a, b, expected, result)
			}
		}
	}
}

func TestMultiplyBCD(t *testing.T) {
	for a := int64(0); a <= 9999; a++ {
		for b := int64(0); b <= 9999; b++ {
			result := BCDToDecimal(MultiplyBCD(DecimalToBCD(a), DecimalToBCD(b)))
			expected := a * b
			if result != expected {
				t.Errorf("MultiplyBCD(%d,%d): expected %d, got %d", a, b, expected, result)
			}
		}
	}
}

func TestMultiplyBCDForBigNumbers(t *testing.T) {
	const maxSafe = 3_037_000_499
	const samples = 1000000 // Liczba par do testowania

	rand.New(rand.NewSource(time.Now().UnixNano()))

	type pair struct {
		a int64
		b int64
	}

	var testPairs []pair

	// Przykładowe wartości
	testPairs = append(testPairs, pair{10_000, 10_000}, pair{maxSafe, maxSafe}, pair{1, 1}, pair{123456789, 987654321})

	// Wylosowane liczby
	for i := 0; i < samples; i++ {
		a := rand.Int63n(maxSafe-10_000) + 10_000
		b := rand.Int63n(maxSafe-10_000) + 10_000
		testPairs = append(testPairs, pair{a, b})
	}

	for _, pair := range testPairs {
		a := pair.a
		b := pair.b

		result := BCDToDecimal(MultiplyBCD(DecimalToBCD(a), DecimalToBCD(b)))
		expected := a * b

		if result != expected {
			t.Errorf("MultiplyBCD(%d,%d): expected %d, got %d", a, b, expected, result)
		}
	}
}
