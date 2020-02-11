// Інколи, необхідно запускати інші процеси. Наприклад, підсвітка
// синтаксису на цьому сайті [реалізована](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)
// стороннім процесом [`pygmentize`](http://pygments.org/) запущеним
// з Go програми. Давайте розглянемо кілька прикладів запуску
// процесів з Go.

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// Розпочнемо ми з простої команди, яка не приймає
	// жодних аргументів або вводу, і яка лише друкує "щось"
	// до `stdout`. Функція `exec.Command` створить об'єкт
	// який представляє собою цей зовнішній процес.
	dateCmd := exec.Command("date")

	// `.Output` - інший метод, який бере на себе управління
	// загальними випадками запуску процесу, очікує його
	// завершення та отримання виводу. Якщо помилок не було,
	// [зріз](slices) `dateOut` буде наповнено байтами відповіді
	// процесу `date`, який ми викликали.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Тепер ми готові розглянути більш просунутий приклад,
	// де ми перенаправимо дані у [`потік введення`](https://uk.wikipedia.org/wiki/Стандартні_потоки#Стандартне_введення) зовнішнього процесу та
	// отримаємо результати з його [`потоку виведення`](https://uk.wikipedia.org/wiki/Стандартні_потоки#Стандартне_введення).
	grepCmd := exec.Command("grep", "hello")

	// Ми чітко вказуємо - "захопи перенаправлення вводу та виводу",
	// запускаємо процес, виконуємо необхідне введення інформації
	// до нього, читаємо результатний вивід і, нарешті, очікуємо
	// завершення процесу.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// Ми пропустили перевірку на помилки у попередньому прикладі,
	// але ви можете використовувати звичайний прийом `if err != nil`
	// для цього. Ми вже скористались результатом отриманим з `StdoutPipe`,
	// у схожим чином, ви можете отримати відповідь що надходить з `StderrPipe`.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Зауважимо - коли ми породжуємо процеси нам необхідно надати
	// чітко розділені команду і масив аргументів, напротивагу передачі
	// всієї команди записаної в один рядок. Якщо ж вам, таки, кортить
	// передати все одним рядком, скористайтесь `bash`'овим параметром `-c`:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
