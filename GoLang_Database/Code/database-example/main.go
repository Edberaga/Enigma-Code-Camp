package main

import (
	"database-example/entity"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Edberaga7"
	dbname = "enigma_camp"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname);


func main() {
	// student1 := entity.Student{
	// 	Id: 7, 
	// 	Name: "Sisil",
	// 	Email: "Sisil@yahoo.com",
	// 	Address: "Malaka",
	// 	BirthDate: time.Date(2000, 11, 20, 0, 0, 0, 0, time.Local),
	// 	Gender: "F",
	// }

	// addStudent(student1)
	// updateStudent(student1)
	// deleteStudent("9")
	// students := getAllStudent()
	// for _, student := range students {
	// 	fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	// }
	// fmt.Println(getStudentById("7"))
	students := searchBy("Ke", "2000-11-01")
	for _, student := range students {
		fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	}
}

func getStudentById(id string) entity.Student{
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM mst_student WHERE id = $1;"

	student := entity.Student{}
	err = db.QueryRow(sqlStatement, id).Scan(&student.Id, &student.Name, &student.Email, &student.Address,
		&student.BirthDate, &student.Gender)
	if err != nil {
		panic(err)
	}
	return student
}

func getAllStudent() []entity.Student {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_student;"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func scanStudent(rows *sql.Rows) []entity.Student {
	students := []entity.Student{}
	var err error

	for rows.Next() {
		student := entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Address,
		&student.BirthDate, &student.Gender)
		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return students
}

func addStudent(s entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO mst_student (id, name, email, address, birth_date, gender) VALUES ($1, $2, $3, $4, $5, $6);"
	
	_ , err = db.Exec(sqlStatement, s.Id, s.Name, s.Email, s.Address, s.BirthDate, s.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

func updateStudent(s entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "UPDATE mst_student SET name = $2, email = $3, address= $4, birth_date = $5, gender = $6 WHERE id = $1;"
	
	_ , err = db.Exec(sqlStatement, s.Id, s.Name, s.Email, s.Address, s.BirthDate, s.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

func deleteStudent(id string) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "DELETE FROM mst_student WHERE id = $1;"
	
	_ , err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Deleted Data From Id: ", id)
	}
}

func searchBy(name string, dob string) []entity.Student{
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_student WHERE name LIKE $1 OR birth_date = $2;"

	rows, err := db.Query(sqlStatement, "%"+name+"%", dob)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Connected!")

	return db
}