package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func loginUser(c *gin.Context) {
	var loginRequest LoginRequest
	var isUserCreated bool

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request": err.Error()})
		return
	}

	user, err := GetUserByEmail(loginRequest.Email)

	if err != nil && err.Error() == "record not found" {
		log.Info("User not found, creating new user")
		user = User{
			Email:    loginRequest.Email,
			Password: loginRequest.Password,
		}
		if err := createUser(&user); err != nil {
			log.Error("Error creating user", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		isUserCreated = true
	} else if err != nil {
		log.Error("Error getting user by email", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if !isUserCreated && user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	accessToken, err := createJwtToken(user)
	if err != nil {
		log.Error("Error creating JWT token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func getUserDetails(c *gin.Context) {
	userID := c.Param("id")

	var requestHeaders RequestHeaders
	if err := c.ShouldBindHeader(&requestHeaders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request": err.Error()})
		return
	}

	if requestHeaders.UID != userID {
		log.Info("User ID mismatch", requestHeaders.UID, userID)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID mismatch"})
		return
	}

	user, err := GetUserByID(userID)

	if err != nil && err.Error() == "record not found" {
		log.Info("User not found, creating new user")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return

	} else if err != nil {
		log.Error("Error getting user by email", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	userDetailResponse := UserDetailResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Platform:  requestHeaders.Platform,
		AppName:   requestHeaders.AppName,
	}

	c.JSON(http.StatusOK, userDetailResponse)
}
