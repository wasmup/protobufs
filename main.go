package main

import (
	"fmt"
	"log"

	pb "github.com/wasmup/protobufs"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	bk := &pb.AddressBook{}
	bk.People = append(bk.People, &p)
	in, err := proto.Marshal(bk)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(in)
	fmt.Println("Length is", len(in))

	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(book.People)
}
