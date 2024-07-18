package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"organization-ranking-backend/models"
)

func GetOrganizationsRanking(c *gin.Context) {

	query, err := models.OrganizationRankingSQL(models.DB, "migrations/organization-ranking.sql")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read SQL file"})
        return
    }

	rows, err := models.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute SQL query"})
		return
	}
	defer rows.Close()
	organizations_data, err := models.GetOrganizationResultJSON(rows,c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get organization data"})
		return
	}
    c.JSON(http.StatusOK, organizations_data)
}