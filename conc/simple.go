package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//runSimpleCase()
	vData := simplePool(17, 4)

	fmt.Println()
	fmt.Printf("vData: %d\n", vData)
}

// 测试用例
func runSimpleCase() {
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
//
// !当线程过大时程序会崩溃（警告分配内存过多而异常）
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

// 协程池，控制使其的并发数无限大
func simplePool(cct int, poolSize uint) int {
	var countData int
	var lock sync.Mutex
	// 现场数添加
	var pt = &poolTask{size: poolSize, task: func(i int) {
		for {
			if lock.TryLock() {
				countData++
				if i%2 == 0 {
					countData += 2
				} else {
					countData -= 3
				}
				lock.Unlock()
				break
			}
		}
	}}
	pt.init()

	// 数据遍历
	for i := range cct {
		pt.add(i)
	}

	// 等待所有任务完成
	pt.run()
	pt.wg.Wait()
	return countData
}

type poolTask struct {
	task     func(int) // 任务
	size     uint
	wg       sync.WaitGroup
	taskList []int
}

// 初始化
func (pt *poolTask) init() {
}

// 添加任务
func (pt *poolTask) add(i int) {
	pt.taskList = append(pt.taskList, i)
}

// 运行
func (pt *poolTask) run() {
	pt.wg.Add(1)
	var indexLock sync.Mutex
	//var queue []int    // 队列
	var index int = -1 // 已执行到的任务索引
	tryAddIndexFn := func() int {
		for {
			if indexLock.TryLock() {
				index += 1
				indexLock.Unlock()
				return index
			}
		}
	}
	// [2]int  => [协程池索引, 任务索引]
	var runGIdx = make(chan [2]int, pt.size)
	var isAllDone = make(chan bool)
	var taskSize = len(pt.taskList)
	fmt.Println("任务数：", taskSize, " 协程池：", pt.size)

	// 任务执行
	runTaskFn := func(rIdx uint, gIdx int) {
		if gIdx >= taskSize {
			isAllDone <- true
			return
		}
		pt.task(gIdx)
		runGIdx <- [2]int{int(rIdx), gIdx}
	}

	// 初始化协程池，并推送
	for i := range pt.size {
		index += 1
		go runTaskFn(i, index)
	}

	// 循环处理
	for {
		isBreak := false
		select {
		case idxArr := <-runGIdx:
			// fmt.Printf("idxArr: %v ...\r", idxArr)
			fmt.Println(idxArr)
			go runTaskFn(uint(idxArr[0]), tryAddIndexFn())
		case <-isAllDone:
			isBreak = true
		}

		// 退出任务
		if isBreak {
			pt.wg.Done()
			break
		}
	}

}
