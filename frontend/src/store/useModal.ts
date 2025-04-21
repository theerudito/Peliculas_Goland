import { create } from "zustand";

type Data = {
  _modal_Auth: boolean;
  _modal_Movie: boolean;
  _modal_Content: boolean;
  OpenModal_Auth: (value: boolean) => void;
  OpenModal_Movie: (value: boolean) => void;
  OpenModal_Content: (value: boolean) => void;
  reset: () => void;
};

export const useModal = create<Data>((set) => ({
  _modal_Auth: false,
  _modal_Movie: false,
  _modal_Content: false,
  OpenModal_Auth: (value) => set({ _modal_Auth: value }),
  OpenModal_Movie: (value) => set({ _modal_Movie: value }),
  OpenModal_Content: (value) => set({ _modal_Content: value }),
  reset: () =>
    set({
      _modal_Auth: false,
      _modal_Movie: false,
      _modal_Content: false,
    }),
}));
