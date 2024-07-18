package models
import(
	"os"
	"database/sql"
	"fmt"
)
func ExecuteRankingSQL(db *sql.DB, filePath string) (string, error) {
    query, err := os.ReadFile(filePath)
    if err != nil {
        return "", fmt.Errorf("could not read SQL file: %v", err)
    }

    return string(query), nil
}