package latihan

import (
	"context"
	"fmt"
	"time"
)

//Soal 3 — context.WithTimeout
//Buat function AmbilData(ctx context.Context) yang simulasi mengambil data dari database (pakai time.Sleep(3 * time.Second)).
//Panggil function ini dengan timeout 2 detik. Program harus print:

//Mengambil data...
//Gagal: context deadline exceeded
//Bukan nunggu 3 detik sampai selesai.

func AmbilData(ctx context.Context) {
	fmt.Println("Mengambil data...")
	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Gagal:", err)
		return
	case <-time.After(3 * time.Second):
		fmt.Println("Done")
	}

}

func Latihan3() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	AmbilData(ctx)
}
