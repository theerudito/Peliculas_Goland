import { create } from "zustand";
import { _episodes, Episodes } from "../models/Episodes";
import { _content, Content, ContentDTO } from "../models/Contents";
import {
  GET_Content,
  GET_Content_Type,
  POST_Content,
} from "../helpers/Fetching_Content";

type Data = {
  // LISTADO
  list_content: ContentDTO[];
  list_content_type: ContentDTO[];

  // DATOS
  type_content: number;

  // FUNCIONES
  changeType: (type: number) => void;
  getContent_Type: () => void;
  getContent: () => void;
  postContent: (obj: Content) => void;
  reset: () => void;

  // FORMULARIOS
  form_content: Content;
  form_content_episode: Episodes;
};

export const useContent = create<Data>((set, get) => ({
  list_content: [],
  list_content_type: [],

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

  postContent: async (obj: Content) => {
    const result = await POST_Content(obj);

    console.log(result);

    if (result.success === true) {
      get().reset();
      get().getContent();
      get().getContent_Type();
      return result.data;
    }

    return result.error;
  },

  reset: () => {
    set({
      form_content: _content,
      form_content_episode: _episodes,
    });
  },

  form_content: _content,
  form_content_episode: _episodes,
}));
