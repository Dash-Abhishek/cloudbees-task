package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Dash-Abhishek/cloudbees-task/booking"
	"google.golang.org/grpc"
)

var bookingReqs = []*booking.BookingRequest{
	{
		From: "BLR",
		To:   "DEL",
		User: &booking.User{
			FirstName: "Abhishek",
			LastName:  "Dash",
			Email:     "abhishek@gmail.com",
		},
		SelectedSeat: &booking.Seat{Section: "A", SeatNumber: 1},
	},
	{
		From: "BLR",
		To:   "HYD",
		User: &booking.User{
			FirstName: "Jhon",
			LastName:  "Doe",
			Email:     "jhon@gmail.com",
		},
		SelectedSeat: &booking.Seat{Section: "A", SeatNumber: 5},
	},
	{
		From: "ROU",
		To:   "BLR",
		User: &booking.User{
			FirstName: "dev",
			LastName:  "dash",
			Email:     "dash@gmail.com",
		},
		SelectedSeat: &booking.Seat{Section: "B", SeatNumber: 3},
	},
}

func main() {

	CreateBookings()

	for _, req := range bookingReqs {
		FetchReceipt(req.User.Email)
	}

	ListBookings("A")
	ListBookings("B")

	UpdateBooking()

	CancelBooking()
}

func CreateBookings() {
	var wg sync.WaitGroup

	for _, req := range bookingReqs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			conn, err := getConnection()
			if err != nil {
				log.Fatalf("failed to not connect: %v", err)
			}
			defer conn.Close()

			ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
			resp, err := booking.NewBookingClient(conn).Create(ctx, req)

			if err != nil {
				log.Println("Error while creating booking: ", err)
			} else {
				fmt.Println("Booking created: ", resp)
			}

		}()

	}
	wg.Wait()
}

func UpdateBooking() {

	conn, err := getConnection()
	if err != nil {
		log.Fatalf("failed to not connect: %v", err)
	}
	defer conn.Close()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	resp, err := booking.NewBookingClient(conn).UpdateSeat(ctx, &booking.UpdateSeatRequest{
		UserEmail: &booking.UserEmail{Value: "abhishek@gmail.com"},
		Seat:      &booking.Seat{Section: "B", SeatNumber: 8},
	})

	if err != nil {
		log.Println("Error while updating booking: ", err)
	} else {
		fmt.Println("Booking updated: ", resp)
	}

}

func CancelBooking() {

	conn, err := getConnection()
	if err != nil {
		log.Fatalf("failed to not connect: %v", err)
	}
	defer conn.Close()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	_, err = booking.NewBookingClient(conn).Cancel(ctx, &booking.UserEmail{Value: "dash@gmail.com"})
	if err != nil {
		log.Println("Error while cancelling booking: ", err)
	} else {
		fmt.Println("Booking cancelled successfully")
	}
}

func FetchReceipt(userEmail string) {

	conn, err := getConnection()
	if err != nil {
		log.Fatalf("failed to not connect: %v", err)
	}
	defer conn.Close()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	resp, err := booking.NewBookingClient(conn).GetReceipt(ctx, &booking.UserEmail{Value: userEmail})

	if err != nil {
		log.Println("Error while fetching receipt: ", err)
	} else {
		fmt.Println("Receipt: ", resp)
	}
}

func ListBookings(section string) {
	conn, err := getConnection()
	if err != nil {
		log.Fatalf("failed to not connect: %v", err)
	}
	defer conn.Close()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	res, err := booking.NewBookingClient(conn).GetSeatAllocations(ctx, &booking.Section{Name: section})

	if err != nil {
		log.Fatalf("Error while fetching bookings: %v", err)
	}

	fmt.Printf("Bookings in Section %s : %v\n", section, res)
}

func getConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn, nil
}
