package models

type Occupancies struct {
	RoomNr     int `json:"rooms"`
	AdultNr    int `json:"adults"`
	ChildrenNr int `json:"children"`
}
