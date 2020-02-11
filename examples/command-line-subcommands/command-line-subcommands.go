// Деякі з утиліт командного рядку, як то `go` чи `git`,
// мають великий вибір додаткових команд, а також власний набір
// прапорців/параметрів. Наприклад `go build` і `go get`
// являються додатковими командами для `go`-утиліти.
// Для роботи з прапорцями/параметрами в
// стандартній бібліотеці `go` є пакет `flag`, який дозволяє
// з легкістю визначати команди а також їх параметри.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Тут, за допомогою функції `NewFlagSet`, визначається нова команда.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// Далі оголошуються два додаткові параметри `enable` і `name`
	// з типами ціле и стрічка відповідно.
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Для різних команд можна задати різні параметри,
	// які вони підтримують.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Очікується, що параметри будуть передані як перший аргумент при старті програми.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Перевірити, яка команда була передана.
	switch os.Args[1] {

	// Якщо була передана команда `foo`, виконати парсинг параметрів,
	// а також вивести всі параметри які не були задіяні при парсингу вхідної стрічки.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	// Якщо була передана команда `bar`, виконати парсинг параметрів,
	// а також вивести всі параметри які не були задіяні при парсингу вхідної стрічки.
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	// Якщо команда, що була передана - являється забороненою,
	// вивести помилку і закінчити виконання.
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
