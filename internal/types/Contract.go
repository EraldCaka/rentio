package types

import "time"

const (
	Active   = 0
	Inactive = 1
)

type Contract struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"roomID"`
	UserID    string    `json:"userID"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"EndDate"`
	Rent      int       `json:"rent"`
	Status    int       `json:"status"`
}
type ContractRequest struct {
	RoomID    string    `json:"roomID"`
	UserID    string    `json:"userID"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"EndDate"`
	Rent      int       `json:"rent"`
	Status    int       `json:"status"`
}
type ContractResponse struct {
	ID string `json:"id"`
}
