package main

import "fmt"

func main(){
		fmt.Println("1 + 1 =", 1 + 1) // 2
		fmt.Println("1.0 + 1.0 =", 1.0 + 1.0) // 2
		stringOperations("Hello World")
}


func stringOperations(x string){
	fmt.Println("\nx:\t\t", x)
	fmt.Println("len(x):\t\t", len(x))
	fmt.Println("x[0]:\t\t",x[0])
	fmt.Println("string(x[0]):\t", string(x[0]))
}