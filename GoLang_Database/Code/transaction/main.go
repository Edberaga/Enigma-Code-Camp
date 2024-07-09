// package main

// import "fmt"

// func init() {
// 	fmt.Println("init function")
// }

// func main() {
// 	fmt.Println("Main Function")
// }

package main

import (
	"fmt"
	"database/sql"
	tx "transaction/entity"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Edberaga7"
	dbname = "enigma_camp"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	student := tx.Enroll{
		Id: 1,
		StudentId: 7,
		Subject: "Database",
		Credit: 4,
	}

	enrollSubject(student)
}

func enrollSubject(s tx.Enroll) {
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	//insert student enrollment
	insertEnrollment(s, tx)
	takenCredit := getStudentCredit(s.StudentId, tx)
	updateStudent(takenCredit, s.StudentId, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Succesfully insert for student ID: ", s.Id)
	}

}

func insertEnrollment(s tx.Enroll, tx *sql.Tx) {
	query := "INSERT INTO tx_enrollment (id, studentId, subject, credit) VALUES ($1, $2, $3, $4);"

	_, err := tx.Exec(query, s.Id, s.StudentId, s.Subject, s.Credit)
	validate(err, "Insert", tx)
}

func getStudentCredit(id int, tx *sql.Tx) int{
	query := "SELECT SUM(credit) FROM tx_enrollment WHERE studentId = $1;"

	takenCredit := 0
	err := tx.QueryRow(query, id).Scan(&takenCredit)
	validate(err, "Select", tx)

	return takenCredit
}

func updateStudent(takenCredit int, studentId int, tx *sql.Tx) {
	query := "UPDATE mst_student SET credit = $1 WHERE id = $2;"

	_, err := tx.Exec(query, takenCredit, studentId)
	validate(err, "Update", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback!")
	} else {
		fmt.Println("Succesfully", message, "data!")
	}
}

func connectDb()*sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Connected")
	}

	return db
}