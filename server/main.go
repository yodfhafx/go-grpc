package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yodfhafx/go-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	var s *grpc.Server

	tls := flag.Bool("tls", false, "Enable TLS Connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Register
	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	reflection.Register(s)

	// Run server
	fmt.Print("gRPC server listening on port 50051")
	if *tls {
		fmt.Println(" with TLS")
	} else {
		fmt.Println()
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
