package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0
// Ao executar o comando go run -race main.go, podemos ver que o programa não está sendo executado corretamente. Problemas relacionados a concorrência.
func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1) // Esta função é usada para garantir o mutex da variavel number
		// m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
