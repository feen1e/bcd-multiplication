package main

import (
	"strings"
)

// BCDToString konwertuje liczbę BCD z informacją o pozycji przecinka na string
// bcd — wejściowa liczba BCD
// decimalPlaces — pozycja przecinka (licząc od prawej strony)
func BCDToString(bcd []byte, decimalPlaces int) string {
	if len(bcd) == 0 {
		return "0"
	}

	var result strings.Builder

	// Dodaj cyfry przed przecinkiem
	for i := 0; i < len(bcd)-decimalPlaces; i++ {
		result.WriteByte('0' + bcd[i])
	}

	// Dodaj przecinek, jeśli są cyfry po przecinku
	if decimalPlaces > 0 {
		result.WriteByte('.')

		// Dodaj cyfry po przecinku
		for i := len(bcd) - decimalPlaces; i < len(bcd); i++ {
			result.WriteByte('0' + bcd[i])
		}
	}

	return result.String()
}

// StringToBCD konwertuje string reprezentujący liczbę na liczbę BCD z informacją o pozycji przecinka
// str — wejściowy string reprezentujący liczbę
// Zwraca tablicę bajtów BCD oraz pozycję przecinka (licząc od prawej strony)
func StringToBCD(str string) ([]byte, int) {
	// Znajdź pozycję przecinka
	decimalPos := strings.Index(str, ".")

	// Usuń przecinek
	strWithoutDecimal := strings.Replace(str, ".", "", 1)

	// Konwertuj na BCD cyfra po cyfrze
	var bcd []byte
	for _, ch := range strWithoutDecimal {
		if ch >= '0' && ch <= '9' {
			digit := byte(ch - '0')
			bcd = append(bcd, digit)
		}
	}

	// Oblicz pozycję przecinka (licząc od prawej strony)
	decimalPlaces := len(strWithoutDecimal) - decimalPos

	if decimalPos == -1 {
		decimalPlaces = 0
	}

	return bcd, decimalPlaces
}

// BinaryToBCD konwertuje liczbę binarną, zwracaną jako wynik mnożenia liczb BCD, z powrotem na liczbę BCD
// bin — wejściowa liczba binarna (wynik mnożenia dwóch liczb BCD)
func BinaryToBCD(bin byte) []byte {
	var bcd uint16 = 0
	for i := 0; i < 8; i++ {
		// Jeśli jakakolwiek cyfra BCD jest >= 5, dodanie korekty 3
		if bcd&0x000F >= 5 {
			bcd += 0x0003
		}
		if (bcd>>4)&0x000F >= 5 {
			bcd += 0x0030
		}

		// Przesunięcie bitowe i dodanie odpowiedniego bitu wejściowego
		bcd = (bcd << 1) | uint16((bin>>(7-i))&0x01)
	}

	// Zamiana na tablicę bajtów
	var bcdTab []byte
	bcdTab = append([]byte{byte(bcd & 0b1111)}, bcdTab...)
	bcdTab = append([]byte{byte(bcd >> 4)}, bcdTab...)

	return bcdTab
}
