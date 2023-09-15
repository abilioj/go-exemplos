package models

import (
	"api-postgres/db"
	"log"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	sql := `INSERT INTO todos (title,description,done) VALUES ($1,$2,$3) RETURNING id`

	// err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	stmt, err := conn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Query(todo.Title, todo.Description, todo.Done)

	log.Println(result)

	for result.Next() {
		err := result.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}
