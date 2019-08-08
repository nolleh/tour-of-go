package main

import (
	"fmt"
	"time"
)

type Goroutines struct {}

func (g *Goroutines) say(s string) {
	for i := 0; i <5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func (g *Goroutines) sayInterval() {
	go g.say("world") // 현재의 고루틴에서 say, world 가 eval 되고 새 고루틴에서 say 가 수행됨
	g.say("hello")
	// [[http://golang.org/pkg/sync/][sync]] 패키지를 통한 동기화 필요.
	
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func (g *Goroutines) goSum() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) // 7, 2, 8
	go sum(a[len(a)/2:], c) // -9, 4, 0
	x := <- c
	y := <- c
	// x, y := <-c, <-c // 채널은 내부에 큐가 있어서, 하나씩 데이터를 전달해 줄 수 있다.

	// 동기화가 필요없는 이유 : 양 편이 준비 될 때까지 블럭
	fmt.Println(x, y, x+y)
}

func (g *Goroutines) buffering() {
	c := make(chan int, 1)
	c <- 1
	// 만약 여기서 이 구문을 실행하면 channel 의 버퍼가 꽉차서 더이상 데이터를 넣을 수 없으므로 블럭된다. 
	// 다른 고루틴에서 1을 가져가주면 이 구문을 실행해도 문제가 안될 것.
	// c <- 2 
	fmt.Println(<-c)
	c <- 2
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func (g *Goroutines) fibonacci() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// 송신측에서 채널을 닫으면 v, ok := <- ch 의 ok 가 false 가 되고, 
	// range c 는 채널이 닫힐때까지 값을 받는다. 
	for i := range c {
		fmt.Println(i)
	}
}

func selectFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // case 의 구문중 하나가 수행 될 수 있을때까지 블럭
		case c <- x: // x 를 채널에 넣을 수 있거나 
			x, y = y, x+y
		case <- quit: // quit 이 들어왔을 때
			fmt.Println("quit")
			return
		default: // 위 채널의 송 or 수신이 block 된 상태인동안 실행
			fmt.Println("   .")
			time.Sleep(5e7)
		}
	}
}

func (g *Goroutines) selectfibonacci() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	selectFibonacci(c, quit)
}