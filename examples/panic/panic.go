// `panic` або `паніка` зазвичай означає те, що щось пішло
// неочікувано погано. Дуже погано. В більшості випадків, ми
// використовуємо її задля швидкого провалу
// на помилках які зазвичай не мали б виникати, або
// на помилках які ми не здадтні опрацювати акуратно.

package main

import "os"

func main() {

	// Ми скористаємось панікою в цьому прикладі для індикації
	// неочікуваної помилки. Єдина мета цього шматка коду
	// запаніуквати.
	panic("маємо проблему")

	// Загальне використання `panic`и - перервати виконання функції,
	// воно повертає помилку яку ми не знаємо як опрацьовувати
	// (або не хочемо). Ось приклад `panic`и де ми отримали
	// неочікувану помилку, під час створення файлу, з якою ми не
	// дуже хочемо мати справу.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
