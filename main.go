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
}