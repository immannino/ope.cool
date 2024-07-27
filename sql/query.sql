-- name: GetListenCount :one
SELECT count(listen_id) FROM listen;

-- name: GetLatestListen :one
SELECT * FROM listen ORDER BY listen_id DESC LIMIT 1;

-- name: GetLastNListens :many
SELECT * FROM listen ORDER BY listen_id DESC LIMIT $1;

-- name: GetTopNListens :many
SELECT spotify_id, artist, track_name, COUNT(spotify_id) FROM listen GROUP BY spotify_id, artist, track_name ORDER BY COUNT(spotify_id) DESC LIMIT $1;

-- -- name: GetTopUniqueMonthlyListens :many
-- SELECT * FROM listen WHERE listened_at BETWEEN (CURDATE() - INTERVAL 1 MONTH ) and CURDATE() ORDER BY listened_at DESC LIMIT $1;

-- name: GetUniqueMonthlyListenCount :one
-- SELECT COUNT(*) FROM listen WHERE listened_at BETWEEN (CURDATE() - INTERVAL 1 MONTH ) and CURDATE() ORDER BY listened_at DESC;