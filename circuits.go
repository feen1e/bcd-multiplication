package main

func fullAdder(a, b, cin byte) (sum, cout byte) {
	sum = a ^ b ^ cin
	cout = (a & b) | (b & cin) | (a & cin)
	return
}

func halfAdder(a, b byte) (sum, carry byte) {
	sum = a ^ b
	carry = a & b
	return
}

func mux2to1(a, b, ctrl byte) byte {
	ctrl = ctrl & 1
	return (ctrl & b) | ((^ctrl & 1) & a)
}

func mul2x2(a1, a0, b1, b0 byte) (m3, m2, m1, m0 byte) {
	// 3.2.2 2x2 Multiplier
	m0 = a0 & b0
	m1 = (a0 & b1 & ((^a1 | ^b0) & 1)) | (a1 & b0 & ((^a0 | ^b1) & 1)) // do operacji NOT dodano AND 1 jako maskę,
	m2 = a1 & b1 & ((^a0 | ^b0) & 1)                                   // która zapewnia, że tylko ostatni bit jest brany pod uwagę
	m3 = a1 & a0 & b1 & b0
	//fmt.Printf("m3: %0b, m2: %0b, m1: %0b, m0: %0b\n", m3, m2, m1, m0)
	return
}

func mul3x3(a, b byte) (p5pp, p4pp, p3pp, p2pp, p1pp, p0pp byte) {
	// 3.2.2 3x3 Multiplier
	a2, a1, a0 := (a>>2)&1, (a>>1)&1, a&1
	b2, b1, b0 := (b>>2)&1, (b>>1)&1, b&1
	m3, m2, m1, m0 := mul2x2(a1, a0, b1, b0)
	p1pp, p0pp = m1, m0

	// FA1
	p2pp, carry1 := fullAdder(a0&b2, a2&b0, m2)
	// FA2
	haIn, carry2 := fullAdder(a1&b2, a2&b1, carry1)
	// HA
	p3pp, carryHa := halfAdder(haIn, m3)
	// FA3
	p4pp, p5pp = fullAdder(a2&b2, carry2, carryHa)

	//fmt.Printf("p''5: %0b, p''4: %0b, p''3: %0b, p''2: %0b, p''1: %0b, p''0: %0b\n", p5pp, p4pp, p3pp, p2pp, p1pp, p0pp)
	return
}

func circa3b3(a, b byte) (p6p, p5p, p4p, p3p byte) {
	// 3.1 Fig. 5. Circuit (a3 v b3) = 1
	a3, a2, a1, a0 := (a>>3)&1, (a>>2)&1, (a>>1)&1, a&1
	b3, b2, b1, b0 := (b>>3)&1, (b>>2)&1, (b>>1)&1, b&1

	p6p = a3 & b3
	p5p = mux2to1(a2, b2, a3)   // mux a2, b2
	p4p = mux2to1(a1, b1, a3)   // mux a1, b1
	a0b0 := mux2to1(a0, b0, a3) // mux a0, b0
	b0a0 := mux2to1(b0, a0, a3) // mux b0, a0
	p3pIN := b0a0 & p6p
	p3p = p3pIN | a0b0

	//fmt.Printf("p'6: %0b, p'5: %0b, p'4: %0b, p'p: %0b\n", p6p, p5p, p4p, p3p)
	return
}

func dataDrivingLogic(a3, b3, p5p, p4p, p3p, p5pp, p4pp, p3pp byte) (p5, r4, r3 byte) {
	// 3.3 Fig. 7. Data Driving Logic
	a3b3 := a3 | b3    // x
	and5 := p5p & a3b3 // x AND p'5
	and4 := p4p & a3b3 // x AND p'4
	and3 := p3p & a3b3 // x AND p'3
	p5 = and5 | p5pp   // and5 OR p''5
	r4 = and4 | p4pp   // and4 OR p''4
	r3 = and3 | p3pp   // and3 OR p''3

	/*fmt.Printf("DDL\na3: %0b, b3: %0b, a3&b3: %0b\n", a3, b3, a3b3)
	fmt.Printf("p'5: %0b, p'4: %0b, p'3: %0b, p''5: %0b, p''4: %0b, p''3: %0b\n", p5p, p4p, p3p, p5pp, p4pp, p3pp)
	fmt.Printf("and5: %0b, and4: %0b, and3: %0b\n", and5, and4, and3)
	fmt.Printf("p5: %0b, r4: %0b, r3: %0b\n", p5, r4, r3)*/
	return
}

