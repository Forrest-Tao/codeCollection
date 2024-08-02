package singleton

import "sync"

type Singleton struct{}

var ins *Singleton
var once sync.Once

var cnt int

func GetInstance() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
		cnt++
	})
	return ins
}
