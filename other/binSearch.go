package other

func BinSearch(nums []int,num int) (index int){

	left,right := 0,len(nums)
	for i := left+1; i < right ; i++ {
		if nums[i] < nums[i-1] {
			return -1
		}
	}
	index = -1
	Loop:
	for {
		if left >= right {
			break Loop
		}
		mid :=  (left+right) / 2
		switch true {
		case nums[mid]<num:
			left=mid
		case nums[mid]>num:
			right=mid
		case nums[mid]==num:
			index=mid
			break Loop
		}
	}

	return
}