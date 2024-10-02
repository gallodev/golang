package main

import (
	"fmt"
	"gallodev/mathutils"
)


func main() {
	myDeferedVariabled := "Initial Variable"
	var name string
	var age int

	name = "Christian"
	age = 30

	const HUMAN_MAX_AGE = 100
	var agePlusTen = mathutils.Sum(age,10)
	var ageMinusFive = mathutils.Subtract(age,5)
	var lifeCycleAge = mathutils.GetPercentOfLifeTime(float32(age),HUMAN_MAX_AGE)

	fmt.Println("Hello World!")
	fmt.Printf("My variable with initialized value is: %v and the typeOf is: %T \n", myDeferedVariabled, myDeferedVariabled)
	fmt.Printf("My name is: %v \n", name)
	fmt.Printf("My age is: %v years old.\n", age)
	fmt.Printf("10 years from now will be: %v years. OMG !! \n", agePlusTen)
	fmt.Printf("5 Years ago i had %v \n", ageMinusFive)
	fmt.Printf("If you live until %v years you are in %.2f percent of your life \n", HUMAN_MAX_AGE, lifeCycleAge)

	
}