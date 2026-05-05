package latihan

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

//Soal 2 — context.WithCancel
//Buat goroutine yang terus bekerja sampai di-cancel. Simulasinya:

//Goroutine print "sedang bekerja..." setiap 500ms
//Di main, tunggu 2 detik lalu cancel context-nya
//Goroutine harus berhenti saat di-cancel dan print "pekerjaan dihentikan"

func Pekerja(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("pekerjaan dihentikan.")
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("sedang bekerja...")
			}
		}
	}()
}

func Latihan2() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	background := context.Background()
	ctx, cancelFunc := context.WithCancel(background)
	wg := &sync.WaitGroup{}

	Pekerja(ctx, wg)
	time.Sleep(2 * time.Second)

	cancelFunc()
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
