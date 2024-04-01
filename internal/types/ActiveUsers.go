package types

import "time"

type ActiveUsers struct {
	ID       string `json:"id"`
	Role     int    `json:"role"`
	JwtToken string `json:"jwtToken"`
}

type ActiveUsersRequest struct {
	Role       int       `json:"role"`
	JwtToken   string    `json:"jwtToken"`
	Username   string    `json:"username"`
	ExpireTime time.Time `json:"expireTime"`
}

type ActiveUsersResponse struct {
	ID string `json:"id"`
}
type ActiveUsersDeleteRequest struct {
	ID string `json:"id"`
}
