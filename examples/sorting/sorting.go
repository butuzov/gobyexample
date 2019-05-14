// Пакет `sort` реалізує сортування для вбудованих
// та визначених користувачами типів. Спочатку познайомимось з
// сортуванням вбудованих типів.

package main

import "fmt"
import "sort"

func main() {

    // Ось один з прикладів роботи GO з вбудованими типами -
    // сортування рядків. Зауважте що сортування проходить
    // "на місці", тобто сортування проходе на переданому
    // зрізі, а новий не створюється і не повертається.
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Рядки:      ", strs)

    // Сортування цілих чисел - не дуже відрізняється від
    // сортування рядків.
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Цілі числа: ", ints)

    // Ми, також, можемо використати `sort` щоб перевірити
    // чи зріз вже відсортований.
    s := sort.IntsAreSorted(ints)
    fmt.Println("Сортовано:  ", s)
}
