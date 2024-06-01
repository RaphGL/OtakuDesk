CREATE TABLE IF NOT EXISTS users (
    username TEXT NOT NULL UNIQUE PRIMARY KEY,
    email TEXT NOT NULL,
    password BLOB NOT NULL,
    profile_pic BLOB
);

CREATE TABLE IF NOT EXISTS manga (
    name TEXT NOT NULL,  
    path TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS anime (
    name TEXT NOT NULL,  
    path TEXT NOT NULL
);
