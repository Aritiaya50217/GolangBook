package main

import (
	"fmt"
	"network-programming/part04service_architecture/chapter01data_serialization/01serializing_objects/03protocol_buffers/userpb"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Object
	user := &userpb.User{
		Id:    1,
		Name:  "test proto",
		Email: "test@example.com",
	}

	fmt.Println(user)

	// Serialize
	data, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Serialized: %x\n", data)

	// Deserialize
	newUser := &userpb.User{}

	if err = proto.Unmarshal(data, newUser); err != nil {
		panic(err)
	}

	fmt.Printf("Deserialized: %+v\n", newUser)
}
