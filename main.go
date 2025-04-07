package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	menu()
}

func menu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("*—————————————————MENU—————————————————*")
		fmt.Println("| [1] Mnożenie dwóch wybranych liczb   |")
		fmt.Println("| [2] Mnożenie dwóch losowych liczb    |")
		fmt.Println("| [0] Wyjście z programu               |")
		fmt.Println("*——————————————————————————————————————*")
		fmt.Print("Wprowadź numer wybranej opcji: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		fmt.Println()

		switch choice {
		case "1":
			{
				multiplyChosen(reader)
			}
		case "2":
			{
				multiplyRandom()
			}
		case "0":
			{
				return
			}
		default:
			{
				fmt.Println("Nieprawidłowy wybór")
			}
		}
	}
}

func multiplyChosen(reader *bufio.Reader) {
	fmt.Print("Wprowadź mnożną: ")
	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)
	a, _ := strconv.Atoi(aStr)
	fmt.Print("Wprowadź mnożnik: ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)
	b, _ := strconv.Atoi(bStr)
	aBCD := DecimalToBCD(a)
	bBCD := DecimalToBCD(b)
	result := MultiplyBCD(aBCD, bBCD)
	printBCDMultiplication(aBCD, bBCD, result)

	for _, d := range result {
		fmt.Printf("%04b ", d)
	}
	fmt.Printf("(BCD) -> %d (DEC)\n", BCDToDecimal(result))
	fmt.Printf("Sprawdzenie: %d * %d = %d\n\n", a, b, a*b)
}

func printBCDMultiplication(a []byte, b []byte, result []byte) {
	maxLen := max(len(a), len(b), len(result))

	fmt.Print("    ")
	printBCDRow(a, maxLen)

	fmt.Print(" *  ")
	printBCDRow(b, maxLen)

	for i := 0; i <= maxLen; i++ {
		fmt.Print("—————")
	}
	fmt.Println()

	fmt.Print("    ")
	printBCDRow(result, maxLen)
	fmt.Println()
}

func printBCDRow(bcd []byte, width int) {
	pad := width - len(bcd)
	for i := 0; i < pad; i++ {
		fmt.Print("     ")
	}

	for _, d := range bcd {
		fmt.Printf("%04b ", d)
	}
	fmt.Println()
}

func multiplyRandom() {
	// tak tymczasowo
	dec1 := []int{1234, 2345, 3456, 4567, 5678, 6789, 7890, 8901, 9012}
	dec2 := []int{4321, 5432, 6543, 7654, 8765, 9876, 1098, 2109, 3210}

	for i := 0; i < len(dec1); i++ {
		bcdA := DecimalToBCD(dec1[i])
		bcdB := DecimalToBCD(dec2[i])

		result2 := MultiplyBCD(bcdA, bcdB)

		fmt.Printf("Mnożenie: %d * %d = %d \n", dec1[i], dec2[i], dec1[i]*dec2[i])
		fmt.Printf("Wynik: %04b = %d \n \n", result2, BCDToDecimal(result2))
	}
}
