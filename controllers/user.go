package controllers

import (
	// "fmt"
	"net/http"
	"organization-ranking-backend/models"
	"organization-ranking-backend/models/githubQuery"
	"organization-ranking-backend/utils"

	"github.com/gin-gonic/gin"
)

type registerInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	GithubId string `json:"githubid" binding:"required"`
}

func Register(c *gin.Context) {
	var input registerInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contributions,err := githubQuery.GetContributions(input.GithubId)
	// fmt.Println(contributions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	user.Email = input.Email
	user.Username = input.Username
	user.Password = input.Password
	user.GithubId = input.GithubId
	user.Contributions = contributions

	err = user.SaveToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateTokenFromId(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user.ToJSONResponse(), "token": token})
}

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input loginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := models.FindUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = existingUser.CheckPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateTokenFromId(existingUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": existingUser.ToJSONResponse(), "token": token})
}