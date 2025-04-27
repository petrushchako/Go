package main

import (
	"fmt"
	"strings"
)

func main(){
		fmt.Println("1 + 1 =", 1 + 1) // 2
		fmt.Println("1.0 + 1.0 =", 1.0 + 1.0) // 2
		stringOperations("Hello World")
		nextLetter("Hello World")
		spliter("Republic of Ireland")
		booleans()
		printMultiLines()
}


// Strings

func stringOperations(x string){
	fmt.Println("\nx:\t\t", x)
	fmt.Println("len(x):\t\t", len(x))
	fmt.Println("x[0]:\t\t",x[0])
	fmt.Println("string(x[0]):\t", string(x[0]))
}

func nextLetter(x string){
	y := x[0]
	fmt.Println(string(y+1))
}

func spliter(x string){
	splited := strings.Split(x, "")
	fmt.Println(splited)
}


// Booleans

func booleans(){
	var x bool = true
	var y bool = false
	fmt.Println("\nBooleans:")
	fmt.Println((!x))
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!(x && y))
	fmt.Println(convertBoolToInteger(!(x || y)))
}

func convertBoolToInteger(x bool) int {
	if x {
		return 1
	}
	return 0
}



func printMultiLines(){
	fmt.Println(`1
2
3
4`)
}