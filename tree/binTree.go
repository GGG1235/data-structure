package main

import "fmt"

type treeNode struct {
	Data interface{}
	Left *treeNode
	Right *treeNode
}

type BinTree interface {
	InorderTraversal()
	PreorderTraversal()
	PostorderTraversal()
	LevelorderTraversal()
	GetHeight() int
}

func NewBinTree() *treeNode {
	return &treeNode{Data:nil,Left:nil,Right:nil}
}

func (tree *treeNode) InorderTraversal() {
	if tree != nil {
		tree.Left.InorderTraversal()
		fmt.Println(tree.Data)
		tree.Right.InorderTraversal()
	}
}

func (tree *treeNode) AddNode(data interface{}) {
	t := tree
	i:=0
	for t!=nil {
		i++
		fmt.Println(i,t.Data)
		left := t.Left
		right := t.Right
		if left.GetHeight() == 1 && right == nil {
			t = right
		}
		if left.GetHeight() != 1 && left.Right == nil {
			t = left.Right
		}
		if left.GetHeight() != 1 && left.Right != nil && right.GetHeight() == 1 {
			t = right.Left
		}
		if left.GetHeight() != 1 && left.Right != nil && right.GetHeight() != 1 && right.Left != nil {
			t = right.Right
		}
	}
	t = &treeNode{}
	t.Data=data
	t.Left=nil
	t.Right=nil
	tree=t
}

func (tree *treeNode) PreorderTraversal() {
	if tree != nil {
		fmt.Println(tree.Data)
		tree.Left.PreorderTraversal()
		tree.Right.PreorderTraversal()
	}
}

func (tree *treeNode) PostorderTraversal() {
	if tree != nil {
		tree.Left.PostorderTraversal()
		tree.Right.PostorderTraversal()
		fmt.Println(tree.Data)
	}
}

func (tree *treeNode) LevelorderTraversal() {
	q := make([]*treeNode,0)
	q = append(q,tree)
	for len(q)!=0 {
		element := q[0]
		q = q[1:]
		fmt.Println(element.Data)
		if element.Left != nil {
			q = append(q,element.Left)
		}
		if element.Right != nil {
			q = append(q,element.Right)
		}
	}
}

func (tree *treeNode) GetHeight() (res int) {
	var HL,HR int

	if tree != nil {
		HL = tree.Left.GetHeight()
		HR = tree.Right.GetHeight()
		res =  HR + 1
		if HL > HR {
			res = HL + 1
		}
		return 
	}
	return 0
}

func main()  {
	tree := NewBinTree()
	tree.Data="a"
	tree.Left=&treeNode{}
	tree.Left.Data="b"
	tree.Right=&treeNode{}
	tree.Right.Data="c"
	tree.Left.Left=&treeNode{}
	tree.Left.Left.Data="d"
	tree.Left.Right=&treeNode{}
	tree.Left.Right.Data="e"
	tree.Right.Left=&treeNode{}
	tree.Right.Left.Data="f"
	tree.AddNode("g")
	//tree.InorderTraversal()
	//tree.PreorderTraversal( )
	//tree.PostorderTraversal()
	//fmt.Println(tree.GetHeight())
	tree.LevelorderTraversal()
}