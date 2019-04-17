package sort

/*
 * 算法导论P109,计数排序
 */
import (
	"math"
)

func CountingSort(arr []int) (res []int) {
	length := len(arr)
	max := math.MinInt64
	for _, a := range arr {
		if max < a {
			max = a
		}
	}
	a := make([]int, max+1)
	res = make([]int, length+1)
	for i := 0; i < length; i++ {
		a[arr[i]] += 1
	}
	for i := 1; i < max+1; i++ {
		a[i] = a[i] + a[i-1]
	}
	for i := length - 1; i >= 0; i-- {
		res[a[arr[i]]] = arr[i]
		a[arr[i]] -= 1
	}
	res = res[1:]
	return
}
