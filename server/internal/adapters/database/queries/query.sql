-- name: listMessages :many
SELECT * FROM message
ORDER BY id;

-- name: persistMessage :one
INSERT INTO message (client_id, topic, payload, timestamp)
VALUES (?, ?, ?, ?) RETURNING id;
