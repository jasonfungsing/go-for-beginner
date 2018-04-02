package main

import "fmt"

func main(){

    rect1 := Retctange{0,50,10,10}
    fmt.Println("Retctange is", rect1.width, "wide")
    
    fmt.Println("Area of Retctange is", rect1.area())
    
}

type Retctange struct{
    leftX float64
    topY float64
    height float64
    width float64
}

func (rect * Retctange) area() float64{
    
    return rect.width * rect.height
}