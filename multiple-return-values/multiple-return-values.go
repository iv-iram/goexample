// Go has built-in support for _multiple return values_.


package multiple_return_functions

import "fmt"

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int) {
	return 3, 7
}

func Multiple_return_functions() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If want a subset of the returned values,
	// use the blank identifier `_`.
	_, c := vals()
	fmt.Println(c)
}
