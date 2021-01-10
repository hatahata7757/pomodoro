package interactor

import (
	"api/domain/model"
	"api/usecase/repository"
)

// TaskInteractor は TaskRepositoryを埋め込んだ型です。
type TaskInteractor struct {
	TaskRepository repository.TaskRepository
}

// Add は TaskRepository を通して Store を実行するためのメソッドです。
func (interactor *TaskInteractor) Add(t model.Task) (task model.Task, err error) {
	identifier, err := interactor.TaskRepository.Store(t)
	if err != nil {
		return
	}
	task, err = interactor.TaskRepository.FindById(identifier)
	return
}
