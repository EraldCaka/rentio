package types

const (
	SingleRoomApartment = 0
	StudioApartment     = 1
	TwoRoomApartment    = 2
)

type Room struct {
	ID       string `json:"id"`
	RoomType int    `json:"roomType"`
	FloorID  string `json:"floorID"`
	RoomSize int    `json:"roomSize"` // meter square (m^2)
}
type RoomRequest struct {
	RoomType int `json:"roomType"`
	RoomSize int `json:"roomSize"`
}
type RoomResponse struct {
	ActiveContract string `json:"activeContract"`
}
