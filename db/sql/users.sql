-- name: CreateUser :exec
INSERT INTO users (email, password, created_at, role)
VALUES 
  ($1, $2, NOW(), $3);