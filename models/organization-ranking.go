package models
import(
	"os"
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
type OrganizationData struct {
    OrganizationID   int    `json:"organization_id"`
    OrganizationName string `json:"organization_name"`
    TotalContributions       int    `json:"total_contributions"`
}

func OrganizationRankingSQL(db *sql.DB, filePath string) (string, error) {
    query, err := os.ReadFile(filePath)
    if err != nil {
        return "", fmt.Errorf("could not read SQL file: %v", err)
    }

    return string(query), nil
}

func GetOrganizationResultJSON(rows *sql.Rows,c *gin.Context) ([]OrganizationData, error) {
	var organizations_data []OrganizationData
    for rows.Next() {
        var data OrganizationData
         err := rows.Scan(&data.OrganizationID, &data.OrganizationName, &data.TotalContributions)
		 if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan result"})
            return nil, err
        }
        organizations_data = append(organizations_data, data)
    }
	return organizations_data, nil
}