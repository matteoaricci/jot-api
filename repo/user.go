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
	RoleID    *uint64        `gorm:"column:role_id" json:"roleId"`
}

type Role struct {
	ID              uint64         `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	RoleName        string         `gorm:"column:role_name" json:"roleName"`
	RoleDescription string         `gorm:"column:role_description" json:"roleDescription"`
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

func CreateUser(ctx context.Context, firstName string, lastName string, email string, password string, roleID uint64) error {
	slog.InfoContext(ctx, "repo.CreateUser")

	u := Usr{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		RoleID:    &roleID,
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

func UpdateUserRole(ctx context.Context, userID uint64, roleID uint64) error {
	slog.InfoContext(ctx, "repo.UpdateUserRole", slog.Uint64("userId", userID), slog.Uint64("roleId", roleID))

	row := db.WithContext(ctx).Model(&Usr{}).Where("id = ?", userID).Update("role_id", roleID)
	if row.Error != nil {
		return row.Error
	}
	if row.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func FindRoleByName(ctx context.Context, name string) (*Role, error) {
	slog.InfoContext(ctx, "repo.FindRoleByName", slog.String("name", name))

	var r Role
	row := db.WithContext(ctx).Where(&Role{RoleName: name}).First(&r)
	if row.Error != nil {
		return nil, row.Error
	}

	return &r, nil
}

func FindRoleByID(ctx context.Context, id uint64) (*Role, error) {
	slog.InfoContext(ctx, "repo.FindRoleByID", slog.Uint64("id", id))

	var r Role
	row := db.WithContext(ctx).First(&r, id)
	if row.Error != nil {
		return nil, row.Error
	}

	return &r, nil
}
