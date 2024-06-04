package models

type HotelRequestRoomGuests struct {
	AdultCount int `json:"adultCount"`
	ChildCount int `json:"childCount"`
}

type HotelRequestDTO struct {
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
	HotelCodes   []int  `json:"hotelCodes"`
	RoomGuests   struct {
		RoomGuests []struct {
			AdultCount int `json:"adultCount"`
			ChildCount int `json:"childCount"`
		} `json:"roomGuests"`
	} `json:"roomGuests"`
	GuestNationality string `json:"guestNationality"`
	Currency         string `json:"currency"`
	LanguageCode     string `json:"languageCode"`
	Timeout          string `json:"timeout"`
}
