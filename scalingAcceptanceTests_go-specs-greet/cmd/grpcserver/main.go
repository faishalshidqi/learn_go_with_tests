package main

import (
	"go-specs-greet/adapters/grpcserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
	type GreetServer struct {
		grpcserver.UnimplementedGreeterServer
	}

	func (g GreetServer) Greet(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
		//return &grpcserver.GreetReply{Message: "Hello " + request.Name}, nil
		return &grpcserver.GreetReply{Message: interactions.Greet(request.Name)}, nil
	}
*/
func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	grpcserver.RegisterGreeterServer(server, &grpcserver.GreetServer{})
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
