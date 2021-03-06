package database

import "api/domain/model"

// SqlHandler を埋め込んだ UserRepository 型を定義
type UserRepository struct {
	SqlHandler
}

// Store は User を保存するメソッドです
func (repo *UserRepository) Store(u model.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (name) VALUES (?)", u.Name,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

// FindById は 引数で渡された id に一致する User を取得するメソッドです
func (repo *UserRepository) FindById(identifier int) (user model.User, err error) {
	row, err := repo.Query("SELECT id, name FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	row.Next()
	if err = row.Scan(&id, &name); err != nil {
		return
	}
	user.ID = id
	user.Name = name
	return
}

// FindAll は全ての User を取得するメソッドです
func (repo *UserRepository) FindAll() (users model.Users, err error) {
	rows, err := repo.Query("SELECT id, name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		user := model.User{
			ID:   id,
			Name: name,
		}
		users = append(users, user)
	}
	return
}

// DeleteById は引数で渡された id に一致する User をレコードから削除するメソッドです
func (repo *UserRepository) DeleteById(identifier int) (rowCnt int, err error) {
	result, err := repo.Execute("DELETE FROM users WHERE id = ?", identifier)
	if err != nil {
		return
	}
	// 影響を受けた行数を返す(返り値が必要だったため, 便宜的対応)
	rowCnt64, err := result.RowsAffected()
	if err != nil {
		return
	}
	rowCnt = int(rowCnt64)
	return
}
