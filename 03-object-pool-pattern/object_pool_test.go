package pool

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestNewPool(t *testing.T) {
	p := NewPool(5)
	fmt.Println(len(*p))
	if len(*p) != 5 {
		t.Error("线程池构造错误")
	}
	for ob := range *p {
		ob.Do()
	}
}
