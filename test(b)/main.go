package main

import "fmt"

func main() {
	ch := make(chan string)
	var send = "发送成功"
	var receive string
	go func() {
		ch  <- send
		fmt.Println("下山的路又堵起了")
	}()
	receive = <- ch
	fmt.Println(receive)
}
