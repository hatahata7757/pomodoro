package main

import "api/infrastructure"

// ルーティングの呼び出し
func main() {
	infrastructure.Router.Run()
}
