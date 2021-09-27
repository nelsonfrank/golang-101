package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main(){
	s := "Hello, World"

	fmt.Println(s) // Hello, World
	fmt.Println(len(s)) // 12
	fmt.Println(s[0]) // 72 ascii value of H
	fmt.Printf("%c", s[0]) // H
	fmt.Println() // new line

	fmt.Println(s[0:6]) // Hello,
	fmt.Println(s[6:]) // World

	s = s + " Again"
	s += " Again"

	fmt.Println(s) // Hello, World Again Again

	// String literal
	fmt.Println("Hello, \nWorld")
	fmt.Println("Hello, \tWorld")
	fmt.Println("Hello, \bWorld")

	// strings are slice of byte
	abc := "abc"
	b := []byte(abc)
	fmt.Println(abc, b) // abc [97 98 99] --> value [ascii values]
	
	fmt.Printf("%s %s", abc, b) // abc abc

	fmt.Println()

	abc2 := string(b)
	fmt.Println(abc2) // abc

	// non-ascii character
	// golang also support non-asscii character like mandarin character
	// <-- Here -->

	// String Library
	fmt.Println(strings.ToUpper("Hello, World!"))
	fmt.Println(strings.ToLower("Hello, World!"))
	fmt.Println(strings.HasPrefix("Hello, World!", "Hello")) // true
	fmt.Println(strings.HasSuffix("Hello, World!", "World!")) // true
	fmt.Println(strings.HasSuffix("Hello, World!", "World")) // false
	fmt.Println(strings.Replace("Hello, World!", "World", "Ninja", 1))
	fmt.Println(strings.Replace("Hello, World World", "World", "Ninja", 2))
	fmt.Println(strings.ReplaceAll("Hello, World World", "World", "Ninja",))

	// String builder
	// is a tool that you can utilize to build your strings

	var sb strings.Builder

	fmt.Println("String Builder --> ", sb.String())

	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.WriteString("Hello")

	fmt.Println("String Builder --> ", sb.String())

	fmt.Println(sb.Cap()) // output : 8 capacity_range:0 2 4 8
	fmt.Println(sb.Len()) // 5 --> "Hello"

	sb.Grow(100) // 2*currentCap + 100

	sb.Reset()
	fmt.Println(sb.Cap()) // output: 116 why?: 2*currentCap + num, is how Grow work
	fmt.Println(sb.Len())
	fmt.Println(sb.String())

	// fmt.Println()

	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())


	x := 123
	y := strconv.Itoa(x)
	fmt.Println(y)
	fmt.Printf("%T",y)
	fmt.Println()

	z, _ := strconv.Atoi(y)
	fmt.Println(z)
	fmt.Printf("%T",z)
}