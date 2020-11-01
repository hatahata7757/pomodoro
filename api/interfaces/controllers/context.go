package controllers

// Context は gin.Contextで利用するメソッドのインターフェースです。
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
