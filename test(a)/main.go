package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Unlock()
	}()
	mu.Lock()
}
//题目里面错位了lock与unlock。因为主协程的执行快于其他的goroutine，原题先来到解锁，而没有设置锁，就出错了。