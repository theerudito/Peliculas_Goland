import { create } from "zustand";
import { _episodes, Episodes } from "../models/Episodes";
import {
  _content,
  Content,
  ContentDataDTO,
  ContentDTO,
} from "../models/Contents";
import {
  GET_Content,
  GET_Content_Epidodes,
  GET_Content_Season,
  GET_Content_Type,
} from "../helpers/Fetching_Content";
import { Content_Type_List } from "../helpers/Data";

type Data = {
  form_content: Content;
  form_Episode: Episodes;
  list_content_type: ContentDTO[];
  list_content: ContentDTO[];
  data_content: ContentDataDTO[];
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
    const response = await GET_Content_Type(get().type_content);
    set({
      list_content_type: response === null ? Content_Type_List : response,
    });
  },

  getContent: async () => {
    const response = await GET_Content();
    set({
      list_content: response === null ? Content_Type_List : response,
    });
  },

  getSeason: async (value: string) => {
    const response = await GET_Content_Season(value);
    set({
      listSeason: Array.isArray(response) ? response : [],
    });
  },

  getEpisode: async (value: string, id: number) => {
    try {
      const response = await GET_Content_Epidodes(value, id);

      set({
        data_content: Array.isArray(response) ? response : [],
      });
    } catch (error) {
      console.error("Error al obtener episodios:", error);
      set({ data_content: [] });
    }
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
