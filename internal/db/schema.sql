CREATE DATABASE medods_db;

CREATE TABLE users (uid SERIAL PRIMARY KEY);

CREATE TABLE refresh_tokens (id SERIAL PRIMARY KEY, user_id INTEGER NOT NULL REFERENCES users(uid), token TEXT, is_active BOOLEAN NOT NULL);

INSERT INTO users(uid) VALUES (1);
