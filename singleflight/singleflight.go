package singleflight

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

var errGoexit = errors.New("runtime.Goexit was called")

type panicError struct {
	stack []byte
	value interface{}
}

func (p *panicError) Error() string {
	return fmt.Sprintf("panic: %v\n%s", p.value, p.stack)
}

func (p *panicError) Unwrap() error {
	err, ok := p.value.(error)
	if !ok {
		return nil
	}
	return err
}

func newPanicError(value interface{}) *panicError {
	stack := debug.Stack()
	if line := bytes.IndexByte(stack, '\n'); line >= 0 {
		stack = stack[:line+1]
	}
	return &panicError{
		value: value,
		stack: stack,
	}
}

type call struct {
	val   interface{}
	err   error
	wg    sync.WaitGroup
	dups  int
	chans []chan<- Result
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

type Result struct {
	val    interface{}
	err    error
	shared bool
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	g.mu.Lock()
	//懒加载
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	//在这之前有更早的调用,等待第一个的返回值即可
	if c, ok := g.m[key]; ok {
		c.dups++
		g.mu.Lock()
		c.wg.Wait()

		if e, ok := c.err.(*panicError); ok {
			panic(e)
		} else if c.err == errGoexit {
			runtime.Goexit()
		}
		return c.val, c.err, true
	}
	c := new(call)
	g.m[key] = c
	g.mu.Unlock()
	c.wg.Add(1)
	g.doCall(c, key, fn)
	return c.val, c.err, c.dups > 0
}

func (g *Group) doCall(c *call, key string, fn func() (interface{}, error)) {

	normalReturn := false
	recoved := false

	defer func() {
		if !normalReturn && !recoved {
			c.err = errGoexit
		}
		g.mu.Lock()
		defer g.mu.Unlock()
		c.wg.Done()
		if g.m[key] == c {
			delete(g.m, key)
		}
		//成功recover到了
		if e, ok := c.err.(*panicError); ok {
			if len(c.chans) > 0 {
				go panic(e)
				select {}
			} else {
				panic(e)
			}
		} else if c.err == errGoexit { //没能正常返回且没能recover到 什么也不处理

		} else { //正常返回
			for _, ch := range c.chans {
				ch <- Result{
					val:    c.val,
					err:    c.err,
					shared: c.dups > 0,
				}
			}
		}
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				c.err = newPanicError(r)
			}
		}()
		c.val, c.err = fn()
		normalReturn = true
	}()

	if !normalReturn {
		recoved = true
	}
}
