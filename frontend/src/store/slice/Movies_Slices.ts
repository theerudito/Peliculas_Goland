import { createSlice } from "@reduxjs/toolkit";
import { _movies } from "../../models/Movies";
import { Movies_List } from "../../helpers/Data";

export const Movies_Slices = createSlice({
  name: "movies",
  initialState: {
    form_Movies: _movies,
    data_Movies: Movies_List,
  },
  reducers: {
    addMovies(state, actions) {
      state.form_Movies = actions.payload;
    },
    getMovies(state, actions) {
      state.data_Movies = actions.payload;
    },
  },
});

export const { addMovies, getMovies } = Movies_Slices.actions;
