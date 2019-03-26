package main

import (
	"fmt"
	"strconv"
)

// 声明
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var newNode *ListNode = nil
	var tmp *ListNode = head
	var valNs []int = []int{}
	for {
		if head == nil {
			break
		}
		isBreak := false
		isContinue := false
		valN := tmp.Val
		for _, oN := range valNs {
			if valN == oN {
				isContinue = true
				break
			}
		}

		if tmp.Next != nil {
			tmp = tmp.Next
		} else {
			isBreak = true
		}
		if isContinue == false {
			valNs = append(valNs, valN)
			/*
			if newNode != nil {
				//newNode = &ListNode{valN, newNode}
				valNOld := newNode.Val
				newNode.Val = valN
				newNode = &ListNode{valNOld, newNode}
			} else {
				newNode = &ListNode{valN, nil}
			}
			*/
		}
		if isBreak {
			break
		}
	}
	//fmt.Println(valNs)
	//for _, nv := range valNs{
	for i := len(valNs) - 1; i >= 0; i--{
		nv := valNs[i]
		if newNode != nil{
			newNode = &ListNode{nv, newNode}
		}else {
			newNode = &ListNode{nv, nil}
		}
	}
	return newNode
}

// 控制台
func rdfsl_out(head *ListNode) string {
	var s string
	var tmp *ListNode = head
	for {
		if nil == tmp{
			break
		}
		valN := tmp.Val
		n2s := strconv.Itoa(valN)
		s += n2s + " -> "

		if tmp.Next != nil {
			tmp = tmp.Next
		} else {

			break
		}
	}
	return s
}

func main() {
	var vin *ListNode
	vin = &ListNode{
		1, &ListNode{
			1, &ListNode{
				2, nil,
			},
		},
	}
	fmt.Println(rdfsl_out(vin), rdfsl_out(deleteDuplicates(vin)))

	vin = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	fmt.Println(rdfsl_out(vin), rdfsl_out(deleteDuplicates(vin)))
}
