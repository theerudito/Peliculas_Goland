import { createSlice } from "@reduxjs/toolkit";
import { _users } from "../../models/Movies";

export const Auth_Slices = createSlice({
  name: "auth",
  initialState: {
    form_Auth: _users,
  },
  reducers: {
    Login(state, actions) {
      state.form_Auth = actions.payload;
    },
  },
});

export const { Login } = Auth_Slices.actions;
