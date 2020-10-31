package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// SqlHandler はSQLハンドラを扱う型です。
type SqlHandler struct {
	Conn *sql.DB
}

// NewSqlHandler は SqlHandlerを生成します。
func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/pomodoro")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
