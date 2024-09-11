package SQLRequests

import "database/sql"

type User struct {
	id       int
	login    string
	password []byte
	notes    []Note `json:"Notes"`
}

type Note struct {
	text string `json:"Text"`
}

func CheckIfLoginOccupied(login string) bool {
	db, err := sql.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row := db.QueryRow("select exists(select * from Users where login = %v)", login)
	var occupied bool
	row.Scan(&occupied)
	return occupied
}

func CheckForLoginData(login, password string) bool {
	db, err := sql.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Query("select * from Users where login = %v")

}
