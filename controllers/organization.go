package controllers

import (
	"net/http"
	"organization-ranking-backend/models"

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

	c.JSON(http.StatusCreated, gin.H{"organization": organization.ToJSONResponse()})
}