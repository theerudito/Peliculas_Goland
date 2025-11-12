CREATE DATABASE peliculas;

CREATE TABLE login
(
    user_id  SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    logo     TEXT
);

CREATE TABLE gender
(
    gender_id   SERIAL PRIMARY KEY,
    gender_name TEXT NOT NULL UNIQUE
);

CREATE TABLE season
(
    season_id   SERIAL PRIMARY KEY,
    season_name TEXT NOT NULL
);

CREATE TABLE movie
(
    movie_id    SERIAL PRIMARY KEY,
    movie_title TEXT    NOT NULL,
    movie_year  INTEGER NOT NULL,
    movie_cover TEXT    NOT NULL,
    movie_url   TEXT    NOT NULL,
    gender_id   INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES gender (gender_id)
);

CREATE TABLE content_type
(
    content_id    SERIAL PRIMARY KEY,
    content_title TEXT    NOT NULL,
    content_type  INTEGER NOT NULL,
    content_cover TEXT    NOT NULL,
    content_year  INTEGER NOT NULL,
    gender_id     INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES gender (gender_id)
);

CREATE TABLE episode
(
    episode_id     SERIAL PRIMARY KEY,
    episode_number INTEGER NOT NULL,
    episode_name   TEXT    NOT NULL,
    episode_url    TEXT    NOT NULL,
    season_id      INTEGER NOT NULL,
    content_id     INTEGER NOT NULL,
    FOREIGN KEY (season_id) REFERENCES season (season_id),
    FOREIGN KEY (content_id) REFERENCES content_type (content_id)
);