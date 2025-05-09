package main

import (
	"EdHub-BE/handlers"
	"EdHub-BE/middleware"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", middleware.CORS_MIDDLEWARE(handlers.SignUp))
	http.HandleFunc("/login", middleware.CORS_MIDDLEWARE(handlers.SignIn))
	http.HandleFunc("/send-otp", middleware.CORS_MIDDLEWARE(handlers.SendOTP))
	http.HandleFunc("/verify", middleware.CORS_MIDDLEWARE(handlers.VerifyOTP))

	log.Println("Server is Running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
