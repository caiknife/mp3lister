package test

import (
	"testing"
)

func TestChan_UseForRange(t *testing.T) {
	intChan := make(chan int, 10)

	// 在协程里写入chan，主进程里读取
	go func(ch chan int) {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}(intChan)

	for v := range intChan {
		t.Log(v)
	}
}

func TestChan_UseSelect(t *testing.T) {
	intChan := make(chan int, 10)

	// 在协程里写入chan，主进程里读取
	go func(ch chan int) {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}(intChan)

LoopFor:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break LoopFor
			}
			t.Log(v)
		}
	}
}
