package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

func createTable() (*sql.DB, error) {
	os.MkdirAll("./storage/", 0755)
	os.Create("./storage/data.db")
	db, err := sql.Open("sqlite3", "./storage/data.db")
	if err != nil {
		log.Println("sql open " + err.Error())
		return nil, err
	}
	//query, err := db.Prepare("CREATE TABLE `sentences` (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`sentence` TEXT")
	query, err := db.Prepare("CREATE TABLE `sentences` (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`sentence` TEXT)")
	if err != nil {
		log.Println("prepare create table " + err.Error())
		return nil, err
	}
	res, err := query.Exec()
	if err != nil {
		log.Println("exec create table " + err.Error())
		return nil, err
	}
	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())
	return db, nil
}

func NewDatabase() *Database {
	db, err := createTable()
	if err != nil {
		log.Println("createTable: " + err.Error())
		return nil
	}
	database := &Database{
		db: db,
	}
	return database
}

type Database struct {
	db *sql.DB
}

func (d *Database) GetSentences() []string {
	rows, err := d.db.Query("SELECT * FROM sentences")

	if rows.Err() != nil {
		log.Println(rows.Err())
		return []string{}
	}
	if err != nil {
		log.Println(err.Error())
		return []string{}
	}
	dbSentences := make([]string, 0)
	for rows.Next() {
		s := sentences{}
		err = rows.Scan(&s.id, &s.sentence)
		if err != nil {
			log.Fatal(err)
			return dbSentences
		}
		dbSentences = append(dbSentences, s.sentence)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
	}
	err = rows.Close()

	if err != nil {
		return []string{}
	}
	return dbSentences
}

func (d *Database) AddSentence(s string) error {
	if d == nil {
		return errors.New("banco de dados é nulo")
	}
	newSentence := sentences{
		sentence: s,
	}
	stmt, err := d.db.Prepare("INSERT INTO sentences(id, sentence) VALUES (?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	res, err := stmt.Exec(nil, newSentence.sentence)
	log.Println(fmt.Sprint(res.LastInsertId()))
	log.Println(res.RowsAffected())
	defer stmt.Close()
	log.Println("Added " + newSentence.sentence)
	return err
}

func (d *Database) RemoveSentence(s string) error {
	stmt, err := d.db.Prepare("")
	if err != nil {
		return err
	}
	err = stmt.Close()
	return err
}

func (d *Database) Close() {
	d.db.Close()
}
