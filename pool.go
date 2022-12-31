package light_pool

import "sync"

type Pool struct {
	workNum int
	params  chan interface{}
	fn      func(interface{})
}

func New(workNum int, fn func(interface{})) *Pool {
	return &Pool{
		workNum: workNum,
		params:  make(chan interface{}),
		fn:      fn,
	}
}

func (p *Pool) Push(param interface{}) {
	p.params <- param
}

func (p *Pool) Close() {
	close(p.params)
}

func (p *Pool) Run() {
	var wg sync.WaitGroup
	wg.Add(p.workNum)
	for i := 0; i < p.workNum; i++ {
		go func() {
			for param := range p.params {
				p.fn(param)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
