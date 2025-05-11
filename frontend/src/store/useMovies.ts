import { create } from "zustand";
import { _movies, Movies, MoviesDTO } from "../models/Movies";
import {
  DELETE_Movie,
  GET_Find_Movies,
  GET_Movies,
  POST_Movie,
  PUT_Movie,
} from "../helpers/Fetching_Movies";
import { Movies_List } from "../helpers/Data";

type Data = {
  // LISTADO
  list_movies: MoviesDTO[];

  // VARABLES
  searhMovie: string;

  // FUNCIONES
  getMovies: () => void;
  findMovies: (value: string) => void;
  setSearchMovie: (value: string) => void;

  postMovies: (obj: Movies) => void;
  deleteMovies: (id: number) => void;
  putMovies: (obj: Movies, id: number) => void;
  reset: () => void;

  // FORMULARIO
  form_movie: Movies;
};

export const useMovies = create<Data>((set, get) => ({
  // LISTADO
  list_movies: [],
  setSearchMovie: (value: string) => set({ searhMovie: value }),

  // VARIABLES
  searhMovie: "",

  // FUNCIONES
  getMovies: async () => {
    const result = await GET_Movies();

    if (result.success === true && Array.isArray(result.data)) {
      set({ list_movies: result.data });
    } else {
      set({ list_movies: Movies_List });
    }

    return result.error;
  },

  findMovies: async (value: string) => {
    set({ searhMovie: value });

    if (value.trim() === "") {
      await get().getMovies();
      set({ searhMovie: "" });
      return;
    }

    const result = await GET_Find_Movies(value);

    if (result.success === true && Array.isArray(result.data)) {
      set({ list_movies: result.data });
    } else {
      set({ list_movies: Movies_List });
    }
    set({ searhMovie: "" });
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

  putMovies: async (obj: Movies, id: number) => {
    const result = await PUT_Movie(obj, id);

    if (result.success === true) {
      get().reset();
      get().getMovies();
      return result.data;
    }

    return result.error;
  },

  deleteMovies: async (id: number) => {
    const result = await DELETE_Movie(id);

    if (result.success === true) {
      get().getMovies();
      return result.data;
    }

    return result.error;
  },

  reset: () => set({ form_movie: _movies }),

  // FORMULARIOS
  form_movie: _movies,
}));
