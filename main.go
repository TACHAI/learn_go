package main

import (
	"awesomeProject/gorotine"
	"awesomeProject/interface_demo"
	"awesomeProject/point_demo"
	"awesomeProject/struct_demo"
	"time"
)

func main()  {
	d:=new(struct_demo.Dog)
    d.Id=1
    d.Age=2
    d.Name="旺财"
   // d.Run()

   // var b interface_demo.Behavior
    b := new(struct_demo.Dog)
	action(b);
    c :=new(struct_demo.Cat)
    action(c)

    go gorotine.Loop()
    go gorotine.Loop()
	time.Sleep(time.Second*1)


    point_demo.TestPoint()
}

func action(b interface_demo.Behavior)string  {
	b.Run()
	b.Eat()
	return ""
}
