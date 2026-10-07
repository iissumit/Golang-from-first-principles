// // package main

// // import "fmt"

// // func main() {
// // 	name := "John Cena"
// // 	age := 42
// // 	city := "West Newbury"

// // 	fmt.Println(name)
// // 	fmt.Println(age)
// // 	fmt.Println(city)

// // 	answer := canDrive(age)
// // 	fmt.Println(answer)

// // 	maximumAnswer1 := getMax(10, 20)
// // 	maximumAnswer2 := getMax(50, 30)
// // 	maximumAnswer3 := getMax(7, 7)

// // 	fmt.Println(maximumAnswer1)
// // 	fmt.Println(maximumAnswer2)
// // 	fmt.Println(maximumAnswer3)

// // 	score := 87
// // 	result := classifyScore(score)
// // 	fmt.Println(result)

// // 	// balance := 100000
// // 	// newBalance := deposit(balance, 5000)
// // 	// fmt.Println("Old balance", balance)
// // 	// fmt.Println("new balance", newBalance)
// // 	balance := 1000
// // 	balance = deposit(balance, 500)
// // 	fmt.Println(balance)
// // 	balance = deposit(balance, 500)
// // 		fmt.Println(balance)

// // 	balance = withdraw(balance, 200)
// // 		fmt.Println(balance)

// // 	balance = deposit(balance, 100)
// // 		fmt.Println(balance)

// // }

// // func canDrive(age int) bool {
// // 	return age >= 18
// // }

// // func getMax(a int, b int) int {
// // 	if a > b {
// // 		return a
// // 	}
// // 	return b
// // }

// // func classifyScore(score int) string {
// // 	if score >= 90 {
// // 		return "A"
// // 	} else if score >= 80 {
// // 		return "B"
// // 	} else if score >= 70 {
// // 		return "C"
// // 	} else if score >= 60 {
// // 		return "D"
// // 	}
// // 	return "F"
// // }

// // func deposit(balance int, amount int) int {
// // 	return balance + amount

// // }

// // func withdraw(balance int, amount int) int{
// // 	return  balance - amount

// //
// package main

// import "fmt"

// func change(x int) {
// 	x = 1000
// }

// func change2(y *int){
// 	*y = 999
// }
// func main() {
// 	x := 10
// 	y :=20

// 	change(x)
// 	fmt.Println(x)

// 	change2(&y)
// 	fmt.Println(y)
// }

// Zero values 

package main

import "fmt"

func main(){

	// In many languages, variables that are not assigned with something or havent been given a meaningful value yet
	// Go takes a different appproach
	// Every declared variable gets a default value based on its type
	// called zero values

	var age int
	fmt.Println(age)
	var age2 float64
	fmt.Println(age2)
	var age3 bool
	fmt.Println(age3)
	// var name string
	// fmt.Println(name) ""
	// empty string so it would look blank

	// Imagine you have balance variable var balance int
	// go guarantees balance has well defined value immediately 

	// No uninitialised integer
	// best to work with structs, pointers, slices, maps, configuration, HTTP REquest data, data base models, concurrent programs
	
}