func correction(p6p, r4, r3, p0pp byte) (p4, p3 byte) {
	// 3.4 Correction of 1001 x 1001
	and1 := p6p & p0pp
	p4 = (and1 & 1) | r4
	p3 = ((^and1) & 1) & r3

	//fmt.Printf("p4: %0b, p3: %0b\n", p4, p3)
	return
}

// FinalTwoDigitMultiplier oblicza iloczyn dwóch cyfr BCD; zwraca 7-bitową liczbę binarną. Implementacja na podstawie Fig 7.
func FinalTwoDigitMultiplier(a, b byte) (p byte) {
	// Fig. 7. Specific architecture for binary multiplier of two BCD digits
	p5pp, p4pp, p3pp, p2pp, p1pp, p0pp := mul3x3(a, b)
	p6p, p5p, p4p, p3p := circa3b3(a, b)
	a3, b3 := (a>>3)&1, (b>>3)&1
	p5, r4, r3 := dataDrivingLogic(a3, b3, p5p, p4p, p3p, p5pp, p4pp, p3pp)
	p4, p3 := correction(p6p, r4, r3, p0pp)
	p6, p2, p1, p0 := p6p, p2pp, p1pp, p0pp
	p = (p6&1)<<6 | (p5&1)<<5 | (p4&1)<<4 | (p3&1)<<3 | (p2&1)<<2 | (p1&1)<<1 | (p0 & 1)
	return
}

// BinaryToBCDConverter konwertuje 7-bitową liczbę binarną na dwie cyfry w reprezentacji BCD. Implementacja na podstawie sekcji 4.2.
func BinaryToBCDConverter(p byte) (bcdResult []byte) {
	// 4.2 Proposed binary-to-bcd converter
	p6, p5, p4, p3, p2, p1, p0 := (p>>6)&1, (p>>5)&1, (p>>4)&1, (p>>3)&1, (p>>2)&1, (p>>1)&1, p&1

	// Drugi artykuł, sekcja 'Arithmetic II' na stronie 37
	cc := ((p3&1)<<3 | (p2&1)<<2 | (p1&1)<<1 | p0&1) + ((p4 & 1) * 6) + ((p6&1)*4 + (p5&1)*2)
	dd := ((p6&1)<<2 | (p5&1)<<1 | p4&1) + ((p6&1)<<1 | p5&1)

	cy1, cy0 := byte(0), byte(0)
	if cc > 0b10011 {
		cy1 = 1
	}
	if 0b1001 < cc && cc < 0b10100 {
		cy0 = 1
	}

	cc3, cc2, cc1 := (cc>>3)&1, (cc>>2)&1, (cc>>1)&1
	dd2, dd1, dd0 := (dd>>2)&1, (dd>>1)&1, dd&1

	c321 := (cc3<<2 | cc2<<1 | cc1) + (cy1<<2 | (cy1|cy0)<<1 | cy0)
	c0 := p0
	d3 := p6 & p0
	d210 := (dd2<<2 | dd1<<1 | dd0) + (cy1<<1 | cy0)

	// Zapis jako tablica bajtów
	c := c321<<1 | c0
	c = c & 0b1111 // zapewnia, że cyfra jest 4-bitowa
	d := d3<<3 | d210
	d = d & 0b1111

	bcdResult = append(bcdResult, d)
	bcdResult = append(bcdResult, c)
	return
}
