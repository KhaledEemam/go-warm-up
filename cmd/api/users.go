package main

import (
	"net/http"
	"github.com/KhaledEemam/go-warm-up/internal/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required,min=2"`
	Password string `json:"password" binding:"required,min=8"`
}

func (app *application) registerUser(c *gin.Context) {
	var registerRequest registerRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}


	generatedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	} 
	registerRequest.Password = string(generatedPassword)

	user := database.User{
		Email:registerRequest.Email ,
		Name: registerRequest.Name,
		Password: registerRequest.Password,
	}

	err = app.models.Users.Insert(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()} )
		return
	}

	c.JSON(http.StatusOK, user)
}