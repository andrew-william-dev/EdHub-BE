package main

import (
	"EdHub-BE/handlers"
	"EdHub-BE/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/register", middleware.CORS_MIDDLEWARE(handlers.SignUp))
	http.HandleFunc("/login", middleware.CORS_MIDDLEWARE(handlers.SignIn))
	http.HandleFunc("/send-otp", middleware.CORS_MIDDLEWARE(handlers.SendOTP))
	http.HandleFunc("/verify", middleware.CORS_MIDDLEWARE(handlers.VerifyOTP))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is Running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
