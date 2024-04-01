package types

type Building struct {
	ID       string `json:"id"`
	ClientID string `json:"clientID"`
	Location string `json:"location"`
}
type BuildingRequest struct {
	ClientID string `json:"clientID"`
	Location string `json:"location"`
}
type BuildingResponse struct {
	ID string `json:"id"`
}
