package main

import (
	"flag"
	"log"

	"github.com/yodfhafx/go-grpc-client/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:50051", "gRPC server address")
	tls := flag.Bool("tls", false, "use TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		creds = insecure.NewCredentials()

	}

	cc, err = grpc.Dial(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	err = calculatorService.Hello("John")
	// err = calculatorService.Fibonacci(5)
	// err = calculatorService.Average(1, 2, 3, 4, 5)
	// err = calculatorService.Sum(1, 2, 3, 4, 5)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			log.Printf("[%v] %v", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}
	}
}
