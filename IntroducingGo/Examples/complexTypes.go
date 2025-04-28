package main
import "fmt"

func main(){
	fmt.Println("Arrays Examples:")
	arraysExample1()
	arraysExample2()
	calculateSum([5]float64{12.23, 23.34, 34.45, 45.56, 56.67})
}


func arraysExample1(){
	var arr [2]string
	arr[0] = "Oleksandr"
	arr[1] = "Petrushchak"
	fmt.Println(arr[0]+ " " + arr[1])
}

func arraysExample2(){
	var arr [8]int
	arr[0] = 1
	arr[2] = 1
	arr[3] = 1
	arr[6] = 1
	fmt.Println(arr, "8-bit binary number is", calculate10Base(arr), "in 10 base")
}

func calculate10Base( bits [8]int) int{
	// (1 << (7 - i)) in Go << is the bitwise left shift operator.
	// 1 << n means: take 1 and shift it left by n positions.
	// Shifting left by n bits is the same as multiplying by 2ⁿ.
	var result int
	for i:=0; i < 8; i++ {
		result += bits[i] * (1 << (7 - i))
	}
	return result
}

func calculateSum(arr [5]float64){
	var total float64
	for i:=0; i<5; i++{
		total += arr[i]
	}
	fmt.Println("Total sum of", arr, "is", total)
}