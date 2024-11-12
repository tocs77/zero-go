package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case <-ctx.Done():
				out <- fmt.Sprintf("%v operation cancelled", msg)
				return
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation complete", msg)
				return
			}
		}
	}()
	return out
}
func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "webserver", 200)
	microservice := sampleOperation(ctx, "microservice", 500)
	database := sampleOperation(ctx, "database", 900)

	for loop := true; loop; {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microservice:
			fmt.Println(val)
			fmt.Println("cancel context")
			cancelCtx()
			loop = false
		case val := <-database:
			fmt.Println(val)
		}
	}

	fmt.Println(<-database)
}
