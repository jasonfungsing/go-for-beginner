package main

import "fmt"

func main(){
    
    x := 0

    changeXVal(x)

    fmt.Println("x = ", x)

    changeXValPointer(&x)

    fmt.Println("x = ", x)
    fmt.Println("Memory address for x =", &x)

    yPointer := new(int)
    changeYValPointer(yPointer)
    fmt.Println("y =", *yPointer)
}

func changeXVal(x int){
    x = 2
}

func changeXValPointer(x *int){
    *x = 2
}

func changeYValPointer(yPointer *int){
    *yPointer = 100
}