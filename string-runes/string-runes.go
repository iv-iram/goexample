//In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.

package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {

    //s is a string assigned a literal value representing the word “hello” in the Thai language. Go string literals are UTF-8 encoded text.
    const s = "สวัสดี"

    fmt.Println("Len:", len(s))


	//Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
        examineRune(runeValue)
    }
}

func examineRune(r rune) {

    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}