package latihan

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Soal 4 — Gabungan: Context di goroutine paralel
//Kamu punya 3 "worker" yang berjalan paralel. Kalau salah satu worker gagal, semua worker lain harus ikut berhenti.
//goworkers := []string{"Worker A", "Worker B", "Worker C"}

//Setiap worker jalan dan print "[nama] bekerja" setiap 300ms
//Worker B gagal setelah 1 detik (simulasi error)
//Saat Worker B gagal, semua worker lain langsung berhenti
//Print "Semua worker dihentikan karena Worker B gagal"

func Latihan4() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	goworkers := []string{"Worker A", "Worker B", "Worker C"}
	for _, worker := range goworkers {
		wg.Add(1)
		go func(worker string) {
			defer wg.Done()

			ticker := time.NewTicker(300 * time.Millisecond)
			defer ticker.Stop()

			//var failTimer *time.Timer
			//if worker == "Worker B" {
			//	failTimer = time.NewTimer(1 * time.Second)
			//	defer failTimer.Stop()
			//}

			if worker == "Worker B" {
				go func() {
					time.Sleep(1 * time.Second)
					fmt.Println("Worker B gagal")
					cancelFunc()
				}()
			}

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					fmt.Println(worker, "bekerja")
					//case <-func() <-chan time.Time {
					//	if failTimer != nil {
					//		return failTimer.C
					//	}
					//	return nil
					//}():
					//	fmt.Println("Worker B gagal")
					//	cancelFunc()
					//	return
				}
			}
		}(worker)
	}
	wg.Wait()
	fmt.Println("Semua worker dihentikan karena Worker B gagal")
}
