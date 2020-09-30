package main

import "fmt"

func main() {
	s := make(chan string, 2) //宣告一個 channel 變數

	s <- "hello" //寫入 channel (sender)
	val := <-s   //讀取 channel (receiver)
	fmt.Println(val)
}
