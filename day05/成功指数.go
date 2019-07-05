package main

import (
	"game"
	"fmt"
)

func main() {
	name,age,money,readBook := game.GetArgs()

	//平均水平：以【40岁有50万标准】=1.0
	const k = 1 / (5.0e5 / 40)
	//成功程度，与money成正比，与age成反比,算出成功指数
	successIndex := k * money / float64(age)

	//富有且读书的，加分，贫穷且不读书的减分
	if money > 1.0e6 && readBook {
		successIndex *= 1.1
	}
	if money < 1.0e4 && !readBook {
		successIndex *= 0.9
	}
	//得到评语
	var remark string
	if successIndex > 1 {
		remark = "棒棒的"
	} else {
		remark = "呵呵"
	}

	//姓名用于输出报告
	fmt.Printf("尊敬的%s阁下：你的成功指数是%.2f,%s\n", name, successIndex, remark)
}
