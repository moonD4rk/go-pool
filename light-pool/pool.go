package light_pool

import (
	"fmt"
	"sync"
)

type Pool struct {
	workNum int
	wg      sync.WaitGroup
	fns     chan func()
}

func NewPool(workNum int) *Pool {
	p := &Pool{
		workNum: workNum,
		wg:      sync.WaitGroup{},
		fns:     make(chan func()),
	}
	p.wg.Add(workNum)
	for i := 0; i < workNum; i++ {
		// 反射捕获错误
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
					p.wg.Done()
				}
			}()
			// 遍历 channel, 取出函数对象并执行
			for fn := range p.fns {
				fn()
			}
			p.wg.Done()
		}()
	}
	return p
}

// 向 chan 推送任务
func (p *Pool) Push(fn func()) {
	p.fns <- fn
}

func (p *Pool) Run() {
	close(p.fns)
	p.wg.Wait()
}

