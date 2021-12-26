package main

import "fmt"

const count = 10

func main() {
	ch := make(chan struct{}, count)
	for i := 0; i < count; i++ {
		go func() {
			fmt.Println("强人所难")
			ch <- struct{}{}
		}()
	}

	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Println("完蛋啦！")
}
