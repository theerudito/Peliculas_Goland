-- GENEROS
CREATE TABLE
  IF NOT EXISTS gender (
    gender_id INTEGER PRIMARY KEY AUTOINCREMENT,
    gender_name TEXT NOT NULL UNIQUE
  );

INSERT INTO
  gender (gender_name)
VALUES
  ('DRAMA'),
  ('COMEDIA'),
  ('ACCION'),
  ('TERROR'),
  ('ROMANCE'),
  ('CIENCIA_FICCION'),
  ('FANTASIA'),
  ('DOCUMENTAL'),
  ('ANIMACION'),
  ('AVENTURA'),
  ('MISTERIO'),
  ('MUSICAL'),
  ('CRIMEN'),
  ('THRILLER'),
  ('HISTORIA'),
  ('GUERRA'),
  ('FAMILIAR'),
  ('BIOGRAFIA'),
  ('WESTERN'),
  ('DEPORTES'),
  ('NOIR'),
  ('SUSPENSO'),
  ('EROTICO'),
  ('CORTO'),
  ('REALITY'),
  ('INFANTIL'),
  ('SUPERNATURAL'),
  ('POLITICO'),
  ('RELIGIOSO'),
  ('EXPERIMENTAL');

-- TEMPORADAS
CREATE TABLE
  IF NOT EXISTS season (
    season_id INTEGER PRIMARY KEY AUTOINCREMENT,
    season_name TEXT NOT NULL
  );

INSERT INTO
  season (season_name)
VALUES
  ('TEMPORADA 1'),
  ('TEMPORADA 2'),
  ('TEMPORADA 3'),
  ('TEMPORADA 4'),
  ('TEMPORADA 5'),
  ('TEMPORADA 6'),
  ('TEMPORADA 7'),
  ('TEMPORADA 8'),
  ('TEMPORADA 9'),
  ('TEMPORADA 10'),
  ('TEMPORADA 11'),
  ('TEMPORADA 12'),
  ('TEMPORADA 13'),
  ('TEMPORADA 14'),
  ('TEMPORADA 15');

-- PELÍCULAS 
CREATE TABLE
  IF NOT EXISTS movie (
    movie_id INTEGER PRIMARY KEY AUTOINCREMENT,
    movie_title TEXT NOT NULL,
    movie_year INTEGER NOT NULL,
    movie_cover TEXT NOT NULL,
    movie_url TEXT NOT NULL,
    gender_id INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES genders (gender_id)
  );

INSERT INTO
  movie (
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
  IF NOT EXISTS content_type (
    content_id INTEGER PRIMARY KEY AUTOINCREMENT,
    content_title TEXT NOT NULL,
    content_type INTEGER NOT NULL,
    content_cover TEXT NOT NULL,
    content_year INTEGER NOT NULL,
    gender_id INTEGER NOT NULL,
    FOREIGN KEY (gender_id) REFERENCES genders (gender_id)
  );

INSERT INTO
  content_type (
    content_title,
    content_type,
    content_cover,
    content_year,
    gender_id
  )
VALUES
  (
    'JUJUTSU KAISEN',
    1,
    'https://latanime.org/thumbs/imagen/jujutsu-kaisen-s2-castellano-1692305959.jpg',
    2023,
    1
  ),
  (
    'LOS 100',
    2,
    'https://mediaproxy.tvtropes.org/width/1200/https://static.tvtropes.org/pmwiki/pub/images/the100.png',
    2010,
    1
  );

-- EPISODIOS DE TEMPORADAS
CREATE TABLE
  IF NOT EXISTS episode (
    episode_id INTEGER PRIMARY KEY AUTOINCREMENT,
    episode_number INTEGER NOT NULL,
    episode_name TEXT NOT NULL,
    episode_url TEXT NOT NULL,
    season_id INTEGER NOT NULL,
    content_id INTEGER NOT NULL,
    FOREIGN KEY (season_id) REFERENCES seasons (season_id),
    FOREIGN KEY (content_id) REFERENCES content_types (content_id)
  );

INSERT INTO
  episode (
    episode_number,
    episode_name,
    episode_url,
    season_id,
    content_id
  )
VALUES
  (
    1,
    'Capitulo 1',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_1.mp4?alt=media&token=6e063fc6-29c2-4522-be07-2e45cdbd81e8',
    1,
    1
  ),
  (
    2,
    'Capitulo 2',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2Fvideo.mp4?alt=media&token=5f0bae50-f142-4052-ab75-7456b4d513eb',
    1,
    1
  ),
  (
    3,
    'Capitulo 3',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_2.mp4?alt=media&token=7c1b4c8e-5b78-48e8-ab4f-e4b43571998a',
    1,
    1
  ),
  (
    1,
    'Capitulo 1',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_1.mp4?alt=media&token=6e063fc6-29c2-4522-be07-2e45cdbd81e8',
    1,
    2
  ),
  (
    2,
    'Capitulo 2',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2Fvideo.mp4?alt=media&token=5f0bae50-f142-4052-ab75-7456b4d513eb',
    1,
    2
  ),
  (
    3,
    'Capitulo 3',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_2.mp4?alt=media&token=7c1b4c8e-5b78-48e8-ab4f-e4b43571998a',
    1,
    2
  ),
  (
    4,
    'Capitulo 4',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_1.mp4?alt=media&token=6e063fc6-29c2-4522-be07-2e45cdbd81e8',
    1,
    2
  ),
  (
    5,
    'Capitulo 5',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2Fvideo.mp4?alt=media&token=5f0bae50-f142-4052-ab75-7456b4d513eb',
    1,
    2
  ),
  (
    6,
    'Capitulo 6',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_2.mp4?alt=media&token=7c1b4c8e-5b78-48e8-ab4f-e4b43571998a',
    1,
    2
  ),
  (
    7,
    'Capitulo 1',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2Fvideo.mp4?alt=media&token=5f0bae50-f142-4052-ab75-7456b4d513eb',
    2,
    2
  ),
  (
    8,
    'Capitulo 2',
    'https://firebasestorage.googleapis.com/v0/b/imagenes-cd065.appspot.com/o/Peliculas%2FVideos%2FCap_2.mp4?alt=media&token=7c1b4c8e-5b78-48e8-ab4f-e4b43571998a',
    2,
    2
  );