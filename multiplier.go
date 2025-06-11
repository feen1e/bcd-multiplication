package main

// MultiplyFloatBCD mnoży liczby BCD z uwzględnieniem części ułamkowej
// a — mnożna BCD
// aDecimalPlaces — pozycja przecinka w a (licząc od prawej strony)
// b — mnożnik BCD
// bDecimalPlaces — pozycja przecinka w b (licząc od prawej strony)
// Zwraca wynik mnożenia BCD oraz pozycję przecinka w wyniku
func MultiplyFloatBCD(a []byte, aDecimalPlaces int, b []byte, bDecimalPlaces int) ([]byte, int) {
	// Mnożymy liczby BCD bez uwzględniania przecinka
	result := MultiplyBCD(a, b)

	// Obliczamy nową pozycję przecinka (suma pozycji przecinków w mnożnej i mnożniku)
	resultDecimalPlaces := aDecimalPlaces + bDecimalPlaces

	return result, resultDecimalPlaces
}

// MultiplyBCD mnoży liczby BCD
// a — mnożna
// b — mnożnik
func MultiplyBCD(a, b []byte) []byte {
	result := make([]byte, len(a)+len(b)) // Maksymalna długość wyniku
	n := len(result)

	// i — indeks cyfry b (mnożnik)
	for i := len(b) - 1; i >= 0; i-- {
		// j — indeks cyfry a (mnożna)
		for j := len(a) - 1; j >= 0; j-- {

			//partial := MultiplySingleDigitBCD(a[j], b[i]) // - algorytm z Fig. 2, poniżej wersja z układem z Fig. 7
			partial := BinaryToBCDConverter(FinalTwoDigitMultiplier(a[j], b[i])) // wynik częściowy
			position := n - 1 - ((len(a) - 1 - j) + (len(b) - 1 - i))            // wyliczenie pozycji w wyniku

			// Dodanie wyniku częściowego do wyniku głównego
			k := len(partial) - 1
			carry := byte(0)

			for k >= 0 && position >= 0 {
				sum := result[position] + partial[k] + carry
				result[position] = sum % 10
				carry = sum / 10
				k--
				position--
			}

			// Obsługa ewentualnego przeniesienia poza wynik częściowy (partial)
			for carry > 0 && position >= 0 {
				sum := result[position] + carry
				result[position] = sum % 10
				carry = sum / 10
				position--
			}
		}
	}

	// Zwracamy tylko część bez zer na początku (jeśli występują)
	start := 0
	for start < len(result)-1 && result[start] == 0 {
		start++
	}
	return result[start:]
}

// MultiplySingleDigitBCD mnoży dwie 4-bitowe cyfry BCD; implementacja na podstawie Fig. 2. w artykule
// a — mnożna
// b — mnożnik
func MultiplySingleDigitBCD(a, b byte) []byte {
	var p byte = 0
	// Wycięcie poszczególnych bitów z A i B
	a3, a2, a1, a0 := (a>>3)&1, (a>>2)&1, (a>>1)&1, a&1
	b3, b2, b1, b0 := (b>>3)&1, (b>>2)&1, (b>>1)&1, b&1

	// Implementacja algorytmu z artykułu (Modified binary multiplication algorithm for BCD digits)
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

	// Przywrócenie wyniku binarnego na liczbę BCD
	pBcd := BinaryToBCDConverter(p)
	return pBcd
}
