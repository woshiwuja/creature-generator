package queries

import (
	"anarchy2036/models"

	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB)(Users []models.User){
	
	db.Raw("SELECT * FROM users").Scan(&Users)

	return Users
}
