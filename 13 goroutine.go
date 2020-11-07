// +build ignore

package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func checkIsPrimeNumber (num int) bool {
	for cur := 2; cur < num; cur++ {
		if num % cur == 0 {
			return false
		}
	}
	return true
}

var flag = 0
func checkNumbers(begin, end int) {
	for idx := begin; idx <= end; idx++ {
		if checkIsPrimeNumber(idx) {
			flag++
		}
	}
	wg.Done()
}

func main() {
	primeNumberSlice := []int{}
	// howMany := 0
	begin := time.Now()
	for idx := 2; idx <= 3 * 10000; idx++ {
		if checkIsPrimeNumber(idx) {
			// howMany++
			primeNumberSlice = append(primeNumberSlice, idx)
		}
	}
	fmt.Println("耗时", time.Since(begin))
	fmt.Println("质数数量：", len(primeNumberSlice))
	// fmt.Println("质数数量：", howMany)
	wg.Add(3)
	begin2 := time.Now()
	go checkNumbers(2, 10000)
	go checkNumbers(10001, 10000 * 2)
	go checkNumbers(20001, 10000 * 3)
	wg.Wait()
	fmt.Println("耗时", time.Since(begin2))
	fmt.Println("质数数量：", flag)
}
