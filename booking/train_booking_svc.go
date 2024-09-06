package booking

import (
	context "context"
	"errors"
	"fmt"
	sync "sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var storage store
var availableSeats seats
var ErrSeatUnavailable = errors.New("Seat not available")
var ErrBookingNotFound = errors.New("Booking not found")

type Booking struct {
	Id    string
	User  *User
	From  string
	To    string
	Seat  *Seat
	Price float32
}

type seats struct {
	secA      []bool
	secB      []bool
	secAIndex int
	secBIndex int
	mutex     *sync.Mutex
}

type store struct {
	bookings map[string]Booking
	mutex    *sync.Mutex
}

type TrainBookingSvc struct {
	UnimplementedBookingServer
}

func init() {
	storage = store{
		bookings: make(map[string]Booking),
		mutex:    &sync.Mutex{},
	}

	availableSeats = seats{
		secA:      make([]bool, 10),
		secB:      make([]bool, 10),
		secAIndex: 0,
		secBIndex: 0,
		mutex:     &sync.Mutex{},
	}
}

// Create a new booking
func (svc *TrainBookingSvc) Create(ctx context.Context, req *BookingRequest) (*Receipt, error) {

	fmt.Println("New Booking request received", req)

	// Allocating seat
	seat, err := AllocateSeat()
	if err != nil {
		return nil, status.Errorf(codes.ResourceExhausted, err.Error())
	}

	// Generating bookingID
	bookingId := uuid.New().String()
	booking := Booking{
		Id:    bookingId,
		User:  req.User,
		From:  req.From,
		To:    req.To,
		Price: 20.0,
		Seat:  seat,
	}

	// Persisting booking
	storage.mutex.Lock()
	storage.bookings[bookingId] = booking
	storage.mutex.Unlock()

	return booking.toReceipt(), nil

}

// Get receipt for a booking
func (svc *TrainBookingSvc) GetReceipt(ctx context.Context, bookingId *BookingId) (*Receipt, error) {

	fmt.Println("fetching receipt for bookingId:", bookingId)
	res, found := storage.bookings[bookingId.Value]
	if !found {
		fmt.Println("Booking not found for ", bookingId)
		return nil, status.Errorf(codes.NotFound, ErrBookingNotFound.Error())
	}
	return res.toReceipt(), nil

}

// TODO - Implement this method
func (svc *TrainBookingSvc) UpdateSeat(ctx context.Context, req *UpdateSeatRequest) (*Receipt, error) {

	return nil, nil
}

// TODO - Implement this method
func (svc *TrainBookingSvc) Cancel(ctx context.Context, id *BookingId) (*Empty, error) {

	return nil, nil
}

// Get seat allocations for a section
// This method is used to get all the seat allocations for a given section
// The section is specified by the section name
func (svc *TrainBookingSvc) GetSeatAllocations(ctx context.Context, section *Section) (*AllocationList, error) {

	fmt.Println("fetching seat allocations for sec:", section)
	results := []*Allocation{}

	for id, booking := range storage.bookings {
		if booking.Seat.Section == section.Name {
			results = append(results, &Allocation{
				BookingId: id,
				User:      booking.User,
				Seat:      booking.Seat,
			})
		}
	}

	return &AllocationList{Allocations: results}, nil

}

// Allocate a seat
// This method is used to allocate a seat for a booking
// The method returns a seat object if the seat is available
// Otherwise it returns an error
func AllocateSeat() (*Seat, error) {

	if availableSeats.secAIndex == 10 && availableSeats.secBIndex == 10 {
		return nil, ErrSeatUnavailable
	}

	var seatNumber int32
	var section string
	availableSeats.mutex.Lock()
	defer availableSeats.mutex.Unlock()

	if availableSeats.secAIndex < 10 {
		availableSeats.secA[availableSeats.secAIndex] = true
		availableSeats.secAIndex++
		seatNumber = int32(availableSeats.secAIndex)
		section = "A"
	} else {
		availableSeats.secB[availableSeats.secBIndex] = true
		availableSeats.secBIndex++
		seatNumber = int32(availableSeats.secBIndex)
		section = "B"
	}

	return &Seat{SeatNumber: seatNumber, Section: section}, nil

}

func (b Booking) toReceipt() *Receipt {
	return &Receipt{
		Id:    b.Id,
		User:  b.User,
		From:  b.From,
		To:    b.To,
		Price: b.Price,
		Seat:  b.Seat,
	}
}
