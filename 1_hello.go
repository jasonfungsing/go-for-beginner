package main

import "fmt"

func main() {

	fmt.Println("Hello World!")

	var age int = 34

	var favNum float64 = 1.324423

	randNum := 1

	fmt.Println(age, " ", favNum, randNum)

	var numOne = 1.000
	var num99 = .9999

	fmt.Println(numOne - num99)

	const pi float64 = 3.14

	var (
		varA = 2
		varB = 3
	)
	fmt.Println(varA, varB)

	// string
	var myName string = "Jason"
	fmt.Println(myName + " Feng")
	fmt.Println("")

	// boolean
	var isOver34 bool = true
	fmt.Println("// boolean")

	fmt.Printf("%t \n", isOver34)
	fmt.Println("true && false =", true && false)

	fmt.Println("//")
	fmt.Println("")

	// for loop
	fmt.Println("// for loop")

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	fmt.Println("//")
	fmt.Println("")

	// if statement
	fmt.Println("// if statement")

	yourAge := 18
	if yourAge > 16 {
		fmt.Println("You can drive since you are", yourAge)
	} else if yourAge >= 18 {
		fmt.Println("You can vote since you are", yourAge)
	} else {
		fmt.Println("You can still enjoy your Life")
	}

	fmt.Println("//")
	fmt.Println("")

	// switch statement
	fmt.Println("// switch statement")

	switch yourAge {
	case 16:
		fmt.Println("You can drive since you are", yourAge)
	case 18:
		fmt.Println("You can vote since you are", yourAge)
	default:
		fmt.Println("You can still enjoy your Life")
	}
	fmt.Println("//")
	fmt.Println("")

	// array
	fmt.Println("// array")

	var favNums [5] float64
	favNums[0] = 163
	favNums[1] = 15.32
	favNums[2] = 155.32
	favNums[3] = 15.3254
	favNums[4] = -15.32
	fmt.Println(favNums)

	favNums2 := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(favNums2)

	for i, value := range favNums2 {
		fmt.Println(i, value)
	}

	for _, value := range favNums2 {
		fmt.Println(value)
	}
	fmt.Println("//")
	fmt.Println("")

	// slice
	fmt.Println("// slice")
	numSlice := []int{5, 4, 3, 2, 1}

	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2[0] =", numSlice2[0])
	fmt.Println("numSlice2[1] =", numSlice2[1])

	fmt.Println("numSlice[:2] =", numSlice[:2])
	fmt.Println("numSlice[2:0] =", numSlice[2:])

	numSlice3 := make([]int, 5, 10) // func make([]T, len, cap) []T
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3)

	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3)

	fmt.Println("//")
	fmt.Println("")

	// map
	fmt.Println("// map")
	ageMap := make(map[string]int)
	ageMap["Jason"] = 34

	fmt.Println(ageMap)
	fmt.Println(len(ageMap))

	ageMap["John"] = 35
	delete(ageMap, "John")

	fmt.Println("//")
	fmt.Println("")
}
