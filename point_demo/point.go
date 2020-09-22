package point_demo

import "fmt"

func TestPoint()  {
	var count int =20
	//定义指针
	var countPoint *int
	//指针赋值
	countPoint =&count

	fmt.Println("count 变量的地址 :%x \n",&count)
	fmt.Println("countPoint 变量存储的地址 :%x \n",countPoint)

	fmt.Println("countPoint 指针指向的值 :%d \n",*countPoint)

}

func TestPointArr()  {
	a,b :=1,2
	//指针数组
	pointArr :=[...]*int{&a,&b}

	fmt.Println("指针数组 pointArr",pointArr)

	//数组指针
	arr :=[...]int{3,4,5}
	arrPoint :=&arr

	fmt.Println("数组指针 arrPoint",arrPoint)
}