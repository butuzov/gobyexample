// Буває, що нашою метою є робота програми з
// колекцією даних, наприклад вибір даних, що задовольняють
// певний предикат (умову) або розмітка даних у вигляді нової колекції
// за допомогою власних (користувацьких) функцій. Деякі мови,
// надають моживіть використовувати нам (в таких випадках)
// [узагальнені (generic)](https://uk.wikipedia.org/wiki/Узагальнене_програмування) структури даних та алгоритми.
// Go не підтримує таких узагальнень, бо існує загальноприйнята
// практика створювати та використовувати функцій, які працюватимуть
// з колекціями - тоді і тільки тоді, коли в них виникає потреба.

// Нижче наведено кілька прикладів, що оперуватимуть над зрізом рядків.
// Скористайтесь ними як заготовками для створення власного функціоналу.
// Зауважте, в деяких випадках набагато ефективніше скористатись
// прямою вставкою коду, без створення функції.

package main

import (
	"fmt"
	"strings"
)

// Index повертає індекс першого елементу еквівалентного `t`,
// або -1 якщо збігів не знайдено.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// `Include` повертає `true`, якщо індекс з значенням `t`
// знаходиться у шуканому зрізі.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// `Any` повертає `true`, якщо хоча б один з рядків задовольняє [предикат](https://uk.wikipedia.org/wiki/Предикат) `f`.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// `All` повертає `true`, якщо всі елементи зрізу задовольняють предикат `f`.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// `Filter` повертає новий зріз, що містить усі елементи
// відфільтровані по критерію предиката `f`
// (такі, що задовольняють предикат).
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// `Map` повертає новий зріз, що містить результати роботи
// предиката `f` над кожним елементом оригінального зрізу.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// Приклади застосування вищенаведених функцій відносно
	// невеликої колекції рядків.
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// Вищенаведенні приклади використовують анонімні функції,
	// але ніщо не заважає передавати звичайні іменовані функції,
	// що задовольнятимуть типу предиката.

	fmt.Println(Map(strs, strings.ToUpper))
}
