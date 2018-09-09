package models

type NotificationType struct {
	ID   uint `gorm:"primary_key"`
	Type string
}
