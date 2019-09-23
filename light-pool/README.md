## 轻量级的线程池

使用 10 个线程模拟请求 url，单个 url 耗时 1s，10个线程 100个 url 总耗时 10s。

```go
package main

import (
	"fmt"
	"time"

	pool "github.com/darkMoon1973/go-pool/light-pool"
)

func main() {
	startTime := time.Now()
	var (
		p        *pool.Pool
		workNum  = 10
		urlCount = 100
	)
  // 初始化 pool
	p = pool.New(workNum)
	for i := 0; i < urlCount; i++ {
    // 推送任务
		p.Push(request(i))
	}
  // 运行
	p.Run()
  // 程序运行时长
	costTime := time.Since(startTime)
	fmt.Printf("程序耗时共 %f 秒", costTime.Seconds())
}

// 模拟发送请求过程
func request(i int) func() {
	return func() {
		time.Sleep(time.Second * 1)
		fmt.Println("finish request, url num is:", i)
	}
}

```
输出

```
finish request, url num is: 7
finish request, url num is: 1
finish request, url num is: 4
finish request, url num is: 8
......
finish request, url num is: 96
finish request, url num is: 97
finish request, url num is: 93
finish request, url num is: 91
程序耗时共 10.015464 秒
Process finished with exit code 0

```
