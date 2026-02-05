package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db *sql.DB
)

type Worter struct {
	ID        int
	Wörter    string
	Translate string
}

func main() {

	_ = godotenv.Load()

	var err error

	db, err = connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vorter := Worter{
		Wörter:    "gut",
		Translate: "Bom",
	}

	if err := insertWörter(vorter); err != nil {
		panic(err)
	}

	words, err := getAllWorter()
	if err != nil {
		panic(err)
	}

	for _, w := range words {
		fmt.Println(*w)
	}
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func getAllWorter() ([]*Worter, error) {

	res, err := db.Query(
		`
	SELECT id, worter, translate 
	FROM deutsche_worter
	`,
	)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	words := []*Worter{}

	for res.Next() {
		var word Worter

		if err := res.Scan(&word.ID, &word.Wörter, &word.Translate); err != nil {
			return nil, err
		}

		words = append(words, &word)
	}

	return words, nil
}

func insertWörter(worter Worter) error {
	_, err := db.Exec(
		"INSERT INTO deutsche_worter (worter, translate) VALUES (?, ?)",
		worter.Wörter,
		worter.Translate,
	)
	if err != nil {
		return err
	}
	fmt.Println("Wort erfolgreich eingegeben")
	return nil

}
