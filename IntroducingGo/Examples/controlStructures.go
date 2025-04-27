package main
import "fmt"

func main(){
    forLoop1()
    forLoop2()

    ifStatement()
    switchStatement()
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

func ifStatement(){
    for i:=1; i <= 10; i++  {
        if i % 2 ==0{
            fmt.Println(i, " : even ")
        }else{
            fmt.Println(i, " : odd")
        }
    }
}

func switchStatement(){
    var input int16
    fmt.Print("Enter number 1-5: ")
    fmt.Scanf("%d", &input)
    switch input{
    case 0: 
        fmt.Println("Zero")
    case 1: 
        fmt.Println("One")
    case 2: 
        fmt.Println("Two")
    case 3: 
        fmt.Println("Three")
    case 4: 
        fmt.Println("Four")
    case 5: 
        fmt.Println("Five")
    default: 
        fmt.Println("Wrong number")
    }
}