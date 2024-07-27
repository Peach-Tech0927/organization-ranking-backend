package models

import "fmt"

type UserOrganizationMembership struct {
	Id             uint `json:"id"`
	UserId 		   uint `json:"user_id"`
	OrganizationId uint `json:"organization_id"`
}

func (u *UserOrganizationMembership) CreateNewRecord() error {
	result, err := DB.Exec("INSERT INTO user_organization_memberships (user_id, organization_id) VALUES (?, ?)", u.UserId, u.OrganizationId)
	if err != nil {
		return fmt.Errorf("%w: %v", Err008, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%w: %v", Err008, err)
	}

	u.Id = uint(id)
	return nil
}