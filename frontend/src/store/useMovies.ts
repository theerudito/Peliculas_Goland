import { create } from "zustand";
import { _movies, Movies, MoviesDTO } from "../models/Movies";
import { GET_Movies, POST_Movie } from "../helpers/Fetching_Movies";

type Data = {
  form_movie: Movies;
  list_movies: MoviesDTO[];
  getMovies: () => void;
  postMovies: (obj: Movies) => void;
  reset: () => void;
};

export const useMovies = create<Data>((set, get) => ({
  form_movie: _movies,
  list_movies: [],

  getMovies: async () => {
    const result = await GET_Movies();

    if (result.data && Array.isArray(result.data)) {
      set({
        list_movies: result.data.length === 0 ? [] : result.data,
      });
    } else {
      set({ list_movies: [] });
    }
  },

  postMovies: async (obj: Movies) => {
    const result = await POST_Movie(obj);

    if (result.success === true) {
      get().reset();
      get().getMovies();
      return result.data;
    }

    return result.error;
  },

  reset: () => set({ form_movie: _movies }),
}));
