package main

import (
	"log"
	"runtime"

	"github.com/nats-io/go-nats"

	pb "nats_go_demo/order"

	"github.com/golang/protobuf/proto"
)

func main() {
	natsConn, _ := nats.Connect(nats.DefaultURL)
	log.Println("connect to the ", nats.DefaultURL)

	natsConn.Subscribe("Discovery.OrderService", func(m *nats.Msg) {
		orderServDiscovery := &pb.ServiceDiscovery{OrderServiceUri: "http://localhost:80/OrderServiceUri"}
		data, err := proto.Marshal(orderServDiscovery)
		if err == nil {
			natsConn.Publish(m.Reply, data)
		}
	})

	natsConn.Subscribe("Discovery.Email", func(m *nats.Msg) {
		orderServDiscovery := &pb.ServiceDiscovery{OrderServiceUri: "http://localhost:80/EmailUri"}
		data, err := proto.Marshal(orderServDiscovery)
		if err == nil {
			natsConn.Publish(m.Reply, data)
		}
	})
	runtime.Goexit()
}
