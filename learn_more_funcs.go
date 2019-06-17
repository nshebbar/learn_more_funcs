package main

import "fmt"

func addOne(a int) int {
	return a + 1
	}

func addTwo(a int) int {
	return a + 2
}
func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b 
	}
}

func makeDoubler(f func(int) func(int) int){
	return func(a int) int{
		b := f(a)
		return b * 2
	}
}


func main() {

	
	myaddone := addOne
	fmt.Println(addOne(1))
	fmt.Println(myaddone(1))

	b := 2
	anotherAddOne := func(a int) int {
		fmt.Println("inside anotherAddone()")
		return a + b
	}
	
		mineAddone := func(a int) int {
		return a + 2
	}
	fmt.Println(mineAddone(1))
	fmt.Println(anotherAddOne(1))
	printOperation(1, addOne)
	printOperation(1, addTwo)

	newAddOne := makeAdder(5)
	newAddTwo := makeAdder(10)

	fmt.Println(newAddOne(5))
	fmt.Println(newAddTwo(10))
	
	AddOne := makeAdder(1)
	doubleAddOne := makeDoubler(AddOne)

	fmt.Println(addOne(100))
	fmt.Println(doubleAddOne(100))
}