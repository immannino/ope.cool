CREATE TABLE IF NOT EXISTS listen (
        listen_id SERIAL PRIMARY KEY,
        spotify_id VARCHAR(255),
        spotify_href VARCHAR(255),
        spotify_uri VARCHAR(255),
        artist VARCHAR(255),
        artist_genres VARCHAR(255),
        track_name VARCHAR(255),
        album_name VARCHAR(255),
        album_image VARCHAR(255),
        album_image_height INTEGER,
        album_image_width INTEGER,
        html TEXT,
        listened_at TIMESTAMP DEFAULT NOW()
);