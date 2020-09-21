package struct_demo

import "fmt"


type Dog struct {
	Id int
	Name string
	Age int
}


type Cat struct {
	Id int
	Name string
	Age int
}

func (d *Dog)Run() string {
	fmt.Println("dog is ",d.Id," runing")
	return "dog runging"
}
func (d *Dog)Eat() string {
	fmt.Println("dog eating")
	return "dog eating"
}





func (c *Cat)Run() string  {
	fmt.Println("cat is ",c.Id," runing")
	return "cat runing"
}

func (c *Cat)Eat()string  {
	fmt.Println("cat is ",c.Id," eating")

	return "cat eating"
}