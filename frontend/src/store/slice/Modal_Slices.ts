import { createSlice } from "@reduxjs/toolkit";

export const Modal_Slices = createSlice({
  name: "modal",
  initialState: {
    openModal_Auth: false,
    openModal_Movie: false,
    openModal_Content: false,
  },
  reducers: {
    openModal(state, actions) {
      switch (actions.payload) {
        case 1:
          state.openModal_Auth = true;
          break;
        case 2:
          state.openModal_Movie = true;
          break;
        case 3:
          state.openModal_Content = true;
          break;
      }
    },
    closeModal(state, actions) {
      switch (actions.payload) {
        case 1:
          state.openModal_Auth = false;
          break;
        case 2:
          state.openModal_Movie = false;
          break;
        case 3:
          state.openModal_Content = false;
          break;
      }
    },
  },
});

export const { openModal, closeModal } = Modal_Slices.actions;
