package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DataStore struct{
	db *sql.DB

}

func createDataStore(path string) (ds *DataStore, err error){
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
		return
	}
	ds = &DataStore{db}

	// Create tables if they don't exist
	ds.createSchema()
	return
}

func (ds *DataStore) createSchema() (err error){
	dateTable := `
        CREATE TABLE IF NOT EXISTS dates (
          id INTEGER NOT NULL PRIMARY KEY,
          date TEXT,
          rev INTEGER)
        `

	dataTable := `
        CREATE TABLE IF NOT EXISTS data (
          id INTEGER NOT NULL PRIMARY KEY,
          rev INTEGER,
          name TEXT,
          downloads INTEGER)
        `
	// create tables in a transaction
	tx, err := ds.db.Begin()
	if err != nil {
		return
	}
	
	_, err = tx.Exec(dateTable)
	if err != nil {
		return
	}

	_, err = tx.Exec(dataTable)
	if err != nil {
		return
	}
	
	err = tx.Commit()
	if err != nil {
		return
	}

	return
	
}

func (ds *DataStore) GetLatestRecords() (result []Record, err error){
	rows, err := ds.db.Query("SELECT rev, name, downloads FROM data WHERE rev = (SELECT MAX(rev) FROM data) ORDER BY downloads DESC;")
	if err != nil {
		return
	}

	for rows.Next() {
		var rev, downloads int
		var name string
		rows.Scan(&rev, &name, &downloads)

		result = append(result, Record{rev, name, downloads})
	}
	return
}

func (ds *DataStore) InsertRecord(rec Record) (err error){
	stmt, err := ds.db.Prepare("INSERT INTO data(rev, name, downloads) VALUES(?, ?, ?)")
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(rec.Revision, rec.Name, rec.Downloads)
	if err != nil {
		return
	}

	return
}

func (ds *DataStore) Close() {
	ds.db.Close()
}
