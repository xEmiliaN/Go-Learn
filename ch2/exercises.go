package main

import "fmt"

func main() {
	task1()
	task2()
	task3()
}

// Задача 1

func task1() {
	var i int = 20
	var f float32 = float32(i)

	fmt.Println("\nЗадача 1")
	fmt.Printf("i = %d, f = %f\n", i, f)
}

// Задача 2
func task2() {
	const value = 10
	var i int = value
	var f float32 = value

	fmt.Println("\nЗадача 2")
	fmt.Printf("i(int) = %d, f(float) = %f\n", i, f)
}

// Задача 3
func task3() {
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	b += 1
	smallI += 1
	bigI += 1
	fmt.Println("\nЗадача 3")
	fmt.Printf("b(byte) = %d, smallI(int32) = %d, bigI(uint64) = %d\n", b, smallI, bigI)
}
