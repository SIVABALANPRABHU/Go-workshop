package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hellow, World!")
	var age int = 22
	var name string = "Sivabalan Prabhu"
	var height float32= 6.01
	fmt.Printf("My name is %s and my age is  %d\n",name,age)
	fmt.Printf("my height is %.2f\n",height)

	place := "Ramanathapuram"
	runs := 100
	fmt.Println("I am from", place)

	if runs > 100{
		fmt.Println("You have scored a century")
	} else{
		fmt.Println("You have scored less than a century")
	}


	for i :=0; i <5; i++{
		fmt.Println(i)
	}

	i := 10
	for i < 15{
		fmt.Println(i)
		i++
	}

	j := 20
	for {
		if j >25 { break }
		fmt.Println("Infinite loop")
		j++
	}

	
	nums := [5]int{ 100, 200, 300, 400, 500}
	for index, value := range nums{
		fmt.Println(index,value)
	}

	var x  int64 = 10
	var y int64 = 20 
	fmt.Println("Sum of the numbers is", add(x,y))
	fmt.Println("The value of x is", x)

	//Pointers
	var xp *int64
	xp = &x
	fmt.Println("The value of xp is", xp)
	xp = &y
	fmt.Println("The value of xp is", xp)



	fmt.Println("Current Value of x is ", x)
	fmt.Println("Sum of the numbers is", addByReference(&x,y))
	fmt.Println("Value of x after add by reference is ", x)

	type Person struct{
		name string
		age int
		height float32
		place string
	}

	p := Person{ name: "Siva", age:30, height: 6.01, place: "ramnad"}
	fmt.Println(p.name, p.age, p.height, p.place)

	q := new(Person)
	q.name = "Kumar"
	q.age = 25
	q.height = 5.11
	q.place = "Chennai"
	fmt.Println(q.name, q.age, q.height, q.place)
	

}

func add(x int64, y int64) int64 {
	x = x + 10
	return x + y
}

func addByReference(x *int64, y int64) int64 {
	*x = *x + 10
	return *x + y
}