// Створення/Запис файлів у Go - слідує вже знайомим нам
// шаблонам розглянутим в [Читання файлів](reading-files).

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var err error

	// Почнемо з запису невеличкого зрізу байтів до файлу.
	d1 := []byte("привіт\ngo\n")
	err = ioutil.WriteFile("/tmp/dat1", d1, 0644)
	if err != nil {
		panic(err)
	}

	// Щоб записати до файлу деякі данні, нам необхідно його спочатку відкрити.
	f, err := os.Create("/tmp/dat2")
	if err != nil {
		panic(err)
	}

	// Зверніть увагу, на ідіоматично вірний відкладений виклик `Close`,
	// що ми ініціюємо відразу після відкриття.
	defer f.Close()

	// Запишемо, за допомогою методу `Write` ще один зріз байтів.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Записано %d байт\n", n2)

	// `WriteString` дозволить записати рядок.
	n3, err := f.WriteString("writes\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Записано %d байт\n", n3)

	// Виконання`Sync` - надішле буфер до файлу.
	f.Sync()

	// Пакет `bufio` надає можливість буферизованих
	// записів, надодачу до буферизованого читання
	// (яке вже було нами розглянуто).
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Записано %d байт\n", n4)

	// Щоб впевнитись у тому, що усі буфери в які ми записували
	// наші данні було відправлено до джерела запису,
	// виконайте метод `Flush`.
	w.Flush()

}
