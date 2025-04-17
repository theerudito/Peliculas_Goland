import { create } from "zustand";

type Data = {
  openModal_Auth: boolean;
  openModal_Movie: boolean;
  openModal_Content: boolean;
  OpenModal_Auth: (value: boolean) => void;
  OpenModal_Movie: (value: boolean) => void;
  OpenModal_Content: (value: boolean) => void;
  reset: () => void;
};

export const useModal = create<Data>((set) => ({
  openModal_Auth: false,
  openModal_Movie: false,
  openModal_Content: false,
  OpenModal_Auth: (value) => set({ openModal_Auth: value }),
  OpenModal_Movie: (value) => set({ openModal_Movie: value }),
  OpenModal_Content: (value) => set({ openModal_Content: value }),
  reset: () =>
    set({
      openModal_Auth: false,
      openModal_Movie: false,
      openModal_Content: false,
    }),
}));
