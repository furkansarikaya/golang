package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Fs253008@tcp(localhost:3306)/demodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	/*
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
	*/
	/*
		create table `users`(
		`ID` int(11) not null AUTO_INCREMENT,
		`Username` varchar(45) NOT NULL,
		`Email` varchar(45) NOT NULL,
		`Password` varchar(45) NOT NULL,
		`FirstName` varchar(45) NOT NULL,
		`LastName` varchar(45) NOT NULL,
		`BirthDate` varchar(45) DEFAULT NULL,
		`IsActive` tinyint(1) DEFAULT NULL,
		PRIMARY KEY (`ID`),
		UNIQUE KEY `ID_UNIQUE` (`ID`)
		) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
	*/
	/*
		createStatement := "`users`( `ID` int(11) not null AUTO_INCREMENT, `Username` varchar(45) NOT NULL, `Email` varchar(45) NOT NULL, `Password` varchar(45) NOT NULL, `FirstName` varchar(45) NOT NULL, `LastName` varchar(45) NOT NULL, `BirthDate` varchar(45) DEFAULT NULL, `IsActive` tinyint(1) DEFAULT NULL, PRIMARY KEY(`ID`), UNIQUE KEY `ID_UNIQUE` (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;"

		_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + createStatement)
		if err != nil {
			log.Fatal(err)
		}

		// Veri Ekleme
		res, err := db.Exec("INSERT INTO users(Username, Email, Password, FirstName,LastName,BirthDate,IsActive) VALUES('Deneme1','deneme@deneme.com','12345!','Furkan','SARIKAYA','1995.08.25',1)")
		if err != nil {
			log.Fatal(err)
		}

		rowCount, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Inserted %d rows", rowCount)

		lastID, err1 := res.LastInsertId()
		if err1 != nil {
			log.Fatal(err1)
		}

		log.Printf("Last Inserted ID %d", lastID)

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
		// Eklenen kayıtları getir.
		rows, err := db.Query("SELECT * FROM users")

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
		}

		//alternatif
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}

		rows.Close()
	*/

	//...........

	/*
		rows, errQ := db.Query("SELECT * FROM users WHERE ID = ?", 5)
		if errQ != nil {
			log.Fatal(errQ)
		}

		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
		}
		errQ = rows.Err()
		if errQ != nil {
			log.Fatal(errQ)
		}
	*/

	//...........

	/*
		err = db.QueryRow("SELECT * FROM users limit 1").Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			if err == sql.ErrNoRows {
				// kayıt yoksa yapılacak işler
			} else {
				log.Fatal(err)
			}
		}
		log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))

	*/

	//...........

	// Multiple Active REsult Set : MARS

	// _ err := db.Exec("DELETE FROM xTable; DELETE FROM xTable2")

	//Preparing Queries

	/*
		stmt, errQ := db.Prepare("SELECT * FROM users WHERE ID =?")
		if errQ != nil {
			log.Fatal(errQ)
		}
		defer stmt.Close()
		rows, errX := stmt.Query(5)
		if errX != nil {
			log.Fatal(errX)
		}
		defer rows.Close()
		for rows.Next() {
			scanErr := rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
			if scanErr != nil {
				log.Fatal(scanErr)
			}
			log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
		}

	*/

	//...........

	//Preparing Query - Single rows
	/*
		stmt, errQ := db.Prepare("SELECT * FROM users WHERE ID =?")
		if errQ != nil {
			log.Fatal(errQ)
		}
		defer stmt.Close()

		errX := stmt.QueryRow(6).Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if errX != nil {
			log.Fatal(errX)
		}

		log.Println(FirstName + " " + LastName)

	*/

	//...........

	//Modifying Data
	/*
		res, errQ := db.Exec("DELETE FROM users LIMIT 1")
		if errQ != nil {
			log.Fatal(errQ)
		}

		rowCount, errY := res.RowsAffected()
		if errY != nil {
			log.Fatal(errY)
		}
		log.Println(rowCount)

		lastID, errL := res.LastInsertId()
		if errL != nil {
			log.Fatal(errL)
		}
		log.Println(lastID)
	*/

	//Insert İşlemi
	/*
		stmt, errS := db.Prepare("INSERT INTO users(Username, Email, Password, FirstName,LastName,BirthDate,IsActive) VALUES(?,?,?,?,?,?,?)")
		if errS != nil {
			log.Fatal(errS)
		}

		defer stmt.Close()

		Username = "Furkan"
		Email = "furkan@furkansarikaya.com"
		Password = "12345"
		FirstName = "Furkan"
		LastName = "SARIKAYA"
		BirthDate = "1995.08.25"
		IsActive = true

		res, errStmt := stmt.Exec(Username, Email, Password, FirstName, LastName, BirthDate, IsActive)
		if errStmt != nil {
			log.Fatal(errStmt)
		}

		insertedID, _ := res.LastInsertId()

		ID = int(insertedID)

		log.Println(ID)
	*/

	//Transaction

	tx, errTx := db.Begin()
	// ....

	_, errTx = db.Exec("Update .....")
	//...
	if errTx != nil {
		log.Fatal(errTx)
		tx.Rollback()
	}
	errTx = tx.Commit()
}
