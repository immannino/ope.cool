-- name: GetListenCount :one
SELECT count(listen_id) FROM listen;

-- name: GetLatestListen :one
SELECT * FROM listen ORDER BY listen_id DESC LIMIT 1;

-- name: GetLastNListens :many
SELECT * FROM listen ORDER BY listen_id DESC LIMIT ?;

-- name: GetTopNListens :many
SELECT spotify_song_id, artist, track_name, COUNT(spotify_song_id) FROM listen GROUP BY spotify_song_id, artist, track_name ORDER BY COUNT(spotify_song_id) DESC LIMIT ?;