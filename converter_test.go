package main

import (
	"bytes"
	"testing"
)

func TestDecimalToBCD(t *testing.T) {
	expected := [][]byte{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	for ind, c := range expected {
		i := int64(ind)
		result := DecimalToBCD(i)
		if !bytes.Equal(c, result) {
			t.Errorf("%d: expected %s, got %s", i, c, result)
		}
	}
}

func TestBCDToDecimal(t *testing.T) {
	expected := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for ind, c := range expected {
		i := int64(ind)
		result := BCDToDecimal(DecimalToBCD(i))
		if result != c {
			t.Errorf("%d: expected %d, got %d", i, c, result)
		}
	}
}

func TestBinaryToBCD(t *testing.T) {
	expected := [][]byte{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	for i, c := range expected {
		result := BinaryToBCD(byte(i))
		if !bytes.Equal(c, result[1:]) {
			t.Errorf("%d: expected %b, got %b", i, c[0], result[0])
		}
	}
}
