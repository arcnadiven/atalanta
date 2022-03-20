package ratelimiter

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"testing"
	"time"
)

func TestNewConcurrencyLimiter(t *testing.T) {
	limiter := NewConcurrencyLimiter(5)
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 5; i++ {
			limiter.Preempt(strconv.Itoa(i))
		}
		ch <- struct{}{}
		time.Sleep(5 * time.Second)
		for i := 0; i < 5; i++ {
			limiter.Release(strconv.Itoa(i))
		}
	}()
	<-ch
	for i := 0; i < 5; i++ {
		fmt.Println(limiter.Preempt(strconv.Itoa(i)))
	}
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Kill, os.Interrupt)
	<-exit
}
