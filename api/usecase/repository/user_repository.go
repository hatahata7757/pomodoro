package repository

import "api/domain/model"

// UserRepository は User 型が持つデータ保存の振る舞いを記述したインターフェースです。
type UserRepository interface {
	Store(model.User) (int, error)
	FindById(int) (model.User, error)
	FindAll() (model.Users, error)
	DeleteById(int) (int, error)
}
