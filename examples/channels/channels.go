// _Канали_ це магістралі для обміну повідомленнями між горутинами,
// які виконуються водночас. Ми можете надсилати
// дані до каналу в однієї горутини та отримаувати у іншій.

package main

import "fmt"

func main() {

    // Канали створюються наступним синтаксисом `make(chan тип-каналу)` де
    // тип-каналу визначає значення якого типу можуть бути по ньому передані.
    messages := make(chan string)

    // _Надіслати_ значення до каналу можна використовуючи наступний
    // `channel <-` синтаксис. Використовуючи анонімну горутину,
    // ми надішлемо значення `ping` до каналу `messages`.
    go func() { messages <- "ping" }()

    // Наступний синтаксис `<-channel` описує _отримання_ значення з каналу.
    // Так ми отримуємо і друкуємо повідомлення `ping`, яке прийшло каналом
    // `messages` з анонімної горутини.
    msg := <-messages
    fmt.Println(msg)
}
