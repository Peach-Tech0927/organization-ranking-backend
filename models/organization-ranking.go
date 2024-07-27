package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func OrganizationRankingSQL(db *sql.DB, filePath string) (string, error) {
    query, err := os.ReadFile(filePath)
    if err != nil {
        return "", fmt.Errorf("could not read SQL file: %v", err)
    }

    return string(query), nil
}

func GetOrganizationResultJSON(rows *sql.Rows,c *gin.Context) ([]Organization, error) {
	var organizations_data []Organization
    for rows.Next() {
        var data Organization
         err := rows.Scan(&data.Id, &data.Name, &data.TotalContributions)
		 if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan result"})
            return nil, err
        }
        organizations_data = append(organizations_data, data)
    }
	return organizations_data, nil
}