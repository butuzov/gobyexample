// Однією з базових операцій у програмуванні є читання та запис файлів.
// Зараз ми познайомимось, з тим - як читати файли у Go.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	var err error
	// Завантаження файлу у пам'ять - найбільш проста задача
	// пов'язана з читанням.
	dat, err := ioutil.ReadFile("/tmp/dat")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))

	// Але, зазвичай, нам необхідно більше контролю над процесом
	// чатання з файлів: напирклад - як і що читати?
	// Почніть з відкриття ( метод `Open`) файлу для отримання
	// дексриптору `os.File`.
	f, err := os.Open("/tmp/dat")
	if err != nil {
		panic(err)
	}

	// Прочитаємо деякі байти розташовані на початку файлу,
	// наприклад перші 5 байт, але зауважте - скільки байт
	// було прочитано насправді.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d байт: %s\n", n1, string(b1))

	// Метод `Seek` дозволяє переносити вказівник читання/запису
	// (або _перемотувати_)  до певної позиції у файлі,
	// а метод `Read` дозволить прочитати з цього місця.
	o2, err := f.Seek(6, 0)
	if err != nil {
		panic(err)
	}

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d байт @ %d: %s\n", n2, o2, string(b2))

	// Пакет `io` пропонує вам певний функціонал, що стане в нагоді
	// під час "читання". Наприклад, прикладі зверху, може
	// бути надійніше імплементовано з `ReadAtLeast`.
	o3, err := f.Seek(6, 0)
	if err != nil {
		panic(err)
	}
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d байт @ %d: %s\n", n3, o3, string(b3))

	// Скористайтесь `Seek(0, 0)` для повернення вказівника читання
	// до початоку файла (назразок "перемотки касети на початок").
	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// Пакет `bufio` впроваджує буферизоване читання,
	// що є більш ефективним - завдяки багатьом невеликим
	// зчитуванням і додатковими методам які реалізує пакет.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("5 байт: %s\n", string(b4))

	// Закривайте файл - у разі, якщо ви з ним закінчили
	// працювати (зазвичай це роблять відразу по-відкритті
	// використовуючи [відкладені виклики](defer)).
	f.Close()
}
