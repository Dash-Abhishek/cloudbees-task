package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/Dash-Abhishek/cloudbees-task/booking"
	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("failed to create listener", err)
	}
	svcRegistrar := grpc.NewServer()

	bookingSvc := booking.TrainBookingSvc{}
	booking.RegisterBookingServer(svcRegistrar, &bookingSvc)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		if err := svcRegistrar.Serve(listener); err != nil {
			log.Fatal("failed to boot server")
		}
	}()
	fmt.Println("Service live on port:9000")
	wg.Wait()

}
