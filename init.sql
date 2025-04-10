-- GÉNEROS
CREATE TABLE
  IF NOT EXISTS genders (
    gender_id INTEGER PRIMARY KEY AUTOINCREMENT,
    descripcion TEXT NOT NULL UNIQUE
  );

INSERT INTO
  genders (descripcion)
VALUES
  ('DRAMA'),
  ('COMEDIA'),
  ('ACCION'),
  ('TERROR'),
  ('ROMANCE'),
  ('CIENCIA_FICCION'),
  ('FANTASIA'),
  ('DOCUMENTAL');

-- MOVIES (con cover y url como BLOBs)
CREATE TABLE
  IF NOT EXISTS movies (
    movie_id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    cover BLOB NOT NULL,
    video BLOB,
    gender_id INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES genders (gender_id)
  );

-- TIPOS DE CONTENIDO (Series, Animes, etc.)
CREATE TABLE
  IF NOT EXISTS content_types (
    content_type_id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    cover BLOB NOT NULL,
    gender_id INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES genders (gender_id)
  );

-- TEMPORADAS DE SERIES O ANIMES
CREATE TABLE
  IF NOT EXISTS seasons (
    season_id INTEGER PRIMARY KEY AUTOINCREMENT,
    content_type_id INTEGER NOT NULL,
    season_number INTEGER NOT NULL,
    title TEXT NOT NULL,
    FOREIGN KEY (content_type_id) REFERENCES content_types (content_type_id)
  );

-- EPISODIOS DE TEMPORADAS
CREATE TABLE
  IF NOT EXISTS episodes (
    episode_id INTEGER PRIMARY KEY AUTOINCREMENT,
    season_id INTEGER NOT NULL,
    episode_number INTEGER NOT NULL,
    title TEXT NOT NULL,
    video BLOB,
    FOREIGN KEY (season_id) REFERENCES seasons (season_id)
  );