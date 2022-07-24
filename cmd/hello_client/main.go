package main

import (
	"context"
	"fmt"
	"log"
	"time"

	h "github.com/WWtLF/hello_grpc/pkg/api/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := h.NewPingClient(conn)

	// Contact the server and print out its response.

	// defer cancel()
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second)
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := c.SayHello(context.Background(), &h.Test{Test: fmt.Sprintf("request number %d", i)})

		if err != nil {
			fmt.Printf("could not ping: %v\n\r", err)
			continue
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
		// cancel()

	}

}
