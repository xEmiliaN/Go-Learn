// модифицированный пример кода из:
// https://github.com/learning-go-book-2e/ch03/blob/main/sample_code/len_cap/main.go

package main

import "fmt"

func main() {
	var a []int16
	var x []int32
	var y []int64

	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("x: ", x, len(x), cap(x))
	fmt.Println("y: ", y, len(y), cap(y))

	a = append(a, 10)
	x = append(x, 10)
	y = append(y, 10)

	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("x: ", x, len(x), cap(x))
	fmt.Println("y: ", y, len(y), cap(y))

	a = append(a, 20)
	x = append(x, 20)
	y = append(y, 20)

	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("x: ", x, len(x), cap(x))
	fmt.Println("y: ", y, len(y), cap(y))

	a = append(a, 30)
	x = append(x, 30)
	y = append(y, 30)

	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("x: ", x, len(x), cap(x))
	fmt.Println("y: ", y, len(y), cap(y))

	a = append(a, 40)
	x = append(x, 40)
	y = append(y, 40)

	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("x: ", x, len(x), cap(x))
	fmt.Println("y: ", y, len(y), cap(y))

	a = append(a, 50)
	x = append(x, 50)
	y = append(y, 50)

	fmt.Println("a: ", a, len(a), cap(a))
	fmt.Println("x: ", x, len(x), cap(x))
	fmt.Println("y: ", y, len(y), cap(y))

}
