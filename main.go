package main

import (
	"context"
	"fmt"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := grpcreflect.NewClient(context.Background(), grpc_reflection_v1alpha.NewServerReflectionClient(conn))
	svc, err := client.ListServices()
	if err != nil {
		panic(err)
	}
	fmt.Println("SERVICES:")
	for _, s := range svc {
		if s == "grpc.reflection.v1alpha.ServerReflection" {
			continue
		}

		fmt.Printf("  %s\n", s)
		fmt.Println("    METHODS:")
		d, err := client.ResolveService(s)
		if err != nil {
			panic(err)
		}
		for _, m := range d.GetMethods() {
			fmt.Printf("      * %s\n", m.GetName())
			fmt.Printf("        +-- IN:  %s\n", m.GetInputType().GetName())
			fmt.Printf("        +-- OUT: %s\n", m.GetOutputType().GetName())
		}
	}
}
