package cal

import "fmt"

var C int = 10

func Add(a int, b int) (int, bool) {
	fmt.Println("staring add")
	c := a + b
	if c < 0 {
		return c, true
	}
	return c, false
}

// local variables
// global variables
