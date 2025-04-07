package main

// DecimalToBCD konwertuje liczbę dziesiętną na liczbę BCD
// decimal — wejściowa liczba dziesiętna
func DecimalToBCD(decimal int) []byte {
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

// BCDToDecimal konwertuje liczbę BCD na liczbę dziesiętną
// bcd — wejściowa liczba BCD
func BCDToDecimal(bcd []byte) int {
	decimal := 0

	for _, digit := range bcd {
		decimal = decimal*10 + int(digit)
	}

	return decimal
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
