package main

import (
	"fmt"
	"go-grpc-demo/service"

	"google.golang.org/protobuf/proto"
)

func main2() {
	user := &service.User{
		Username: "lizy",
		Age:      25,
	}

	data, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	userRecv := &service.User{}

	if err = proto.Unmarshal(data, userRecv); err != nil {
		panic(err)
	}

	fmt.Println(user.String())

}
