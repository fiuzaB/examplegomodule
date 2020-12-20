package hello
// When you add a package it will automatically download it in your project, if you dont have it, pretty cool.
// Check the go.mod to see what packages a project requires. When you use import, it will also automatically,
// write the require in the go.mod file
import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}
