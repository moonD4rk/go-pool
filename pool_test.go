package light_pool

import (
	"fmt"
	"testing"
	"time"
)

// TestNewPool，测试 Pool 功能
func TestNewPool(t *testing.T) {
	var (
		workNum = 10
		jobNum  = 100
	)
	p := New(workNum, request)

	go func() {
		for i := 0; i < jobNum; i++ {
			p.Push(i)
		}
		p.Close()
	}()

	p.Run()
}

func request(i interface{}) {
	time.Sleep(time.Second * 1)
	fmt.Println("finish request, url num is:", i.(int))
}
