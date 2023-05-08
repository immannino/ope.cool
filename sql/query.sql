-- name: GetListenCount :one
SELECT count(listen_id) FROM listen;

-- name: GetLatestListen :one
SELECT * FROM listen ORDER BY listen_id DESC LIMIT 1;

-- name: GetLastNListens :many
SELECT * FROM listen ORDER BY listen_id DESC LIMIT ?;