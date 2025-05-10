import { create } from "zustand";
import { _episodes, EpisodeDTO } from "../models/Episodes";
import {
  _content,
  _content_episodes,
  Content,
  Content_Full_Data,
  ContentDTO,
  ContentDTO_EpisodeDTO,
} from "../models/Contents";
import {
  GET_Content,
  GET_Content_Episodes,
  GET_Content_Type,
  POST_Content,
  POST_Content_Episodes,
} from "../helpers/Fetching_Content";

type Data = {
  // LISTADO
  list_content: ContentDTO[];
  list_content_type: ContentDTO[];
  list_episodes: EpisodeDTO[];
  list_full_data: Content_Full_Data | null;
  loading: boolean;
  contend_id: number;

  // DATOS
  episode_id: number;
  type_content: number;
  openContent: boolean;

  // FUNCIONES
  openContent_Type: (open: boolean, id: number) => void;
  changeType: (type: number) => void;
  getContent_Type: () => void;
  getContent: () => void;
  getContent_full: () => void;
  postContent: (obj: Content) => void;
  postEpisodes: (obj: ContentDTO_EpisodeDTO) => void;
  add_episodes: (obj: EpisodeDTO) => void;
  remove_episodes: () => void;
  obtener_episode: (episode_id: number) => void;
  reset: () => void;

  // FORMULARIOS
  form_content: Content;
  form_episode: EpisodeDTO;
  form_content_episodes: ContentDTO_EpisodeDTO;
};

export const useContent = create<Data>((set, get) => ({
  list_content: [],
  list_content_type: [],
  list_episodes: [],
  list_full_data: null,
  loading: false,
  openContent: false,

  episode_id: 0,
  type_content: 1,
  contend_id: 0,

  openContent_Type: (open: boolean, id: number) => {
    set({ openContent: open, contend_id: id });
  },

  changeType: (type) => {
    set({ type_content: type });
    get().getContent_Type();
  },

  getContent: async () => {
    const result = await GET_Content();

    if (result.success === true && Array.isArray(result.data)) {
      set({ list_content: result.data });
    } else {
      set({ list_content: [] });
    }
  },

  getContent_Type: async () => {
    const result = await GET_Content_Type(get().type_content);

    if (result.success === true && Array.isArray(result.data)) {
      set({ list_content_type: result.data });
    } else {
      set({ list_content_type: [] });
    }
  },

  getContent_full: async () => {
    const result = await GET_Content_Episodes(get().contend_id);
    if (result.success === true) {
      set({ list_full_data: result.data, loading: false });
    }
  },

  postContent: async (obj: Content) => {
    const result = await POST_Content(obj);

    if (result.success === true) {
      get().reset();
      get().getContent();
      get().getContent_Type();
      return result.data;
    }

    return result.error;
  },

  postEpisodes: async (obj: ContentDTO_EpisodeDTO) => {
    const result = await POST_Content_Episodes(obj);

    console.log(result);

    if (result.success === true) {
      get().reset();
      get().getContent();
      get().getContent_Type();
      return result.data;
    }
    return result.error;
  },

  add_episodes: (obj: EpisodeDTO) => {
    const list = get().list_episodes;

    let episode_id = obj.episode_id;

    if (!episode_id || list.some((ep) => ep.episode_id === episode_id)) {
      const maxId = list.reduce((max, ep) => Math.max(max, ep.episode_id), 0);
      episode_id = maxId + 1;
    }

    const newEpisode: EpisodeDTO = {
      ...obj,
      episode_id,
    };

    set((state) => ({
      list_episodes: [...state.list_episodes, newEpisode],
    }));

    get().reset();
  },

  remove_episodes: () => {
    const { list_episodes, episode_id } = get();

    if (list_episodes.length <= 0) {
      set({ episode_id: 0 });
      return;
    }

    const updatedList = list_episodes.filter(
      (ep) => ep.episode_id !== episode_id
    );

    set({
      list_episodes: updatedList,
      episode_id: updatedList.length === 0 ? 0 : episode_id,
    });
  },

  obtener_episode: (episode_id: number) => {
    set({ episode_id });
  },

  reset: () => {
    set({ form_content: _content, form_episode: _episodes });
  },

  form_content: _content,
  form_episode: _episodes,
  form_content_episodes: _content_episodes,
}));
