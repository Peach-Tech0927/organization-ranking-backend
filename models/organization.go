package models

import "fmt"

type Organization struct {
    Id                 uint    `json:"id"`
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
    o.Id = uint(id)

    return nil
}

func (o *Organization) Join(user *User) error {
    var membership UserOrganizationMembership
    membership.UserId = user.Id
    membership.OrganizationId = o.Id

    err := membership.CreateNewRecord()
    if err != nil {
        return fmt.Errorf("%w: %v", Err006, err)
    }

    return nil
}

// 隠匿するべき情報があるモデルをそのまま返さないために、各モデルにToJSONResponseを使うことを慣習にする
func (o *Organization) ToJSONResponse() map[string]interface{} {
    return map[string]interface{}{
        "id":                o.Id,
        "name":              o.Name,
        "total_contributions": o.TotalContributions,
    }
}