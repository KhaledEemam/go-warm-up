package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (app *application) login(c *gin.Context) {
	var auth loginRequest
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	existingUser, err := app.models.Users.GetByEmail(auth.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email of password."})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(auth.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password."})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": existingUser.Id,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	stringToken, err := token.SignedString([]byte(app.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: stringToken})
}