package 单元测试

import (
	"testing"
	"fmt"
)

func TsetAdd(t *testing.T){
	ret := add(3,5)
	if ret != 8{
		t.Errorf("want 8 , got %d\n",ret)
		return
	}
	fmt.Println("测试通过")
}