/*Values
Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
Strings, which can be added together with +.   
Integers and floats.*/

package values
import "fmt"

func Values(){
	fmt.Println("go"+"lang")
	fmt.Println("1+1=",1+1)
	fmt.Println("7.0/3.0 =",7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}