package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// Membuat child context
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	// Melihat hierarki parent dan child
	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Mengambil isi value dari context dengan keynya
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))

	// Context hanya mencari ke parentnya, tidak ke childnya
	fmt.Println(contextA.Value("b"))
}

func CreateCounterGoroutinesLeak() chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestContextGoroutinesLeak(t *testing.T) {
	fmt.Println("Total Goroutines Start:", runtime.NumGoroutine())
	destination := CreateCounterGoroutinesLeak()
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	fmt.Println("Total Goroutines Finish:", runtime.NumGoroutine())
}

func CreateCounterWithCancel(ctx context.Context, group *sync.WaitGroup) chan int {
	destination := make(chan int)
	group.Add(1)
	go func() {
		defer close(destination)
		defer group.Done()
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			case destination <- counter:
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutines Start:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	group := &sync.WaitGroup{}

	destination := CreateCounterWithCancel(ctx, group)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	cancel()
	group.Wait()

	fmt.Println("Total Goroutines Finish:", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutines Start:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 10*time.Second)
	group := &sync.WaitGroup{}
	defer cancel()

	destination := CreateCounterWithCancel(ctx, group)
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	group.Wait()

	fmt.Println("Total Goroutines Finish:", runtime.NumGoroutine())
}

func TestContextWithDeadLine(t *testing.T) {
	fmt.Println("Total Goroutines Start:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10*time.Second))
	group := &sync.WaitGroup{}
	defer cancel()

	destination := CreateCounterWithCancel(ctx, group)
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	group.Wait()

	fmt.Println("Total Goroutines Finish:", runtime.NumGoroutine())
}
