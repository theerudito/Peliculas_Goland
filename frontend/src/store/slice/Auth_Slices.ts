import { createSlice } from "@reduxjs/toolkit";
import { _users } from "../../models/Movies";

export const Auth_Slices = createSlice({
  name: "auth",
  initialState: {
    userData: _users,
  },
  reducers: {
    Login(state, actions) {
      state.userData = actions.payload;
    },
  },
});

export const { Login } = Auth_Slices.actions;
