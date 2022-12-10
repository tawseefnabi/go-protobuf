package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	elliot := &Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			Youtube: 250,
			Twitter: 140,
		},
	}

	// data, err proto.Marshal(elliot)
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshalling error", err)
	}
	fmt.Println("Marshalling data", data)

	// let's go the other way and unmarshal
	// our protocol buffer into an object we can modify
	// and use

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshalling error", err)
	}
	fmt.Println("Unmarshalling data Name:", newElliot.GetName())
	fmt.Println("Unmarshalling data Age:", newElliot.GetAge())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())

}
