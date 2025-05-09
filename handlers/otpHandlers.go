package handlers

import (
	"EdHub-BE/configs"
	"EdHub-BE/models"
	"EdHub-BE/utils"
	"encoding/json"
	"log"
	"net/http"
)

func SendOTP(w http.ResponseWriter, r *http.Request) {
	var otp models.OtpService
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	db := configs.ConnectDB()
	defer db.Close()

	generatedOtp := utils.GenerateNumericOTP(6)
	err = utils.SendEmailOTP(otp.Email, generatedOtp)
	if err != nil {
		log.Println("Failed to send OTP:", err)
		utils.Response(w, http.StatusInternalServerError, "Failed to send the verification code")
		return
	}

	_, err = db.Exec(`INSERT INTO otp (email, otp) 
                  VALUES ($1, $2) 
                  ON CONFLICT (email) 
                  DO UPDATE SET otp = EXCLUDED.otp, created_at = CURRENT_TIMESTAMP`, otp.Email, generatedOtp)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, "Failed to Save the OTP")
		return
	}

	utils.Response(w, http.StatusOK, "Code has been created and sent")
}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req models.OtpVerify
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	db := configs.ConnectDB()
	defer db.Close()

	if err := utils.VerifyOTP(db, req.Email, req.Otp); err != nil {
		if err.Error() == "user not found" {
			utils.Response(w, http.StatusUnauthorized, "User Not Found")
		} else {
			utils.Response(w, http.StatusUnauthorized, "Incorrect OTP")
		}
		return
	}

	utils.Response(w, http.StatusOK, "OTP Verified Successfully")
}
