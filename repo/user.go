package repo

import (
	"context"
	"log/slog"
	"time"

	"gorm.io/gorm"
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
	Role      string         `gorm:"type:text;default:'user'" json:"role"`
}

func FindUser(ctx context.Context, email string) (*Usr, error) {
	slog.InfoContext(ctx, "repo.FindUser")

	var u Usr

	row := db.WithContext(ctx).Where(&Usr{Email: email}).First(&u)
	if row.Error != nil {
		return nil, row.Error
	}

	return &u, nil
}

func CreateUser(ctx context.Context, firstName string, lastName string, email string, password string, role string) error {
	slog.InfoContext(ctx, "repo.CreateUser")

	u := Usr{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
	}

	row := db.WithContext(ctx).Create(&u)
	if row.Error != nil {
		return row.Error
	}

	return nil
}

func DeleteUser(ctx context.Context, id string) error {
	slog.InfoContext(ctx, "repo.DeleteUser", slog.String("id", id))

	row := db.WithContext(ctx).Delete(&Usr{}, id)
	if row.Error != nil {
		return row.Error
	}

	return nil
}
