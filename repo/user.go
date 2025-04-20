package repo

import (
	"gorm.io/gorm"
	"time"
)

type Usr struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Email     string         `gorm:"email" json:"email"`
	Password  string         `gorm:"password" json:"password"`
	FirstName string         `gorm:"first_name" json:"firstName"`
	LastName  string         `gorm:"last_name" json:"lastName"`
}

func FindUser(email string, password string) (*Usr, error) {
	m := make(map[string]any)
	m["email"] = email
	m["password"] = password

	var u Usr

	row := db.Where(m).Find(&u)
	if row.Error != nil {
		return nil, row.Error
	}

	return &u, nil
}

func CreateUser(firstName string, lastName string, email string, password string) error {
	u := Usr{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	row := db.Create(&u)
	if row.Error != nil {
		return row.Error
	}

	return nil
}
