package database

type User struct {
	TelegramID string `gorm:"primaryKey;autoIncrement;column:id"`
	City string `gorm:"column:city;size:150"`
	Notify bool `gorm:"column:notify;default:false"`
}