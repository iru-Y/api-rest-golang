package schemas

import (
	"gorm.io/gorm"
)

type OPening struct {
	gorm.Model
	NAME 	 			string
	PASSWORD 			string
	EMAIL				string
	ROLE 	 			string

}
