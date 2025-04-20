import { create } from "zustand";
import { _episodes, Episodes } from "../models/Episodes";
import { _content, Content, ContentDTO } from "../models/Contents";
import { GET_Content } from "../helpers/Fetching_Content";
import { Content_Type_List } from "../helpers/Data";

type Data = {
  form_content: Content;
  form_Episode: Episodes;
  list_content: ContentDTO[];
  listEpisode: Episodes[];
  type_content: number;
  changeType: (type: number) => void;
  getContent: () => void;
  addEpisode: (episode: Episodes) => void;
  removeEpisode: (id: number) => void;
  postContent: () => void;
  reset: () => void;
};

export const useContent = create<Data>((set, get) => ({
  form_content: _content,
  form_Episode: _episodes,
  list_content: [],
  listEpisode: [],
  type_content: 1,

  changeType: (type) => {
    set({ type_content: type });
    get().getContent();
  },

  getContent: async () => {
    const response = await GET_Content(get().type_content);
    set({
      list_content: response === null ? Content_Type_List : response,
    });
  },

  postContent: () => {},

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
