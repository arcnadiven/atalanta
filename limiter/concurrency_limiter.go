package ratelimiter

import (
	"sync"
)

type ConcurrencyLimiter interface {
	Resize(int)
	Preempt(string) bool
	Release(string)
}

type ConcurrencyLimiterImpl struct {
	// protect for concurrency
	*sync.Mutex
	// max concurrency num
	maxNum int
	// working concurrency num
	workingNum int
	// cache the working concurrency by keyword
	cache map[string]struct{}
}

func NewConcurrencyLimiter(capacity int) ConcurrencyLimiter {
	return &ConcurrencyLimiterImpl{
		Mutex:      new(sync.Mutex),
		maxNum:     capacity,
		workingNum: 0,
		cache:      map[string]struct{}{},
	}
}

func (cl *ConcurrencyLimiterImpl) Resize(capacity int) {
	cl.Lock()
	cl.maxNum = capacity
	cl.Unlock()
}

func (cl *ConcurrencyLimiterImpl) Preempt(key string) bool {
	cl.Lock()
	defer cl.Unlock()
	if cl.workingNum == cl.maxNum {
		//fmt.Println("too many concurrency")
		return false
	}
	if _, ok := cl.cache[key]; ok {
		//fmt.Println("concurrency is reentrant")
		return false
	}
	cl.cache[key] = struct{}{}
	cl.workingNum++
	return true
}

func (cl *ConcurrencyLimiterImpl) Release(key string) {
	cl.Lock()
	defer cl.Unlock()
	if _, ok := cl.cache[key]; ok {
		delete(cl.cache, key)
	}
	if cl.workingNum != 0 {
		cl.workingNum--
	}
}
