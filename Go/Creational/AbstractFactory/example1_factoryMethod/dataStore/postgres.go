package dataStore

import (
	"database/sql"
	"fmt"
	"errors"
	"log"
	_ "github.com/lib/pq"
)

type PostgreSQLDataStore struct {
	DSN string
	Table string
	UserNameField string
	IDField string
	DB *sql.DB
}

func (pds *PostgreSQLDataStore) Name() string {
	return "PostgreDataStore"
}

func (pds *PostgreSQLDataStore) FindUserNameById(id int64) (string, error){
	var username string
	q_s := fmt.Sprintf("SELECT %s FROM %s WHERE %s=%v", pds.UserNameField, pds.Table, pds.IDField, id)
	rows, err := pds.DB.Query(q_s)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", UserNotFoundError
		}
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			if err == sql.ErrNoRows {
				return "", UserNotFoundError
			}
			return "", err
		}
	}
	return username, nil
}

func NewPostgreSQLDataStore(conf Configuration) (DataStore, error) {
	dsn := conf.Get("DATASTORE_POSTGRES_DSN", "")
	if dsn == "" {
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DATASTORE_POSTGRES_DSN"))
	}
	table := conf.Get("DB_TABLE", "")
	if table == "" {
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DB_TABLE"))
	}
	usernameField := conf.Get("DB_USERNAME_FIELD", "")
	if usernameField == "" {
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DB_USERNAME_FIELD"))
	}
	idField := conf.Get("DB_ID_FIELD", "")
	if idField == "" {
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DB_ID_FIELD"))
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf("Failed to connect to datastore: %s", err.Error())
		return nil, errors.New("Unnable connect to postresql database")
	}

	return &PostgreSQLDataStore{
		DSN:            dsn,
		DB:             db,
		Table:          table,
		IDField:        idField,
		UserNameField:  usernameField,
	}, nil
}
