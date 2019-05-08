// У Go існує підтримка _повернення кількох значень_ з функцій.
// Ця особилвість часто ідіоматично використовується в Go,
// для повернення результату роботи та помилки з функції.

package main

import "fmt"

// Конструкція `(int, int)` - цією декларацією функція
// показує, що вона зобов’‎язується повернути два цілих числа.
func vals() (int, int) {
    return 3, 7
}

func main() {

    // В цьому прикладі - ми декларуємо дві різні змінні шляхом
    // присвоєння результату функції (що повертає два значення).
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // Якщо ви плануєте викристати лише частину повернених
    // значень, скористайтесь пустим ідентифікатором (`_`).
    _, c := vals()
    fmt.Println(c)
}
