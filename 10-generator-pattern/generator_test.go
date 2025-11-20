package generator

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	ch := Count(1, 50)

	// 从只读通道获取，直到通道关闭后自动结束
	for i := range ch {
		fmt.Println(i)
	}
}
