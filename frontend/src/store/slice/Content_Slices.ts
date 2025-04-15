import { createSlice } from "@reduxjs/toolkit";
import { _form } from "../../models/Movies";
import { Content_List } from "../../helpers/Data";

export const Content_Slices = createSlice({
  name: "content",
  initialState: {
    form_Content: _form,
    data_Content: Content_List,
  },
  reducers: {
    addContent(state, actions) {
      state.data_Content = actions.payload;
    },
  },
});

export const { addContent } = Content_Slices.actions;
