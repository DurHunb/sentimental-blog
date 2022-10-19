package model

type User struct {
	*Model
	Nickname           string `json:"nickname"`
	Username           string `json:"username"`
	Phone              string `json:"phone"`
	Password           string `json:"password"`
	Encrypted_password string `json:"encrypted_password"`
	Portrait           string `json:"portrait"`
	Status             int    `json:"status"`
	Character          int    `json:"character"`
}
