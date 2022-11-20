package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num uint = 299                                 // explicit declaration
	var number = 3000                                  // implicit declaration
	walrus := "Expression Assignnment operator walrus" // declaration + assignment
	fmt.Println(walrus)
	fmt.Printf("%T", number)                  // type check
	fmt.Println("Hello, World!", num, number) // print line
	sprintF := fmt.Sprintf("I am number %.f", 1.1)
	fmt.Println(sprintF)

	scannerObj := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a number: ")
	scannerObj.Scan()
	value, _ := strconv.ParseInt(scannerObj.Text(), 10, 64)
	for value <= 10 {
		fmt.Printf("Your value %d \n", value)
		value++
	}
	switchNum := 7
	switch switchNum {
	case 11:
		fmt.Println("condition 1")
	case 12:
		fmt.Println("condition 2")
	default:
		fmt.Println("default")
	}
	fmt.Println("Goodbye!")
}
