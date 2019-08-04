package controller

import (
	"net/http"

	"gotutorials/app/model"

	"github.com/jinzhu/gorm"

	"encoding/json"
)

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.UserObj{}
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}
