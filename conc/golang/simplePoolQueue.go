package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程队列

// Task 表示一个需要执行的任务，包含唯一ID和处理函数
type Task struct {
	ID     int
	Handle func() // 任务处理逻辑
}

// Queue 协程队列，按索引管理任务
type Queue struct {
	capacity   int            // 队列容量（同时运行的最大协程数）
	taskCh     chan Task      // 任务通道（用于传递待执行任务）
	doneCh     chan int       // 完成通知通道（传递完成的队列索引）
	active     []bool         // 记录每个队列位置是否活跃（true表示正在执行任务）
	wg         sync.WaitGroup // 等待所有任务完成
	totalTasks int            // 总任务数（用于判断是否全部完成）
	processed  int            // 已处理的任务数
	mu         sync.Mutex     // 保护共享变量（active/processed）
}

// NewQueue 创建一个新的协程队列
// capacity: 队列容量（初始化的协程数）
// totalTasks: 总任务数（需要处理的任务总量）
func NewQueue(capacity, totalTasks int) *Queue {
	return &Queue{
		capacity:   capacity,
		taskCh:     make(chan Task, capacity), // 缓冲大小与容量一致
		doneCh:     make(chan int, capacity),  // 用于接收完成的索引
		active:     make([]bool, capacity),    // 初始化所有位置为非活跃
		totalTasks: totalTasks,
	}
}

// Start 启动队列，开始处理任务
func (q *Queue) Start() {
	// 1. 初始化队列：为每个位置启动一个协程
	for idx := 0; idx < q.capacity; idx++ {
		q.wg.Add(1)
		go q.worker(idx) // 每个worker绑定一个队列索引
	}

	// 2. 启动调度器：监控任务完成，动态填充新任务
	go q.scheduler()
}

// worker 工作协程，绑定固定队列索引，循环处理任务
func (q *Queue) worker(idx int) {
	defer q.wg.Done()
	for task := range q.taskCh {
		// 标记当前位置为活跃
		q.mu.Lock()
		q.active[idx] = true
		q.mu.Unlock()

		// 执行任务
		fmt.Printf("队列索引 %d 开始处理任务 %d\n", idx, task.ID)
		task.Handle() // 执行任务逻辑
		fmt.Printf("队列索引 %d 完成任务 %d\n", idx, task.ID)

		// 标记当前位置为非活跃，通知调度器
		q.mu.Lock()
		q.active[idx] = false
		q.processed++
		q.mu.Unlock()
		q.doneCh <- idx // 发送完成信号（携带队列索引）
	}
}

// scheduler 调度器：监控任务完成，动态生成并分发新任务
func (q *Queue) scheduler() {
	taskID := 0
	// 初始填充任务（填满队列）
	for i := 0; i < q.capacity && taskID < q.totalTasks; i++ {
		q.taskCh <- Task{
			ID: taskID,
			Handle: func() {
				time.Sleep(500 * time.Millisecond) // 模拟任务耗时
			},
		}
		taskID++
	}

	// 循环等待任务完成，动态补充新任务
	for {
		// 检查是否所有任务都已处理完毕
		q.mu.Lock()
		if q.processed >= q.totalTasks {
			q.mu.Unlock()
			close(q.taskCh) // 关闭任务通道，让worker退出
			return
		}
		q.mu.Unlock()

		// 等待某个队列位置完成任务
		idx := <-q.doneCh
		fmt.Printf("调度器：队列索引 %d 已空闲，准备分配新任务\n", idx)

		// 生成并分配新任务（如果还有剩余）
		if taskID < q.totalTasks {
			q.taskCh <- Task{
				ID: taskID,
				Handle: func() {
					time.Sleep(500 * time.Millisecond) // 模拟任务耗时
				},
			}
			taskID++
		}
	}
}

// Wait 等待所有任务处理完毕
func (q *Queue) Wait() {
	q.wg.Wait()
}

func main() {
	// 创建队列：容量为3，总任务数为10
	queue := NewQueue(3, 10)
	queue.Start()

	// 等待所有任务完成
	queue.Wait()
	fmt.Println("所有任务处理完毕")
}
