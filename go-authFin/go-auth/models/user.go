package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Password []byte `json:"-"`
	Email    string `json:"email" gorm:"unique"`
}
