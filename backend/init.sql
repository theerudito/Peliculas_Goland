-- GÉNEROS
CREATE TABLE
  IF NOT EXISTS genders (
    gender_id INTEGER PRIMARY KEY AUTOINCREMENT,
    gender_name TEXT NOT NULL UNIQUE
  );

INSERT INTO
  genders (gender_name)
VALUES
  ('DRAMA'),
  ('COMEDIA'),
  ('ACCION'),
  ('TERROR'),
  ('ROMANCE'),
  ('CIENCIA_FICCION'),
  ('FANTASIA'),
  ('DOCUMENTAL');

-- PELÍCULAS (con cover y url como BLOBs)
CREATE TABLE
  IF NOT EXISTS movies (
    movie_movie_id INTEGER PRIMARY KEY AUTOINCREMENT,
    movie_title TEXT NOT NULL,
    movie_year INTEGER NOT NULL,
    movie_cover TEXT NOT NULL,
    movie_url TEXT NOT NULL,
    gender_id INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES genders (gender_id)
  );

INSERT INTO
  movies (
    movie_title,
    movie_year,
    movie_cover,
    movie_url,
    gender_id
  )
VALUES
  (
    'MAN OF HONOR',
    2000,
    'https://m.media-amazon.com/images/I/91M9JytQi2L._AC_UF894,1000_QL80_.jpg',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Medias%2FVideos%2FPete%20Bellis%20%26%20Tommy%20-%20You%20Should%20Know%20(Costa%20Mee%20Remix).mp4?alt=media&token=c1348dc4-ba43-4af5-9f65-1880584b7dd1',
    1
  ),
  (
    'THE SHAWSHANK REDEMPTION',
    1994,
    'https://cuevana.biz/_next/image?url=https%3A%2F%2Fimage.tmdb.org%2Ft%2Fp%2Foriginal%2FuRRTV7p6l2ivtODWJVVAMRrwTn2.jpg&w=640&q=75',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Medias%2FVideos%2FPete%20Bellis%20%26%20Tommy%20-%20You%20Should%20Know%20(Costa%20Mee%20Remix).mp4?alt=media&token=c1348dc4-ba43-4af5-9f65-1880584b7dd1',
    2
  ),
  (
    'GLADIATOR',
    2000,
    'https://cuevana.biz/_next/image?url=https%3A%2F%2Fimage.tmdb.org%2Ft%2Fp%2Foriginal%2Fo6XhzKghQFliN49iE4M4RX94PTB.jpg&w=640&q=75',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Medias%2FVideos%2FPete%20Bellis%20%26%20Tommy%20-%20You%20Should%20Know%20(Costa%20Mee%20Remix).mp4?alt=media&token=c1348dc4-ba43-4af5-9f65-1880584b7dd1',
    3
  );

-- TIPOS DE CONTENIDO (Series, Animes, etc.)
CREATE TABLE
  IF NOT EXISTS content_types (
    content_id INTEGER PRIMARY KEY AUTOINCREMENT,
    content_type INTEGER NOT NULL,
    content_cover TEXT NOT NULL,
    content_url TEXT NOT NULL,
    content_year TEXT NOT NULL,
    gender_id INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES genders (gender_id)
  );

INSERT INTO
  content_types (
    content_type,
    content_cover,
    content_url,
    content_year,
    gender_id
  )
VALUES
  (
    1,
    'JUJUTSU KAISEN',
    'https://latanime.org/thumbs/imagen/jujutsu-kaisen-s2-latino-1690490946.jpg',
    2023,
    7
  ),
  (
    2,
    'VIKING',
    'https://latanime.org/thumbs/imagen/jujutsu-kaisen-s2-castellano-1692305959.jpg',
    2024,
    1
  );

-- TEMPORADAS DE SERIES O ANIMES
CREATE TABLE
  IF NOT EXISTS seasons (
    season_id INTEGER PRIMARY KEY AUTOINCREMENT,
    content_id INTEGER NOT NULL,
    season_name INTEGER NOT NULL,
    FOREIGN KEY (content_id) REFERENCES content_types (content_id)
  );

INSERT INTO
  seasons (content_id, season_name)
VALUES
  (1, 'Temporada 1'),
  (2, 'Temporada 2');

-- EPISODIOS DE TEMPORADAS
CREATE TABLE
  IF NOT EXISTS episodes (
    episode_id INTEGER PRIMARY KEY AUTOINCREMENT,
    season_id INTEGER NOT NULL,
    episode_number INTEGER NOT NULL,
    episode_name TEXT NOT NULL,
    episode_url TEXT NOT NULL,
    FOREIGN KEY (season_id) REFERENCES seasons (season_id)
  );

INSERT INTO
  episodes (
    season_id,
    episode_number,
    episode_name,
    episode_url
  )
VALUES
  (
    1,
    1,
    'Capitulo 1',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_1.mp4?alt=media&token=6e063fc6-29c2-4522-be07-2e45cdbd81e8'
  ),
  (
    2,
    1,
    'Capitulo 1',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2Fvideo.mp4?alt=media&token=5f0bae50-f142-4052-ab75-7456b4d513eb'
  ),
  (
    2,
    2,
    'Capitulo 2',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_2.mp4?alt=media&token=7c1b4c8e-5b78-48e8-ab4f-e4b43571998a'
  );