package main

import "fmt"

func ExecEX1() {
	fmt.Println("execute base examples..")
	var ex1 Example1
	ex1.Start()

}

func ExecServer() {
	const host = "localhost:4000"
	fmt.Println("execute basic servers..")
	var h Hello
	// h.Start(host)
	h.Start2(host)
}

func ExecGoroutine() {
	var gr Goroutines
	// gr.sayInterval()
	// gr.goSum()
	// gr.fibonacci()
	gr.selectfibonacci()
}

func main() {
	// ExecEX1()
	// ExecServer()
	ExecGoroutine()
}