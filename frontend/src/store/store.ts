import { combineReducers, configureStore } from "@reduxjs/toolkit";
import { Modal_Slices } from "./slice/Modal_Slices";
import { Auth_Slices } from "./slice/Auth_Slices";
import { Movies_Slices } from "./slice/Movies_Slices";
import { Content_Slices } from "./slice/Content_Slices";
import { Data_Slices } from "./slice/Data_Slices";

const rootReducer = combineReducers({
  modal: Modal_Slices.reducer,
  auth: Auth_Slices.reducer,
  movie: Movies_Slices.reducer,
  content: Content_Slices.reducer,
  data: Data_Slices.reducer,
});

export type RootState = ReturnType<typeof rootReducer>;

export const store = configureStore({
  reducer: rootReducer,
});
