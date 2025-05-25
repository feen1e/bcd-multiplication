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

// main jest funkcją startową programu
func main() {
	menu()
}

// menu wyświetla menu główne programu
func menu() {
	reader := bufio.NewReader(os.Stdin)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for {
		fmt.Println()
		fmt.Println("*——————————————————————————MENU——————————————————————————*")
		fmt.Println("| [1] Mnożenie dwóch wybranych liczb                     |")
		fmt.Println("| [2] Mnożenie dwóch losowych liczb całkowitych          |")
		fmt.Println("| [3] Mnożenie dwóch losowych max 4-cyfrowych liczb int  |")
		fmt.Println("| [4] Mnożenie dwóch losowych liczb zmiennoprzecinkowych |")
		fmt.Println("| [5] Mnożenie dwóch losowych 4-cyfrowych liczb float    |")
		fmt.Println("| [0] Wyjście z programu                                 |")
		fmt.Println("*————————————————————————————————————————————————————————*")
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
				multiplyRandomInt(r)
			}
		case "3":
			{
				multiplyRandom4DigitInt(r)
			}
		case "4":
			{
				multiplyRandomFloat(r)
			}
		case "5":
			{
				multiplyRandom4DigitFloat(r)
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

// multiplyRandomInt generuje dwie losowe liczby całkowite i wykonuje mnożenie
// r — generator liczb losowych
func multiplyRandomInt(r *rand.Rand) {
	a := r.Int63n(int64(math.Pow10(r.Intn(20))))
	b := r.Int63n(int64(math.Pow10(r.Intn(20))))
	fmt.Printf("Wylosowane liczby: %d i %d\n", a, b)

	multiplyAndPrint(strconv.FormatInt(a, 10), strconv.FormatInt(b, 10))
}

// multiplyRandom4DigitInt generuje dwie losowe max 4-cyfrowe liczby całkowite i wykonuje mnożenie
// r — generator liczb losowych
func multiplyRandom4DigitInt(r *rand.Rand) {
	a := r.Int63n(9999)
	b := r.Int63n(9999)
	fmt.Printf("Wylosowane liczby: %d i %d\n", a, b)

	multiplyAndPrint(strconv.FormatInt(a, 10), strconv.FormatInt(b, 10))
}

// multiplyRandomFloat generuje dwie losowe liczby zmiennoprzecinkowe i wykonuje mnożenie
// r — generator liczb losowych
func multiplyRandomFloat(r *rand.Rand) {
	// Generowanie losowych liczb zmiennoprzecinkowych
	aPrec := r.Intn(10)
	bPrec := r.Intn(10)
	a := 1 + r.Float64()*float64(r.Int63n(int64(math.Pow10(aPrec))))
	b := 1 + r.Float64()*float64(r.Int63n(int64(math.Pow10(bPrec))))

	// Formatowanie liczb do stringów
	aStr := strconv.FormatFloat(a, 'f', aPrec, 64)
	bStr := strconv.FormatFloat(b, 'f', bPrec, 64)

	fmt.Printf("Wylosowane liczby: %s i %s\n", aStr, bStr)

	multiplyAndPrint(aStr, bStr)
}

// multiplyRandom4DigitFloat generuje dwie losowe liczby 4-cyfrowe i wykonuje mnożenie
// r — generator liczb losowych
func multiplyRandom4DigitFloat(r *rand.Rand) {
	// Generowanie losowych liczb 4-cyfrowych (1000-9999)
	a := 1000 + r.Float64()*9000
	b := 1000 + r.Float64()*9000

	// Formatowanie liczb do stringów
	aPrec := r.Intn(10)
	bPrec := r.Intn(10)
	aStr := strconv.FormatFloat(a, 'f', aPrec, 64)
	bStr := strconv.FormatFloat(b, 'f', bPrec, 64)

	fmt.Printf("Wylosowane liczby: %s i %s\n", aStr, bStr)

	multiplyAndPrint(aStr, bStr)
}

// multiplyChosen pozwala użytkownikowi wprowadzić dwie liczby i wykonuje mnożenie
// reader — bufor do odczytu danych wprowadzanych przez użytkownika
func multiplyChosen(reader *bufio.Reader) {
	fmt.Print("Wprowadź pierwszą liczbę: ")
	aStr, _ := reader.ReadString('\n')
	aStr = strings.TrimSpace(aStr)

	fmt.Print("Wprowadź drugą liczbę: ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)

	// Zamiana przecinków na kropki (dla polskiej notacji)
	aStr = strings.Replace(aStr, ",", ".", -1)
	bStr = strings.Replace(bStr, ",", ".", -1)

	multiplyAndPrint(aStr, bStr)
}

// printBCDMultiplication wyświetla mnożenie pisemne liczb BCD
// a — mnożna BCD
// aDecPos — pozycja przecinka w a (licząc od prawej strony)
// b — mnożnik BCD
// bDecPos — pozycja przecinka w b (licząc od prawej strony)
// result — wynik mnożenia BCD
// resultDecPos — pozycja przecinka w wyniku (licząc od prawej strony)
func printBCDMultiplication(a []byte, aDecPos int, b []byte, bDecPos int, result []byte, resultDecPos int) {
	maxLen := max(len(a), len(b), len(result))

	fmt.Print("\n    ")
	printBCDRow(a, aDecPos, maxLen)

	fmt.Print(" *  ")
	printBCDRow(b, bDecPos, maxLen)

	for i := 0; i <= maxLen; i++ {
		fmt.Print("—————")
	}
	fmt.Println()

	fmt.Print("    ")
	printBCDRow(result, resultDecPos, maxLen)
	fmt.Println()
}

// printBCDRow wyświetla jeden wiersz liczby BCD z uwzględnieniem pozycji przecinka
// bcd — liczba BCD do wyświetlenia
// decPos — pozycja przecinka (licząc od prawej strony)
// width — szerokość wiersza (dla wyrównania)
func printBCDRow(bcd []byte, decPos int, width int) {
	pad := width - len(bcd)
	for i := 0; i < pad; i++ {
		fmt.Print("     ")
	}

	for i, d := range bcd {
		// Sprawdź, czy należy wyświetlić przecinek
		if decPos > 0 && i == len(bcd)-decPos-1 {
			fmt.Printf("%04b.", d)
		} else {
			fmt.Printf("%04b ", d)
		}
	}
	fmt.Println()
}

// multiplyAndPrint wykonuje mnożenie dwóch liczb podanych jako stringi i wyświetla wynik
// aStr — pierwsza liczba jako string
// bStr — druga liczba jako string
func multiplyAndPrint(aStr string, bStr string) {
	// Konwersja stringów na BCD
	aBCD, aDecPos := StringToBCD(aStr)
	bBCD, bDecPos := StringToBCD(bStr)

	// Mnożenie BCD z uwzględnieniem części ułamkowej
	resultBCD, resultDecPos := MultiplyFloatBCD(aBCD, aDecPos, bBCD, bDecPos)

	// Wyświetlenie mnożenia pisemnego
	printBCDMultiplication(aBCD, aDecPos, bBCD, bDecPos, resultBCD, resultDecPos)

	// Konwersja wejściowych stringów na float64 dla porównania
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)
	regularResult := a * b
	resultStr := BCDToString(resultBCD, resultDecPos)

	// Wyświetlenie wyniku mnożenia dziesiętnego dla porównania z BCD
	fmt.Printf("Mnożenie BCD:        %s * %s = %s\n", aStr, bStr, resultStr)
	fmt.Printf("Mnożenie dziesiętne: %s * %s = %.30g\n", aStr, bStr, regularResult)
}
