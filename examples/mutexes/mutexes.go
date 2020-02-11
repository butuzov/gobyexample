// У попередньому прикладі ми розглянули як керувати простим
// лічильником за допомогою [атомарних лічильників](atomic-counters).
// У більш складних випадках ми можемо скористатись [_м’ютексами_](https://uk.wikipedia.org/wiki/М%27ютекс)
// для безпечного доступу до даних в горутинах.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// В нашому прикладі `state` буде мапою.
	var state = make(map[int]int)

	// А це - `mutex`, що керуватиме синхронізацією доступу до `state`.
	var mutex = &sync.Mutex{}

	// Ми будемо слідкувати - скільки операцій запису та
	// читання ми робитимемо.
	var readOps uint64
	var writeOps uint64

	// Ось ми запускаємо 100 горутин - які повторно читатимуть дані
	// стану нашої мапи раз в мілісекунду у кожній горутині.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// Для кожного читання ми: генеруємо ключ для
				// доступа, замикаємо (`Lock()`) м’ютекс
				// (щоб впевнитись в унікальному доступі до `state`),
				// читаємо значення обраного ключа,
				// відмикаємо (`Unlock()`) м’ютекс, і збільшуємо
				// лічильник `readOps` на одиницю.
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// Чекаємо мілісекунду на наступну ітерацію.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Ми також стартуємо 10 горутин, для записів,
	// використовуючи такий же шаблон, що ми використали
	// для читання.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Дамо рівно 1 секунду нашим горутинам попрацювати з `state` та
	// `mutex`.
	time.Sleep(time.Second)

	// Нарешті прочитаємо звіти, щодо наших атоманих операцій
	// читання та запису.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// І покажемо як же `state` став виглядати після усіх наших операцій.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
