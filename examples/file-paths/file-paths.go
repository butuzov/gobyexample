// Пакет `filepath` забезпечує функціонал для роботи з *файловими шляхами*
// таким чином, що його робота є універсальною для операційних систем
// сімейства UNIX або Windows.

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` використовується для побудови шляхів.
	// Функція прийме будь-яку кількість строкових аргументів
	// і побудує ієрархічний шлях.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Використовуйте `Join` замість об'єднання `/` або `\` вручну.
	// На додачу, функція `Join` нормалізує шляхи, видаляючи зайві
	// розділювачі та каталоги.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Скористаємось `Dir` та `Base` для визначення каталогу та
	// файлу з наданого шляху. Альтернативою є `Split`- який буде
	// повертати обидва значення за одни виклик.
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	dir, file := filepath.Split(p)
	fmt.Printf("Dir(p): %s та Base(p): %s\n", dir, file)

	// Перевірка - чи є шлях абсолютним?
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// Більшість файлів мають в назві розширення (після крапки).
	// Ми можемо дізнатись розширення файлу за допомогою функції `Ext`.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Щоб дізнатись ім'я файлу без розширення,
	// скористайтесь `strings.TrimSuffix`.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` знаходить шлях між *base* і *target*.
	// Він повертає помилку, якщо це не можна зробити.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
