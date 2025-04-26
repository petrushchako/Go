package main

import "fmt"

func main(){
	fmt.Print("Enter feet value: \t")
	var feet float64
	fmt.Scanf("%f", &feet)
	fmt.Println("Meters equivalent: \t", feet * 0.3048)
}