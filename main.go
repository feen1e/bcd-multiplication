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

const maxAllowed = 3_037_000_499

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
				multiplyRandom(r, maxAllowed)
			}
		case "3":
			{
				multiplyRandom(r, 9999)
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

// multiplyRandom generuje dwie losowe liczby całkowite i wykonuje mnożenie
// r — generator liczb losowych
// maxRandom — maksymalna wartość losowanej liczby
func multiplyRandom(r *rand.Rand, maxRandom int64) {
	a := r.Int63n(maxRandom)
	b := r.Int63n(maxRandom)
	fmt.Printf("Wylosowane liczby: %d i %d\n\n", a, b)

	multiplyAndPrint(strconv.FormatInt(a, 10), strconv.FormatInt(b, 10))
}

// multiplyRandomFloat generuje dwie losowe liczby zmiennoprzecinkowe z maksymalnie 5 miejscami po przecinku i wykonuje mnożenie
func multiplyRandomFloat(r *rand.Rand) {
	// Generowanie losowych liczb zmiennoprzecinkowych
	a := r.Float64() * 100 // Liczby od 0 do 100
	b := r.Float64() * 100

	// Formatowanie liczb do stringów z ograniczeniem do 5 miejsc po przecinku
	aStr := strconv.FormatFloat(a, 'f', 5, 64)
	bStr := strconv.FormatFloat(b, 'f', 5, 64)

	fmt.Printf("Wylosowane liczby: %s i %s\n\n", aStr, bStr)

	multiplyAndPrint(aStr, bStr)
}

// multiplyRandom4DigitFloat generuje dwie losowe liczby 4-cyfrowe z maksymalnie 3 miejscami po przecinku i wykonuje mnożenie
func multiplyRandom4DigitFloat(r *rand.Rand) {
	// Generowanie losowych liczb 4-cyfrowych (1000-9999) z maksymalnie 3 miejscami po przecinku
	a := 1000 + r.Float64()*9000
	b := 1000 + r.Float64()*9000

	// Formatowanie liczb do stringów z ograniczeniem do 3 miejsc po przecinku
	aStr := strconv.FormatFloat(a, 'f', 3, 64)
	bStr := strconv.FormatFloat(b, 'f', 3, 64)

	fmt.Printf("Wylosowane liczby: %s i %s\n\n", aStr, bStr)

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

	// Zamień przecinki na kropki (dla polskiej notacji)
	aStr = strings.Replace(aStr, ",", ".", -1)
	bStr = strings.Replace(bStr, ",", ".", -1)

	a, errA := strconv.ParseFloat(aStr, 64)
	b, errB := strconv.ParseFloat(bStr, 64)

	if errA != nil || errB != nil {
		fmt.Println("Błąd: Nieprawidłowa liczba")
		return
	}

	if a > maxAllowed || b > maxAllowed {
		fmt.Printf("Liczby muszą być mniejsze niż %d\n", maxAllowed)
		return
	}

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

	fmt.Print("    ")
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
