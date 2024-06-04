package models

type Occupancies struct {
	RoomNr     int `json:"rooms"`
	AdultNr    int `json:"adults"`
	ChildrenNr int `json:"children"`
}

type HotelResponse struct {
	HotelId  int    `json:"hotelId"`
	Currency string `json:"currency"`
	Price    string `json:"price"`
}
