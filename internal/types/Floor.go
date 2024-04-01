package types

type Floor struct {
	ID          string `json:"id"`
	BuildingID  string `json:"buildingID"`
	FloorNumber int    `json:"floorNumber"`
}
type FloorRequest struct {
	BuildingID  string `json:"buildingID"`
	FloorNumber int    `json:"floorNumber"`
}
type FloorResponse struct {
	ID string `json:"id"`
}
