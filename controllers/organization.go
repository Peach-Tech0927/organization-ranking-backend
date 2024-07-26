package controllers

import (
	"net/http"
	"organization-ranking-backend/models"
	"organization-ranking-backend/utils"

	"github.com/gin-gonic/gin"
)

type CreateOrganizationInput struct {
	OrganizationName string `json:"name"`
}

func CreateOrganization(c *gin.Context) {
	var input CreateOrganizationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var organization models.Organization
	organization.Name = input.OrganizationName
	if err := organization.CreateNewRecord(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := utils.ExtractTokenId(c) // エラー処理はミドルウェアに任せる
	user, err := models.FindUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := organization.Join(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"organization": organization.ToJSONResponse()})
}