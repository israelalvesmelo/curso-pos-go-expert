package main

import "github.com/israelalvesmelo/curso-pos-go-expert/APIs/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
