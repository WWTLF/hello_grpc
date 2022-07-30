module github.com/WWtLF/hello_grpc

go 1.18

require (
	github.com/google/uuid v1.1.2
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.33.0
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/exporters/jaeger v1.8.0
	go.opentelemetry.io/otel/sdk v1.8.0
	go.opentelemetry.io/otel/trace v1.8.0
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
