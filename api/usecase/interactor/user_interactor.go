package interactor

import (
	"api/domain/model"
	"api/usecase/repository"
)

// UserInteractor は UserRepositoryを持つ型です。
type UserInteractor struct {
	UserRepository repository.UserRepository
}

// Add は UserRepository を通して Store を実行するためのメソッドです。
func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

// Users は UserRepository を通して FindAll を実行するためのメソッドです。
func (interactor *UserInteractor) Users() (user model.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

// UserById は UserRepository を通して FindByIdを実行するためのメソッドです。
func (interactor *UserInteractor) UserById(identifier int) (user model.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

// DeleteById はUserRepository を通して DeleteByIdを実行するためのメソッドです。
func (interactor *UserInteractor) DeleteById(identifier int) (err error) {
	_, err = interactor.UserRepository.DeleteById(identifier)
	return
}
