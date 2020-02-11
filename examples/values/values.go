// Go має різні типи значень, включаючи: рядки (strings),
// цілі числа (int), числа з рухомою комою (float),
// логічні (bool) абощо. Ось кілька базових прикладів.

package main

import "fmt"

func main() {

	// Рядки можуть бути об'єднані поміж собою за допомогою
	// оператора `+`.
	fmt.Println("go" + "lang")

	// Цілі числа та числа з плаваючою комою.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Логічний тип та логічні оператор.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
