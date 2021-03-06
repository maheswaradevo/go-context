package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

//Context with deadline used when we want to finished the process
//with the time that we decided to
func TestContextWithDeadline(t *testing.T){
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	parent := context.Background()
	
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
	defer cancel()
	
	destination := CreateCounter(ctx)
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter : ", n)
	}

	fmt.Println("Goroutines : ", runtime.NumGoroutine())
}