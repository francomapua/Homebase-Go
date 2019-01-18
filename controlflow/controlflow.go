package main

import (
	"errors"
	"fmt"
)

func main() {
	DeferPanicRecover()
}
func DeferPanicRecover() {
	/*
		Defer
		- moves execution of statements till after function scope closes
		- creates a stack to which deferred statements go to (FILO)

		PANIC
		- ends the function, executes deferred functions, then returns to caller with arguments
		- puts the thread on panic mode, can only be stopped if someone calls recover()
		- use Panic if you are ABSOLUTELY SURE the program cannot continue; Useful for programmer-to-programmer communication

		RECOVER
		- stops panic
		- recover() returns what panic threw
	*/
	f()
	fmt.Println("Resolving after f()")
}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(i, errors.New())
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
