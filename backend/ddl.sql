CREATE DATABASE movies;

-- USER
CREATE TABLE
    login
(
    user_id  INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- GENEROS
CREATE TABLE
    gender
(
    gender_id   INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    gender_name TEXT NOT NULL UNIQUE
);

-- TEMPORADAS
CREATE TABLE
    season
(
    season_id   INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    season_name TEXT NOT NULL
);

-- PELÍCULAS
CREATE TABLE
    movie
(
    movie_id    INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    movie_title TEXT    NOT NULL,
    movie_year  INTEGER NOT NULL,
    movie_cover TEXT    NOT NULL,
    movie_url   TEXT    NOT NULL,
    gender_id   INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES gender (gender_id)
);

-- TIPOS DE CONTENIDO (Series, Animes, etc.)
CREATE TABLE
    content_type
(
    content_id    INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    content_title TEXT    NOT NULL,
    content_type  INTEGER NOT NULL,
    content_cover TEXT    NOT NULL,
    content_year  INTEGER NOT NULL,
    gender_id     INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES gender (gender_id)
);

-- EPISODIOS DE TEMPORADAS
CREATE TABLE
    episode
(
    episode_id     INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    episode_number INTEGER NOT NULL,
    episode_name   TEXT    NOT NULL,
    episode_url    TEXT    NOT NULL,
    season_id      INTEGER NOT NULL,
    content_id     INTEGER NOT NULL,
    FOREIGN KEY (season_id) REFERENCES season (season_id),
    FOREIGN KEY (content_id) REFERENCES content_type (content_id)
);
