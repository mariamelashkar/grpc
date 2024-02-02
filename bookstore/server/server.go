package main

import (
	// "fmt"
	"log"
	"context"
	"net"
	 pb "github.com/maryoma/proto"
	"google.golang.org/grpc"
	//"google.golang.org/protobuf/proto"
)

var adrr string = "0.0.0.0:50051"

type Server struct {
	pb.BookServiceServer 
}
/*
books := &pb.BookRequest{
	Name:   "Mariam",
	Author: "Mariam",
	Id:     5,
}
*/
func main() {
	lis,err := net.Listen("tcp",adrr)
	if err != nil {
        log.Fatalf("Failed to Listen on %v\n",err)
	}
	log.Printf("Listening on %v\n",adrr)
	/*
	data, err := proto.Marshal(books)
	if err != nil {
		log.Fatal("Marshaling error", err)
	}
	fmt.Println(data)

	newbooks := &proto3.Book{}
	err = proto.Unmarshal(data, newbooks)
	if err != nil {
		log.Fatal("unmarshaling error", err)
	}
	fmt.Println(newbooks.GetAuthor())
	fmt.Println(newbooks.GetId())
	*/
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s,&Server{})
		if err = s.Serve(lis) ; err != nil {
		log.Fatalf("Failed to Serve on %v \n",err)
	} 


}
func (s *Server) Book(ctx context.Context,in *pb.BookRequest) (*pb.BookResponse, error){
	log.Printf("Book function invoked with %v \n",in)
	return &pb.BookResponse{
	BookDetails: "The book name is" + in.Name,
	},nil
	}
