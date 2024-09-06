package booking

import (
	context "context"
	"fmt"
)

type TrainBookingSvc struct {
	UnimplementedBookingServer
}

func (svc *TrainBookingSvc) Create(ctx context.Context, req *BookingRequest) (*Receipt, error) {
	fmt.Println(req)
	return nil, nil
}

func (svc *TrainBookingSvc) GetReceipt(ctx context.Context, id *BookingId) (*Receipt, error) {

	return nil, nil
}
