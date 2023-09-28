// models/user_model.go
package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
}
