package main

import (
	users "database-practice/entity"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Edberaga7"
	dbname = "users"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname);

func main() {
	// user := users.User{
	// 	Id: 4,
	// 	Name: "Ray",
	// }

	// updateUser(user)
	// fmt.Println(getStudentById(1))
	// userData := getUserAll()

	// fmt.Println(userData)
	// for i := 0; i < len(userData); i++ {
	// 	fmt.Println(userData[i].Id, userData[i].Name)
	// }

	userData := searchUser("S")
	for _, userData := range userData {
		fmt.Println(userData.Id, userData.Name)
	}
}

func getStudentById(id int) users.User {
	db := connectDb()
	defer db.Close()

	query := "SELECT * FROM users WHERE userId = $1"
	userList := users.User{}
	var err error

	err = db.QueryRow(query, id).Scan(&userList.Id, &userList.Name)
	if err != nil {
		fmt.Println(err)
	}
	return userList
}

func getUserAll() []users.User {
	db := connectDb()
	defer db.Close()

	query := "SELECT * FROM users"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	userList := scanRow(rows)
	return userList
}

func scanRow(rows *sql.Rows) []users.User {
	userList := []users.User{}

	for rows.Next() {
		userRows := users.User{}
		err := rows.Scan(&userRows.Id, &userRows.Name)
		if(err != nil) {
			fmt.Println(err)
		}
		userList = append(userList, userRows)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}
	return userList
}

func addUser(s users.User) {
	db := connectDb()
	defer db.Close()
	var err error

	query := "INSERT INTO users(userId, userName) VALUES($1, $2);"

	_, err = db.Exec(query, s.Id, s.Name)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Succesfully Insert Data")
	}
}

func searchUser(name string) []users.User {
	db := connectDb()
	defer db.Close()
	query := "SELECT * FROM users WHERE userName LIKE $1"

	rows, err := db.Query(query, "%"+name+"%")
	if err != nil {
		panic(err)
	}
	userList := scanRow(rows)

	return userList
}

func updateUser(s users.User) {
	db := connectDb()
	defer db.Close()

	query := "UPDATE users SET userName = $2 WHERE userId = $1;"

	_, err := db.Exec(query, s.Id, s.Name)
	if(err != nil) {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Succesfully Update data!")
	}
}

func deletUser(id int) {
	db := connectDb()
	defer db.Close()
	query := "DELETE FROM users WHERE userId = $1;"

	_, err := db.Exec(query, id)
	if(err != nil) {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Succesfully Delete User Id:", id)
	}
}

func connectDb()*sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if(err != nil) {
		fmt.Println(err)
	}
	fmt.Println("Succesfully connected.")
	return db
}