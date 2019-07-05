package 单元测试

import (
	"testing"
	"fmt"
)

func TestTrunc(t *testing.T) {

	var red = color{"red", [4]int{255, 0, 0, 0}}
	truncfile(red)
	fmt.Println(red)

	var isRed color
	loadfile(isRed)

	if red.Name == isRed.Name && red.RGBAnum == isRed.RGBAnum {
		t.Log("成功")
	}else{
		t.Error("出错")
	}
}
