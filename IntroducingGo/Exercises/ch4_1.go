package main

import "fmt"

func main(){
	fmt.Println("Numbers between 1 and 100 that are evenly divisible by 3:")

	for i:=1; i<=100; i++{
		if i%3 == 0{
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}