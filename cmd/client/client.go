package main

import (
    greetv1 "connect-rpc/gen/greet/v1"
    "connect-rpc/gen/greet/v1/greetv1connect"
    "connectrpc.com/connect"
    "context"
    "github.com/bufbuild/httplb"
    "github.com/bufbuild/httplb/picker"
    "log"
)

func main() {
    cl := httplb.NewClient(
        // Switch from the default round-robin policy to power-of-two.
        httplb.WithPicker(picker.NewPowerOfTwo),
    )
    defer cl.Close()
    client := greetv1connect.NewGreetServiceClient(
        cl,
        "h2c://localhost:8080/",
        connect.WithGRPC(),
    )
    res, err := client.Greet(
        context.Background(),
        connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
    )
    if err != nil {
        log.Println(err)
        return
    }
    log.Println(res.Msg.GetGreeting())
}
