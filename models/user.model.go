package models

type PermissionLevel string

const (
	USER               PermissionLevel = "USER"
	ADMINISTRATION     PermissionLevel = "ADMINISTRATION"
	SUPERADMINISTRATOR PermissionLevel = "SUPERADMINISTRATOR"
)

type User struct {
	ID              uint            `gorm:"primary_key;auto_increment" json:"id"`
	Username        string          `gorm:"size:255;not null;unique" json:"username"`
	Password        string          `gorm:"size:255;not null;" json:"password"`
	Name            string          `gorm:"size:255;not null;" json:"name"`
	PermissionLevel PermissionLevel `gorm:"type:permission_level" json:"permission_level"`
}
