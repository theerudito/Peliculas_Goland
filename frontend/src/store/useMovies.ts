import { create } from "zustand";
import { _movies, Movies, MoviesDTO } from "../models/Movies";
import { GET_Movies, POST_Movie } from "../helpers/Fetching_Movies";
import { Movies_List } from "../helpers/Data";

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

  getMovies: async () => {
    const response = await GET_Movies();
    set({ list_movies: response === null ? Movies_List : response });
  },

  postMovies: async (obj: Movies) => {
    await POST_Movie(obj);
  },

  reset: () => set({ form_movie: _movies }),
}));
