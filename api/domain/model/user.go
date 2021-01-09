package model

// User は単一ユーザーの型
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Users は複数のユーザーを扱う型
type Users []User
