package models
import (
	"fmt"
)
type Organization struct {
	Id 	   uint	  `json:"id"`
	Name string `json:"name"`

}

func (o *Organization) SaveToDatabase() error {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM organizations WHERE name = ?", o.Name).Scan(&count)
	if err != nil {
		return fmt.Errorf("%w: %v", Err001, err)
	}

	if count > 0 {
		return fmt.Errorf("%w: %v", Err001, "organization already exists")
	}

	result, err := DB.Exec("INSERT INTO organizations (name) VALUES (?)", o.Name)
	if err != nil {
		return fmt.Errorf("%w: %v", Err001, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%w: %v", Err001, err)
	}

	o.Id = uint(id)
	return nil
}

func (o *Organization) ToJSONResponse() map[string]interface{} {
	return map[string]interface{}{
		"id": o.Id,
		"name": o.Name,
	}
}