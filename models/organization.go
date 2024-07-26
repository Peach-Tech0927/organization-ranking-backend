package models

import "fmt"

type Organization struct {
    Id                 int    `json:"id"`
    Name               string `json:"name"`
    TotalContributions int    `json:"total_contributions"`
}

func (o *Organization) CreateNewRecord() error { // contributionの取得は行わない
    var count int
    err := DB.QueryRow("SELECT COUNT(*) FROM organizations WHERE id = ?", o.Id).Scan(&count)
    if err != nil {
        return fmt.Errorf("%w: %v", Err005, err)
    }
    if count > 0 {
        return fmt.Errorf("%w: %v", Err005, "organization already exists")
    }

    result, err := DB.Exec("INSERT INTO organizations (name) VALUES (?)", o.Name)
    if err != nil {
        return fmt.Errorf("%w: %v", Err005, err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("%w: %v", Err005, err)
    }
    o.Id = int(id)

    return nil
}

func (o *Organization) UpdateTotalContributions() error {
    _, err := DB.Exec("UPDATE organizations SET total_contributions = ? WHERE id = ?", o.TotalContributions, o.Id)
    if err != nil {
        return fmt.Errorf("%w: %v", Err006, err)
    }
    return nil
}
