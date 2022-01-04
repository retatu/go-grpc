package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/retatu/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect with the gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUsersStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Matheus",
		Email: "mrehbein45@gmail.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Matheus",
		Email: "mrehbein45@gmail.com",
	}
	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not recieve the msg: %v", err)
		}
		fmt.Println("Status: ", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Matheus",
			Email: "mrehbein0@gmail.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Matheus 1",
			Email: "mrehbein1@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Matheus 2",
			Email: "mrehbein2@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Matheus 3",
			Email: "mrehbein3@gmail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Matheus 4",
			Email: "mrehbein4@gmail.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Matheus 5",
			Email: "mrehbein5@gmail.com",
		},
	}
	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not recieve gRPC Response: %v", err)
	}
	fmt.Println(res)
}

func AddUsersStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUsersStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Matheus",
			Email: "mrehbein0@gmail.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Matheus 1",
			Email: "mrehbein1@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Matheus 2",
			Email: "mrehbein2@gmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Matheus 3",
			Email: "mrehbein3@gmail.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Matheus 4",
			Email: "mrehbein4@gmail.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Matheus 5",
			Email: "mrehbein5@gmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Could not recieve the msg: %v", err)
				break
			}
			fmt.Printf("Receiving %v com status %v \n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()
	<-wait
}
