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
var SectionIndexMapping = map[string]int{"A": 0, "B": 1}
var ErrSeatUnavailable = errors.New("Seat not available")
var ErrBookingNotFound = errors.New("Booking not found")
var ErrInvalidSeat = errors.New("Invalid seat")
var ErrDuplicateBooking = errors.New("Booking already exists for user")
var ErrInvalidSection = errors.New("Invalid section")

const (
	MaxSeatNumber = 9
)

type Booking struct {
	Id    string
	User  *User
	From  string
	To    string
	Seat  *Seat
	Price float32
}

type seats struct {
	seatGrid [2][10]string
	mutex    *sync.RWMutex
}

type store struct {
	bookings map[string]Booking
	mutex    *sync.RWMutex
}

type TrainBookingSvc struct {
	UnimplementedBookingServer
}

func init() {
	storage = store{
		bookings: make(map[string]Booking),
		mutex:    &sync.RWMutex{},
	}

	availableSeats = seats{
		seatGrid: [2][10]string{},
		mutex:    &sync.RWMutex{},
	}
}

// Create a new booking
func (svc *TrainBookingSvc) Create(ctx context.Context, req *BookingRequest) (*Receipt, error) {

	fmt.Println("New Booking request received", req)
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	if _, found := storage.bookings[req.User.Email]; found {
		return nil, status.Errorf(codes.AlreadyExists, ErrDuplicateBooking.Error())
	}

	// Generating bookingID
	bookingId := uuid.New().String()
	newBooking := Booking{
		Id:    bookingId,
		User:  req.User,
		From:  req.From,
		To:    req.To,
		Price: 20.0,
		Seat:  req.SelectedSeat,
	}

	// Allocating seat
	err := allocateSeat(req.SelectedSeat, newBooking.User.Email)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	// Persisting booking
	storage.bookings[newBooking.User.Email] = newBooking
	fmt.Printf("Booking successful: %s for user: %s\n", newBooking.Id, newBooking.User.Email)
	return newBooking.GenerateReceipt(), nil

}

// Get receipt for a booking
func (svc *TrainBookingSvc) GetReceipt(ctx context.Context, userEmail *UserEmail) (*Receipt, error) {

	fmt.Println("fetching receipt for user:", userEmail)
	res, found := storage.bookings[userEmail.Value]
	if !found {
		fmt.Println("Booking not found for ", userEmail.Value)
		return nil, status.Errorf(codes.NotFound, ErrBookingNotFound.Error())
	}
	return res.GenerateReceipt(), nil

}

// Update seat for a booking
// The booking is specified by the booking id
// The method returns a receipt if the seat is successfully updated
// Otherwise it returns an error
func (svc *TrainBookingSvc) UpdateSeat(ctx context.Context, req *UpdateSeatRequest) (*Receipt, error) {

	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	booking, found := storage.bookings[req.UserEmail.Value]
	if !found {
		return nil, status.Errorf(codes.NotFound, ErrBookingNotFound.Error())
	}

	err := allocateSeat(req.Seat, req.UserEmail.Value)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	deAllocateSeat(booking.Seat)

	booking.Seat = req.Seat
	storage.bookings[req.UserEmail.Value] = booking

	return booking.GenerateReceipt(), nil

}

// Cancel a booking
// This method is used to cancel a booking
// The booking is specified by the booking id
// The method returns an empty response if the booking is successfully cancelled
// The method returns an error if the booking is not found
func (svc *TrainBookingSvc) Cancel(ctx context.Context, userEmail *UserEmail) (*Empty, error) {

	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	// Check if booking exists
	booking, found := storage.bookings[userEmail.Value]
	if !found {
		return nil, status.Errorf(codes.NotFound, ErrBookingNotFound.Error())
	}

	delete(storage.bookings, userEmail.Value)
	deAllocateSeat(booking.Seat)

	return &Empty{}, nil
}

// Get seat allocations for a section
// This method is used to get all the seat allocations for a given section
// The section is specified by the section name
func (svc *TrainBookingSvc) GetSeatAllocations(ctx context.Context, section *Section) (*AllocationList, error) {

	fmt.Println("fetching seat allocations for sec:", section)
	results := []*Allocation{}

	sectionIndex, found := SectionIndexMapping[section.Name]
	if !found {
		return nil, status.Errorf(codes.InvalidArgument, ErrInvalidSection.Error())
	}

	for _, bookingId := range availableSeats.seatGrid[sectionIndex] {

		if bookingId == "" {
			continue
		}

		bookingDetails := storage.bookings[bookingId]
		results = append(results, &Allocation{
			BookingId: bookingDetails.Id,
			User:      bookingDetails.User,
			Seat:      bookingDetails.Seat,
		})
	}

	return &AllocationList{Allocations: results}, nil

}

// Allocate a seat
// The method returns an error if the section is invalid
// The method returns an error if the seat is invalid
// The method returns an error if the seat is already allocated
// The method returns nil if the seat is successfully allocated
func allocateSeat(selectedSeat *Seat, userId string) error {

	availableSeats.mutex.Lock()
	defer availableSeats.mutex.Unlock()
	if selectedSeat.Section != "A" && selectedSeat.Section != "B" {
		return ErrInvalidSection
	}

	if selectedSeat.SeatNumber < 0 || selectedSeat.SeatNumber > MaxSeatNumber {
		return ErrInvalidSeat
	}

	sectionIndex := SectionIndexMapping[selectedSeat.Section]

	if availableSeats.seatGrid[sectionIndex][selectedSeat.SeatNumber] != "" {
		return ErrSeatUnavailable
	}

	availableSeats.seatGrid[sectionIndex][selectedSeat.SeatNumber] = userId

	return nil

}

// Deallocate a seat
func deAllocateSeat(seat *Seat) {

	availableSeats.mutex.Lock()
	defer availableSeats.mutex.Unlock()

	sectionIndex := SectionIndexMapping[seat.Section]

	availableSeats.seatGrid[sectionIndex][seat.SeatNumber] = ""

}

func (b Booking) GenerateReceipt() *Receipt {
	return &Receipt{
		Id:    b.Id,
		User:  b.User,
		From:  b.From,
		To:    b.To,
		Price: b.Price,
		Seat:  b.Seat,
	}
}
