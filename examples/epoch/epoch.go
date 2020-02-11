// Нерідко потрібно взнати час що пройшов з моменту
// [початку часу Unix](https://uk.wikipedia.org/wiki/Час_Unix).
// Цей час можна представити в секундах, мілісекундах,
// або наносекундах - і ось як ми це робимо в Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Для отримання секунд або наносекунд з моменту початку Unix
	// скористаємось методом `Unix` або `UnixNano` об'єкта `time.Now`
	// відповідно.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Зауважте - не існує `UnixMillis`, а отже, щоб отримати
	// час в мілісекундах - необхідно поділити наносекунди
	// на 100 000 (кількість наносекунд в мілісекунді).
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// Дозволяється конвертувати цілочисельні секунди
	// або наносекунди у відповідний `time`.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
