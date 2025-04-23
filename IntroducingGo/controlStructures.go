package main
import "fmt"

func main(){
    i := 1
    for i <=10 {
        fmt.Print(i, " ")
        i++
    }
    fmt.Println()

    for i:=10; i >0; i-- {
        fmt.Print(i, " ")
    }
    fmt.Println()
}
