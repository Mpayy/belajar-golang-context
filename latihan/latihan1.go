package latihan

import (
	"context"
	"fmt"
)

//Soal 1 — context.WithValue
//Simulasikan request masuk ke sistem. Buat function ProsesPesanan(ctx context.Context) yang:

//Baca nilai "user_id" dan "request_id" dari context
//Kalau salah satunya tidak ada, print "context tidak lengkap" dan return
//Kalau ada, print "Memproses pesanan untuk user [user_id] dengan request [request_id]"

// Panggil dari main dengan context yang sudah diisi kedua nilainya.
type contextKey string

const userIDKey contextKey = "user_id"
const requestIDKey contextKey = "request_id"

func ProsesPesanan(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)

	if userID == nil || requestID == nil {
		fmt.Println("context tidak lengkap")
		return
	}

	fmt.Println("Memproses pesanan untuk user", ctx.Value(userIDKey), "dengan request", ctx.Value(requestIDKey))
}

func Latihan1() {

	contextParent := context.Background()
	contextUserId := context.WithValue(contextParent, userIDKey, "Achmad Rifaih")
	contextRequestId := context.WithValue(contextUserId, requestIDKey, "Djarum Coklat")

	ProsesPesanan(contextRequestId)
}
