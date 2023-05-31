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
	if d == nil || d.db == nil {
		return []string{}
	}
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
	if d == nil || d.db == nil {
		return errors.New("banco de dados é nulo")
	}
	if len(s) == 0 {
		return nil
	}
	newSentence := sentences{
		sentence: s,
	}
	stmt, err := d.db.Prepare("INSERT INTO sentences(id, sentence) VALUES (?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	res, err := stmt.Exec(nil, newSentence.sentence)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(fmt.Sprint(res.LastInsertId()))
	log.Println(res.RowsAffected())
	log.Println("Added " + newSentence.sentence)
	return stmt.Close()
}

func (d *Database) RemoveSentence(s string) error {
	if d == nil || d.db == nil {
		return errors.New("banco de dados é nulo")
	}
	stmt, err := d.db.Prepare("DELETE FROM sentences WHERE id=?")
	if err != nil {
		return err
	}
	index := d.getIndexFromStr(s)
	res, err := stmt.Exec(index)
	if err != nil {
		return err
	}
	log.Println(fmt.Sprint(res.LastInsertId()))
	log.Println(res.RowsAffected())
	log.Println("Removed " + s)
	return stmt.Close()
}

func (d *Database) Close() {
	if d == nil || d.db == nil {
		return
	}
	d.db.Close()
}

func (d *Database) getIndexFromStr(s string) int {
	sentences := d.GetSentences()
	index := -1
	if d == nil || d.db == nil {
		return -1
	}
	for i, sentence := range sentences {
		if sentence == s {
			index = i
		}
	}
	return index
}
