package sort

/*
 * 冒泡排序
 */
func BubbleSort(arr []int,left,right int) {
	if left>right{
		return
	}
	flag := true
	for i:=0;i<right-left;i++ {
		flag=true
		for j:=right-left-i;j>i;j-- {
			if arr[j]<arr[i]{
				arr[i],arr[j]=arr[j],arr[i]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
