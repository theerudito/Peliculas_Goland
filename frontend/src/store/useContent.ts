import { create } from "zustand";
import { _episodes, Episodes } from "../models/Episodes";
import { _content, Content, ContentDTO } from "../models/Contents";

type Data = {
  form_content: Content;
  form_Episode: Episodes;
  list_content_type: ContentDTO[];
  listEpisode: Episodes[];
  getContent_Types: () => void;
  addEpisode: (episode: Episodes) => void;
  removeEpisode: (id: number) => void;
  postContent_Types: () => void;
  reset: () => void;
};

export const useContent = create<Data>((set) => ({
  form_content: _content,
  form_Episode: _episodes,
  list_content_type: [],
  listEpisode: [],
  getContent_Types: () => {
    set({ list_content_type: [] });
  },

  postContent_Types: () => {},

  addEpisode: (episode: Episodes) => {
    set((state) => ({
      listEpisode: [...state.listEpisode, episode],
    }));
  },

  removeEpisode: (id: number) => {
    set((state) => ({
      listEpisode: state.listEpisode.filter((ep) => ep.episode_id !== id),
    }));
  },

  reset: () => {
    set({
      form_content: _content,
      form_Episode: _episodes,
      listEpisode: [],
    });
  },
}));
