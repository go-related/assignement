package models

type StayDTO struct {
	CheckIn  string `json:"checkIn"`
	CheckOut string `json:"checkOut"`
}

type OccupancyDTO struct {
	Rooms    int `json:"rooms"`
	Adults   int `json:"adults"`
	Children int `json:"children"`
}

type HotelDTOs struct {
	Hotel []uint `json:"hotel"`
}

type HotelRequestDTO struct {
	Stay        StayDTO        `json:"stay"`
	Occupancies []OccupancyDTO `json:"occupancies"`
	Hotels      HotelDTOs      `json:"hotels"`
}
