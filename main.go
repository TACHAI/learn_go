package main

import (
	"awesomeProject/interface_demo"
	"awesomeProject/struct_demo"
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

}

func action(b interface_demo.Behavior)string  {
	b.Run()
	b.Eat()
	return ""
}
