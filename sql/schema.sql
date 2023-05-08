CREATE TABLE IF NOT EXISTS listen (
        listen_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
        feed_slug VARCHAR(255),
        spotify_song_id VARCHAR(255),
        spotify_href VARCHAR(255),
        spotify_uri VARCHAR(255),
        artist VARCHAR(255),
        artist_genres VARCHAR(255),
        track_name VARCHAR(255),
        album_name VARCHAR(255),
        album_image VARCHAR(255),
        album_image_height INTEGER,
        album_image_width INTEGER,
        listened_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS jsonfeed (
        slug        VARCHAR(255) NOT NULL PRIMARY KEY UNIQUE,
        url         VARCHAR(255) NOT NULL,
        title       VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        item_count  INT NOT NULL,
        createdAt   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updatedAt   DATETIME NOT NULL
);
