package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("simple 并发测试")

	var now = time.Now()
	var cct = 1_987_654
	var countData = simple(cct)

	fmt.Println()
	fmt.Println()
	fmt.Printf("并发数：%d, countData: %d， 用时： %v\n", cct, countData, time.Since(now))

}

// Mutex  互斥锁，保证同一时刻只有一个 goroutine 访问资源。
// RWMutex  读写锁，支持多读单写，提升读多写少场景的效率。
//
// simple 并发测试
func simple(cct int) int {
	var countData int

	var lock sync.Mutex
	var wg sync.WaitGroup

	for i := range cct {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("goroutine: %d ……\r", idx)
			for {
				if lock.TryLock() {
					countData++
					if idx%2 == 0 {
						countData += 2
					} else {
						countData -= 3
					}
					lock.Unlock()
					break
				}
			}
		}(i)
	}

	wg.Wait()
	return countData
}
