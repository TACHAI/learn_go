package main

import (
	"awesomeProject/gorotine"
	"awesomeProject/interface_demo"
	"awesomeProject/json_demo"
	"awesomeProject/point_demo"
	"awesomeProject/struct_demo"
	"fmt"
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

    json_demo.Serialize()
    json_demo.SerializeMap()
    json_demo.DeSerializeStruct()
    json_demo.DeSerializeMap()

	// 协程通信
    go gorotine.Send()
    //go gorotine.Receive()
    go gorotine.Receive2()
    time.Sleep(time.Second*10)

    // 协程同步
    gorotine.Read()
    go gorotine.Write()
    gorotine.WG.Wait()
    fmt.Println("All done!")
    time.Sleep(time.Second*60)

}

func action(b interface_demo.Behavior)string  {
	b.Run()
	b.Eat()
	return ""
}
