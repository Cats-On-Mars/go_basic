package main

import "fmt"

func main() {
	var nums = [10]int{6,2,7,5,8,4,9,0,3,1}

	//第一种:选择排序
	for i:=0;i<len(nums)-1;i++  {
		for j:=i+1;j<len(nums);j++  {
			if nums[i]<nums[j]{
				nums[i],nums[j]=nums[j],nums[i]    //不需要借助中间变量，go语言特色
			}
		}
	}
	fmt.Println(nums)

	nums = [10]int{6,2,7,5,8,4,9,0,3,1}
	//第二种：冒泡排序
	for i:=len(nums)-1;i>0;i-- {
		for j:=0;j<len(nums)-1;j++{
			if nums[j]<nums[j+1]{
				nums[j],nums[j+1]=nums[j+1],nums[j]
			}
		}
	}
	fmt.Println(nums)

}
