// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package pg

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createAgent = `-- name: CreateAgent :one
INSERT INTO agents (name, email)
VALUES ($1, $2)
RETURNING id, name, email
`

type CreateAgentParams struct {
	Name  string
	Email string
}

func (q *Queries) CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error) {
	row := q.db.QueryRowContext(ctx, createAgent, arg.Name, arg.Email)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, website, agent_id)
VALUES ($1, $2, $3)
RETURNING id, name, website, agent_id
`

type CreateAuthorParams struct {
	Name    string
	Website sql.NullString
	AgentID int64
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Website, arg.AgentID)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}

const createBook = `-- name: CreateBook :one
INSERT INTO books (title, description, cover)
VALUES ($1, $2, $3)
RETURNING id, title, description, cover
`

type CreateBookParams struct {
	Title       string
	Description string
	Cover       string
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook, arg.Title, arg.Description, arg.Cover)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Cover,
	)
	return i, err
}

const deleteAgent = `-- name: DeleteAgent :one
DELETE FROM agents
WHERE id = $1
RETURNING id, name, email
`

func (q *Queries) DeleteAgent(ctx context.Context, id int64) (Agent, error) {
	row := q.db.QueryRowContext(ctx, deleteAgent, id)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :one
DELETE FROM authors
WHERE id = $1
RETURNING id, name, website, agent_id
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, deleteAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING id, title, description, cover
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, deleteBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Cover,
	)
	return i, err
}

const getAgent = `-- name: GetAgent :one
SELECT id, name, email FROM agents
WHERE id = $1
`

func (q *Queries) GetAgent(ctx context.Context, id int64) (Agent, error) {
	row := q.db.QueryRowContext(ctx, getAgent, id)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, website, agent_id FROM authors
WHERE id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}

const getBook = `-- name: GetBook :one
SELECT id, title, description, cover FROM books
WHERE id = $1
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Cover,
	)
	return i, err
}

const listAgents = `-- name: ListAgents :many
SELECT id, name, email FROM agents
ORDER BY name
`

func (q *Queries) ListAgents(ctx context.Context) ([]Agent, error) {
	rows, err := q.db.QueryContext(ctx, listAgents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Agent
	for rows.Next() {
		var i Agent
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAgentsByAuthorIDs = `-- name: ListAgentsByAuthorIDs :many
SELECT agents.id, agents.name, agents.email, authors.id AS author_id FROM agents, authors
WHERE agents.id = authors.agent_id AND authors.id = ANY($1::bigint[])
`

type ListAgentsByAuthorIDsRow struct {
	ID       int64
	Name     string
	Email    string
	AuthorID int64
}

func (q *Queries) ListAgentsByAuthorIDs(ctx context.Context, dollar_1 []int64) ([]ListAgentsByAuthorIDsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAgentsByAuthorIDs, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAgentsByAuthorIDsRow
	for rows.Next() {
		var i ListAgentsByAuthorIDsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.AuthorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, website, agent_id FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsByAgentID = `-- name: ListAuthorsByAgentID :many
SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, agents
WHERE agents.id = authors.agent_id AND authors.agent_id = $1
`

func (q *Queries) ListAuthorsByAgentID(ctx context.Context, agentID int64) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsByAgentID, agentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsByBookID = `-- name: ListAuthorsByBookID :many
SELECT authors.id, authors.name, authors.website, authors.agent_id FROM authors, book_authors
WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1
`

func (q *Queries) ListAuthorsByBookID(ctx context.Context, bookID int64) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsByBookID, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Website,
			&i.AgentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooks = `-- name: ListBooks :many
SELECT id, title, description, cover FROM books
ORDER BY title
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Cover,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooksByAuthorID = `-- name: ListBooksByAuthorID :many
SELECT books.id, books.title, books.description, books.cover FROM books, book_authors
WHERE books.id = book_authors.book_id AND book_authors.author_id = $1
`

func (q *Queries) ListBooksByAuthorID(ctx context.Context, authorID int64) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooksByAuthorID, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Cover,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setBookAuthor = `-- name: SetBookAuthor :exec
INSERT INTO book_authors (book_id, author_id)
VALUES ($1, $2)
`

type SetBookAuthorParams struct {
	BookID   int64
	AuthorID int64
}

func (q *Queries) SetBookAuthor(ctx context.Context, arg SetBookAuthorParams) error {
	_, err := q.db.ExecContext(ctx, setBookAuthor, arg.BookID, arg.AuthorID)
	return err
}

const unsetBookAuthors = `-- name: UnsetBookAuthors :exec
DELETE FROM book_authors
WHERE book_id = $1
`

func (q *Queries) UnsetBookAuthors(ctx context.Context, bookID int64) error {
	_, err := q.db.ExecContext(ctx, unsetBookAuthors, bookID)
	return err
}

const updateAgent = `-- name: UpdateAgent :one
UPDATE agents
SET name = $2, email = $3
WHERE id = $1
RETURNING id, name, email
`

type UpdateAgentParams struct {
	ID    int64
	Name  string
	Email string
}

func (q *Queries) UpdateAgent(ctx context.Context, arg UpdateAgentParams) (Agent, error) {
	row := q.db.QueryRowContext(ctx, updateAgent, arg.ID, arg.Name, arg.Email)
	var i Agent
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, website = $3, agent_id = $4
WHERE id = $1
RETURNING id, name, website, agent_id
`

type UpdateAuthorParams struct {
	ID      int64
	Name    string
	Website sql.NullString
	AgentID int64
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor,
		arg.ID,
		arg.Name,
		arg.Website,
		arg.AgentID,
	)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Website,
		&i.AgentID,
	)
	return i, err
}

const updateBook = `-- name: UpdateBook :one
UPDATE books
SET title = $2, description = $3, cover = $4
WHERE id = $1
RETURNING id, title, description, cover
`

type UpdateBookParams struct {
	ID          int64
	Title       string
	Description string
	Cover       string
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Cover,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Cover,
	)
	return i, err
}
