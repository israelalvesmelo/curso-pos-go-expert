package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha123")
	boolHotel(ctx)
}

func boolHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
