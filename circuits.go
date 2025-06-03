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
	m1 = (a0 & b1 & ((^a1 | ^b0) & 1)) | (a1 & b0 & ((^a0 | ^b1) & 1)) // do operacji NOT dodano AND 1 jako maske,
	m2 = a1 & b1 & ((^a0 | ^b0) & 1)                                   // ktora zapewnia ze tylko ostatni bit jest brany pod uwage
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

func finalTwoDigitMultiplier(a, b byte) (p byte) {
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
