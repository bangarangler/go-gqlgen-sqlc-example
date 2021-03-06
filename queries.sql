-- name: GetAgent :one
SELECT * FROM agents
WHERE id = $1;

-- name: ListAgents :many
SELECT * FROM agents
ORDER BY name;

-- name: CreateAgent :one
INSERT INTO agents (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateAgent :one
UPDATE agents
SET name = $2, email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAgent :one
DELETE FROM agents
WHERE id = $1
RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name, website, agent_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, website = $3, agent_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :one
DELETE FROM authors
WHERE id = $1
RETURNING *;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: CreateBook :one
INSERT INTO books (title, description, cover)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET title = $2, description = $3, cover = $4
WHERE id = $1
RETURNING *;

-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING *;

-- name: SetBookAuthor :exec
INSERT INTO book_authors (book_id, author_id)
VALUES ($1, $2);

-- name: UnsetBookAuthors :exec
DELETE FROM book_authors
WHERE book_id = $1;

-- -- name: ListAuthorsByAgentID :many
-- SELECT authors.* FROM authors, agents
-- WHERE agents.id = authors.agent_id AND authors.agent_id = $1;

-- name: ListBooksByAuthorID :many
SELECT books.* FROM books, book_authors
WHERE books.id = book_authors.book_id AND book_authors.author_id = $1;

-- -- name: ListAuthorsByBookID :many
-- SELECT authors.* FROM authors, book_authors
-- WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1;

-- name: ListAgentsByAuthorIDs :many
SELECT agents.*, authors.id AS author_id FROM agents, authors
WHERE agents.id = authors.agent_id AND authors.id = ANY($1::bigint[]);

-- name: ListAuthorsByAgentIDs :many
SELECT authors.* FROM authors, agents
WHERE authors.agent_id = agents.id AND agents.id = ANY($1::bigint[]);

-- name: ListAuthorsByBookIDs :many
SELECT authors.*, book_authors.book_id FROM authors, book_authors
WHERE book_authors.author_id = authors.id AND book_authors.book_id = ANY($1::bigint[]);

-- name: ListBooksByAuthorIDs :many
SELECT books.*, book_authors.author_id FROM books, book_authors
WHERE book_authors.book_id = books.id AND book_authors.author_id = ANY($1::bigint[]);
