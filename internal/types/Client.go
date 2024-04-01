package types

import "fmt"

type Client struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"Password"`
}

type ClientRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

type ClientRequest struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

type ClientResponse struct {
	ID string `json:"id"`
}

func (params *ClientRegisterRequest) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Username) < minUsernameLen {
		errors["username"] = fmt.Sprintf("username length should be at least %d characters", minUsernameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	return errors
}
