package main

import "fmt"
import "math"

func main(){

    rect := Retctange{10,20}
    circ := Circle{4}

    fmt.Println("Retctange area = ", getArea(rect))
    fmt.Println("Circle area = ", getArea(circ))
}

type Shape interface {
    area() float64
}

type Retctange struct{
    height float64
    width float64
}

type Circle struct{
    radius float64
}

func (r Retctange) area() float64 {
    return r.height * r.width
}

func (c  Circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
    return shape.area()
}
