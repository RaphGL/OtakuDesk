CREATE TABLE IF NOT EXISTS users (
    username TEXT NOT NULL UNIQUE PRIMARY KEY,
    email TEXT NOT NULL,
    password BLOB NOT NULL,
    profile_pic BLOB
);

CREATE TABLE IF NOT EXISTS manga (
    name TEXT NOT NULL,  
    path TEXT NOT NULL,
    episode INTEGER,
    curr_episode INTEGER,
    total_episodes INTEGER,
    chapter INTEGER, 
    curr_chapter INTEGER,
    total_chapters INTEGER
);

CREATE TABLE IF NOT EXISTS anime (
    name TEXT NOT NULL,  
    path TEXT NOT NULL,
    episode INTEGER,
    curr_episode INTEGER,
    total_episodes INTEGER,
    chapter INTEGER, 
    curr_chapter INTEGER,
    total_chapters INTEGER
);
