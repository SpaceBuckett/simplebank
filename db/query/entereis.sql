-- name: CreateEntry :one
INSERT INTO entereis(
                     account_id,
                     amount
) VALUES (
          $1, $2
         ) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entereis
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entereis
WHERE
    account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;