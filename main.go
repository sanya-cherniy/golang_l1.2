package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	for i := range arr {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)
		}(arr[i])
	}
	wg.Wait()
}
