package schemas

import (
	"time"

)

type User struct {
	ID 			string 		`json:"id"`
	NAME 		string 		`json:"name"`
	PASSWORD 	string		`json:"-"`
	EMAIL    	string 		`json:"email"`
	ROLE      	string    	`json:"role"`
}

type UserResponse struct {
	ID        	string   	`json:"id"`
	CREATEAT 	time.Time 	`json:"createdAt"`
	UPDATEDAT 	time.Time   `json:"updatedAt"`
	DELETEDAT 	time.Time 	`json:"deteledAt,omitempty"`
	NAME     	string 		`json:"name"`
	EMAIL    	string 		`json:"email"`
	ROLE      	string      `json:"role"`
}