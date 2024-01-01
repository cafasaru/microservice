--------- USERS SAVE ------------------

-- name: SaveUser :one
INSERT INTO users (
    user_id,
    email,
    username,
    password_hash,
    verification_code,
    verification_expiration,
    verified,
    active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users 
WHERE user_id = $1;

-- name: GetUserByUsername :one
SELECT *
FROM users 
WHERE UPPER(username) = UPPER($1) LIMIT 1;
