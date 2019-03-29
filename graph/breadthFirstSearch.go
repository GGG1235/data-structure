package graph

/*
 * 邻接表广度度优先遍历
 */

func BFS(node int) {
	q := make([]int,0)

	q=append(q,node)

	for len(q) != 0 {
		font := q[0]
		q=q[1:]
		for _,i := range edge[font] {
			if !bOK[i]{
				bOK[i]=true
				q = append(q,i)
			}
		}
	}
}