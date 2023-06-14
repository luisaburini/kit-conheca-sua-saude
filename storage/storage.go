package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

func createTables() (*sql.DB, error) {
	// err := os.MkdirAll("./storage/", 0755)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	_, err := os.Create("./data.db")
	if err != nil {
		log.Println(err.Error())
	}
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println("sql open " + err.Error())
		return nil, err
	}
	query, err := db.Prepare("CREATE TABLE `sentences` (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`sentence` TEXT)")
	if err != nil {
		log.Println("prepare create sentences " + err.Error())
		return nil, err
	}
	res, err := query.Exec()
	if err != nil {
		log.Println("exec create sentences " + err.Error())
		return nil, err
	}
	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())

	query, err = db.Prepare("CREATE TABLE `board` (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`word` TEXT)")
	if err != nil {
		log.Println("prepare create board " + err.Error())
		return nil, err
	}
	res, err = query.Exec()
	if err != nil {
		log.Println("exec create board " + err.Error())
		return nil, err
	}
	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())
	return db, nil
}

func NewDatabase() *Database {
	db, err := createTables()
	if err != nil {
		log.Println("createTables: " + err.Error())
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

func (d *Database) SetBoard(words []string) {
	if d == nil || d.db == nil {
		return
	}
	if len(words) == 0 {
		return
	}
	for _, w := range words {
		newWord := word{
			text: w,
		}
		stmt, err := d.db.Prepare("INSERT INTO board(id, word) VALUES (?,?)")
		if err != nil {
			log.Println(err.Error())
		}
		res, err := stmt.Exec(nil, newWord.text)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(fmt.Sprint(res.LastInsertId()))
		log.Println(res.RowsAffected())
		log.Println("Added " + newWord.text)
	}
}

func (d *Database) GetBoard() []string {
	// return []string{
	// 	"Anticoncepcional",
	// 	"Boca",
	// 	"Coração",
	// 	"Saúde Mental",
	// 	"Sexo",
	// 	"Vulva",
	// }
	if d == nil || d.db == nil {
		return []string{}
	}
	rows, err := d.db.Query("SELECT * FROM board")
	if rows.Err() != nil {
		log.Println(rows.Err())
		return []string{}
	}
	if err != nil {
		log.Println(err.Error())
		return []string{}
	}
	boardWords := []string{}
	for rows.Next() {
		w := word{}
		err = rows.Scan(&w.id, &w.text)
		if err != nil {
			log.Fatal(err)
			return boardWords
		}
		boardWords = append(boardWords, w.text)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
	}
	err = rows.Close()

	if err != nil {
		return []string{}
	}
	return boardWords
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
	dbSentences := []string{}
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
