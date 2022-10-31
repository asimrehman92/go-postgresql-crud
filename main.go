<<<<<<< HEAD
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgre@123" //psql shell password
	dbname   = "mydatabase"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}



func main() {

	//connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err := sql.Open("postgres", psqlconn)

	//call function to check error
	CheckError(err)

	//close database
	defer db.Close()

	//check db
	//It is vitally important that you call the Ping() method
	//becuase the sql.Open() function call does not ever create a connection to the database.
	//Instead, it simply validates the arguments provided.
	//By calling db.Ping() we force our code to actually open up a connection to the database
	//which will validate whether or not our connection string was 100% correct.

	err = db.Ping()
	CheckError(err)
	fmt.Println("\nMydatabase is Successfully connected")
	fmt.Println()

	//------------insert-------------------
	//hardcoded
	// insertStmt := `INSERT INTO students (name, roll) VALUES ('Cloud',8)`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	// //insert dynamic code
	// insrtDynStmt := `INSERT INTO Students (Name, Roll) VALUES ($1,$2)`
	// _, e := db.Exec(insrtDynStmt, "Jane", 2)
	// CheckError(e)

	//---------------update------------------
	// updateStmt := `UPDATE students SET name = $1 WHERE roll = $2'`
	// _, e := db.Exec(updateStmt, "Bowser",7)
	// CheckError(e)

	//-------DELETE-------
	// deleteStmt := `DELETE FROM students WHERE roll = $1`
	// _, e := db.Exec(deleteStmt, 7)
	// CheckError(e)

	//--------SELECT------------
	rows, err := db.Query(`SELECT * FROM students`)

	CheckError(err)

	//print 1 record
	// var name string
	// var roll int
	// err = rows.Scan(&name, &roll)
	// CheckError(err)
	// fmt.Println(name, roll)

	//Printing multiple rows
	fmt.Println("------THIS IS A STUDENT TABLE-----")
	fmt.Println()
	for rows.Next() {
		var name string
		var roll int
		err = rows.Scan(&name, &roll)
		CheckError(err)
		fmt.Println(name, roll)
	}
	fmt.Println()
}
=======
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgre@123" //psql shell password
	dbname   = "mydatabase"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err := sql.Open("postgres", psqlconn)

	//call function to check error
	CheckError(err)

	//close database
	defer db.Close()

	//check db
	//It is vitally important that you call the Ping() method
	//becuase the sql.Open() function call does not ever create a connection to the database.
	//Instead, it simply validates the arguments provided.
	//By calling db.Ping() we force our code to actually open up a connection to the database
	//which will validate whether or not our connection string was 100% correct.

	err = db.Ping()
	CheckError(err)
	fmt.Println("Successfully connected")

	//insert
	//hardcoded
	insertStmt := `INSERT INTO students (name, roll) VALUES ('Koopa',4)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// //insert dynamic code
	// insrtDynStmt := `INSERT INTO Students (Name, Roll) VALUES ($1,$2)`
	// _, e := db.Exec(insrtDynStmt, "Jane", 2)
	// CheckError(e)
}
>>>>>>> 55d448b013069d2583a4a9dc49b2b9348077f3b9
