package linear_structure

import (
	"fmt"
)

type ListNode struct {
	Val interface{}
	Next *ListNode
}

func NewLinkedList() *ListNode {
	return &ListNode{nil,nil}
}

func (list *ListNode) AddNode(val interface{}) {
	if list.Val == nil {
		list.Val = val
		list.Next = nil
	}else {
		l := list
		for l.Next != nil {
			l = l.Next
		}
		l.Next = &ListNode{Val: val, Next: nil}
		list = l
	}
}

func (list *ListNode) GetSize() (res int) {
	l := list
	res = 0
	for l != nil {
		res ++
		l = l.Next
	}
	return
}

func (list ListNode) FindByVal(val interface{}) (index int) {
	l := &list
	index = 0
	var Val interface{}
	for l != nil {
		index ++
		Val = l.Val
		if l.Val == val {
			break
		}
		l = l.Next
	}
	if val != Val && index == list.GetSize(){
		index = -1
	}
	return
}

func (list ListNode) GetValByIndex(index int) (res interface{}) {
	res =  nil
	if index <= 0 || index > list.GetSize() {
		return
	}
	l,i := &list,0
	for l != nil {
		i ++
		if i == index {
			res = l.Val
			break
		}
		l = l.Next
	}
	return
}

func (list *ListNode) RmHead() (res interface{}) {
	l := list
	res = l.Val
	list = l.Next
	return
}

func (list ListNode) Print() {
	l := &list
	for l != nil{
		fmt.Println(l.Val)
		l = l.Next
	}
}