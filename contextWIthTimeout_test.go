package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

//Context With Timeout used when a channel didn't get any data for any seconds
//So the Goroutines will be stopped or canceled
func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Goroutine : ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 20 * time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Goroutine : ", runtime.NumGoroutine())
	for n := range destination{
		fmt.Println("Counter", n)
	}

	fmt.Println("Goroutine : ", runtime.NumGoroutine())
}