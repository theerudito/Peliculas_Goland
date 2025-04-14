import { createSlice } from "@reduxjs/toolkit";

export const ModalSlices = createSlice({
  name: "modal",
  initialState: {
    isOpen: false,
  },
  reducers: {
    openModal(state, actions) {
      switch (actions.payload) {
        case 1:
          state.isOpen = true;
          break;
        case 2:
          state.isOpen = true;
          break;
        case 3:
          state.isOpen = true;
          break;
      }
    },
    closeModal(state, actions) {
      switch (actions.payload) {
        case 1:
          state.isOpen = false;
          break;
        case 2:
          state.isOpen = false;
          break;
        case 3:
          state.isOpen = false;
          break;
      }
    },
  },
});

export const { openModal, closeModal } = ModalSlices.actions;
