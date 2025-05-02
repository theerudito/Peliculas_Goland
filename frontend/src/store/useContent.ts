import { create } from "zustand";
import { _episodes, Episodes } from "../models/Episodes";
import { _content, Content, ContentDTO } from "../models/Contents";
import {
  GET_Content,
  GET_Content_Season,
  GET_Content_Type,
} from "../helpers/Fetching_Content";

type Data = {
  form_content: Content;
  form_Episode: Episodes;
  list_content_type: ContentDTO[];
  list_content: ContentDTO[];
  listEpisode: Episodes[];
  listSeason: ContentDTO[];
  type_content: number;

  changeType: (type: number) => void;
  getContent_Type: () => void;
  getContent: () => void;
  getSeason: (value: string) => void;
  getEpisode: (value: string, id: number) => void;
  addEpisode: (episode: Episodes) => void;
  removeEpisode: (id: number) => void;
  postContent: () => void;

  reset: () => void;
};

export const useContent = create<Data>((set, get) => ({
  form_content: _content,
  form_Episode: _episodes,
  list_content: [],
  list_content_type: [],
  listEpisode: [],
  listSeason: [],
  data_content: [],

  type_content: 1,

  changeType: (type) => {
    set({ type_content: type });
    get().getContent_Type();
  },

  getContent_Type: async () => {
    const result = await GET_Content_Type(get().type_content);

    if (result.data && Array.isArray(result.data)) {
      set({
        list_content_type: result.data.length === 0 ? [] : result.data,
      });
    } else {
      set({ list_content_type: [] });
    }
  },

  getContent: async () => {
    const result = await GET_Content();

    if (result.data && Array.isArray(result.data)) {
      set({
        list_content: result.data.length === 0 ? [] : result.data,
      });
    } else {
      set({ list_content: [] });
    }
  },

  getSeason: async (value: string) => {
    const response = await GET_Content_Season(value);
    set({
      listSeason: Array.isArray(response) ? response : [],
    });
  },

  getEpisode: async () => {},

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
