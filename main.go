package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	fmt.Println("Sum:", indexSumTrick(99999999999))

	rand.Seed(time.Now().Unix())
	num := rand.Int()
	num = num >> 56
	fmt.Println("Hello,", num)
	os.Exit(num)
}

func indexSum(size int) int {
	var s int
	for i := 0; i < size; i++ {
		s += i
	}
	return s
}

func indexSumTrick(size int) int {
	a := size - 1
	b := size - 2
	mul := a * b // this is a commen t
	return size + (mul >> 1) - 1
}
