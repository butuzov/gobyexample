// У Go є кілька корисних функцій для роботи з
// *директоріями* файлової системи.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Створюємо директорію за назвою *subdir*. Другим параметром
	// є налаштування параметрів [доступу](https://en.wikipedia.org/wiki/File_system_permissions#Numeric_notation)
	// до директорії.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// При створенні тимчасової директорії, гарною ідеєю є
	// викликати її видалення за допомогою [`відкладення`](./defer).
	// Функція `os.RemoveAll` видаляє директорію повністю (разом
	// з вкладеними), за аналогією до`rm -rf subdir` в
	// UNIX-подібних системах.
	defer os.RemoveAll("subdir")

	// Допоміжна функція що створює тимчасовий файл.
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Також можливо створити повну ієрархію директорій за допомогою `MkdirAll`.
	// Ця функція є аналогом до команди `mkdir -p subdir/parent/child`.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` зчитає зміст директорії *parent*,
	// і поверне [зріз](./slices) об'єктів `os.FileInfo`.
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Перегляд subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` змінює поточну робочу директорію (так само, як і
	// команда `cd`).
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Ось, 'ioutil.ReadDir' повертає [зріз](./slices) об'єктів `os.FileInfo`
	// для поточної директорії (якою на даний момент є `subdir/parent/child`).
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Перегляд subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Повертаємось до початкової директорії.
	err = os.Chdir("../../..")
	check(err)

	// Ми можемо обійти директорію *рекурсивно*, відвідуючи всі
	// вкладені директорії. `Walk` другим параметром приймає
	// функцію зворотного виклику, яка викликається для кожного
	// знайденого файлу та директорії.
	fmt.Println("Відвідуємо subdir")
	err = filepath.Walk("subdir", visit)
}

// Функція `visit` буде викликана для кожного файлу чи директорії
// знайденого в процесі обходу *subdir*.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
