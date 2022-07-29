package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	h "github.com/WWtLF/hello_grpc/pkg/api/hello"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	my_uuid string = uuid.New().String()
)

type PingServer struct {
	h.UnimplementedPingServer
}

func (s *PingServer) SayHello(test *h.Test, p h.Ping_SayHelloServer) error {
	fmt.Println("NEW CONNECTION")
	send(test.Test, my_uuid, p)
	fmt.Println("Stream finished")
	return nil
}

func send(req string, text string, p h.Ping_SayHelloServer) {
	for i := 0; i < 100; i++ {
		fmt.Printf("Request %s, I send %s: %d \n\r", req, text, i)
		err := p.Send(&h.Test{Test: fmt.Sprintf("%s: %d", text, i)})
		if err != nil {
			fmt.Println("ERROR FINISHED ", err.Error())
			return
		}
		time.Sleep(time.Second)
	}
	fmt.Println("NORMAL FINISHED")
}

func main() {

	fmt.Println("My POD IP is ", os.Getenv("POD_IP"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	h.RegisterPingServer(s, &PingServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
