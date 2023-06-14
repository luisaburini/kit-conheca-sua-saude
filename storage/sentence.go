package storage

import (
	_ "github.com/mattn/go-sqlite3"
)

type sentences struct {
	id       int
	sentence string
}

type word struct {
	id   int
	text string
}
