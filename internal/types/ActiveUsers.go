package types

type ActiveUsers struct {
	ID       string `json:"id"`
	Role     int    `json:"role"`
	JwtToken string `json:"jwtToken"`
}

type ActiveUsersRequest struct {
	Role     int    `json:"role"`
	JwtToken string `json:"jwtToken"`
}

type ActiveUsersResponse struct {
	ID string `json:"id"`
}
type ActiveUsersDeleteRequest struct {
	ID string `json:"id"`
}
