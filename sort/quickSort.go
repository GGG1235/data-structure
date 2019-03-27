package sort

/*
 * 快速排序
 */
func QuickSort(arr []int, left,right int) {
	if left>right {
		return
	}
	explodeIndex:=left

	for i := left + 1;i <= right ; i++ {
		if arr[left] >= arr[i] {
			explodeIndex ++
			arr[i], arr[explodeIndex] = arr[explodeIndex], arr[i]
		}
	}

	arr[left] , arr[explodeIndex] = arr[explodeIndex] , arr[left]

	QuickSort(arr,left,explodeIndex - 1)
	QuickSort(arr,explodeIndex + 1,right)

}