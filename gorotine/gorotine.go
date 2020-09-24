package gorotine

import (
	"fmt"
	"time"
)


var chanInt chan int = make(chan  int,10)

func Loop()  {
	for i:=0;i<11;i++{
		time.Sleep(time.Microsecond*1)
		fmt.Println(i)
	}
}

// 发送数据
func Send()  {
	chanInt <- 1
	chanInt <- 2
	chanInt <- 3
}

// 接收数据
func Receive()  {
	num := <- chanInt
	fmt.Println("num is ",num)
	//time.Sleep()
	num = <- chanInt
	fmt.Println("num is ",num)

	num = <- chanInt
	fmt.Println("num is ",num)
}