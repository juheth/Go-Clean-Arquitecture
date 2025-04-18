package entities

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"uniqueIndex"`
}
