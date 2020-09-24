package gorotine

import (
	"fmt"
	"sync"
	"time"
)


var chanInt chan int = make(chan  int,10)

var timeout chan bool = make(chan bool)

var WG sync.WaitGroup

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

	time.Sleep(time.Second*2)
	timeout <- true
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

func Receive2()  {
	for {
		select {
		case num:= <-chanInt:
			fmt.Println("num is ",num)
		case <- timeout:
			fmt.Println("timeout...")
		}
	}
}
//sync.waitgroup
// Add(delta int)添加协程记录
// Done()移除协程记录
// Wait()同步等待所有记录的协程全部结束

//读取数据
func Read()  {
	for i:=0;i<3;i++{
		WG.Add(1)
	}
}


// 写入
func Write()  {
	for i:=0;i<3;i++{
		time.Sleep(time.Second*2)
		fmt.Println("Done->",i)
		WG.Done()
	}
}