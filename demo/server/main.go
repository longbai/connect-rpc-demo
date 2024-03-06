package main

import (
    greetv1 "connect-rpc/greet/v1"
    "context"
    "google.golang.org/grpc"
    "net"
)

type GreeterService struct {
    greetv1.UnimplementedGreetServiceServer
}

func (g *GreeterService) Greet(c context.Context, r *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {
    name := r.GetName()
    return &greetv1.GreetResponse{Greeting: "Hello GR, " + name + "!"}, nil
}

func main() {
    // Create a new GreeterService
    greeter := &GreeterService{}
    server := grpc.NewServer()
    server.RegisterService(&greetv1.GreetService_ServiceDesc, greeter)
    list, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    err = server.Serve(list)
    if err != nil {
        panic(err)
    }
}
