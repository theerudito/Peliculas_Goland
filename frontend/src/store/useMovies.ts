import { create } from "zustand";
import { _movies, Movies, MoviesDTO } from "../models/Movies";
import { GET_Movies, POST_Movie } from "../helpers/Fetching_Movies";
import { Movies_List } from "../helpers/Data";

type Data = {
  // LISTADO
  list_movies: MoviesDTO[];

  // FUNCIONES
  getMovies: () => void;
  postMovies: (obj: Movies) => void;
  reset: () => void;

  // FORMULARIO
  form_movie: Movies;
};

export const useMovies = create<Data>((set, get) => ({
  list_movies: [],

  getMovies: async () => {
    const result = await GET_Movies();

    if (result.success === true && Array.isArray(result.data)) {
      set({ list_movies: result.data });
    } else {
      set({ list_movies: Movies_List });
    }

    return result.error;
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

  form_movie: _movies,
}));
