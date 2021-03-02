package models

import (
	"database/sql"
	"time"
)

var DB *sql.DB

type Client struct {
	ID int
	First_name string
	Last_name string
	Ci string
	Birthday time.Time
}

type TotalClient struct {
	ClientTitle string
	Clients []Client
}

func AllClient() ([]Client, error) {
	rows, err := DB.Query("SELECT id, first_name, last_name, ci, birthday FROM clients LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cls []Client
	for rows.Next() {
		var cl Client
		err := rows.Scan(&cl.ID, &cl.First_name, &cl.Last_name, &cl.Ci, &cl.Birthday)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cls, nil
}