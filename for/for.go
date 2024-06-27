//for is Go’s only looping construct. Here are some basic types of for loops.
package forloop
import "fmt"

func Forloop(){
	i:=1
	for i <= 3 {
		fmt.Println(i)  
		i = i + 1
	}

	for j:=0;j<3;j++{
		fmt.Println(j)
	}

	//Another way of accomplishing the basic “do this N times” iteration is range over an integer.

	for i:= range []int {0,1,2} {
		fmt.Println("range", i)
	}

	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

	for {
		fmt.Println("loop")
		break
	}


	for n  := range []int {0,1,2,3,4,5}{
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}


