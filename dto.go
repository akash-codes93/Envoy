package main

import "time"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type UserDetailResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Platform  string    `json:"platform"`
	AppName   string    `json:"app_name"`
}

type RequestHeaders struct {
	UID      string `header:"x-auth-uid"      binding:"required"`
	DeviceID string `header:"x-auth-deviceid"`
	Platform string `header:"platform"`
	AppName  string `header:"app-name"`
}
