package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		//{2, 6, 9, 24},
		//{7, 3, 94, 3, 0},
		//{4, 2, 8, 35},
		{4, 6},
		{7, 9},
	}
	// Ваша реалізація

	wg := sync.WaitGroup{}
	wg.Add(len(n))
	ch := make(chan int)
	res := 0

	for i := range n {
		go sum(n[i], ch)
	}

	for _ = range n {
		res += <-ch
		wg.Done()
	}

	wg.Wait()

	fmt.Println(res)
}

func sum(sl []int, ch chan int) {
	res := 0
	for i := range sl {
		res += sl[i]
	}

	ch <- res

}
