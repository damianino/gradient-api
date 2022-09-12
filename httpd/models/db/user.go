package dbm

import (
	"gradientApi/httpd/db"
)

type User struct {
	Id       int
	Username string
	Password string
}

func (u *User) Insert() error {
	query := `insert into users (username, password) values ($1, $2)`
	
	_, err := db.DB.Exec(query, u.Username, u.Password)
	if err != nil {
		return err
	} 
	return nil
}

func (u *User) Unique() bool {
	query := `select * from users where username = $1`
	rows, err := db.DB.Query(query, u.Username)
	if err != nil {
		return false
	} 
	defer rows.Close()
	var exists bool
	err = rows.Scan(&exists)
	return err == nil
}

