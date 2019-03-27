package sort

/*
 * 选择排序
 */
func SelectionSort(arr []int,left,rigth int){

	for i:=0;i<rigth-left+1;i++ {
		min:=i
		for j:=i+1;j<rigth-left+1;j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i],arr[min]=arr[min],arr[i]
	}

}
