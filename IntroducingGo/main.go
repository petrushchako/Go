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
}


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