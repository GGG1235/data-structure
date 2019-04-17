package sort

/*
 * 算法导论P112,桶排序
 */
import (
	"math"
	"sort"
)

func BucketSort(arr []int) (res []int) {
	length := len(arr)
	buckets := make([][]int,length)
	res = make([]int,length)
	max := math.MinInt64
	for _,a := range arr {
		if max < a {
			max=a
		}
	}
	index := 0
	for i := 0;i<length;i++{
		index = arr[i]*(length-1)/max
		buckets[index]=append(buckets[index],arr[i])
	}
	var t=0
	for i:=0;i<length;i++{
		bucketsNum := len(buckets[i])
		if bucketsNum > 0 {
			sort.Ints(buckets[i])
			copy(res[t:],buckets[i])
			t+=bucketsNum
		}
	}
	return
}