package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxAllowed = 3_037_000_499

func main() {
	menu()
}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for {
		fmt.Println("*———————————————————————MENU————————————————————————*")
		fmt.Println("| [1] Mnożenie dwóch wybranych liczb                |")
		fmt.Println("| [2] Mnożenie dwóch losowych liczb                 |")
		fmt.Println("| [3] Mnożenie dwóch losowych max 4-cyfrowych liczb |")
		fmt.Println("| [0] Wyjście z programu                            |")
		fmt.Println("*———————————————————————————————————————————————————*")
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
				multiplyRandom(r, maxAllowed)
			}
		case "3":
			{
				multiplyRandom(r, 9999)
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
	a, b := int64(math.MaxInt64), int64(math.MaxInt64)

	for a > maxAllowed || b > maxAllowed {
		fmt.Print("Wprowadź mnożną: ")
		aStr, _ := reader.ReadString('\n')
		aStr = strings.TrimSpace(aStr)
		a, _ = strconv.ParseInt(aStr, 10, 64)
		fmt.Print("Wprowadź mnożnik: ")
		bStr, _ := reader.ReadString('\n')
		bStr = strings.TrimSpace(bStr)
		b, _ = strconv.ParseInt(bStr, 10, 64)
		if a > maxAllowed || b > maxAllowed {
			fmt.Printf("Liczby muszą być mniejsze niż %d\n", maxAllowed)
		}
	}

	multiplyAndPrint(a, b)
}

func multiplyRandom(r *rand.Rand, maxRandom int64) {
	a := r.Int63n(maxRandom)
	b := r.Int63n(maxRandom)
	fmt.Printf("Wylosowane liczby: %d i %d\n\n", a, b)

	multiplyAndPrint(a, b)
}

func multiplyAndPrint(a int64, b int64) {
	aBCD := DecimalToBCD(a)
	bBCD := DecimalToBCD(b)
	result := MultiplyBCD(aBCD, bBCD)
	printBCDMultiplication(aBCD, bBCD, result)

	for _, d := range result {
		fmt.Printf("%04b ", d)
	}
	resultDec := BCDToDecimal(result)
	fmt.Printf("(BCD) -> %d (DEC)\n", resultDec)
	fmt.Printf("Sprawdzenie: %d * %d = %d -> ", a, b, a*b)
	if resultDec == a*b {
		fmt.Printf("Wynik prawidłowy\n\n")
	} else {
		fmt.Printf("Wynik jest błędny. Coś poszło nie tak\n\n")
	}
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
