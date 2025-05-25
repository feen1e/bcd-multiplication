package main

import (
	"bytes"
	"testing"
)

func TestBinaryToBCD(t *testing.T) {
	expected := [][]byte{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	for i, c := range expected {
		result := BinaryToBCD(byte(i))
		if !bytes.Equal(c, result[1:]) {
			t.Errorf("%d: expected %b, got %b", i, c[0], result[0])
		}
	}
}

func TestStringToBCD(t *testing.T) {
	testCases := []struct {
		input          string
		expectedBCD    []byte
		expectedDecPos int
	}{
		{"123.45", []byte{1, 2, 3, 4, 5}, 2},
		{"0.5", []byte{0, 5}, 1},
		{"10.0", []byte{1, 0, 0}, 1},
		{"3.14159", []byte{3, 1, 4, 1, 5, 9}, 5},
		{"123456.789", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3},
		{"9.9999999999999999999", []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 19},
	}

	for _, tc := range testCases {
		bcd, decPos := StringToBCD(tc.input)

		// Sprawdzenie pozycji przecinka
		if decPos != tc.expectedDecPos {
			t.Errorf("StringToBCD(%s): expected decimal position %d, got %d",
				tc.input, tc.expectedDecPos, decPos)
		}

		// Sprawdzenie liczby BCD
		if len(bcd) != len(tc.expectedBCD) {
			t.Errorf("StringToBCD(%s): expected BCD length %d, got %d",
				tc.input, len(tc.expectedBCD), len(bcd))
			continue
		}

		for i := 0; i < len(bcd); i++ {
			if bcd[i] != tc.expectedBCD[i] {
				t.Errorf("StringToBCD(%s): at position %d expected %d, got %d",
					tc.input, i, tc.expectedBCD[i], bcd[i])
			}
		}
	}
}

func TestBCDToString(t *testing.T) {
	testCases := []struct {
		bcd           []byte
		decimalPlaces int
		expected      string
	}{
		{[]byte{1, 2, 3, 4, 5}, 2, "123.45"},
		{[]byte{0, 5}, 1, "0.5"},
		{[]byte{1, 0, 0}, 1, "10.0"},
		{[]byte{3, 1, 4, 1, 5, 9}, 5, "3.14159"},
		{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, "123456.789"},
		{[]byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 19, "9.9999999999999999999"},
	}

	for _, tc := range testCases {
		result := BCDToString(tc.bcd, tc.decimalPlaces)
		if result != tc.expected {
			t.Errorf("BCDToString(%v, %d): expected %s, got %s",
				tc.bcd, tc.decimalPlaces, tc.expected, result)
		}
	}
}

func TestStringToBCDToBCDToString(t *testing.T) {
	testCases := []string{
		"123.45",
		"0.5",
		"10.0",
		"3.14159",
		"123456.789",
		"9.9999999999999999999",
		"1.23456789012345678901234567890",
		"0.00000000000000000001",
	}

	for _, tc := range testCases {
		bcd, decPos := StringToBCD(tc)
		result := BCDToString(bcd, decPos)

		if result != tc {
			t.Errorf("Round trip conversion of %s: got %s", tc, result)
		}
	}
}
