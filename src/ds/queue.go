/**
 * 2017年6月28日 星期三
 * Joshua Conero
 * 队列
 * 特殊的线性表, FIFO(先进先出)
 * 入队/出队
 */

package ds

// 队列对象
type Queue struct {
	Front int      // 队头指针 - 前端
	Data  []string // 数据内容
	Rear  int      // 队尾指针 - 后端
}

/**
 * 队列初始化
 */
func (q *Queue) Init() *Queue {
	q.Front = 0
	q.Rear = 0
	q.Data = []string{}
	return q
}

/**
 * 入队
 */
func (q *Queue) In(value string) bool {
	q.Front = q.Rear + 1
	q.Data = append(q.Data, value)
	return true
}

/**
 * 出队
 */
func (q *Queue) Out() (bool, string) {
	if len(q.Data) > 0 {
		dtLen := len(q.Data) - 1
		dtLastVal := q.Data[dtLen]
		q.Data = q.Data[0:dtLen]
		q.Front = q.Front + 1
		return true, dtLastVal
	}
	return false, ""
}

/**
 * 获取读队头元素
 */
func (q *Queue) GetFront() (bool, string) {
	if len(q.Data) > 0 {
		return true, q.Data[len(q.Data)-1]
	}
	return false, ""
}
