package defer_go

import "fmt"

func Defer_go() {                 //defer works on LIFO principle
    defer fmt.Println("d1")
    defer fmt.Println("d2")
    defer fmt.Println("d3")
    defer fmt.Println("d4")

    fmt.Println("main function")
}
