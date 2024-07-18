package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"organization-ranking-backend/models"
	"fmt"
)

type OrganizationData struct {
    OrganizationID   int    `json:"organization_id"`
    OrganizationName string `json:"organization_name"`
    TotalScore       int    `json:"total_score"`
}

func GetOrganizationsRanking(c *gin.Context) {

	query, err := models.ExecuteRankingSQL(models.DB, "migrations/organization-ranking.sql")
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
	fmt.Println(rows)
    var organizations_data []OrganizationData
    for rows.Next() {
        var data OrganizationData
         err := rows.Scan(&data.OrganizationID, &data.OrganizationName, &data.TotalScore)
		 if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan result"})
            return
        }
        organizations_data = append(organizations_data, data)
    }

    c.JSON(http.StatusOK, organizations_data)
}