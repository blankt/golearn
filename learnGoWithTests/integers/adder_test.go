package integers

import "fmt"

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// out put 6
}
