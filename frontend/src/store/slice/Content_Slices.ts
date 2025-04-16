import { createSlice } from "@reduxjs/toolkit";
import { _form, Episodes } from "../../models/Movies";
import { Content_List } from "../../helpers/Data";

export const Content_Slices = createSlice({
  name: "content",
  initialState: {
    form_Content: _form,
    data_Content: Content_List,
    episodes_list: [] as Episodes[],
  },
  reducers: {
    addContent(state, actions) {
      state.data_Content = actions.payload;
    },
    getContent(state, actions) {
      state.data_Content = actions.payload;
    },
    removeEpisode(state, action) {
      state.episodes_list = state.episodes_list.filter(
        (ep) => ep.episode_id !== action.payload
      );
    },
  },
});

export const { addContent, getContent, removeEpisode } = Content_Slices.actions;
