package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	h "github.com/WWtLF/hello_grpc/pkg/api/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func main() {

	resolver.SetDefaultScheme("dns")
	conn, err := grpc.Dial(os.Getenv("SERVER_HOST_PORT"), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := h.NewPingClient(conn)

	// Contact the server and print out its response.

	// defer cancel()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		time.Sleep(time.Second)
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		wg.Add(1)
		go func(wg *sync.WaitGroup, gr_num int) {
			fmt.Println("Starting reqiuest goroutin ", gr_num)
			client, err := c.SayHello(context.Background(), &h.Test{Test: fmt.Sprintf("request number %d", gr_num)})

			if err != nil {
				fmt.Printf("could not ping: %v\n\r", err)
				return
			}

			for {
				test, recvErr := client.Recv()
				if recvErr != nil {
					fmt.Println("Error ", recvErr.Error())
					client.CloseSend()
					break
				}
				if test != nil {
					fmt.Println(test.Test)
				}

			}
			wg.Done()
		}(&wg, i)
		// cancel()
	}
	wg.Wait()
}
