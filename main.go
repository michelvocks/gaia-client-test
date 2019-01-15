package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/gaia-pipeline/protobuf"
	"google.golang.org/grpc"
)

const (
	address = "localhost:55200"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewPluginClient(conn)

	// Get all jobs
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.GetJobs(ctx, &pb.Empty{})
	if err != nil {
		panic(err)
	}
	for {
		job, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("Got following job: %#v\n", job)
	}
}
