package mygoroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	//go say("world")
	//say("hello")
}

func sum(a []int, c chan int) {
	//fmt.Printf("before sum called %d \n", <-c)

	total := 0
	for _, v := range a {
		fmt.Printf("currrent num is %d \n", v)
		total += v
	}
	c <- total // send total to c
	fmt.Printf("after sum called %d \n", c)
}

func TestGoroutine2(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("sum1 is called")
	go sum(a[:len(a)/2], c)
	fmt.Println("sum2 is called")
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		fmt.Printf("call fibonacci,%d \n", i)
	}
	close(c)
}

func TestGoroutine3(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci4(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("select default case is called")
			runtime.Gosched()
		}
	}
}

func TestGoroutine4(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci4(c, quit)
}

func TestGoroutine5(t *testing.T) {
	println("TestGoroutine5 start .\n")
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(1 * time.Second):
				println("TestGoroutine5 go timeout1.\n")
				o <- true
				break
			}
			println("TestGoroutine5 go timeout2.\n")
			break
		}
	}()
	<-o
}

func TestGoroutineRuntime(t *testing.T) {

	numCpu := runtime.NumCPU()
	fmt.Printf("TestGoroutineRuntime, numcpu is %d \n", numCpu)
	runtime.GOMAXPROCS(numCpu)

	c := make(chan int, 2)
	fmt.Printf("TestGoroutineRuntime,1 num goroutine is %d \n", runtime.NumGoroutine())
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("TestGoroutineRuntime printf is called, %d \n", i)
		}
		defer func() {
			fmt.Printf("TestGoroutineRuntime end is called \n")
		}()
		c <- 0
		c <- 1
	}()
	fmt.Printf("TestGoroutineRuntime,2 num goroutine is %d \n", runtime.NumGoroutine())
	for {
		isTimeOut := false
		select {
		case v := <-c:
			println(v)
		case <-time.After(1 * time.Second):
			println("TestGoroutineRuntime 1 sec timeout.\n")
			isTimeOut = true
			break
		}
		if isTimeOut {
			break
		}
	}
}
