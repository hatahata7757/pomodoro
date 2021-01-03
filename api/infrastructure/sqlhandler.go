package infrastructure

import (
	"database/sql"

	"api/interfaces/database"

	_ "github.com/go-sql-driver/mysql"
)

// SqlHandler はSQLハンドラを扱う型です。
type SqlHandler struct {
	Conn *sql.DB
}

// NewSqlHandler は、生成したSqlHandlerのインターフェースを返します。
func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "root:password@tcp(db:3306)/pomodoro")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

// Execute は、SqlHandler インターフェースにあるExecute の振る舞いの実装です。
func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

// Query は、SqlHandler インターフェースにある Query の振る舞いの実装です。
func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

// LastInsertId は、SqlResult インターフェースにある LastInsertId の振る舞いの実装です。
func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected は、SqlResult インターフェースにある RowsAffected の振る舞いの実装です。
func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

// Scan は、SqlRow インターフェースにある Scan の振る舞いの実装です。
func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next は、SqlRow インターフェースにある Next の振る舞いの実装です。
func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

// Close は、SqlRow インターフェースにある Close の振る舞いの実装です。
func (r SqlRow) Close() error {
	return r.Rows.Close()
}
