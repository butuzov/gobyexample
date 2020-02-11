// Go пропонує підтримку XML та XML-подібних
// форматів із пакетом `encoding.xml`.

package main

import (
	"encoding/xml"
	"fmt"
)

// Цей тип буде відображено в XML. Аналогічно до
// прикладу з JSON, теги полів містять директиви для
// кодера та декодера. Тут ми використовуємо деякі особливості
// пакета XML: назва поля `XMLName` вказує на
// назви елемента в XML, що представляє цю структуру;
// `id,attr` означає, що поле `id` є _атрибутом_ тегу `plant`.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Тепер можна створити XML, із структури `Plant`. Використовуйте
	// `MarshalIndent` для отримання більш зрозумілого для людини результату.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Щоб додати стандартний XML-заголовок до виводу, приєднайте його явно.
	fmt.Println(xml.Header + string(out))

	// Використовуйте `Unmarhshal` для розбору потоку
	// байтів з XML в структуру даних. Помилка повертається, якщо XML неправильно
	// сформований або не може бути перетворений в Plant.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// `parent>child>plant` тег вказує кодеру шлях
	// `<parent><child>...` за яким будуть розміщенні `plant`'s
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
