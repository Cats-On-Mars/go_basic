package mathutil

func GetNums(i int) (nums [3]int) {
	nums[0] = i/100
	nums[1] = i%100/10
	nums[2] = i%10

	return

}

