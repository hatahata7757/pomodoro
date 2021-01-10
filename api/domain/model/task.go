package model

// Task は単一タスクの型
type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Done   bool   `json:"done"`
	UserId int    `json:"user_id"`
}

// Tasks は複数タスクのユーザーを扱う型
type Tasks []Task
