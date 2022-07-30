package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	h "github.com/WWtLF/hello_grpc/pkg/api/hello"
	"github.com/WWtLF/hello_grpc/pkg/tracing"
	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
)

var (
	my_uuid string = uuid.New().String()
)

type PingServer struct {
	h.UnimplementedPingServer
	// tr trace.Tracer
}

func (s *PingServer) SayHello(test *h.Test, p h.Ping_SayHelloServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// ctx, span := s.tr.Start(ctx, "SayHello_server")
	// defer span.End()

	fmt.Println("NEW CONNECTION")
	s.send(ctx, test.Test, my_uuid, p)
	fmt.Println("Stream finished")
	return nil
}

func (s *PingServer) send(ctx context.Context, req string, text string, p h.Ping_SayHelloServer) {
	for i := 0; i < 100; i++ {
		// _, span := s.tr.Start(ctx, "SayHello_SEND")
		// span.End()
		fmt.Printf("Request %s, I send %s: %d \n\r", req, text, i)
		err := p.Send(&h.Test{Test: fmt.Sprintf("%s: %d", text, i)})
		if err != nil {
			fmt.Println("ERROR FINISHED ", err.Error())
			return
		}
		// time.Sleep(time.Second)
		// span.End()
	}
	fmt.Println("NORMAL FINISHED")
}

func main() {
	ctx := context.Background()
	tp, err := tracing.TracerProvider("http://jaeger-all-in-one:14268/api/traces")
	defer func() { _ = tp.Shutdown(ctx) }()
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	// tr := tp.Tracer("component-main")

	fmt.Println("My POD IP is ", os.Getenv("POD_IP"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	// h.RegisterPingServer(s, &PingServer{tr: tr})
	h.RegisterPingServer(s, &PingServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
