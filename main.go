package main

import (
	"data-structure/sort"
	"fmt"
)

func main()  {
	var a=[]int{32,432,12,31,54,0,3,12,54,12,2}
	sort.QuickSort(a,0,len(a)-1)
	for i,e := range a {
		fmt.Println(i,e)
	}
	//sort.BubbleSort(a,0,len(a)-1)
	//for i,e := range a {
	//	fmt.Println(i,e)
	//}
	//sort.SelectionSort(a,0,len(a)-1)
	//for i,e := range a {
	//	fmt.Println(i,e)
	//}
}

//func main(){
//
//	list := linear_structure.NewLinkedList()
//	list.AddNode("a")
//	list.AddNode("b")
//	list.AddNode("c")
//	list.Print()
//	fmt.Println(list.GetSize())
//	fmt.Println(list.FindByVal("d"))
//	fmt.Println(list.GetValByIndex(0))
//	//fmt.Println(list.RmByIndex(1))
//	list.RmHead()
//	list.Print()
//}