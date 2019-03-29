package graph

import "fmt"

/*
 * 邻接表深度优先遍历
 */

const MAX  = 0x3f3f3f3f
var n,e int
var edge = make([][]int,MAX)
var bOK = make([]bool,MAX)

func CreateGraph() {
	_, _ = fmt.Scanf("%d%d", &n, &e)
	var x,y int
	for i := 1;i <= e;i++ {
		_, _ = fmt.Scanf("%d%d", &x,&y)
		edge[x] = append(edge[x],y)
		edge[y] = append(edge[y],x)
	}
}

func DFS(node int) {
	for _,i := range edge[node] {
		if !bOK[i]{
			bOK[i]=true
			DFS(i)
		}
	}
}

func Check() bool {

	for i:=1;i<=n;i++ {
		flag:=bOK[i]
		if !flag {
			return false
		}
	}
	return true

}

func PrintGraph(){
	for i := 1;i <= n ; i++ {
		fmt.Println(edge[i])
	}
}