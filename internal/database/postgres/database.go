// Package postgres provides data storage in a postgres database
package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func Conn(driverName, dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn can not be missing")
	}

	if driverName == "" {
		return nil, fmt.Errorf("driver name can not be missing")
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return db, err
	}

	log.Println("Connect to database")

	return db, nil
}

func SetUpDataBase(ctx context.Context, db *sql.DB) error {

	sqlCreateDB := `CREATE TABLE IF NOT EXISTS urls (
								id serial PRIMARY KEY,
								user_id VARCHAR NOT NULL, 	
								origin_url VARCHAR NOT NULL, 
								short_url VARCHAR NOT NULL UNIQUE,
                                is_deleted BOOLEAN NOT NULL DEFAULT FALSE
					);`
	res, err := db.ExecContext(ctx, sqlCreateDB)

	log.Println("Create table", err, res)

	return nil
}
