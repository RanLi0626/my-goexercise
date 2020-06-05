package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancle := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("over!")
				fmt.Println(ctx.Err())
				return
			default:
				fmt.Println("working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancle()
	time.Sleep(5 * time.Second)

}
