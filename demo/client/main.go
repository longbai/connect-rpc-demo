package main

import (
    greetv1 "connect-rpc/greet/v1"
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "log"
)

func main() {
    addr := "localhost:8080"
    opt := grpc.WithTransportCredentials(insecure.NewCredentials())
    conn, err := grpc.Dial(addr, opt)
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    c := greetv1.NewGreetServiceClient(conn)
    r, err := c.Greet(
        context.Background(),
        &greetv1.GreetRequest{Name: "Jane"})
    if err != nil {
        log.Println(err)
    }
    log.Println(r.GetGreeting())
}
