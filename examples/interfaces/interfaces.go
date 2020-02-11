// _Інтерфейси_ - це іменовані коллекції описів методів.

package main

import (
	"fmt"
	"math"
)

// Базовий приклад інтерфейсу геометричних фігур.
type geometry interface {
	area() float64
	perim() float64
}

// Задля прикладу - ми реалізуємо цей інтерфейс для типів
// `rect` (квадрат) та `circle` (круг).
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Щоб реалізувати інтерфейс в Go, нам лише потрібно
// створити усі методи описані цим інтерфейсом. Для прикладу
// ми реалізуємо `geometry` для типу `rect` (квадрата).
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// А також для типу  `circle` (круга).
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// У разі, якщо змінна має тип інтерфейс, нам дозволяється
// викликати оголошені в цьому інтерфейсі методи.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Обидві типи структур - `circle` (круг) і `rect` (квадрат)
	// реалізують інтерфейс `geometry`, отож ми можемо використати
	// їх інстанцій як аргумент функції `measure`.
	measure(r)
	measure(c)
}
