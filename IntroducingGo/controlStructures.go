package main
import "fmt"

func main(){
    forLoop1()
    forLoop2()

    ifSwitch()
}

func forLoop1 (){
    i := 1
    for i <=10 {
        fmt.Print(i, " ")
        i++
    }
    fmt.Println()
}

func forLoop2(){
    for i:=10; i >0; i-- {
        fmt.Print(i, " ")
    }
    fmt.Println()
}

func ifSwitch(){
    for i:=1; i <= 10; i++  {
        if i % 2 ==0{
            fmt.Println(i, " : even ")
        }else{
            fmt.Println(i, " : odd")
        }
    }
}