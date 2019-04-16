package main

import (
	"fmt"
	"time"
)

func main() {
	arr := SleepSort([]int{1,3,1,42,32,12})
	fmt.Println(arr)
}

func SleepSort(arr []int) (res []int) {
	var Channel = make(chan int)
	res = make([]int,0)
	for _,e := range arr{
		go func(e int) {
			time.Sleep(time.Millisecond * time.Duration(e))
			Channel <- e
		}(e)
	}
	for{
		select {
		case c := <- Channel:
			res = append(res, c)
		case <- time.After(time.Second):
			return
		}
	}
}