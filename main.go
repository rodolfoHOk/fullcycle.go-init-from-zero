package main

import (
	"fmt"
	"math/rand"
	"time"
)

func webServerWorkerOld(msg int) {
	fmt.Println("Mensagem processada: ", msg)
	time.Sleep(time.Second)
}

func mainOld() { // ineficiente
	for i := 0; i < 10; i++ {
		webServerWorkerOld(i)
	}
}

func webServerWorker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("WorkerId: ", workerId,"Mensagem processada: ", res)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func mainOld2() {
	msg := make(chan int)
	for i := 0; i < 3; i++ {
		go webServerWorker(i, msg)
	}
	for i := 0; i < 10; i++ {
		msg <- i
	}
	time.Sleep(time.Second * 2)
}

// Thread 1
func mainOld3() {
	hello := make(chan string)

	// Thread 2
	go func() {
		hello <- "Hello World"
	}()

	// Thread 1
	fmt.Println(<-hello)
}

func worker1(done chan string) {
	fmt.Println("worker1 start...")
	time.Sleep(time.Second * 2)
	fmt.Println("worker1 done")
	done <- "Worker1 é f*"
}

func worker2(done chan string) {
	fmt.Println("worker2 start...")
	time.Sleep(time.Second * 2)
	fmt.Println("worker2 done")
	done <- "Worker2 é f*"
}

func main() {
	done := make(chan string)
	go worker1(done)
	go worker2(done)
	fmt.Println(<- done)
}
