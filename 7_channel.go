package main

import (
	"fmt"
	"time"
	"strconv"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make dough and Send for sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add Sauce and Send", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addTopping(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add Toppings", pizza, "and ship")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func main() {

	stringChan := make(chan string)

	for i := 0; i <= 3; i ++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addTopping(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}
}