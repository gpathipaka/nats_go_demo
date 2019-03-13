package main

import (
	"log"
	pb "nats_go_demo/order"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

func main() {
	//Create NATS Server Connection
	natsConn, _ := nats.Connect(nats.DefaultURL)
	defer natsConn.Close()
	log.Println("Connected to the nats server....", nats.DefaultURL)
	msg, err := natsConn.Request("Discovery.OrderService", nil, 1000*time.Microsecond)
	log.Println(err)
	if err == nil && msg != nil {
		log.Println("*************")
		orderServiceDiscovery := pb.ServiceDiscovery{}
		err := proto.Unmarshal(msg.Data, &orderServiceDiscovery)
		if err != nil {
			log.Println("error ", err)
		}
		log.Println("Message Received ", orderServiceDiscovery.OrderServiceUri)
	}

	msg, err = natsConn.Request("Discovery.Email", nil, 1000*time.Microsecond)
	if err == nil && msg != nil {
		log.Println("*************")
		orderServiceDiscovery := pb.ServiceDiscovery{}
		err := proto.Unmarshal(msg.Data, &orderServiceDiscovery)
		if err != nil {
			log.Println("error ", err)
		}
		log.Println("Email Service Message Received ", orderServiceDiscovery.OrderServiceUri)
	}
	log.Println("Exiting client....")
}
