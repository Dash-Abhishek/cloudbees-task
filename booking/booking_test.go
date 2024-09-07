package booking

import (
	"context"
	"testing"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

var bookingTestTable = []struct {
	BookingRequest   *BookingRequest
	ExpectedResponse *Receipt
	ExpectedError    error
}{
	{
		BookingRequest: &BookingRequest{
			From:         "BLR",
			To:           "DEL",
			SelectedSeat: &Seat{Section: "A", SeatNumber: 1},
			User:         &User{FirstName: "John", LastName: "Doe", Email: "jhon@gmail.com"},
		},
		ExpectedResponse: &Receipt{
			Id:    "f678cb81-f128-4a17-a7f5-04793ae4e53a",
			From:  "BLR",
			To:    "DEL",
			Seat:  &Seat{Section: "A", SeatNumber: 1},
			User:  &User{FirstName: "John", LastName: "Doe", Email: "jhon@gmail.com"},
			Price: 20,
		},
		ExpectedError: nil,
	},
	{
		BookingRequest: &BookingRequest{
			From:         "BLR",
			To:           "DEL",
			SelectedSeat: &Seat{Section: "B", SeatNumber: 2},
			User:         &User{FirstName: "John", LastName: "Doe", Email: "jhon@gmail.com"},
		},
		ExpectedResponse: nil,
		ExpectedError:    status.Errorf(codes.AlreadyExists, ErrDuplicateBooking.Error()),
	},
	{
		BookingRequest: &BookingRequest{
			From:         "BLR",
			To:           "HYD",
			SelectedSeat: &Seat{Section: "B", SeatNumber: 2},
			User:         &User{FirstName: "abhishek", LastName: "Dash", Email: "abhishek@gmail.com"},
		},
		ExpectedResponse: &Receipt{
			Id:    "f678cb81-f128-4a17-a7f5-04793ae4egfa",
			From:  "BLR",
			To:    "HYD",
			User:  &User{FirstName: "abhishek", LastName: "Dash", Email: "abhishek@gmail.com"},
			Price: 20,
			Seat:  &Seat{Section: "B", SeatNumber: 2},
		},
		ExpectedError: nil,
	},
	{
		BookingRequest: &BookingRequest{
			From:         "BLR",
			To:           "HYD",
			SelectedSeat: &Seat{Section: "C", SeatNumber: 2},
			User:         &User{FirstName: "test", LastName: "user", Email: "test@gmail.com"},
		},
		ExpectedResponse: nil,
		ExpectedError:    status.Errorf(codes.InvalidArgument, ErrInvalidSeat.Error()),
	},
}

var receiptTestTable = []struct {
	UserEmail        *UserEmail
	ExpectedResponse *Receipt
	ExpectedError    error
}{
	{UserEmail: &UserEmail{Value: "abhishek@gmail.com"},
		ExpectedResponse: &Receipt{
			Id:    "f678cb81-f128-4a17-a7f5-04793ae4egfa",
			From:  "BLR",
			To:    "HYD",
			User:  &User{FirstName: "abhishek", LastName: "Dash", Email: "abhishek@gmail.com"},
			Price: 20,
			Seat:  &Seat{Section: "B", SeatNumber: 2},
		},
		ExpectedError: nil,
	},
	{UserEmail: &UserEmail{Value: "test@gmail.com"},
		ExpectedResponse: nil,
		ExpectedError:    status.Errorf(codes.NotFound, ErrBookingNotFound.Error()),
	},
}

func TestNewBooking(t *testing.T) {
	bookingSvc := TrainBookingSvc{}

	t.Run("Create Bookings", func(t *testing.T) {

		for _, test := range bookingTestTable {

			res, err := bookingSvc.Create(context.Background(), test.BookingRequest)

			if err != nil && err.Error() != test.ExpectedError.Error() {
				t.Errorf("expected error %v, got %v", test.ExpectedError, err)
			}
			if res != nil && res.Id == "" {
				t.Errorf("expected booking id, got empty string")

			}

		}

		t.Run("Get Receipts", func(t *testing.T) {

			for _, test := range receiptTestTable {
				res, err := bookingSvc.GetReceipt(context.Background(), test.UserEmail)

				if err != nil && err.Error() != test.ExpectedError.Error() {
					t.Errorf("expected error %v, got %v", test.ExpectedError, err)
				}
				if res != nil && res.Id == "" {
					t.Errorf("expected booking id, got empty string")
				}
			}
		})

		t.Run("List booking of sections", func(t *testing.T) {

			secAres, err1 := bookingSvc.GetSeatAllocations(context.Background(), &Section{Name: "A"})
			secBres, err2 := bookingSvc.GetSeatAllocations(context.Background(), &Section{Name: "B"})
			_, err := bookingSvc.GetSeatAllocations(context.Background(), &Section{Name: "C"})

			if err != nil && err.Error() != status.Errorf(codes.InvalidArgument, ErrInvalidSection.Error()).Error() {
				t.Errorf("expected error %v, got %v", status.Errorf(codes.InvalidArgument, ErrInvalidSection.Error()), err)
			}

			if err1 != nil {
				t.Errorf("should not have returned error for case1, got %v", err1)
			}

			if err2 != nil {
				t.Errorf("should not have returned error for case2, got %v", err2)
			}

			if len(secAres.Allocations) != 1 {
				t.Errorf("should have returned 1 booking Section A, got %v", len(secAres.Allocations))
			}
			if len(secBres.Allocations) != 1 {
				t.Errorf("should have returned 1 booking for Section B, got %v", len(secBres.Allocations))
			}

		})

	})

}
