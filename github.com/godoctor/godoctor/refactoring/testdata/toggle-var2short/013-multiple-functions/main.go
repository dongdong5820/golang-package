// <<<<< toggle,19,2,19,27,pass
package main

import "fmt"

func f() int {
	return 1
}

func g() string {
	return "hello"
}

func h() float64 {
	return 5.78
}

func main() {
var i, j, k = f(), g(), h()
fmt.Println("Value of i,j,k:", i, j, k)
}
