package main

/**
@date 2018年10月24日 星期三
@link
*/

import (
	"fmt"
	"strconv"
)

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 获取值
func (ln *ListNode) Value() int64 {
	var str string
	for tnd := ln; tnd != nil; {
		str += strconv.Itoa(tnd.Val)
		tnd = tnd.Next
	}
	n, er := strconv.ParseInt(str, 10, 64)
	if er != nil {
		n = 0
	}
	return n
}

// 获取值
func (ln *ListNode) Value64() int64 {
	var num int64
	var b int64 = 1
	for tnd := ln; tnd != nil; {
		num += b * int64(tnd.Val)
		b *= 10
		tnd = tnd.Next
	}
	return num
}

// 获取值
func (ln *ListNode) Value2() int {
	var num int
	b := 1
	for tnd := ln; tnd != nil; {
		num += b * tnd.Val
		b *= 10
		tnd = tnd.Next
	}
	return num
}

/**
[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]		=> ListNode
[5,6,4]   int(465)
*/
func GetLNdFromQue(que []int) *ListNode {
	var lnd *ListNode
	for i := len(que) - 1; i > -1; i-- {
		if lnd == nil {
			lnd = &ListNode{que[i], nil}
		} else {
			tmpLnd := &ListNode{que[i], lnd}
			lnd = tmpLnd
		}
	}
	return lnd
}

// 测试用例
type Case struct {
}

// 测试1
// 示例：
//		输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//		输出：7 -> 0 -> 8
//		原因：342 + 465 = 807
func (c Case) Test1() {
	lnd1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	lnd2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	//fmt.Println(lnd1.Value())
	//fmt.Println(lnd2.Value())

	atn := addTwoNumbers(lnd1, lnd2)
	fmt.Println(lnd1.Value(), lnd2.Value2(), atn.Value())
}

func (c Case) Test2() {
	lnd1 := &ListNode{1, &ListNode{0, &ListNode{0,
		&ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}
	lnd2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	//fmt.Println(lnd1.Value())
	//fmt.Println(lnd2.Value())

	atn := addTwoNumbers(lnd1, lnd2)
	fmt.Println(lnd1.Value(), lnd2.Value2(), atn.Value())
}

func main() {
	var cs Case
	cs.Test1()
	cs.Test2()

	//fmt.Println(GetLNdFromQue([]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,9,1}).Value64())

	// 边界测试
	//fmt.Println(addTwoNumbers(GetLNdFromQue([]int{0}), GetLNdFromQue([]int{7, 3})).Value())
	//fmt.Println(addTwoNumbers(GetLNdFromQue([]int{0}), GetLNdFromQue([]int{1})).Value())
	//fmt.Println(addTwoNumbers(GetLNdFromQue([]int{5}), GetLNdFromQue([]int{5})).Value())
	fmt.Println(addTwoNumbers(GetLNdFromQue([]int{1, 8}), GetLNdFromQue([]int{0})).Value())
	fmt.Println(addTwoNumbers(&ListNode{1, &ListNode{8, nil}}, &ListNode{0, nil}).Value())

	t := addTwoNumbers(&ListNode{1, &ListNode{8, nil}}, nil)
	fmt.Println(t)
	//fmt.Println(t, t.Next)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 链表的值取到数组中
	q1, q2 := []int{}, []int{}
	for t1 := l1; t1 != nil; {
		q1 = append(q1, t1.Val)
		t1 = t1.Next
	}

	for t2 := l2; t2 != nil; {
		q2 = append(q2, t2.Val)
		t2 = t2.Next
	}

	maxLen := len(q1)
	minLen := len(q2)
	if minLen > maxLen {
		// 值交换
		maxLen, minLen = minLen, maxLen
		q1, q2 = q2, q1
	}
	//fmt.Println(q1, q2, maxLen, minLen)
	// 遍历请且生成链表
	mBit := 0 // 进位
	que := []int{}
	for i := 0; i < maxLen; i++ {
		v := q1[i]
		if i < minLen {
			v += q2[i] + mBit
			if v > 9 {
				v = v - 10
				mBit = 1
			} else {
				mBit = 0
			}
			que = append(que, v)
			continue
		}

		if mBit > 0 {
			v = v + mBit
		}

		if v > 9 {
			v = v - 10
			mBit = 1
		} else {
			mBit = 0
		}

		que = append(que, v)
	}

	if mBit > 0 {
		que = append(que, mBit)
	}
	//fmt.Println(que)
	var value *ListNode
	//vlen := len(que)
	//for j := 0; j < vlen; j++{
	for j := len(que) - 1; j > -1; j-- {
		v := que[j]
		//fmt.Println(v, j)
		if value == nil {
			value = &ListNode{v, nil}
		} else {
			value = &ListNode{v, value}
		}
	}
	return value
}
