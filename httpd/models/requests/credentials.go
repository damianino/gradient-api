package requests

import (
	"fmt"
	"gradientApi/httpd/db"
	dbm "gradientApi/httpd/models/db"
)

type Credentials struct {
	Username string
	Password string
}

func (c *Credentials) Validate() (bool, error) {
	// TODO implement
	fmt.Println("credentials.go >>> TODO: validate")
	return true, nil
}

func (u *Credentials) Authenticate() (bool, error) {
	query := `select * from users where username = $1`
	rows, err := db.DB.Query(query, u.Username)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	var user dbm.User
	rows.Next()
	err = rows.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil{
		return false, err
	}
	
	fmt.Println("credentials.go >>> CHECK PASSWORD HASH")
	return true, nil
}