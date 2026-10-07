// package main
// // Go file belongs to a package - everyone of them
// // package is group of code
// // main - executable program

// // fmt - package for formatted I/O
// import "fmt"


// // main function - entry point pf the program
// func main(){

// 	//using Println function from fmt package to print Hello
// 	// fmt.Println("Hello, World!")
// 	// fmt.Println("I am learning Go programming language")
// 	// fmt.Println(10+20)
// 	// fmt.Println(40)

// 	// var age int = 25
// 	// fmt.Println(age)
// 	// // a much shorter form can be
// 	// age := 27

// 	// name:="Henry"
// 	// age:=30
// 	// height:=5.10
// 	// isLearning:=false
// 	// Go automatically infers the type of variable based on the value assigned to it

// 	// lets see operators

// 	// a:=10
// 	// b:=3
// 	// fmt.Println(a+b)
// 	// fmt.Println(a-b)
// 	// fmt.Println(a/b)
// 	// fmt.Println(a*b)
// 	a:=40
// 	b:=50
// 	fmt.Println(add(a,b))
// 	// conditions
// 	if a>=40{
// 		fmt.Println("Pizza sucks")
// 	}else{
// 		fmt.Println("Pizza is awesome")
// 	}

// 	// Loops
// 	for i:=1;i<=5;i++{
// 		fmt.Println(i)
// 	}
// 	// or

// 	j:=1
// 	for j<=5{
// 		fmt.Println(j)
// 		j++
// 	}
// }

// func add(a int, b int) int{
// 	return a+b
// 	// why does functions matter?
// 	// we build each peices and boundaries of the program to have clear understadings
// 	// basically break problems into small units of behaviour
// }

package main

import "fmt"

func main(){
	name := "Sumit"
	age := 24

	fmt.Println(name)
	fmt.Println(age)

	result:=isAdult(age)
	fmt.Println(result)
}

func isAdult(age int) bool{
	return age>=18
}