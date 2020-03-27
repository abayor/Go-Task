package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to DectoBin ")
	fmt.Println("This program converts Decimal numbers to Binary")
	fmt.Println("Kindly input your Decimal Numner")

	var decimal int64

	fmt.Scan(&decimal)
	fmt.Println("Decimal Number is %s", decimal)
	binary := strconv.FormatInt(decimal, 2)
	fmt.Printf("decimal = %d, in binary format = %d\n", decimal, binary)

}
