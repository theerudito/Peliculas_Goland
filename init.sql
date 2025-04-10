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

-- PELICULAS (con cover y url como BLOBs)
CREATE TABLE
  IF NOT EXISTS movies (
    movie_id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    cover TEXT NOT NULL,
    url TEXT NOT NULL,
    gender_id INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES genders (gender_id)
  );

INSERT INTO
  movies (title, year, cover, url, gender_id)
VALUES
  (
    'MAN OF HONOR',
    2000,
    'https://pics.filmaffinity.com/men_of_honor-584186623-large.jpg',
    './videos/Cap_1.mp4',
    1
  );

-- TIPOS DE CONTENIDO (Series, Animes, etc.)
CREATE TABLE
  IF NOT EXISTS content_types (
    content_type_id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    descripcion TEXT NOT NULL,
    cover TEXT NOT NULL,
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
    url TEXT NOT NULL,
    FOREIGN KEY (season_id) REFERENCES seasons (season_id)
  );