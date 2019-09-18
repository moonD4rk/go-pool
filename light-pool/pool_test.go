package light_pool

import (
	"fmt"
	"testing"
	"time"
)

// 模拟 10 个线程同时发送请求
// 总共为 100 个 url 任务
func TestNewPool(t *testing.T) {
	var (
		p        *Pool
		workNum  = 10
		urlCount = 100
	)
	p = NewPool(workNum)
	for i := 0; i < urlCount; i++ {
		p.Push(request(i))
	}
	p.Run()
}

// 模拟发送请求过程
func request(i int) func() {
	return func() {
		time.Sleep(time.Second * 1)
		fmt.Println("finish request, url num is:", i)
	}
}

