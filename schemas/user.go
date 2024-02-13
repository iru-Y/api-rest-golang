package schemas

import (
	"time"
)

type User struct {
	ID 			string
	NAME 		string
	PASSWORD 	string
	EMAIL    	string 
	ROLE      	string    
}

type UserResponse struct {
	ID        	string     	`json:"id"`
	CREATEAT 	time.Time 	`json:"createdAt"`
	UPDATEDAT 	time.Time   `json:"updatedAt"`
	DELETEDAT 	time.Time 	`json:"deteledAt,omitempty"`
	NAME     	string 		`json:"name"`
	PASSWORD 	string 		`json:"password"`
	EMAIL    	string 		`json:"email"`
	ROLE      	string      `json:"role"`
}