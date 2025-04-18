import { create } from "zustand";
import { Movies_List } from "../helpers/Data";
import { _movies, Movies, MoviesDTO } from "../models/Movies";

type Data = {
  form_movie: Movies;
  list_movies: MoviesDTO[];
  getMovies: () => void;
  postMovies: (obj: Movies) => void;
  reset: () => void;
};

export const useMovies = create<Data>((set) => ({
  form_movie: _movies,
  list_movies: [],
  getMovies: () => {
    set({ list_movies: Movies_List });
  },
  postMovies: (obj: Movies) => {
    console.log(obj);
  },
  reset: () => set({ form_movie: _movies }),
}));
