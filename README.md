# Golang 线程池

此项目只是用来学习 Golang 并发的小项目，想要更完备的线程池可以使用 [conc](https://github.com/sourcegraph/conc)。

## 轻量级的线程池

使用 10 个线程模拟请求 url，单个 url 耗时 1s，10个线程 100个 url 总耗时 10s。

```go
package main

import (
	"fmt"
	"time"

	pool "github.com/moond4rk/go-pool"
)

func main() {
	var (
		workNum = 10
		jobNum  = 100
	)
	p := pool.New(workNum, request)

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
