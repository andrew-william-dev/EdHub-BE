package handlers

import (
	"EdHub-BE/configs"
	"EdHub-BE/models"
	"EdHub-BE/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lib/pq"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, "Invalid Request Body")
		return 
	}

	db := configs.ConnectDB()
	defer db.Close()

	hashPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, "Hashing of Password is Failing")
		return
	}

	_, err = db.Exec("INSERT INTO users (userName, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, hashPwd)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" && strings.Contains(pqErr.Message, "email") {
				utils.Response(w, http.StatusUnauthorized, "Email already exists")
				return
			}
		}
		utils.Response(w, http.StatusInternalServerError, "DB Operation has been failed")
		return
	}

	utils.Response(w, http.StatusCreated, "User has been created")
}

