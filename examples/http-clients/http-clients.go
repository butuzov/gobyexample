// Стандартна бібліотека Go поставляється з чудовою підтримкою
// HTTP клієнтів і серверів. Дана підтримка забезпечується пакетом `net/http`.
// Даний приклад продемонструє створення простого HTTP запиту.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Створити *µHTTP* запит з методом *GET*. `http.Get` являється
	// зручною обгорткою для створення *HTTP*-клієнта, і відправки *GET* запиту.
	// Дана обгортка використовує `http.DefaultClient` під капотом,
	// що має подібні стандартні налаштування.
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	// Необхідно також не забувати _відкласти_ закриття рідера(Reader)
	// за допомогою [`defer`](./defer), для того щоб не допустити *витік ресурсів*.
	defer resp.Body.Close()

	// Вивести *HTTP* статус відповіді.
	fmt.Println("Response status:", resp.Status)

	// Вивести перші 5 ліній з тіла відповіді
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
