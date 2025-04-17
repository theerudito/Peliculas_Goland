export interface Movies {
  movie_title: string;
  movie_year: number;
  movie_cover: string;
  movie_url: string;
  gender_id: number;
}
export interface MoviesDTO {
  movie_movie_id: number;
  movie_title: string;
  movie_year: number;
  movie_cover: string;
  movie_url: string;
  movie_gender: string;
}

export const _movies: Movies = {
  movie_title: "",
  movie_year: 0,
  movie_cover: "",
  movie_url: "",
  gender_id: 0,
};
