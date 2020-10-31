package controllers

import (
	"api/domain/model"
	"api/interfaces/database"
	"api/usecase/interactor"
	"strconv"
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

// Index :ユーザー一覧取得に関するUserControllerのメソッド
func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

// Show :ユーザー情報取得に関するUserControllerのメソッド
func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}
