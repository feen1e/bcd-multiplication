package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	menu()
}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

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
				multiplyRandom(r)
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
	a, b := 10000, 10000
	for a > 9999 || b > 9999 {
		fmt.Print("Wprowadź mnożną: ")
		aStr, _ := reader.ReadString('\n')
		aStr = strings.TrimSpace(aStr)
		a, _ = strconv.Atoi(aStr)
		fmt.Print("Wprowadź mnożnik: ")
		bStr, _ := reader.ReadString('\n')
		bStr = strings.TrimSpace(bStr)
		b, _ = strconv.Atoi(bStr)
		if a > 9999 || b > 9999 {
			fmt.Println("Liczby mogą być maksymalnie 4-cyfrowe.")
		}
	}

	multiplyAndPrint(a, b)
}

func multiplyRandom(r *rand.Rand) {
	a := r.Intn(10000)
	b := r.Intn(10000)
	fmt.Printf("Wylosowane liczby: %d i %d\n\n", a, b)

	multiplyAndPrint(a, b)
}

func multiplyAndPrint(a int, b int) {
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
