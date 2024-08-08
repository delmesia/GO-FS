package models

import (
	"database/sql"
	"time"
)

// this Snippet struct will hold the data for each individual snippet.
// the struct should correspond to the fields in the MySQL snippets table.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel struct type will wrap a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// This method will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	// this is the SQL statement that I want to execute.
	// Splitted into to lines for readability.
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Exec() method on the embedded connection pool will execute the statement.
	// First parameter is the SQL statement, followed by the title, content, and
	// expiry values for placeholder parameters. This method will
	// return a sql.Result type, which contains some basic information about
	// what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
}

// this will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// this will return the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
