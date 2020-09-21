package gorotine

import (
	"fmt"
	"time"
)

func Loop()  {
	for i:=0;i<11;i++{
		time.Sleep(time.Microsecond*1)
		fmt.Println(i)
	}
}
