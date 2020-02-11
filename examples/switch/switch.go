// Оператор розгалуження `switch` використовують,
// як правило, для представлення у вигляді розгалуження
// умов.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Розглянемо базовий приклад `switch`.
	i := 2
	fmt.Print("Напишемо ", i, " як ")
	switch i {
	case 1:
		fmt.Println("один")
	case 2:
		fmt.Println("два")
	case 3:
		fmt.Println("три")
	}

	// Дозволяється використовувати коми (",") - так ми розділяємо
	// кілька умов в одній гілці `case`. Також існує умова `default` -
	// вона відбудеться, за умови що жодна інша умова ще не виконалась.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Це вихідні")
	default:
		fmt.Println("Це робочий день")
	}

	// `switch` без виразу - представляє собою альтернативний спосіб
	// реалізувати `if/else` логіку. Цей приклад демонструє,
	// що гілки `case` не обмежується [константами](constants).
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Над полудень")
	default:
		fmt.Println("По полудні")
	}

	// Так само - як і в попередньому прикладі з `if`,
	// оператор `switch` дозволяє декларувати змінну,
	// яка буде доступна на усіх гілках перевірки умови.
	// Наразі, ми демонструємо це - визначенням типу
	// вхідної змінної (тип [`interface{}`](interfaces)
	// дозволяє працювати з будь-яким типом даних).
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Я логічна змінна")
		case int:
			fmt.Println("Я ціле число")
		default:
			fmt.Printf("Не знаю нічого про тип %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Гей!")
}
