package main

import (
	"context"
	"log"

	pb "github.com/maryoma/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	//creds,err := credentials.NewClientTLSFromCert()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %v \n", err)
	}
	defer conn.Close()
	c := pb.NewBookServiceClient(conn)
	doBook(c)

}
func doBook(c pb.BookServiceClient) {
	log.Println("The function was invoked")
	res, err := c.Book(context.Background(), &pb.BookRequest{
		Name:   "Mariam's book",
		Author: "Maryomaa",
		Id:     5,
	})
	if err != nil {
		log.Fatalf("couldn't %v", err)
	}
	log.Printf("Suceeded %s", res.BookDetails)
}
