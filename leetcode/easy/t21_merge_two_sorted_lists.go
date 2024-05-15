package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "[]"
	}
	var b strings.Builder
	b.WriteString("[")
	cur := l
	for cur != nil {
		b.WriteString(strconv.Itoa(cur.Val))
		cur = cur.Next
		if cur != nil {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

// 21. 合并两个有序链表: https://leetcode.cn/problems/merge-two-sorted-lists/description/
//
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list ListNode
	cur := &list

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return list.Next
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	fmt.Println(mergeTwoLists(l1, l2)) // expect: [1, 1, 2, 3, 4, 4]

	l1 = nil
	l2 = nil
	fmt.Println(mergeTwoLists(l1, l2)) // expect: []

	l1 = nil
	l2 = &ListNode{0, nil}
	fmt.Println(mergeTwoLists(l1, l2)) // expect: [0]
}
