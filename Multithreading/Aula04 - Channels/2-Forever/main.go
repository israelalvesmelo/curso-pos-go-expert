package main

// Thread 1
func main() {
	forever := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true // Envia a mensagem para o canal evitando o deadlock
	}()
	<-forever // Thread 1 fica bloqueada aqui
}
