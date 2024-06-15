package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Write([]byte("Hello, world!\n"))
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		//log imprime no command line stdou - log no servidor
		log.Println("Request processada com sucesso")
		//Imprime no browser
		w.Write([]byte("Request preocessada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}

}
