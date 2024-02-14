package handler

import (
	"fmt"

)

func errParamIsRequired(n, t string) error {
	return fmt.Errorf("param: %s (type: %s) is required", n, t)
}

type CreateUserRequest struct {
	NAME     	string 		`json:"name"`
	PASSWORD 	string 		`json:"password"`
	EMAIL    	string 		`json:"email"`
	ROLE      	string      `json:"role"`
}

func (r *CreateUserRequest) Validate() error {
	if r.ROLE == "" && r.EMAIL == "" && r.NAME == "" && r.PASSWORD == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.ROLE == "" {
		return errParamIsRequired("role", "string")
	}
	if r.EMAIL == "" {
		return errParamIsRequired("email", "string")
	}
	if r.NAME == "" {
		return errParamIsRequired("name", "string")
	}
	if r.PASSWORD == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}
