package main

import "fmt"

func main() {
	fmt.Println(KmpSearch("hello world","world"))
}

func getNext(p string) []int {
	pLen:=len(p)
	next:=make([]int,pLen,pLen)
	next[0],next[1]=-1,0
	i,j:=0,1
	for j < pLen-1 {
		if i==-1 || p[i]==p[j]{
			i++
			j++
			next[j]=i
		} else {
			i = next[i]
		}
	}
	return next
}

func getNextOptimize(p string) []int {
	pLen := len(p)
	next := make([]int, pLen, pLen)
	next[0] = -1
	next[1] = 0
	i := 0
	j := 1
	for j < pLen-1 {
		if i == -1 || p[i] == p[j] {
			i++
			j++
			if p[i] != p[j] {
				next[j] = i
			} else {
				next[j] = next[i]
			}

		} else {
			i = next[i]
		}
	}
	return next
}

func KmpSearch(s, p string) int {
	i, j := 0, 0
	pLen := len(p)
	sLen := len(s)
	next := getNext(p)
	for i < sLen && j < pLen {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}

}