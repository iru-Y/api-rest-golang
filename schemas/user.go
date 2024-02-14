package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 			uuid.Domain `json:"id"`
	NAME 		string 		`json:"name"`
	PASSWORD 	string		`json:"password"`
	EMAIL    	string 		`json:"email"`
	ROLE      	string    	`json:"role"`
}

type UserResponse struct {
	ID        	uuid.UUID   `json:"id"`
	CREATEAT 	time.Time 	`json:"createdAt"`
	UPDATEDAT 	time.Time   `json:"updatedAt"`
	DELETEDAT 	time.Time 	`json:"deteledAt,omitempty"`
	NAME     	string 		`json:"name"`
	PASSWORD 	string 		`json:"password"`
	EMAIL    	string 		`json:"email"`
	ROLE      	string      `json:"role"`
}