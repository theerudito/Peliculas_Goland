CREATE TABLE
  IF NOT EXISTS genders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    descripcion TEXT NOT NULL
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

CREATE TABLE
  IF NOT EXISTS movies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    gender_id INTEGER NOT NULL
  );

INSERT INTO
  movies (title, year, gender_id)
VALUES
  ('HOMBRES DE HONOR', 1998, 1)