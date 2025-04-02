package main

import (
	"fmt"
)

// Konwersja liczby dziesiętnej na BCD
func decimalToBCD(decimal int) []byte {
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

// Konwersja BCD na liczbę dziesiętną
func bcdToDecimal(bcd []byte) int {
	decimal := 0

	for _, digit := range bcd {
		decimal = decimal*10 + int(digit)
	}

	return decimal
}

// Konwersja liczby binarnej na BCD
func binaryToBCD(binary byte) []byte {
	decimal := int(binary)

	var bcd []byte

	bcd = decimalToBCD(decimal)

	return bcd
}

func doubleDabbleBinaryToBCD(binary byte) []byte {
	// TODO
	bcd := make([]byte, 2)

	for i := 7; i > 0; i-- {
		bcd[0] += binary & 0b10000000
		binary <<= 1

		bcd[0] <<= 1

		if bcd[0]&0b10000 == 0b10000 {
			bcd[0] = bcd[0] ^ 0b10000
			bcd[1] = (bcd[1] << 1) + 1
		}

	}

	return bcd
}

// Funkcja do mnożenia dwóch liczb BCD, na podstawie Fig. 2. w artykule
func multiplySingleBCD(a, b byte) []byte {
	var p byte = 0

	// Wycięcie poszczególnych bitów z A i B
	a3, a2, a1, a0 := (a>>3)&1, (a>>2)&1, (a>>1)&1, a&1
	b3, b2, b1, b0 := (b>>3)&1, (b>>2)&1, (b>>1)&1, b&1

	if (a3 | b3) == 1 {
		if a3 == 1 {
			if a == 0b1000 {
				p = b << 3
			} else if a == 0b1001 {
				p = b3<<6 + b2<<5 + b1<<4 + (b0|b3)<<3 + b2<<2 + b1<<1 + b0
			}
		} else if b3 == 1 {
			if b == 0b1000 {
				p = a << 3
			} else if b == 0b1001 {
				p = a3<<6 + a2<<5 + a1<<4 + (a0|a3)<<3 + a2<<2 + a1<<1 + a0
			}
		}
	} else if (a3 | b3) == 0 {
		p = (a & 0b111) * (b & 0b111)
	}
	if (a == 0b1001) && (b == 0b1001) {
		p6, p5, p2, p1, p0 := (p>>6)&1, (p>>5)&1, (p>>2)&1, (p>>1)&1, p&1
		p = p6<<6 + p5<<5 + 0b10000 + p2<<2 + p1<<1 + p0 // korekcja wyniku
	}

	// następnie trzeba zamienić liczbę binarną p z powrotem na bcd

	p_bcd := binaryToBCD(p)
	fmt.Printf("a: %04b * b: %04b = p(BIN): %07b, p(BCD):%04b, %d\n", a, b, p, p_bcd, bcdToDecimal(p_bcd))

	return p_bcd

	// return p
}

func main() {
	decimal1 := 123
	decimal2 := 456
	decimal3 := 789

	bcd1 := decimalToBCD(decimal1)
	bcd2 := decimalToBCD(decimal2)
	bcd3 := decimalToBCD(decimal3)

	fmt.Printf("BCD representation of %d: %04b\n", decimal1, bcd1)
	fmt.Printf("BCD representation of %d: %04b\n", decimal2, bcd2)
	fmt.Printf("BCD representation of %d: %04b\n", decimal3, bcd3)

	var k byte = 0b10100
	bin1 := binaryToBCD(k)

	fmt.Printf("BCD representation of %bb: %04b\n", k, bin1)

	fmt.Printf("Decimal 1: %d\n", bcdToDecimal(bcd1))
	fmt.Printf("Decimal 2: %d\n", bcdToDecimal(bcd2))
	fmt.Printf("Decimal 3: %d\n", bcdToDecimal(bcd3))

	bcd4 := decimalToBCD(4)
	bcd5 := decimalToBCD(5)

	result := multiplySingleBCD(bcd4[0], bcd5[0])
	fmt.Printf("Result of %04b * %04b = %04b (%d)\n\n", bcd4[0], bcd5[0], result, bcdToDecimal(result))

	/*	bcdResult := multiplyBCD(bcd1, bcd2)
		decimalResult := bcdToDecimal(bcdResult)

		fmt.Printf("BCD result: %v\n", bcdResult)
		fmt.Printf("Decimal result: %d\n", decimalResult)*/

	var multResults []byte

	for i := range 10 {
		for j := range 10 {
			//fmt.Printf("%d * %d", i, j)
			multResult := multiplySingleBCD(decimalToBCD(i)[0], decimalToBCD(j)[0])
			multResults = append(multResults, multResult...)
		}
	}

	/*for _, r := range multResults {
		fmt.Printf("%07b == %d\n", r, r)
	}*/
}
