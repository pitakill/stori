-- users
-- name: CreateUser :one
INSERT INTO users (
  id, first_name, last_name, email
) VALUES (
  ?,?,?,?
)
RETURNING *;
-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ? LIMIT 1;
-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;
-- name: GetUserByAccountID :one
SELECT * FROM users
LEFT JOIN accounts
ON users.ID = accounts.user_id
WHERE accounts.ID = ? LIMIT 1;

-- account
-- name: CreateAccount :one
INSERT INTO accounts (
  id, user_id, bank, number
) VALUES (
  ?,?,?,?
)
RETURNING *;
-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;
-- name: GetAccountByUserID :one
SELECT * FROM accounts
WHERE user_id = ? LIMIT 1;

-- transactions
-- name: CreateTransaction :one
INSERT INTO transactions (
  id, account_id, date, credit, amount
) VALUES (
  ?,?,?,?,?
)
RETURNING *;
-- name: GetTransactionsByAccountID :many
SELECT * FROM transactions
WHERE account_id = ?;
