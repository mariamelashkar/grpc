/*
package main

import (
	"context"
	"log"

	pb "github.com/maryoma/proto"
)

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
*/