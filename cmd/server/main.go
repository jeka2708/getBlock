package main

import (
	"flag"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"getBlockTest/pkg/api"
	pb "getBlockTest/proto/go_proto"
)

var (
	port   = flag.String("port", ":8080", "The server port")
	apiKey = flag.String("api_key", "", "The required api key")
)

func main() {

	flag.Parse()
	if *apiKey == "" {
		flag.Usage()
		os.Exit(1)
	}
	listener, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterControllerServer(server, api.NewGetBlock(*apiKey))
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
