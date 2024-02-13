package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NAME 		string
	PASSWORD 	string
	EMAIL    	string 
	ROLE      	string    
}

type UserResponse struct {
	ID        	uint      	`json:"id"`
	CREATEAT 	time.Time 	`json:"createdAt"`
	UPDATEDAT 	time.Time   `json:"updatedAt"`
	DELETEDAT 	time.Time 	`json:"deteledAt,omitempty"`
	NAME     	string 		`json:"name"`
	PASSWORD 	string 		`json:"password"`
	EMAIL    	string 		`json:"email"`
	ROLE      	string      `json:"role"`
}