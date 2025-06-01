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
	m1 = (a0 & b1 & (^a1 | ^b0)) | (a1 & b0 & (^a0 | ^b1))
	m2 = a1 & b1 & (^a0 | ^b0)
	m3 = a1 & a0 & b1 & b0
	return
}

func mul3x3(a, b byte) (p210, p543 byte) {
	// 3.2.2 3x3 Multiplier
	a2, a1, a0 := (a>>2)&1, (a>>1)&1, a&1
	b2, b1, b0 := (b>>2)&1, (b>>1)&1, b&1
	m3, m2, m1, m0 := mul2x2(a1, a0, b1, b0)
	p1pp, p0pp := m1, m0

	// FA1
	p2pp, carry1 := fullAdder(a0&b2, a2&b0, m2)
	// FA2
	haIn, carry2 := fullAdder(a1&b2, a2&b1, carry1)
	// HA
	p3pp, carryHa := halfAdder(haIn, m3)
	// FA3
	p5pp, p4pp := fullAdder(a2&b2, carry2, carryHa)

	p210 = p2pp<<2 | p1pp<<1 | p0pp
	p543 = p5pp<<2 | p4pp<<1 | p3pp
	return
}

func circa3b3() {

}

func dataDrivingLogic() {

}

func correction() {

}

func finalTwoDigitMultiplier() {

}
