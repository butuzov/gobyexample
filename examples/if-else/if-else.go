// Оператори розгалуження `if/else` в Go це просто.

package main

import "fmt"

func main() {

	// Ось базовий приклад, де ми бичимо - що нам <br/> не потрібні
	// круглі дужки навколо умови в Go, хоча потреба фігурних
	// дужках нікуди не зникла.
	if 7%2 == 0 {
		fmt.Println("7 парне")
	} else {
		fmt.Println("7 не парне")
	}

	// Використовуємо ключове слово `if`, без `else`
	if 8%4 == 0 {
		fmt.Println("8 ділиться на 4")
	}

	// Ми можемо декларувати зміні що будуть доступні у
	// всьому розгалуженні (не тільки `if`, але й у `else if`
	// та `else`).
	if num := 9; num < 0 {
		fmt.Println(num, "є негативним числом")
	} else if num < 10 {
		fmt.Println(num, "має одну цифру")
	} else {
		fmt.Println(num, "має багато цифр")
	}
}
