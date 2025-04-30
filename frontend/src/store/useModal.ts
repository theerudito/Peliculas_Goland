import { create } from "zustand";

type Data = {
  _modal_login: boolean;
  _modal_movie: boolean;
  _modal_content: boolean;
  _modal_gender: boolean;
  _modal_season: boolean;
  _modal_episodes: boolean;
  OpenModal_Login: (value: boolean) => void;
  OpenModal_Movie: (value: boolean) => void;
  OpenModal_Content: (value: boolean) => void;
  OpenModal_Gender: (value: boolean) => void;
  OpenModal_Season: (value: boolean) => void;
  OpenModal_Episodes: (value: boolean) => void;
  reset: () => void;
};

export const useModal = create<Data>((set) => ({
  _modal_login: false,
  _modal_movie: false,
  _modal_content: false,
  _modal_gender: false,
  _modal_season: false,
  _modal_episodes: false,
  OpenModal_Login: (value) => set({ _modal_login: value }),
  OpenModal_Movie: (value) => set({ _modal_movie: value }),
  OpenModal_Content: (value) => set({ _modal_content: value }),
  OpenModal_Gender: (value) =>
    set({ _modal_gender: value, _modal_movie: !value, _modal_content: !value }),
  OpenModal_Season: (value) =>
    set({ _modal_season: value, _modal_episodes: !value }),
  OpenModal_Episodes: (value) => set({ _modal_episodes: value }),
  reset: () =>
    set({
      _modal_login: false,
      _modal_movie: false,
      _modal_content: false,
      _modal_gender: false,
      _modal_season: false,
      _modal_episodes: false,
    }),
}));
