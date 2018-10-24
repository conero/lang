package main

/**
	@date 2018年10月24日 星期三
	@link
 */

import (
	"fmt"
	"strconv"
	"strings"
)
// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}
// 获取值
func (ln *ListNode) Value() int {
	var num int
	b := 1
	for tnd := ln; tnd != nil;{
		num += b * tnd.Val
		b *= 10
		tnd = tnd.Next
	}
	return num
}
// 获取值
func (ln *ListNode) Value64() int64 {
	var num int64
	var b int64 = 1
	for tnd := ln; tnd != nil;{
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
	for tnd := ln; tnd != nil;{
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
func GetLNdFromQue(que []int) *ListNode  {
	var lnd *ListNode
	for i := len(que) - 1; i> -1; i--{
		if lnd == nil{
			lnd = &ListNode{que[i], nil}
		}else {
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
func (c Case) Test1()  {
	lnd1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	lnd2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	//fmt.Println(lnd1.Value())
	//fmt.Println(lnd2.Value())

	atn := addTwoNumbers(lnd1, lnd2)
	fmt.Println(lnd1.Value(), lnd2.Value2(), atn.Value())
}


func (c Case) Test2()  {
	lnd1 := &ListNode{1, &ListNode{0, &ListNode{0,
	&ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}
	lnd2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	//fmt.Println(lnd1.Value())
	//fmt.Println(lnd2.Value())

	atn := addTwoNumbers64(lnd1, lnd2)
	fmt.Println(lnd1.Value(), lnd2.Value2(), atn.Value())
}

func main()  {
	var cs Case
	cs.Test1()
	cs.Test2()

	fmt.Println(GetLNdFromQue([]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,9,1}).Value64())
}







/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 将链表转化为整形
	fn := func(ln *ListNode) int {
		var n1 int
		b := 1
		for i1 := ln; i1 != nil; {
			n1 += b * i1.Val
			b *= 10
			i1 = i1.Next
		}
		return n1
	}

	// 整形求和
	n := fn(l1) + fn(l2)

	// 生成性的链表
	var lnd *ListNode
	for _, s := range strings.Split(strconv.Itoa(n), ""){
		n, err := strconv.ParseInt(s, 10, 0)
		if err != nil{
			n = 0
		}
		if lnd == nil{
			lnd = &ListNode{int(n), nil}
		}else{
			tmpLnd := &ListNode{int(n), nil}
			tmpLnd.Next = lnd
			lnd = tmpLnd
		}
	}

	return lnd
}

func addTwoNumbers64(l1 *ListNode, l2 *ListNode) *ListNode {
	// 将链表转化为整形
	fn := func(ln *ListNode) int64 {
		var n1 int64
		var b int64 = 1
		for i1 := ln; i1 != nil; {
			n1 += b * int64(i1.Val)
			b *= 10
			i1 = i1.Next
		}
		return n1
	}

	// 整形求和
	n := fn(l1) + fn(l2)

	// 生成性的链表
	var lnd *ListNode
	for _, s := range strings.Split(strconv.FormatInt(n, 10), ""){
		n, err := strconv.ParseInt(s, 10, 0)
		if err != nil{
			n = 0
		}
		if lnd == nil{
			lnd = &ListNode{int(n), nil}
		}else{
			tmpLnd := &ListNode{int(n), nil}
			tmpLnd.Next = lnd
			lnd = tmpLnd
		}
	}

	return lnd
}