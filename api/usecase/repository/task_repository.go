package repository

import "api/domain/model"

// TaskRepository は Task 型が持つデータ保存の振る舞いを記述したインターフェースです。
type TaskRepository interface {
	Store(model.Task) (int, error)
	FindById(int) (model.Task, error)
}
