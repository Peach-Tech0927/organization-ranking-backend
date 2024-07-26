package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id	   	   uint	  `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	GithubId   string `json:"github_id"`
	Contributions int `json:"contributions"`
}

func (u *User) CreateNewRecord() error {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", u.Email).Scan(&count)
	if err != nil {
		return fmt.Errorf("%w: %v", Err001, err)
	}

	if count > 0 {
		return fmt.Errorf("%w: %v", Err001, "user already exists")
	}

	u.HashPassword()

	result, err := DB.Exec("INSERT INTO users (email, username, password, github_id, contributions) VALUES (?, ?, ?, ?, ?)", u.Email, u.Username, u.Password,u.GithubId,u.Contributions)
	if err != nil {
		return fmt.Errorf("%w: %v", Err001, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%w: %v", Err001, err)
	}

	u.Id = uint(id)
	return nil
}

func (u *User) HashPassword() error {
	if u.Password == "" || len(u.Password) == 0 {
		return fmt.Errorf("%w: %v", Err002, u.Password)
	}
	
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return fmt.Errorf("%w: %v", Err002, err)
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return fmt.Errorf("%w: %v", Err003, err)
	}
	return nil
}

func (u *User) ToJSONResponse() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.Id,
		"email":    u.Email,
		"username": u.Username,
		"github_id": u.GithubId,
		"contributions": u.Contributions,
	}
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	err := DB.QueryRow("SELECT id, email, username, password, github_id, contributions FROM users WHERE email = ?", email).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.GithubId, &user.Contributions)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", Err004, err)
	}
	return &user, nil
}
