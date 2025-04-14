import { combineReducers, configureStore } from "@reduxjs/toolkit";
import { ModalSlices } from "./slice/ModalSlices";

const rootReducer = combineReducers({
  modal: ModalSlices.reducer,
});

export type RootState = ReturnType<typeof rootReducer>;

export const store = configureStore({
  reducer: rootReducer,
});
