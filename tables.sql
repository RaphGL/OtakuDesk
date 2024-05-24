CREATE TABLE IF NOT EXISTS users (
    username TEXT NOT NULL UNIQUE PRIMARY KEY,
    email TEXT NOT NULL,
    password BLOB NOT NULL,
    profile_pic BLOB
);
