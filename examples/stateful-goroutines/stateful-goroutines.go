// Ми скористались [mutex](mutexes)'ами для синхронізації
// доступу до спільного стану в минулому прикладі. Іншим варіантом
// буде використання [горутин](goroutines) та [каналів](channels)
// задля отримання того ж ефекту. Цей спосіб також є ідіоматичним
// у Go - так ми можемо працювати з спільною пам'яттю
// віддаючи її у володіння лише 1-й горутині.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// У цьому прикладі наш стан буде переданий у володіння
// лише одній горутині. Це гарантуватиме, що дані ніколи не будуть
// зіпсовані одночасним доступом. Щоб читати та писати стан
// - інші оперуючі горутини будуть надсилати повідомлення
// головній горутині і отримувати належні відповіді.
// Ці структури - `readOp` та `writeOp`, приховують запити
// і те як головна горутина відповідає.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Як і раніше, ми будемо рахувати операції які ми
	// виконуємо.
	var readOps uint64
	var writeOps uint64

	// Наші канали - `reads` та `writes`, будуть використані
	// іншими горутинами для створення запитів на запис або на
	// читання відповідно.
	writes := make(chan *writeOp)
	reads := make(chan *readOp)

	// Ця горутина - є власником стану (який був мапою в
	// попередньому прикладі), який, наразі є приватним для
	// горутини зі станом. Ця горутина у циклі отримує запити
	// на читання та запис з відповідних каналів (зверніть увагу
	// на `select`) і відповідає на запити як тільки ті надходять.
	// Відповідь відбувається шляхом виконання операції (яку було
	// запитано) та відправкою відповіді до каналу `resp` (індикація
	// успіху або бажане значення у випадку `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Тут ми запускаємо 100 горутин, що будуть надсилати запити
	// зчитування до головної горутини через канал `reads`.
	// Кожна горутина створює `readOp` та надсилає її (змінну) до
	// каналу `reads` і отримує відповідь з каналу `resp`.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Використовуючи такий же метод - ми запускаємо
	// 10 горутин, які писатимуть дані.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Даємо нашим горутинам попрацювати близько секунди.
	time.Sleep(time.Second)

	// Нарешті, дивимось скільки було зчитувань та записів:
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
