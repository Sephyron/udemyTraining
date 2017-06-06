package main

import "fmt"

var a string = "this is stored in a variable a"
var b, c string = "stored in b", "stored in c"
var d string

func main() {

	d = "stored in d"
	var e int = 42
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in j"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := `o`

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
}
