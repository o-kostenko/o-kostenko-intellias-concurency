package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”

func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(n))
	for i := range n {
		go func(i int, v []int) {
			defer wg.Done()
			sum(fmt.Sprintf("slice %d:", i), n[i])
		}(i, n[i])
	}
	wg.Wait()
}

func sum(name string, sl []int) {
	s := 0
	for i := range sl {
		s += sl[i]
	}

	fmt.Println(name, s)
}
