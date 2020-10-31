package controllers

import (
	"api/domain/model"
	"api/interfaces/database"
	"api/usecase/interactor"
)

// UserController は usecase.interactor.UserInteractor を内包した型です。
type UserController struct {
	Interactor interactor.UserInteractor
}

// NewUserController は UserController を生成します。
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: interactor.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// Create :ユーザー登録に関するUserControllerのメソッド
func (controller *UserController) Create(c Context) {
	u := model.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201)
}
