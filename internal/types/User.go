package types

import "fmt"

const (
	minUsernameLen = 3
	minPasswordLen = 7
)
const (
	AdminRole = 0
	UserRole  = 1
	Customer  = 2
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserResponse struct {
	ID string `json:"id"`
}

func (params *UserCreateRequest) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Username) < minUsernameLen {
		errors["username"] = fmt.Sprintf("username length should be at least %d characters", minUsernameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	return errors
}
