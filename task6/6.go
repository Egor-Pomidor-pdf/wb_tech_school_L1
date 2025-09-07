package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func worker(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func workerWitnCtx(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("workerWitnCtx - stop")
			return
		default:
			fmt.Println("workerWitnCtx - work")
			time.Sleep(1*time.Second)
			
		}
	}

}

func workerWithNotification(stopCh chan int) {
	for {
		select {
		case <- stopCh:
			fmt.Println("workerWithNotification - stop")
			return
		default:
			fmt.Println("workerWithNotification - work")
			time.Sleep(1*time.Second)
		}
	}
	
}

func workerWithCondition(stopFlag *atomic.Bool) {
	for {
		if stopFlag.Load() {
			fmt.Println("workerWithCondition - stop")
			return
		}
		fmt.Println("workerWithCondition worked")
		time.Sleep(1 * time.Second)
	}
}

func workerRunTimeExit() {
	for i := 0; i < 10; i++ {
		fmt.Println("workerRunTimeExit - work")
		if i == 7 {
			fmt.Println("workerRunTimeExit - stopOver")
			runtime.Goexit()
		}
	}
	fmt.Println("workerRunTimeExit - stop")
}


func main() {

	// остановка горутины закрытием канала
	ch1 := make(chan int)
	go worker(ch1)
	for i := 0; i<10; i++ {
		ch1 <- i
	}
	close(ch1)


	// остановка горутины с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())
	go workerWitnCtx(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	// остановка горутины через канал уведомления
	stopCh := make(chan int, 1)
	go workerWithNotification(stopCh) 
	time.Sleep(3 * time.Second)
	stopCh <- 1


	// остановка горутины с помощью условия
	var stopFlag atomic.Bool
	go workerWithCondition(&stopFlag)

	time.Sleep(2 * time.Second)
	stopFlag.Store(true)

	// остановка с помощью runtime.Goexit
	go workerRunTimeExit()

	time.Sleep(1*time.Second)
}