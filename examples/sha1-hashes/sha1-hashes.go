// Алгоритм хешування [_SHA1_](https://uk.wikipedia.org/wiki/SHA-1)
// використовують для визначення коротких ідентифікаторів
// для бінарних або текстових даних. Він використовується наприклад
// в системі контролю версій [git](http://git-scm.com/), що
// використовує SHA1 для ідентифікації змін у файлах та директоріях.
// В Go визначити хеш SHA1 хеш можна наступним чином.

package main

// Стандартна бібліотека Go включає реалізацію кількох різноманітних
// алгоритмів хешування, вони доступні нам як пакети `crypto/*`.
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s1 := "це рядок sha1"

	// Шаблон для генерації хешу наступний - `sha1.New()`,
	// `sha1.Write(bytes)`, і наостанок `sha1.Sum([]byte{})`.
	// Почнеємо - з генерації нового хешу.
	h := sha1.New()

	// `Write` очікує байти і у випадку `s1` це рядок,
	// скористаємось `[]byte(s)` для конвертації його у байти.
	h.Write([]byte(s1))

	// Результатом цієї операції - стане SHA1 хеш у вигляді
	// зрізу байтів, який ми можемо представити як рядок пізніше.
	// Аргумент ( що передається методому `Sum`) може бути використано,
	// з ціллю додати його напочаток отриманого зрізу нашого хешу,
	// але насправді така операція зазвичай не потрібна (це не
	// "[солінння](https://uk.wikipedia.org/wiki/Сіль_(криптографія))"
	// як можна було б подумати спочатку).
	bs1 := h.Sum(nil)

	// SHA1 значення часто доступні у шістнадцятирічному
	// вигляді (наприклад у git commit'ах). Скористайтесь
	// дієсловом форматування `%x` - для конвертації хеш результату
	// в шістнадцятирічний рядок.
	fmt.Println(s1)
	fmt.Printf("%x\n", bs1)

	// А ще, можна скористатись функцієї-обгорткою `sha1.Sum`,
	// яка відразу надасть вам потрібний хеш у вигляді байтового зрізу.
	s2 := "використай sha1.Sum([]byte)"
	bs2 := sha1.Sum([]byte(s2))
	fmt.Println(s2)
	fmt.Printf("%x\n", bs2)
}
