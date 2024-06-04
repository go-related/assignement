package services

import (
	"github.com/stretchr/testify/assert"
	"githumb/go-related/nuitteassignment/internal/models"
	"testing"
	"time"
)

type HotelRateRequest struct {
	checkinTime, checkoutTime  *time.Time
	currency, guestNationality string
	hotelIds                   []uint
	occupancies                []models.Occupancies
}

func TestHotelServiceValidations(t *testing.T) {
	testCases := []struct {
		name         string
		input        HotelRateRequest
		expectError  bool
		errorMessage string
	}{
		{
			name: "validateEmptyCheckinTime",
			input: HotelRateRequest{
				checkinTime:  nil,
				checkoutTime: getPointer(time.Now()),
			},
			expectError:  true,
			errorMessage: "invalid data for checkin-time or checkout-time",
		},
		{
			name: "validateEmptyCheckoutTime",
			input: HotelRateRequest{
				checkinTime:  getPointer(time.Now()),
				checkoutTime: nil,
			},
			expectError:  true,
			errorMessage: "invalid data for checkin-time or checkout-time",
		},
		{
			name: "validateCheckinBeforeCurrentTime",
			input: HotelRateRequest{
				checkinTime:  getPointer(time.Now().UTC().AddDate(0, 0, -1)),
				checkoutTime: getPointer(time.Now()),
			},
			expectError:  true,
			errorMessage: "invalid checkin/checkout time, should be in the future",
		},
		{
			name: "validateCheckoutBeforeCurrentTime",
			input: HotelRateRequest{
				checkoutTime: getPointer(time.Now().UTC().AddDate(0, 0, -1)),
				checkinTime:  getPointer(time.Now().UTC().AddDate(0, 0, 1)),
			},
			expectError:  true,
			errorMessage: "invalid checkin/checkout time, should be in the future",
		},
		{
			name: "validateCheckoutBeforeCheckin",
			input: HotelRateRequest{
				checkoutTime: getPointer(time.Now().UTC().AddDate(0, 0, 1)),
				checkinTime:  getPointer(time.Now().UTC().AddDate(0, 0, 2)),
			},
			expectError:  true,
			errorMessage: "invalid checkin-time should be in before checkout-time",
		},
		{
			name: "validateEmptyHotelIds",
			input: HotelRateRequest{
				checkoutTime: getPointer(time.Now().UTC().AddDate(0, 0, 2)),
				checkinTime:  getPointer(time.Now().UTC().AddDate(0, 0, 1)),
				hotelIds:     []uint{},
			},
			expectError:  true,
			errorMessage: "invalid hotel-ids",
		},
		{
			name: "validateEmptyOccupancies",
			input: HotelRateRequest{
				checkoutTime: getPointer(time.Now().UTC().AddDate(0, 0, 2)),
				checkinTime:  getPointer(time.Now().UTC().AddDate(0, 0, 1)),
				hotelIds:     []uint{1},
				occupancies:  []models.Occupancies{},
			},
			expectError:  true,
			errorMessage: "invalid occupancies",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// setup
			var customError *ServiceError
			hotelService, err := NewHotelService("", "", "")
			assert.NoError(t, err)

			// arrange
			_, _, _, err = hotelService.CheckHotelRate(test.input.checkinTime, test.input.checkoutTime, test.input.currency, test.input.guestNationality, test.input.hotelIds, test.input.occupancies)

			// assert
			assert.Error(t, err)
			assert.ErrorAs(t, err, &customError)
			if test.expectError {
				assert.ErrorContains(t, err, test.errorMessage)
			}
		})
	}
}

func getPointer[T any](currentData T) *T {
	return &currentData
}
