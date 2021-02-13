// Структури в Go - це колекції полів з визначеним типом.
// Надзвичайну корисність структур можна оцінити не тільки для
// групування даних, а й тому що вони служать основним рушієм мови
// що орієнтується на данні - на структури.

package main

import "fmt"

// Структура `person` має поля для імені (`name`) та віку (`age`).
type person struct {
	name string
	age  int
}

// `newPerson` constructs a new person struct with the given name.
func newPerson(name string) *person {
	// You can safely return a pointer to local variable
	// as a local variable will survive the scope of the function.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Наступний синтаксис створить нову структуру.
	// Дозволяється іменувати поля при ініціалізації структури.
	fmt.Println(person{name: "Alice", age: 30})

	// А можна і уникнути, іменованої ініціалізації.
	fmt.Println(person{"Bob", 20})

	// Пропущені поля будуть створені відповідно нульового.
	// значення типу поля, що ми оминаємо
	fmt.Println(person{name: "Fred"})

	// Префікс `&` поверне вказівник на структуру.
	fmt.Println(&person{name: "Ann", age: 40})

	// Приховувати процес створення стуктури іншою функцією - це
	// ідіоматичний підхід у Go.
	fmt.Println(newPerson("Jon"))

	// Доступ до полів надається через синтаксис крапки `.`
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Дозволяється використовувати крапки з вказівниками структур,
	// вказівники, в такому разі, автоматично розіменовуються.
	sp := &s
	fmt.Println(sp.age)

	// Дані у структурі можна змінювати (тобто вони `mutable`,
	// такі що дозволяється змінювати).
	sp.age = 51
	fmt.Println(sp.age)
}
