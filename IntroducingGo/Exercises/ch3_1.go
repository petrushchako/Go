package main

import "fmt"

func main(){
	fmt.Println("Enter Fahrenheit:")
	var fahrenheit float32
	fmt.Scanf("%f", &fahrenheit)
	fmt.Println("\nThe Celsium equivalent to", fahrenheit,"is", (fahrenheit-32)*5/9)
}