package dboperations

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@/demo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	/*
			CREATE TABLE `users` (
		  `ID` int NOT NULL,
		  `Username` varchar(45) DEFAULT NULL,
		  `Email` varchar(45) DEFAULT NULL,
		  `Password` varchar(45) DEFAULT NULL,
		  `FirstName` varchar(45) DEFAULT NULL,
		  `LastName` varchar(45) DEFAULT NULL,
		  `BirthDate` varchar(45) DEFAULT NULL,
		  `IsActive` varchar(45) DEFAULT NULL,
		  PRIMARY KEY (`ID`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
	*/

	createStatement := "`users` (`ID` int NOT NULL,`Username` varchar(45) DEFAULT NULL, `Email` varchar(45) DEFAULT NULL,`Password` varchar(45) DEFAULT NULL,`FirstName` varchar(45) DEFAULT NULL, `LastName` varchar(45) DEFAULT NULL, `BirthDate` varchar(45) DEFAULT NULL,`IsActive` varchar(45) DEFAULT NULL, PRIMARY KEY (`ID`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci"
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + createStatement)
	if err != nil {
		log.Fatal(err)
	}

	// insert records
	// stmt prepare + stmt Exec() şeklinde de yapılabilir
	res, err := db.Exec("INSERT INTO users (Username, Email, Password, FirstName, LastName, BirthDate, IsActive) VALUES ('ocakdans', 's@gmail.com', '123', 'selim', 'ocakdan', '17.02.1085', '1')")
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Inserted %d rows", rowCount)

	var (
		ID        int
		Username  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		BirthDate string
		IsActive  bool
	)

	// select records
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Username: %s, Email: %s, Password: %s, FirstName: %s, LastName: %s, BirthDate: %s, IsActive: %t", ID, Username, Email, Password, FirstName, LastName, BirthDate, IsActive)
	}

	createdUsersInfo, err := db.Query("SELECT * FROM users WHERE ID = ?", 3)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for createdUsersInfo.Next() {
		err := createdUsersInfo.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Username: %s, Email: %s, Password: %s, FirstName: %s, LastName: %s, BirthDate: %s, IsActive: %t", ID, Username, Email, Password, FirstName, LastName, BirthDate, IsActive)
	}

	fmt.Println("xxxxxxxx1")

	//var x string
	err = db.QueryRow("SELECT * FROM users limit 1").Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("bir satır bulundu", Username)
	log.Printf("ID: %d, Username: %s, Email: %s, Password: %s, FirstName: %s, LastName: %s, BirthDate: %s, IsActive: %t", ID, Username, Email, Password, FirstName, LastName, BirthDate, IsActive)

	// alternative usage
	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// rows.Close()

	// Multiple Active Result Set : MARS
	//_, err := db.Query("DELETE FROM tableX; DELETE FROM tableY")

	//Preparing Statements
	fmt.Println("xxxxxxxx2")
	stmt, errPS := db.Prepare("SELECT * FROM users WHERE ID = ?")
	if errPS != nil {
		log.Fatal(errPS)
	}
	defer stmt.Close()
	rows, errx := stmt.Query(3) // bunun tek row dönen yapısı QueryRow()
	if errx != nil {
		log.Fatal(errx)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Username: %s, Email: %s, Password: %s, FirstName: %s, LastName: %s, BirthDate: %s, IsActive: %t", ID, Username, Email, Password, FirstName, LastName, BirthDate, IsActive)
	}

	// Moıdify records

	res, _ = db.Exec("DELETE FROM users LIMIT 1")
	rowCount, _ = res.RowsAffected()
	log.Printf("Deleted %d rows", rowCount)
	lastID, _ := res.LastInsertId()
	log.Printf("Last inserted ID: %d", lastID)

}
