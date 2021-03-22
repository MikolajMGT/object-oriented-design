package manager

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type Manager struct {
	db *sql.DB
}

var instance *Manager

func GetInstance() *Manager {
	if instance == nil {
		instance = &Manager{}
		instance.Start()
	}

	return instance
}

func (m *Manager) Start()  {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Panic().Err(err).Msg("unable to open database")
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, email TEXT UNIQUE NOT NULL, username TEXT NOT NULL, password TEXT NOT NULL)")
	if err != nil {
		log.Panic().Err(err).Msg("error while initialising database")
	}


	instance.db = db
}

func (m *Manager) Search(query string) (*sql.Rows, error) {
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, err
}

func (m *Manager) Query(query string) (int64, error) {
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return -1, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Err(err).Msg("error while closing previous statement")
		}
	}()

	res, err := stmt.Exec()
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *Manager) Stop()  {
	if err := m.db.Close(); err != nil {
		log.Panic().Err(err).Msg("error while closing database")
	}
}