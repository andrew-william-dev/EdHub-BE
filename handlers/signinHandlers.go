package handlers

import (
	"EdHub-BE/configs"
	"EdHub-BE/models"
	"EdHub-BE/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateJWT (email string) (string, int64, error) {
	expiryTime := time.Now().Add(60 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"email": email,
		"exp": expiryTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString,err := token.SignedString([]byte("your-secret-key"))
	if(err != nil) {
		return "",0,err;
	}

	return tokenString, expiryTime, nil
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, "Invalid Request Body")
		return 
	}

	db := configs.ConnectDB()
	defer db.Close()

	var hashPwd string
	err = db.QueryRow("SELECT password FROM users WHERE userName=$1",user.Username).Scan(&hashPwd)
	if err != nil {
		utils.Response(w, http.StatusUnauthorized, "User Not Found")
		return
	}

	if(!utils.CheckPassword(hashPwd, user.Password)){
		utils.Response(w, http.StatusUnauthorized, "Incorrect Credentials")
		return
	}

	token, expiry, err := generateJWT(user.Email);
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, "Failed to generate JWT")
		return
	}

	w.Header().Set("Conetnt-Type","application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token" : token,
		"expiry": expiry,
	})
}
