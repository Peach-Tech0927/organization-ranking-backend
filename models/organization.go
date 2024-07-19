package models

type Organization struct {
    OrganizationID   int    `json:"organization_id"`
    OrganizationName string `json:"organization_name"`
    TotalContributions       int    `json:"total_contributions"`
}
