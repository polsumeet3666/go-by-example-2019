package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("inital i: ", i)

	zeroval(i)
	fmt.Println("zeroval i:", i)

	zeroptr(&i)
	fmt.Println("zerptr i:", i)

	fmt.Println("pointer:", &i)

}
