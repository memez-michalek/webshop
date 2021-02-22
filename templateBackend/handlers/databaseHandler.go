package handlers

import (
	"database/sql"
	"log"
	"templateBackend/models"

	"golang.org/x/crypto/bcrypt"
)

func initDB() *sql.DB {
	psqlconnection := "postgres://postgres:password@localhost/db?sslmode=disable"
	db, err := sql.Open("postgres", psqlconnection)
	if err != nil {
		log.Print("postgres connection error", err)
	}
	err = db.Ping()
	if err != nil {
		log.Print("could not ping database ", err)

	}
	return db

}


func queryUser(model models.User, db *sql.DB) (string, string, error) {
	var (
		tempEmail    string
		tempPassword string
	)

	result := db.QueryRow("SELECT * FROM USERS WHERE email=$1", model.Email).Scan(&tempEmail, &tempPassword)
	switch {
	case result == sql.ErrNoRows:
		log.Printf("could not find particular user")
		return "", "", result
	default:
		log.Printf("found user")
		return tempEmail, tempPassword, result
	}

}

func (model models.User) CheckLoginData() error {
	var(
		database := initDB()
		
	)
	defer db.Close()
	_, password, err := queryUser(model, database)
	if err != nil {
		log.Print("could not query user", err)
		return err
	}
	match := bcrypt.CompareHashAndPassword([]byte(password, []byte(model.Password))
	
	if match != nil {
		log.Print("error wrong credentials")
		return match
	}
	return nil

}
func (model models.User) RegisterUser() error{
	var(
		db := initDB()
	)
	defer db.Close()
	_, _, err := queryUser(model, db)
	if err == sql.ErrNoRows{
		hashedpw, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil{
		log.Print("could not hash password")
		return errors.Error("could not hash password")
		}
		_, err := db.Exec("INSERT INTO USERS(email, password)VALUES($1, $2)", model.Email, hashedpw)
		if err != nil {
			log.Print("could not insert data into database")
			return errors.Error("could not insert data")
		}
		return nil

	}else{
		log.Print("user already exists")
		return errors.Error("particular user already exists")
	}
}